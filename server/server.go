package server

import (
	"context"
	"fmt"
	"github.com/evanw/esbuild/pkg/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
	_ "google.golang.org/grpc"
	"strings"
	"time"
)

type Task struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedDt time.Time `json:"createdDt"`
}

type Report struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedDt time.Time `json:"createdDt"`
}

type Server struct {
	app *fiber.App
}

func NewServer(ctx context.Context, devMode bool) (*Server, error) {

	config := &fiber.Config{DisableStartupMessage: false}

	app := fiber.New(*config)
	app.Use(logger.New())
	app.Use(helmet.New())

	server := &Server{
		app: app,
	}

	app.Get("/api/tasks", func(c *fiber.Ctx) error {
		return c.JSON(
			[]Task{
				{ID: uuid.New(), Title: "Buy Food", Status: "New", CreatedDt: time.Now()},
				{ID: uuid.New(), Title: "Go To Sleep", Status: "In Progress", CreatedDt: time.Now()},
				{ID: uuid.New(), Title: "Eat Lunch", Status: "Completed", CreatedDt: time.Now()},
				{ID: uuid.New(), Title: "Work", Status: "Canceled", CreatedDt: time.Now()},
			})
	})
	app.Get("/api/reports", func(c *fiber.Ctx) error {
		return c.JSON(
			[]Report{
				{ID: uuid.New(), Title: "Buy Food", Author: "Jim", CreatedDt: time.Now()},
				{ID: uuid.New(), Title: "Go To Sleep", Author: "Bob", CreatedDt: time.Now()},
				{ID: uuid.New(), Title: "Eat Lunch", Author: "Steve", CreatedDt: time.Now()},
				{ID: uuid.New(), Title: "Work", Author: "Tim", CreatedDt: time.Now()},
			})
	})

	if !devMode {
		app.Static("/", "./public/")
		app.Static("/ui/*", "./public/")
		app.Static("/assets", "./public/assets")
	} else {
		log.Info().Msg("Setting up dev mode")
		buildContext, err := api.Context(api.BuildOptions{
			EntryPoints: []string{"web/index.js"},
			Bundle:      true,
			JSX:         api.JSXAutomatic,
			Loader: map[string]api.Loader{
				".js": api.LoaderJSX,
			},
			Outdir: "public/assets",
		})

		if err != nil {
			return nil, fmt.Errorf("error creating esbuild builder: %v", err)
		}

		ctxError := buildContext.Watch(api.WatchOptions{})
		if ctxError != nil {
			return nil, fmt.Errorf("error creating esbuild watch: %v", ctxError)
		}

		result, ctxError := buildContext.Serve(api.ServeOptions{
			Servedir: "public",
			OnRequest: func(args api.ServeOnRequestArgs) {

			},
		})

		if ctxError != nil {
			return nil, fmt.Errorf("error creating esbuild serve: %v", ctxError)
		}

		buildContext.Cancel()

		log.Info().Msgf("esbuild server listening on port: %d", result.Port)

		app.Get("/*", func(c *fiber.Ctx) error {
			uri := string(c.Request().RequestURI()[:])

			var stream = false

			if uri == "/esbuild" {
				stream = true
			} else if strings.HasPrefix(uri, "/ui/") {
				uri = "/"
			}

			url := fmt.Sprintf("http://localhost:%d%s", result.Port, uri)
			client := &fasthttp.Client{StreamResponseBody: stream}

			if err := proxy.Do(c, url, client); err != nil {
				return err
			}

			return nil
		})
	}

	return server, nil
}

func (s *Server) Start(ctx context.Context, addr string) error {

	return s.app.Listen(addr)
}

func (s *Server) Stop(ctx context.Context) error {

	err := s.app.ShutdownWithContext(ctx)

	if err != nil {
		return fmt.Errorf("error stopping server: %v, err")
	}

	return nil
}
