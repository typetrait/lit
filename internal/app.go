package internal

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	awsS3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"
	mediaAPI "github.com/typetrait/lit/internal/api/media"
	postAPI "github.com/typetrait/lit/internal/api/post"
	"github.com/typetrait/lit/internal/app/media"
	"github.com/typetrait/lit/internal/app/post"
	settings2 "github.com/typetrait/lit/internal/app/settings"
	"github.com/typetrait/lit/internal/infrastructure"
	"github.com/typetrait/lit/internal/infrastructure/content"
	"github.com/typetrait/lit/internal/infrastructure/model"
	"github.com/typetrait/lit/internal/infrastructure/s3"
	"github.com/typetrait/lit/internal/web"
	"github.com/typetrait/lit/internal/web/about"
	"github.com/typetrait/lit/internal/web/home"
	"github.com/typetrait/lit/internal/web/posts"
	"github.com/typetrait/lit/internal/web/rendering"
	"github.com/typetrait/lit/internal/web/sign_in"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	echo        *echo.Echo
	environment *Environment
}

func NewApp(environment *Environment) *App {
	return &App{
		echo:        echo.New(),
		environment: environment,
	}
}

func (app *App) Start(address string) {
	app.echo.Debug = app.environment.IsDebugEnabled
	app.registerRoutes()

	app.echo.Logger.Fatal(
		app.echo.Start(address),
	)
}

func (app *App) registerRoutes() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		app.environment.DBHost,
		app.environment.DBUser,
		app.environment.DBPassword,
		app.environment.DBName,
		app.environment.DBPort,
	)
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		app.echo.Logger.Fatal(err)
	}

	err = db.AutoMigrate(&model.Settings{}, &model.Role{}, &model.User{}, &model.Post{}, &model.Media{})
	if err != nil {
		app.echo.Logger.Fatal(err)
	}

	contentRenderer := rendering.NewContentRenderer(
		rendering.NewMarkdownRenderer(),
	)

	mediaDetector := content.NewDetector()

	ctx := context.Background()
	s3Config, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		app.echo.Logger.Fatal(err)
		return
	}

	s3Client := awsS3.NewFromConfig(s3Config, func(o *awsS3.Options) {
		o.BaseEndpoint = aws.String(app.environment.LocalstackHost)
		o.UsePathStyle = true
	})

	uploader := manager.NewUploader(s3Client)

	settingsRepository := infrastructure.NewSettingsRepository(db)
	settingsProvider := settings2.NewProvider(settingsRepository)
	app.echo.Renderer = web.NewTemplate(settingsProvider)

	postRepository := infrastructure.NewPostRepository(db)
	mediaRepository := infrastructure.NewMediaRepository(db)
	mediaStorage := s3.NewMediaStorage(s3Client, uploader, app.environment.S3Bucket)

	getPost := post.NewGetPost(postRepository)
	getPosts := post.NewGetPosts(postRepository)
	createPost := post.NewCreatePost(postRepository)

	getMedia := media.NewGetMedia(mediaStorage, mediaRepository)
	uploadMedia := media.NewUploadMedia(mediaStorage, postRepository, mediaRepository, mediaDetector)

	homeH := home.NewHandler(getPosts, contentRenderer)
	aboutH := about.NewHandler()
	postsH := posts.NewHandler(getPost, getPosts, contentRenderer)
	signInH := sign_in.NewHandler()

	postAPIH := postAPI.NewAPIHandler(createPost)
	mediaAPIH := mediaAPI.NewAPIHandler(getMedia, uploadMedia)

	// Blog
	app.echo.GET("/", homeH.Get())
	app.echo.GET("/about", aboutH.Get())

	app.echo.GET("/posts", postsH.List())
	app.echo.GET("/posts/:slug", postsH.View())

	app.echo.GET("/sign-in", signInH.Get())
	app.echo.POST("/sign-in", signInH.Post())

	// API (/api)
	apiGroup := app.echo.Group("api")
	apiGroup.POST("/posts", postAPIH.Draft())
	apiGroup.PATCH("/posts/:id", postAPIH.Publish())

	apiGroup.GET("/posts/:post_id/media/:media_id", mediaAPIH.Get())
	apiGroup.POST("/posts/:id/media", mediaAPIH.Upload())
}
