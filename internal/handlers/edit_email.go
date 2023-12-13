package handlers

import (
	"context"
	"database/sql"
	"net/mail"
	"strconv"

	fiber "github.com/gofiber/fiber/v2"

	"petrichormud.com/app/internal/queries"
	"petrichormud.com/app/internal/shared"
)

func EditEmail(i *shared.Interfaces) fiber.Handler {
	type request struct {
		Email string `form:"email"`
	}

	return func(c *fiber.Ctx) error {
		r := new(request)
		if err := c.BodyParser(r); err != nil {
			c.Append("HX-Retarget", "profile-email-error")
			c.Append("HX-Reswap", "outerHTML")
			c.Append(shared.HeaderHXAcceptable, "true")
			c.Status(fiber.StatusBadRequest)
			return c.Render("views/partials/profile/email/edit/err-internal", &fiber.Map{}, "")
		}

		ne, err := mail.ParseAddress(r.Email)
		if err != nil {
			c.Append("HX-Retarget", "#add-email-error")
			c.Append("HX-Reswap", "innerHTML")
			c.Append(shared.HeaderHXAcceptable, "true")
			c.Status(fiber.StatusBadRequest)
			return c.Render("views/partials/profile/email/edit/err-internal", &fiber.Map{}, "")
		}

		pid := c.Locals("pid")
		if pid == nil {
			c.Append("HX-Retarget", "profile-email-error")
			c.Append("HX-Reswap", "outerHTML")
			c.Append(shared.HeaderHXAcceptable, "true")
			c.Status(fiber.StatusUnauthorized)
			return c.Render("views/partials/profile/email/edit/err-internal", &fiber.Map{}, "")
		}

		eid := c.Params("id")
		if len(eid) == 0 {
			c.Append("HX-Retarget", "profile-email-error")
			c.Append("HX-Reswap", "outerHTML")
			c.Append(shared.HeaderHXAcceptable, "true")
			c.Status(fiber.StatusBadRequest)
			return c.Render("views/partials/profile/email/edit/err-internal", &fiber.Map{}, "")
		}

		id, err := strconv.ParseInt(eid, 10, 64)
		if err != nil {
			c.Append("HX-Retarget", "profile-email-error")
			c.Append("HX-Reswap", "outerHTML")
			c.Append(shared.HeaderHXAcceptable, "true")
			c.Status(fiber.StatusBadRequest)
			return c.Render("views/partials/profile/email/edit/err-internal", &fiber.Map{}, "")
		}

		tx, err := i.Database.Begin()
		if err != nil {
			c.Append("HX-Retarget", "profile-email-error")
			c.Append("HX-Reswap", "outerHTML")
			c.Append(shared.HeaderHXAcceptable, "true")
			c.Status(fiber.StatusInternalServerError)
			return c.Render("views/partials/profile/email/edit/err-internal", &fiber.Map{}, "")
		}
		defer tx.Rollback()

		qtx := i.Queries.WithTx(tx)

		e, err := qtx.GetEmail(context.Background(), id)
		if err != nil {
			if err == sql.ErrNoRows {
				c.Append("HX-Retarget", "profile-email-error")
				c.Append("HX-Reswap", "outerHTML")
				c.Append(shared.HeaderHXAcceptable, "true")
				c.Status(fiber.StatusNotFound)
				return c.Render("views/partials/profile/email/edit/err-internal", &fiber.Map{}, "")
			}
			c.Append("HX-Retarget", "profile-email-error")
			c.Append("HX-Reswap", "outerHTML")
			c.Append(shared.HeaderHXAcceptable, "true")
			c.Status(fiber.StatusInternalServerError)
			return c.Render("views/partials/profile/email/edit/err-internal", &fiber.Map{}, "")
		}

		if !e.Verified {
			c.Append("HX-Retarget", "profile-email-error")
			c.Append("HX-Reswap", "outerHTML")
			c.Append(shared.HeaderHXAcceptable, "true")
			c.Status(fiber.StatusForbidden)
			return c.Render("views/partials/profile/email/edit/err-internal", &fiber.Map{}, "")
		}

		if e.PID != pid.(int64) {
			c.Append("HX-Retarget", "profile-email-error")
			c.Append("HX-Reswap", "outerHTML")
			c.Append(shared.HeaderHXAcceptable, "true")
			c.Status(fiber.StatusForbidden)
			return c.Render("views/partials/profile/email/edit/err-internal", &fiber.Map{}, "")
		}

		_, err = qtx.DeleteEmail(context.Background(), id)
		if err != nil {
			if err == sql.ErrNoRows {
				c.Append("HX-Retarget", "profile-email-error")
				c.Append("HX-Reswap", "outerHTML")
				c.Append(shared.HeaderHXAcceptable, "true")
				c.Status(fiber.StatusNotFound)
				return c.Render("views/partials/profile/email/edit/err-internal", &fiber.Map{}, "")
			}
			c.Append("HX-Retarget", "profile-email-error")
			c.Append("HX-Reswap", "outerHTML")
			c.Append(shared.HeaderHXAcceptable, "true")
			c.Status(fiber.StatusInternalServerError)
			return c.Render("views/partials/profile/email/edit/err-internal", &fiber.Map{}, "")
		}

		result, err := qtx.CreateEmail(context.Background(), queries.CreateEmailParams{
			Address: ne.Address,
			PID:     pid.(int64),
		})
		if err != nil {
			c.Append("HX-Retarget", "profile-email-error")
			c.Append("HX-Reswap", "outerHTML")
			c.Append(shared.HeaderHXAcceptable, "true")
			c.Status(fiber.StatusInternalServerError)
			return c.Render("views/partials/profile/email/edit/err-internal", &fiber.Map{}, "")
		}

		id, err = result.LastInsertId()
		if err != nil {
			c.Append("HX-Retarget", "profile-email-error")
			c.Append("HX-Reswap", "outerHTML")
			c.Append(shared.HeaderHXAcceptable, "true")
			c.Status(fiber.StatusInternalServerError)
			return c.Render("views/partials/profile/email/edit/err-internal", &fiber.Map{}, "")
		}

		err = tx.Commit()
		if err != nil {
			c.Append("HX-Retarget", "profile-email-error")
			c.Append("HX-Reswap", "outerHTML")
			c.Append(shared.HeaderHXAcceptable, "true")
			c.Status(fiber.StatusInternalServerError)
			return c.Render("views/partials/profile/email/edit/err-internal", &fiber.Map{}, "")
		}

		return c.Render("views/partials/profile/email/unverified", &fiber.Map{
			"ID":       id,
			"Address":  r.Email,
			"Verified": false,
		}, "")
	}
}
