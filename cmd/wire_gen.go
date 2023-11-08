// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"go.uber.org/zap"
	"harmoni/internal/cron"
	"harmoni/internal/handler"
	"harmoni/internal/infrastructure/config"
	"harmoni/internal/infrastructure/data"
	"harmoni/internal/infrastructure/mq/publisher"
	"harmoni/internal/pkg/app"
	"harmoni/internal/pkg/filesystem"
	"harmoni/internal/pkg/logger"
	"harmoni/internal/pkg/middleware"
	"harmoni/internal/pkg/snowflakex"
	"harmoni/internal/repository/auth"
	"harmoni/internal/repository/comment"
	"harmoni/internal/repository/email"
	file2 "harmoni/internal/repository/file"
	"harmoni/internal/repository/follow"
	"harmoni/internal/repository/like"
	"harmoni/internal/repository/post"
	"harmoni/internal/repository/tag"
	"harmoni/internal/repository/unique"
	user2 "harmoni/internal/repository/user"
	"harmoni/internal/server/http"
	"harmoni/internal/server/mq"
	"harmoni/internal/service"
	comment2 "harmoni/internal/usecase/comment"
	"harmoni/internal/usecase/comment/events"
	email2 "harmoni/internal/usecase/email"
	"harmoni/internal/usecase/file"
	follow2 "harmoni/internal/usecase/follow"
	like2 "harmoni/internal/usecase/like"
	events2 "harmoni/internal/usecase/like/events"
	post2 "harmoni/internal/usecase/post"
	events3 "harmoni/internal/usecase/post/events"
	tag2 "harmoni/internal/usecase/tag"
	"harmoni/internal/usecase/timeline"
	"harmoni/internal/usecase/user"
	events4 "harmoni/internal/usecase/user/events"
)

// Injectors from wire.go:

