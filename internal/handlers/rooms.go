package handlers

import (
	"context"
	"database/sql"

	fiber "github.com/gofiber/fiber/v2"

	"petrichormud.com/app/internal/layouts"
	"petrichormud.com/app/internal/partials"
	"petrichormud.com/app/internal/permissions"
	"petrichormud.com/app/internal/queries"
	"petrichormud.com/app/internal/rooms"
	"petrichormud.com/app/internal/routes"
	"petrichormud.com/app/internal/shared"
	"petrichormud.com/app/internal/util"
	"petrichormud.com/app/internal/views"
)

func RoomsPage(i *shared.Interfaces) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: IsLoggedIn helper?
		_, err := util.GetPID(c)
		if err != nil {
			c.Status(fiber.StatusUnauthorized)
			return c.Render(views.Login, views.Bind(c), layouts.Standalone)
		}

		perms, err := util.GetPermissions(c)
		if err != nil {
			c.Status(fiber.StatusForbidden)
			return c.Render(views.Forbidden, views.Bind(c), layouts.Standalone)
		}

		if !perms.HasPermission(permissions.PlayerViewAllRoomsName) {
			c.Status(fiber.StatusForbidden)
			return c.Render(views.Forbidden, views.Bind(c), layouts.Standalone)
		}

		records, err := i.Queries.ListRooms(context.Background())
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.Render(views.InternalServerError, views.Bind(c), layouts.Standalone)
		}

		pageRooms := []fiber.Map{}
		for _, record := range records {
			pageRoom := fiber.Map{
				"Title":      record.Title,
				"Size":       record.Size,
				"SizeString": rooms.SizeToString(record.Size),
				"Path":       routes.RoomPath(record.ID),
			}

			if perms.HasPermission(permissions.PlayerCreateRoomName) {
				pageRoom["EditPath"] = routes.EditRoomPath(record.ID)
			}

			pageRooms = append(pageRooms, pageRoom)
		}

		b := views.Bind(c)
		b["Rooms"] = pageRooms
		b["PageHeader"] = fiber.Map{
			"Title":    "Rooms",
			"SubTitle": "Individual rooms, where their exits and individual properties are assigned",
		}
		return c.Render(views.Rooms, b)
	}
}

func RoomPage(i *shared.Interfaces) fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, err := util.GetPID(c)
		if err != nil {
			c.Status(fiber.StatusUnauthorized)
			return c.Render(views.Login, views.Bind(c), layouts.Standalone)
		}

		perms, err := util.GetPermissions(c)
		if err != nil {
			c.Status(fiber.StatusForbidden)
			return c.Render(views.Forbidden, views.Bind(c), layouts.Standalone)
		}

		if !perms.HasPermission(permissions.PlayerViewAllRoomsName) {
			c.Status(fiber.StatusForbidden)
			return c.Render(views.Forbidden, views.Bind(c), layouts.Standalone)
		}

		rmid, err := util.GetID(c)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.Render(views.Forbidden, views.Bind(c), layouts.Standalone)
		}

		record, err := i.Queries.GetRoom(context.Background(), rmid)
		if err != nil {
			if err == sql.ErrNoRows {
				c.Status(fiber.StatusNotFound)
				return c.Render(views.NotFound, views.Bind(c), layouts.Standalone)
			}
			c.Status(fiber.StatusInternalServerError)
			return c.Render(views.InternalServerError, views.Bind(c), layouts.Standalone)
		}

		b := views.Bind(c)
		b["NavBack"] = fiber.Map{
			"Path":  routes.Rooms,
			"Label": "Back to Rooms",
		}
		b["PageHeader"] = fiber.Map{
			"Title":    rooms.TitleWithID(record.Title, record.ID),
			"SubTitle": "Room",
		}
		b["Name"] = "ImageName"
		b["Title"] = record.Title
		b["Size"] = rooms.SizeToString(record.Size)
		b["Description"] = record.Description
		return c.Render(views.Room, b, layouts.Main)
	}
}

