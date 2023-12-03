package bind

import (
	"time"

	fiber "github.com/gofiber/fiber/v2"
)

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("bind", fiber.Map{
			"CSRF":           c.Locals("csrf"),
			"PID":            c.Locals("pid"),
			"CopyrightYear":  time.Now().Year(),
			"Title":          "Petrichor",
			"MetaContent":    "Petrichor MUD - a modern take on a classic MUD style of game.",
			"CharactersView": c.Path() == "/characters",
			"ProfileView":    c.Path() == "/profile",
			"HomeView":       c.Path() == "/",
		})

		return c.Next()
	}
}
