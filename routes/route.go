package routes

import (
	"apiantrean/controllers"
	"apiantrean/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(r *fiber.App) {
	antrean := r.Group("/api")

	antrean.Get("/", middlewares.AuthMiddleware, controllers.Index)
	antrean.Get("/:id", middlewares.AuthMiddleware, controllers.Show)
	antrean.Post("/", middlewares.AuthMiddleware, controllers.Create)
	antrean.Put("/:id", middlewares.AuthMiddleware, controllers.Update)
}