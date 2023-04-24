package post

import (
	"bytes"
	"context"
	"database/sql/driver"
	"fmt"
	"harmoni/internal/entity/paginator"
	tagentity "harmoni/internal/entity/tag"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Int64toString []int64

func (is Int64toString) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	b.Grow(64)
	b.WriteByte('[')
	for i, v := range is {
		b.WriteByte('"')
		b.Write([]byte(strconv.FormatInt(v, 10)))
		b.WriteByte('"')
		if i+1 != len(is) {
			b.WriteByte(',')
		}
	}
	b.WriteByte(']')
	return b.Bytes(), nil
}

func (is *Int64toString) UnmarshalJSON(data []byte) error {
	str := strings.Split(strings.Trim(string(data), "[]"), ",")
	for _, v := range str {
		num, err := strconv.ParseInt(strings.Trim(v, " \""), 10, 64)
		if err != nil {
			fmt.Println(err)
			return err
		}
		*is = append(*is, num)
	}
	return nil
}

// Scan 方法实现了 sql.Scanner 接口
func (is *Int64toString) Scan(v interface{}) error {
	if s, ok := v.([]uint8); !ok {
		return fmt.Errorf("断言失败")
	} else {
		str := strings.Split(string(s), ",")
		for _, v := range str {
			num, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return err
			}
			*is = append(*is, num)
		}
	}
	return nil
}

func (is Int64toString) Value() (driver.Value, error) {
	if len(is) == 0 {
		return nil, nil
	}
	s := make([]string, 0, 4)
	for _, v := range is {
		s = append(s, strconv.FormatInt(v, 10))
	}
	return strings.Join(s, ","), nil
}

type Post struct {
	gorm.Model
	Status    int32         `gorm:"not null"`
	PostID    int64         `gorm:"uniqueIndex"`
	AuthorID  int64         `gorm:"index"`
	TagID     Int64toString `gorm:"type:varchar(128)"`
	TagName   string        `gorm:"type:varchar(512);index:,class:FULLTEXT,option:WITH PARSER ngram"`
	Title     string        `gorm:"type:varchar(128)"`
	Content   string        `gorm:"type:varchar(512)"`
	LikeCount int64         `gorm:"-"`
}

type PostDetail struct {
	Status    int32               `json:"status"`
	Like      int64               `json:"like"`
	PostID    int64               `json:"post_id,string"`
	AuthorID  int64               `json:"author_id,string"`
	Tags      []tagentity.TagInfo `json:"tags"`
	Title     string              `json:"title"`
	Content   string              `json:"content"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
}

func ConvertPostToDisplayDetail(post *Post) PostDetail {
	tagNames := strings.Split(post.TagName, ",")
	tags := make([]tagentity.TagInfo, len(post.TagID))
	for i := 0; i < len(post.TagID); i++ {
		tags[i] = tagentity.TagInfo{
			TagName: tagNames[i],
			TagID:   post.TagID[i],
		}
	}

	pd := PostDetail{
		Status:    post.Status,
		PostID:    post.PostID,
		AuthorID:  post.AuthorID,
		Tags:      tags,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}

	return pd
}

type PostRepository interface {
	Create(ctx context.Context, post *Post) error
	GetByPostID(ctx context.Context, postID int64) (Post, error)
	BatchByIDs(ctx context.Context, postIDs []int64) ([]Post, error)
	GetPage(ctx context.Context, pageSize, pageNum int64, orderCond string) (paginator.Page[Post], error)
	LikePost(ctx context.Context, postID int64, userID int64, direction int8) error
}