func NewRoom(i *shared.Interfaces) fiber.Handler {
	type input struct {
		Direction string `form:"direction"`
		LinkID    int64  `form:"id"`
		TwoWay    bool   `form:"two-way"`
	}

	return func(c *fiber.Ctx) error {
		_, err := util.GetPID(c)
		if err != nil {
			c.Status(fiber.StatusUnauthorized)
			c.Append(shared.HeaderHXAcceptable, "true")
			return c.Render(partials.NoticeSectionError, partials.BindNoticeSection(partials.BindNoticeSectionParams{
				SectionID:    "new-room-error",
				SectionClass: "pt-2",
				NoticeText: []string{
					"It looks like your session may have expired.",
				},
				RefreshButton: true,
				NoticeIcon:    true,
			}), layouts.None)
		}

		perms, err := util.GetPermissions(c)
		if err != nil {
			c.Status(fiber.StatusForbidden)
			c.Append(shared.HeaderHXAcceptable, "true")
			return c.Render(partials.NoticeSectionError, partials.BindNoticeSection(partials.BindNoticeSectionParams{
				SectionID:    "new-room-error",
				SectionClass: "pt-2",
				NoticeText: []string{
					"Something's gone terribly wrong.",
				},
				RefreshButton: true,
				NoticeIcon:    true,
			}), layouts.None)
		}

		if !perms.HasPermission(permissions.PlayerCreateRoomName) {
			c.Status(fiber.StatusForbidden)
			c.Append(shared.HeaderHXAcceptable, "true")
			return c.Render(partials.NoticeSectionError, partials.BindNoticeSection(partials.BindNoticeSectionParams{
				SectionID:    "new-room-error",
				SectionClass: "pt-2",
				NoticeText: []string{
					"You don't have the permission(s) necessary to create a Room Image.",
				},
				RefreshButton: true,
				NoticeIcon:    true,
			}), layouts.None)
		}

		tx, err := i.Database.Begin()
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			c.Append(shared.HeaderHXAcceptable, "true")
			return c.Render(partials.NoticeSectionError, partials.BindNoticeSection(partials.BindNoticeSectionParams{
				SectionID:    "new-room-image-error",
				SectionClass: "pt-2",
				NoticeText: []string{
					"Something's gone terribly wrong.",
				},
				RefreshButton: true,
				NoticeIcon:    true,
			}), layouts.None)
		}
		defer tx.Rollback()
		qtx := i.Queries.WithTx(tx)

		result, err := qtx.CreateRoom(context.Background(), queries.CreateRoomParams{
			Title:       rooms.DefaultTitle,
			Description: rooms.DefaultDescription,
			Size:        rooms.DefaultSize,
		})
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			c.Append(shared.HeaderHXAcceptable, "true")
			return c.Render(partials.NoticeSectionError, partials.BindNoticeSection(partials.BindNoticeSectionParams{
				SectionID:    "new-room-image-error",
				SectionClass: "pt-2",
				NoticeText: []string{
					"Something's gone terribly wrong.",
				},
				RefreshButton: true,
				NoticeIcon:    true,
			}), layouts.None)
		}

		rid, err := result.LastInsertId()
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			c.Append(shared.HeaderHXAcceptable, "true")
			return c.Render(partials.NoticeSectionError, partials.BindNoticeSection(partials.BindNoticeSectionParams{
				SectionID:    "new-room-image-error",
				SectionClass: "pt-2",
				NoticeText: []string{
					"Something's gone terribly wrong.",
				},
				RefreshButton: true,
				NoticeIcon:    true,
			}), layouts.None)
		}

		in := new(input)
		if err := c.BodyParser(in); err != nil && err != fiber.ErrUnprocessableEntity {
			c.Status(fiber.StatusBadRequest)
			c.Append(shared.HeaderHXAcceptable, "true")
			return c.Render(partials.NoticeSectionError, partials.BindNoticeSection(partials.BindNoticeSectionParams{
				SectionID:    "new-room-image-error",
				SectionClass: "pt-2",
				NoticeText: []string{
					"Something's gone terribly wrong.",
				},
				RefreshButton: true,
				NoticeIcon:    true,
			}), layouts.None)
		}

		if in.LinkID != 0 {
			// TODO: Can run validations on the room to be linked here, ensuring that:
			// 1. The link-to exit isn't already filled
			// 2. There isn't a setpiece that leads to the proposed destination room
			// etc

			record, err := qtx.GetRoom(context.Background(), in.LinkID)
			if err != nil {
				if err == sql.ErrNoRows {
					c.Status(fiber.StatusNotFound)
					return nil
				}
				c.Status(fiber.StatusInternalServerError)
				return nil
			}

			if !rooms.IsDirectionValid(in.Direction) {
				c.Status(fiber.StatusBadRequest)
				return nil
			}

			if err := rooms.Link(rooms.LinkParams{
				Queries:   qtx,
				ID:        in.LinkID,
				To:        rid,
				Direction: in.Direction,
				TwoWay:    in.TwoWay,
			}); err != nil {
				c.Status(fiber.StatusInternalServerError)
				return nil
			}

			if err := tx.Commit(); err != nil {
				c.Status(fiber.StatusInternalServerError)
				c.Append(shared.HeaderHXAcceptable, "true")
				return c.Render(partials.NoticeSectionError, partials.BindNoticeSection(partials.BindNoticeSectionParams{
					SectionID:    "new-room-image-error",
					SectionClass: "pt-2",
					NoticeText: []string{
						"Something's gone terribly wrong.",
					},
					RefreshButton: true,
					NoticeIcon:    true,
				}), layouts.None)
			}

			// TODO: Get a factory for this
			exit := fiber.Map{
				"ID":              rid,
				"RoomID":          record.ID,
				"Exit":            in.Direction,
				"ExitLetter":      rooms.DirectionLetter(in.Direction),
				"ExitTitle":       rooms.DirectionTitle(in.Direction),
				"EditElementID":   rooms.ExitEditElementID(in.Direction),
				"SelectElementID": rooms.ExitSelectElementID(in.Direction),
				"Title":           rooms.DefaultTitle,
				"Description":     rooms.DefaultDescription,
				"ExitPath":        routes.RoomPath(rid),
				"ExitEditPath":    routes.EditRoomPath(rid),
				"RoomsPath":       routes.Rooms,
			}

			c.Status(fiber.StatusCreated)
			return c.Render(partials.EditRoomExitEdit, exit, layouts.None)
		}

		if err := tx.Commit(); err != nil {
			c.Status(fiber.StatusInternalServerError)
			c.Append(shared.HeaderHXAcceptable, "true")
			return c.Render(partials.NoticeSectionError, partials.BindNoticeSection(partials.BindNoticeSectionParams{
				SectionID:    "new-room-image-error",
				SectionClass: "pt-2",
				NoticeText: []string{
					"Something's gone terribly wrong.",
				},
				RefreshButton: true,
				NoticeIcon:    true,
			}), layouts.None)
		}

		c.Status(fiber.StatusCreated)
		c.Append("HX-Redirect", routes.EditRoomPath(rid))
		c.Append("HX-Reswap", "none")
		return nil
	}
}

