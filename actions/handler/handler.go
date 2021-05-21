package handler

import (
	"fmt"
	"net/http"

	"math/rand"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid"
)

type RegisterUserParam struct {
	Action struct {
		Name string `json:"name"`
	} `json:"action"`
	Input struct {
		Email string `json:"email"`
		Uid   string `json:"uid"`
	} `json:"input"`
}

func generateULID() ulid.ULID {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	return id
}

func RegisterUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(RegisterUserParam)
		if err := c.Bind(param); err != nil {
			return err
		}
		fmt.Println(param.Input)

		/*
			invitations := getTeamInvitations(email)
			if (invitations.empty?) {
				return error
			}
			validInvitations := invitations.map(invitation => {!invitation.expired?})
			if (validinvitations.empty?) {
				return error
			}
		*/

		ulid := generateULID()
		/*
			user:= createUser(ulid, param.Input.Uid, param.Input.Email)
			validinvitations.map(invitation => {
				createTeamMembers(invitation.team_id, user.id, invitation.invitee_name)
			})
		*/

		return c.JSON(http.StatusOK, map[string]interface{}{
			"user_id": ulid.String(),
		})
	}
}
