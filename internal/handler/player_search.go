package handler

import (
	"context"
	"fmt"

	fiber "github.com/gofiber/fiber/v2"

	"petrichormud.com/app/internal/interfaces"
	"petrichormud.com/app/internal/layout"
	"petrichormud.com/app/internal/partial"
	"petrichormud.com/app/internal/view"
)

func SearchPlayer(i *interfaces.Shared) fiber.Handler {
	type input struct {
		Search string `form:"search"`
	}
	return func(c *fiber.Ctx) error {
		pid := c.Locals("pid")
		if pid == nil {
			c.Status(fiber.StatusUnauthorized)
			return c.Render(view.Login, view.Bind(c), layout.Standalone)
		}

		r := new(input)
		if err := c.BodyParser(r); err != nil {
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		searchStr := fmt.Sprintf("%%%s%%", r.Search)
		players, err := i.Queries.SearchPlayersByUsername(context.Background(), searchStr)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		dest := c.Params("dest")
		if len(dest) == 0 {
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if dest == "player-permissions" {
			// TODO: Move this to a constant and inject it
			b := view.Bind(c)
			b["Players"] = players
			c.Status(fiber.StatusOK)
			return c.Render(partial.PlayerPermissionsSearchResults, b, layout.None)
		}

		c.Status(fiber.StatusBadRequest)
		return nil
	}
}
