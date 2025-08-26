package internal

import (
	"fmt"

	"github.com/labstack/echo/v4"
	mediaAPI "github.com/typetrait/lit/internal/api/media"
	postAPI "github.com/typetrait/lit/internal/api/post"
	"github.com/typetrait/lit/internal/app/media"
	"github.com/typetrait/lit/internal/app/post"
	"github.com/typetrait/lit/internal/domain/settings"
	"github.com/typetrait/lit/internal/store"
	"github.com/typetrait/lit/internal/store/model"
	"github.com/typetrait/lit/internal/web"
	"github.com/typetrait/lit/internal/web/about"
	"github.com/typetrait/lit/internal/web/home"
	"github.com/typetrait/lit/internal/web/posts"
	"github.com/typetrait/lit/internal/web/rendering"
	"github.com/typetrait/lit/internal/web/sign_in"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MockSettingsProvider struct {
}

func (m MockSettingsProvider) Settings() settings.Settings {
	return settings.Settings{
		BlogName:           "Bruno C.",
		BlogSubtitle:       "software, games, music production",
		BlogAbout:          "I'm Bruno, a 26yo software developer living in the EU",
		AdditionalSettings: map[string]any{},
	}
}

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
	settingsProvider := MockSettingsProvider{}
	app.echo.Renderer = web.NewTemplate(settingsProvider)

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

	err = db.AutoMigrate(&model.Role{}, &model.User{}, &model.Post{}, &model.Media{})
	if err != nil {
		app.echo.Logger.Fatal(err)
	}

	contentRenderer := rendering.NewContentRenderer(
		rendering.NewMarkdownRenderer(),
	)

	postRepository := store.NewPostRepository(db)
	getPost := post.NewGetPost(postRepository)
	getPosts := post.NewGetPosts(postRepository)
	createPost := post.NewCreatePost(postRepository)
	uploadMedia := media.NewUploadMedia()

	homeH := home.NewHandler(getPosts, contentRenderer)
	aboutH := about.NewHandler()
	postsH := posts.NewHandler(getPost, getPosts, contentRenderer)
	signInH := sign_in.NewHandler()

	postAPIH := postAPI.NewAPIHandler(createPost)
	mediaAPIH := mediaAPI.NewAPIHandler(uploadMedia)

	// Blog
	app.echo.GET("/", homeH.Get())
	app.echo.GET("/about", aboutH.Get())

	app.echo.GET("/posts", postsH.List())
	app.echo.GET("/posts/:slug", postsH.View())

	app.echo.GET("/sign-in", signInH.Get())
	app.echo.POST("/sign-in", signInH.Post())

	// API
	apiGroup := app.echo.Group("api")
	apiGroup.POST("/posts", postAPIH.Draft())
	apiGroup.PATCH("/posts/:id", postAPIH.Publish())
	apiGroup.POST("/posts/:id/media", mediaAPIH.Upload())
}
