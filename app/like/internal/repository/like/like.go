package like

import (
	"context"
	"errors"

	objectv1 "harmoni/api/common/object/v1"
	entitylike "harmoni/app/like/internal/entity/like"
	polike "harmoni/app/like/internal/infrastructure/po/like"
	reasonlike "harmoni/app/like/internal/pkg/reason"
	"harmoni/internal/pkg/data"
	"harmoni/internal/pkg/errorx"
	"harmoni/internal/pkg/reason"
	"harmoni/internal/types/iface"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type LikeRepo struct {
	uniqueRepo iface.UniqueIDRepository
	data       *data.DB
	logger     *log.Helper
}

var _ entitylike.LikeRepository = (*LikeRepo)(nil)

func withUserID(userID int64) data.ScopeFunc {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", userID)
	}
}

func withTargetUserID(targetUserID int64) data.ScopeFunc {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("target_user_id = ?", targetUserID)
	}
}

func withLikeType(likeType entitylike.LikeType) data.ScopeFunc {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("like_type = ?", likeType)
	}
}

func withObjectID(objectID int64) data.ScopeFunc {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("object_id = ?", objectID)
	}
}

func findLike(like *entitylike.Like) data.ScopeFunc {
	return func(db *gorm.DB) *gorm.DB {
		return db.Scopes(
			withUserID(like.User.GetId()),
			withTargetUserID(like.TargetUser.GetId()),
			withLikeType(like.LikeType),
			withObjectID(like.ObjectID),
		)
	}
}

func NewLikeRepo(
	uniqueRepo iface.UniqueIDRepository,
	data *data.DB,
	logger log.Logger,
) *LikeRepo {
	return &LikeRepo{
		uniqueRepo: uniqueRepo,
		data:       data,
		logger: log.NewHelper(
			log.With(logger, "module", "repository/like", "service", "like")),
	}
}

func (r *LikeRepo) Save(ctx context.Context, like *entitylike.Like, isCancel bool) error {
	return r.data.DB(ctx).Transaction(func(tx *gorm.DB) error {
		if isCancel {
			err := tx.Scopes(findLike(like)).
				Delete(&polike.Like{}).Error
			if err != nil {
				return errorx.InternalServer(reason.DatabaseError).WithError(err).WithStack()
			}
			return nil
		}

		po := polike.FromDomain(like)
		model := &polike.Like{}
		err := tx.Debug().Model(model).Unscoped().Scopes(findLike(like)).First(model).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			po.LikingID, err = r.uniqueRepo.GenUniqueID(ctx)
			if err != nil {
				return err
			}
			err = tx.Table("like").Create(po).Error
			if err != nil {
				return errorx.InternalServer(reason.DatabaseError).WithError(err).WithStack()
			}
			return nil
		} else if err != nil {
			return errorx.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		}
		err = tx.Table(model.TableName()).
			Where("liking_id = ?", model.LikingID).
			Update("deleted_at", 0).Error
		if err != nil {
			return errorx.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		}

		return nil
	})
}

func (r *LikeRepo) IsExist(ctx context.Context, like *entitylike.Like) (bool, error) {
	var count int64
	err := r.data.DB(ctx).Model(&polike.Like{}).
		Scopes(findLike(like)).
		Count(&count).Error
	if err != nil {
		return false, errorx.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return count > 0, nil
}

func (r *LikeRepo) Get(ctx context.Context, like *entitylike.Like) (*entitylike.Like, error) {
	l := &entitylike.Like{}
	err := r.data.DB(ctx).Model(&polike.Like{}).
		Scopes(findLike(like)).
		First(l).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.NotFound(reasonlike.LikeNotExist)
		}
		return nil, errorx.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return l, nil
}

func (r *LikeRepo) ListLikeObjectByUserID(ctx context.Context, query *entitylike.ListLikeObjectQuery) ([]*entitylike.Like, int64, error) {
	likeList := make([]*polike.Like, 0, query.Size)
	err := r.data.DB(ctx).Scopes(
		withUserID(query.UserID),
		withLikeType(query.LikeType),
	).Find(&likeList).Error
	if err != nil {
		r.logger.Error(err)
		return nil, 0, errorx.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}

	var count int64
	err = r.data.DB(ctx).Model(&polike.Like{}).Scopes(
		withUserID(query.UserID),
		withLikeType(query.LikeType),
	).Count(&count).Error
	if err != nil {
		r.logger.Error(err)
		return nil, 0, errorx.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}

	return lo.Map(
		likeList,
		func(like *polike.Like, _ int) *entitylike.Like {
			return like.ToDomain()
		}), count, nil
}

func (r *LikeRepo) ListObjectLikedUserByObjectID(
	ctx context.Context,
	query *entitylike.ListObjectLikedUserQuery,
) ([]*entitylike.Like, int64, error) {
	likeList := make([]*polike.Like, 0, query.Size)
	err := r.data.DB(ctx).Scopes(
		withObjectID(query.ObjectID),
		withLikeType(query.LikeType),
	).Find(&likeList).Error
	if err != nil {
		r.logger.Error(err)
		return nil, 0, errorx.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}

	var count int64
	err = r.data.DB(ctx).Scopes(
		withObjectID(query.ObjectID),
		withLikeType(query.LikeType),
	).Count(&count).Error
	if err != nil {
		return nil, 0, errorx.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}

	return lo.Map(
		likeList,
		func(like *polike.Like, _ int) *entitylike.Like {
			return like.ToDomain()
		}), count, nil
}

func (r *LikeRepo) ObjectLikeCount(ctx context.Context, object *objectv1.Object) (*entitylike.LikeCount, error) {
	lc := &polike.LikeCount{}
	err := r.data.DB(ctx).Model(lc).
		Scopes(withObjectID(object.GetId())).
		Where("object_type = ?", object.GetType()).
		First(lc).Error
	if err != nil {
		return nil, errorx.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return &entitylike.LikeCount{
		Count:  lc.Counts,
		Object: object,
	}, nil
}

func (r *LikeRepo) ListObjectLikeCount(ctx context.Context, objectIDs []int64, objectType objectv1.ObjectType) (entitylike.LikeCountList, error) {
	lcList := make([]*polike.LikeCount, 0, 10)
	err := r.data.DB(ctx).Model(lcList).
		Where("object_id IN ?", objectIDs).
		Where("object_type = ?", objectType).
		Find(lcList).Error
	if err != nil {
		return nil, errorx.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return lo.Map(lcList, func(lc *polike.LikeCount, _ int) *entitylike.LikeCount {
		return &entitylike.LikeCount{
			Count: lc.Counts,
			Object: &objectv1.Object{
				Id:   lc.ObjectID,
				Type: lc.OjbectType,
			},
		}
	}), nil
}