func EditRoomPage(i *shared.Interfaces) fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, err := util.GetPID(c)
		if err != nil {
			c.Status(fiber.StatusUnauthorized)
			return c.Render(views.Login, views.Bind(c), layouts.Standalone)
		}

		rmid, err := util.GetID(c)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.Render(views.InternalServerError, views.Bind(c), layouts.Standalone)
		}

		perms, err := util.GetPermissions(c)
		if err != nil {
			c.Status(fiber.StatusForbidden)
			return c.Render(views.Forbidden, views.Bind(c), layouts.Standalone)
		}

		if !perms.HasPermission(permissions.PlayerCreateRoomName) {
			c.Status(fiber.StatusForbidden)
			return c.Render(views.Forbidden, views.Bind(c), layouts.Standalone)
		}

		tx, err := i.Database.Begin()
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.Render(views.InternalServerError, views.Bind(c), layouts.Standalone)
		}
		defer tx.Rollback()
		qtx := i.Queries.WithTx(tx)

		record, err := qtx.GetRoom(context.Background(), rmid)
		if err != nil {
			if err == sql.ErrNoRows {
				c.Status(fiber.StatusNotFound)
				return c.Render(views.NotFound, views.Bind(c), layouts.Standalone)
			}
			c.Status(fiber.StatusInternalServerError)
			return c.Render(views.InternalServerError, views.Bind(c), layouts.Standalone)
		}

		exitRooms, err := rooms.LoadExitRooms(qtx, &record)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.Render(views.InternalServerError, views.Bind(c), layouts.Standalone)
		}

		if err := tx.Commit(); err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.Render(views.InternalServerError, views.Bind(c), layouts.Standalone)
		}

		exits := rooms.BuildExits(&record, exitRooms)

		// TODO: Defer this to a load function
		roomGrid := []fiber.Map{
			{
				"ID": "test-room-grid-row-one",
				"Rooms": []fiber.Map{
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
				},
			},
			{
				"ID": "test-room-grid-row-two",
				"Rooms": []fiber.Map{
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
					{"ID": 1},
					{"ID": 5},
				},
			},
			{
				"ID": "test-room-grid-row-three",
				"Rooms": []fiber.Map{
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
					{"ID": 2},
					{"ID": 0},
				},
			},
			{
				"ID": "test-room-grid-row-four",
				"Rooms": []fiber.Map{
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
				},
			},
			{
				"ID": "test-room-grid-row-five",
				"Rooms": []fiber.Map{
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
				},
			},
		}

		b := views.Bind(c)
		// TODO: Get a bind function for this
		b["NavBack"] = fiber.Map{
			"Path":  routes.Rooms,
			"Label": "Back to Rooms",
		}
		// TODO: Get a bind function for this too
		b["PageHeader"] = fiber.Map{
			"Title":    rooms.TitleWithID(record.Title, record.ID),
			"SubTitle": "Update room properties here",
		}
		b["RoomGrid"] = roomGrid
		b["Title"] = record.Title
		b["Description"] = record.Description
		b["Size"] = record.Size
		// TODO: Put this in a helper function
		b["SizeRadioGroup"] = []fiber.Map{
			{
				"ID":       "edit-room-image-size-tiny",
				"Name":     "size",
				"Variable": "size",
				"Value":    "0",
				"Active":   record.Size == 0,
				"Label":    "Tiny",
			},
			{
				"ID":       "edit-room-image-size-small",
				"Name":     "size",
				"Variable": "size",
				"Value":    "1",
				"Active":   record.Size == 1,
				"Label":    "Small",
			},
			{
				"ID":       "edit-room-image-size-medium",
				"Name":     "size",
				"Variable": "size",
				"Value":    "2",
				"Active":   record.Size == 2,
				"Label":    "Medium",
			},
			{
				"ID":       "edit-room-image-size-large",
				"Name":     "size",
				"Variable": "size",
				"Value":    "3",
				"Active":   record.Size == 3,
				"Label":    "Large",
			},
			{
				"ID":       "edit-room-image-size-huge",
				"Name":     "size",
				"Variable": "size",
				"Value":    "4",
				"Active":   record.Size == 4,
				"Label":    "Huge",
			},
		}
		// TODO: I don't think these individual dirs are needed
		b["North"] = record.North
		b["Northeast"] = record.Northeast
		b["East"] = record.East
		b["Southeast"] = record.Southeast
		b["South"] = record.South
		b["Southwest"] = record.Southwest
		b["West"] = record.West
		b["Northwest"] = record.Northwest
		b["Exits"] = exits
		return c.Render(views.EditRoom, b)
	}
}

