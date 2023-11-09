package mq

import (
	"context"
	"harmoni/app/harmoni/internal/infrastructure/config"
	"harmoni/app/harmoni/internal/server/mq/group/comment"
	"harmoni/app/harmoni/internal/server/mq/group/like"
	"harmoni/app/harmoni/internal/server/mq/group/post"
	"harmoni/app/harmoni/internal/server/mq/group/user"
	"harmoni/app/harmoni/internal/types/iface"
	commentevent "harmoni/app/harmoni/internal/usecase/comment/events"
	likeevent "harmoni/app/harmoni/internal/usecase/like/events"
	postevent "harmoni/app/harmoni/internal/usecase/post/events"
	userevent "harmoni/app/harmoni/internal/usecase/user/events"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/garsue/watermillzap"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var _ iface.Executor = (*MQExecutor)(nil)

type MQExecutor struct {
	*message.Router
}

var ProviderSetMQ = wire.NewSet(
	NewMQRouter,
	NewExecutor,
)

func (r *MQExecutor) Start() error {
	return r.Run(context.Background())
}

func (r *MQExecutor) Shutdown() error {
	return r.Close()
}

func NewExecutor(r *message.Router) *MQExecutor {
	return &MQExecutor{
		Router: r,
	}
}

func NewMQRouter(
	conf *config.MessageQueue,
	commentEventsHandler *commentevent.CommentEventsHandler,
	likeevent *likeevent.LikeEventsHandler,
	postEventsHandler *postevent.PostEventsHandler,
	userevent *userevent.UserEventsHandler,
	logger *zap.Logger,
) (*message.Router, error) {
	r, err := message.NewRouter(message.RouterConfig{}, watermillzap.NewLogger(logger))
	if err != nil {
		return nil, err
	}
	r.AddMiddleware(middleware.Retry{
		MaxRetries: 3,
	}.Middleware)
	{
		err = comment.NewCommentGroup(conf, r, commentEventsHandler, logger)
		if err != nil {
			return nil, err
		}
	}
	{
		err = like.NewLikeGroup(conf, r, likeevent, logger)
		if err != nil {
			return nil, err
		}
	}
	{
		err = post.NewPostGroup(conf, r, postEventsHandler, logger)
		if err != nil {
			return nil, err
		}
	}
	{
		err = user.NewUserGroup(conf, r, userevent, logger)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}