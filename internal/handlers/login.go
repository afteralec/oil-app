package handlers

import (
	"context"
	"fmt"
	"slices"

	fiber "github.com/gofiber/fiber/v2"

	"petrichormud.com/app/internal/password"
	"petrichormud.com/app/internal/permissions"
	"petrichormud.com/app/internal/shared"
	"petrichormud.com/app/internal/username"
)

const (
	LoginRoute = "/login"
)

func Login(i *shared.Interfaces) fiber.Handler {
	type request struct {
		Username string `form:"username"`
		Password string `form:"password"`
	}

	return func(c *fiber.Ctx) error {
		r := new(request)
		if err := c.BodyParser(r); err != nil {
			c.Append("HX-Retarget", "#login-error")
			c.Append("HX-Reswap", "outerHTML")
			c.Append(shared.HeaderHXAcceptable, "true")
			c.Status(fiber.StatusUnauthorized)
			return c.Render("web/views/partials/login/err-invalid", &fiber.Map{}, "")
		}

		p, err := i.Queries.GetPlayerByUsername(context.Background(), r.Username)
		if err != nil {
			c.Append("HX-Retarget", "#login-error")
			c.Append("HX-Reswap", "outerHTML")
			c.Append(shared.HeaderHXAcceptable, "true")
			c.Status(fiber.StatusUnauthorized)
			return c.Render("web/views/partials/login/err-invalid", &fiber.Map{}, "")
		}

		v, err := password.Verify(r.Password, p.PwHash)
		if err != nil {
			c.Append("HX-Retarget", "#login-error")
			c.Append("HX-Reswap", "outerHTML")
			c.Append(shared.HeaderHXAcceptable, "true")
			c.Status(fiber.StatusUnauthorized)
			return c.Render("web/views/partials/login/err-invalid", &fiber.Map{}, "")
		}
		if !v {
			c.Append("HX-Retarget", "#login-error")
			c.Append("HX-Reswap", "outerHTML")
			c.Append(shared.HeaderHXAcceptable, "true")
			c.Status(fiber.StatusUnauthorized)
			return c.Render("web/views/partials/login/err-invalid", &fiber.Map{}, "")
		}

		// TODO: Get tests up around this
		// TODO: Extract this behavior to a function in permissions?
		pid := p.ID
		perms, err := permissions.List(i, pid)
		if err != nil {
			c.Append("HX-Retarget", "#login-error")
			c.Append("HX-Reswap", "outerHTML")
			c.Append(shared.HeaderHXAcceptable, "true")
			c.Status(fiber.StatusUnauthorized)
			return c.Render("web/views/partials/login/err-invalid", &fiber.Map{}, "")
		}
		if !slices.Contains(perms, permissions.Login) {
			c.Append("HX-Retarget", "#login-error")
			c.Append("HX-Reswap", "outerHTML")
			c.Append(shared.HeaderHXAcceptable, "true")
			c.Status(fiber.StatusUnauthorized)
			return c.Render("web/views/partials/login/err-invalid", &fiber.Map{}, "")
		}

		err = username.Cache(i.Redis, pid, p.Username)
		if err != nil {
			c.Append("HX-Retarget", "#login-error")
			c.Append("HX-Reswap", "outerHTML")
			c.Append(shared.HeaderHXAcceptable, "true")
			c.Status(fiber.StatusUnauthorized)
			return c.Render("web/views/partials/login/err-invalid", &fiber.Map{}, "")
		}

		sess, err := i.Sessions.Get(c)
		if err != nil {
			c.Append("HX-Retarget", "#login-error")
			c.Append("HX-Reswap", "outerHTML")
			c.Append(shared.HeaderHXAcceptable, "true")
			c.Status(fiber.StatusUnauthorized)
			return c.Render("web/views/partials/login/err-invalid", &fiber.Map{}, "")
		}

		sess.Set("pid", pid)
		if err = sess.Save(); err != nil {
			c.Append("HX-Retarget", "#login-error")
			c.Append("HX-Reswap", "outerHTML")
			c.Append(shared.HeaderHXAcceptable, "true")
			c.Status(fiber.StatusUnauthorized)
			return c.Render("web/views/partials/login/err-invalid", &fiber.Map{}, "")
		}

		redirect := c.Query("redirect")
		if redirect == "home" {
			c.Append("HX-Redirect", "/")
			c.Status(fiber.StatusOK)
			return nil
		}

		if redirect == "verify" {
			c.Append("HX-Redirect", fmt.Sprintf("/verify?t=%s", c.Query("t")))
			c.Status(fiber.StatusOK)
			return nil
		}

		c.Status(fiber.StatusOK)
		// TODO: Put this event name behind a shared constant
		c.Append("HX-Trigger-After-Swap", "ptrcr:login-success")
		// TODO: Return the success button here
		return nil
	}
}

func LoginPage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: Put this "bind" key into a shared constant
		return c.Render("web/views/login", c.Locals("bind"), "web/views/layouts/standalone")
	}
}