func RoomGrid() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, err := util.GetPID(c)
		if err != nil {
			return nil
		}

		roomGridOne := []fiber.Map{
			{
				"ID": "test-room-grid-row-one",
				"Rooms": []fiber.Map{
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
				},
			},
			{
				"ID": "test-room-grid-row-two",
				"Rooms": []fiber.Map{
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
					{"ID": 1},
					{"ID": 5},
				},
			},
			{
				"ID": "test-room-grid-row-three",
				"Rooms": []fiber.Map{
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
					{"ID": 2},
					{"ID": 0},
				},
			},
			{
				"ID": "test-room-grid-row-four",
				"Rooms": []fiber.Map{
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
				},
			},
			{
				"ID": "test-room-grid-row-five",
				"Rooms": []fiber.Map{
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
				},
			},
		}

		roomGridTwo := []fiber.Map{
			{
				"ID": "test-room-grid-row-one",
				"Rooms": []fiber.Map{
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
				},
			},
			{
				"ID": "test-room-grid-row-two",
				"Rooms": []fiber.Map{
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
					{"ID": 1},
					{"ID": 6},
				},
			},
			{
				"ID": "test-room-grid-row-three",
				"Rooms": []fiber.Map{
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
					{"ID": 2},
					{"ID": 0},
				},
			},
			{
				"ID": "test-room-grid-row-four",
				"Rooms": []fiber.Map{
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
				},
			},
			{
				"ID": "test-room-grid-row-five",
				"Rooms": []fiber.Map{
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
					{"ID": 0},
				},
			},
		}

		id, err := util.GetID(c)
		if err != nil {
			return nil
		}

		selected, err := util.GetParamID(c, "selected")
		if err != nil {
			return nil
		}

		if id != 0 {
			return nil
		}

		b := views.Bind(c)
		if selected == 2 {
			b["RoomGrid"] = roomGridTwo
		} else {
			b["RoomGrid"] = roomGridOne
		}

		return c.Render(partials.RoomGrid, b, layouts.None)
	}
}