func initApplication(conf *config.Config, appConf *config.App, dbconf *config.DB, rdbconf *config.Redis, authConf *config.Auth, emailConf *config.Email, messageConf *config.MessageQueue, fileConf *config.FileStorage, logConf *config.Log) (*app.Application, func(), error) {
	zapLogger, err := logger.NewZapLogger(logConf)
	if err != nil {
		return nil, nil, err
	}
	client, cleanup, err := data.NewRedis(rdbconf)
	if err != nil {
		return nil, nil, err
	}
	sugaredLogger := sugar(zapLogger)
	authRepo := auth.NewAuthRepo(client, sugaredLogger)
	authUseCase := user.NewAuthUseCase(authConf, authRepo, sugaredLogger)
	db, cleanup2, err := data.NewDB(dbconf, zapLogger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	node, err := snowflakex.NewSnowflakeNode(appConf)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	uniqueIDRepo := unique.NewUniqueIDRepo(node)
	userRepo := user2.NewUserRepo(db, client, uniqueIDRepo, sugaredLogger)
	emailRepo := email.NewEmailRepo(client)
	emailUsecase, err := email2.NewEmailUsecase(emailConf, emailRepo, sugaredLogger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	likeRepo := like.NewLikeRepo(db, client, userRepo, sugaredLogger)
	policy := file.NewPolicy(fileConf)
	fileSystem, err := filesystem.NewFileSystem(policy, client)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	fileRepo := file2.NewFileRepository(db, client, uniqueIDRepo, sugaredLogger)
	fileUseCase := file.NewFileUseCase(appConf, client, fileConf, fileSystem, fileRepo, sugaredLogger)
	userUseCase := user.NewUserUseCase(likeRepo, userRepo, authUseCase, fileUseCase, sugaredLogger)
	accountUsecase := user.NewAccountUsecase(emailConf, authUseCase, userRepo, emailUsecase, userUseCase, zapLogger)
	accountService := service.NewAccountService(accountUsecase, sugaredLogger)
	jwtAuthMiddleware := middleware.NewJwtAuthMiddleware(authUseCase)
	accountHandler := handler.NewAccountHandler(accountService, jwtAuthMiddleware, sugaredLogger)
	tagRepo := tag.NewTagRepo(db, client, uniqueIDRepo, sugaredLogger)
	followRepo := follow.NewFollowRepo(db, userRepo, tagRepo, uniqueIDRepo, sugaredLogger)
	tagUseCase := tag2.NewTagUseCase(tagRepo, sugaredLogger)
	followUseCase := follow2.NewFollowUseCase(followRepo, userUseCase, tagUseCase, sugaredLogger)
	followService := service.NewFollowService(followUseCase, userUseCase, sugaredLogger)
	followHandler := handler.NewFollowHandler(followService, sugaredLogger)
	fileService := service.NewFileService(fileUseCase, userUseCase, sugaredLogger)
	fileHandler := handler.NewFileHandler(fileService, sugaredLogger)
	userService := service.NewUserService(userUseCase, authUseCase, accountUsecase, sugaredLogger)
	userHandler := handler.NewUserHandler(userService)
	postRepo := post.NewPostRepo(db, client, tagRepo, uniqueIDRepo, sugaredLogger)
	postUseCase := post2.NewPostUseCase(postRepo, likeRepo, userUseCase, tagUseCase, sugaredLogger)
	postService := service.NewPostService(postUseCase, tagUseCase, sugaredLogger)
	postHandler := handler.NewPostHandler(postService)
	tagService := service.NewTagService(tagUseCase, sugaredLogger)
	tagHandler := handler.NewTagHandler(tagService)
	commentRepo := comment.NewCommentRepo(db, client, uniqueIDRepo, sugaredLogger)
	commentUseCase := comment2.NewCommentUseCase(commentRepo, likeRepo, userRepo, fileUseCase, sugaredLogger)
	commentService := service.NewCommentService(commentUseCase, sugaredLogger)
	commentHandler := handler.NewCommentHandler(commentService)
	messagePublisher, err := publisher.NewPublisher(messageConf, zapLogger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	jsonPublisher := publisher.NewJSONPublisher(messagePublisher)
	likeUsecase, cleanup3, err := like2.NewLikeUsecase(messageConf, likeRepo, postUseCase, commentRepo, userRepo, sugaredLogger, jsonPublisher)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	likeService := service.NewLikeUsecase(likeUsecase, sugaredLogger)
	likeHandler := handler.NewLikeHandler(likeService, sugaredLogger)
	timeLinePullUsecase := timeline.NewTimeLineUsecase(followRepo, postUseCase, sugaredLogger)
	timeLineService := service.NewTimeLineService(timeLinePullUsecase, sugaredLogger)
	timeLineHandler := handler.NewTimeLineHandler(timeLineService)
	harmoniAPIRouter := http.NewHarmoniAPIRouter(accountHandler, followHandler, fileHandler, userHandler, postHandler, tagHandler, commentHandler, likeHandler, timeLineHandler)
	fiberExecutor := http.NewHTTPServer(appConf, zapLogger, harmoniAPIRouter, jwtAuthMiddleware)
	scheduledTaskManager, cleanup4, err := cron.NewScheduledTaskManager(jsonPublisher, likeUsecase, sugaredLogger)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	commentEventsHandler := events.NewCommentEventsHandler(commentRepo)
	likeEventsHandler := events2.NewLikeEventsHandler(likeRepo)
	postEventsHandler := events3.NewPostEventsHandler(postRepo)
	userEventsHandler := events4.NewUserEventsHandler(userRepo)
	router, err := mq.NewMQRouter(messageConf, commentEventsHandler, likeEventsHandler, postEventsHandler, userEventsHandler, zapLogger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	mqExecutor := mq.NewExecutor(router)
	application := newApplication(conf, fiberExecutor, scheduledTaskManager, mqExecutor, sugaredLogger)
	return application, func() {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

func sugar(l *zap.Logger) *zap.SugaredLogger {
	return l.Sugar()
}
