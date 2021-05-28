package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"math/rand"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/machinebox/graphql"
	"github.com/oklog/ulid"
)

/*
Query:
mutation RegisterUser($email: String!, $uid: String!) {
  registerUser(uid: $uid, email: $email) {
    user_id
  }
}

Variables:
{
	"email": "email@example.com",
	"uid": "<firebase_authentication_user_id>"
}
*/

type RegisterUserParam struct {
	Action struct {
		Name string `json:"name"`
	} `json:"action"`
	Input struct {
		Email string `json:"email"`
		Uid   string `json:"uid"`
	} `json:"input"`
}

func RegisterUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(RegisterUserParam)
		if err := c.Bind(param); err != nil {
			return err
		}

		invitations, err := getTeamInvitations(param.Input.Email)
		if err != nil {
			return err
		}

		t := time.Now()
		var validInvitations []Inviation
		for _, invitation := range invitations.TeamInvitations {
			// if expires?
			if t.Before(invitation.ExpiresAt) {
				validInvitations = append(validInvitations, invitation)
			}
		}
		fmt.Println("validInvitations Count:", len(validInvitations))
		if len(validInvitations) == 0 {
			return errors.New("Valid invitation not found.")
		}

		userId := generateULID()
		user, err := createUser(userId, param.Input.Email, param.Input.Uid)
		if err != nil {
			return err
		}
		fmt.Println(user)

		var inputObjects []InputTeamMember
		for _, invitation := range validInvitations {
			inputObjects = append(inputObjects, InputTeamMember{
				TeamId: invitation.TeamId,
				Name:   invitation.InviteeName,
				UserId: userId,
			})
		}
		_, err = createTeamMembers(inputObjects)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"user_id": userId.String(),
		})
	}
}

func generateULID() ulid.ULID {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	return id
}

type Inviation struct {
	ExpiresAt    time.Time `json:"expires_at"`
	HostUserId   string    `json:"host_user_id"`
	Id           int       `json:"id"`
	InviteeEmail string    `json:"invitee_email"`
	InviteeName  string    `json:"invitee_name"`
	TeamId       int       `json:"team_id"`
}
type RespInvitations struct {
	TeamInvitations []Inviation `json:"team_invitations"`
}

func getTeamInvitations(email string) (*RespInvitations, error) {
	client := graphql.NewClient("http://host.docker.internal:8080/v1/graphql")
	req := graphql.NewRequest(`
		query InvitationsByEmail($email: String!) {
			team_invitations(where: {invitee_email: {_eq: $email}}) {
				id
				invitee_email
				invitee_name	
				team_id
				host_user_id
				expires_at
			}
		}
	`)

	req.Var("email", email)

	// TODO: make it secret env
	// Request as Admin Role
	req.Header.Set("x-hasura-admin-secret", "sparring-hasura-secret")

	ctx := context.Background()

	var response RespInvitations
	if err := client.Run(ctx, req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

type RespCreateUser struct {
	InsertUsersOne struct {
		Id string `json:"id"`
	} `json:"insert_users_one"`
}

func createUser(id ulid.ULID, email string, authId string) (*RespCreateUser, error) {
	client := graphql.NewClient("http://host.docker.internal:8080/v1/graphql")
	req := graphql.NewRequest(`
		mutation CreateUser ($id: String, $email: String!, $auth_id: String!) {
			insert_users_one(object: {id: $id, email: $email, auth_id: $auth_id}) {
				id
			}
		}
	`)

	req.Var("id", id)
	req.Var("email", email)
	req.Var("auth_id", authId)

	// TODO: make it secret env
	// Request as Admin Role
	req.Header.Set("x-hasura-admin-secret", "sparring-hasura-secret")

	ctx := context.Background()

	var response RespCreateUser
	if err := client.Run(ctx, req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

type InputTeamMember struct {
	TeamId int       `json:"team_id"`
	Name   string    `json:"name"`
	UserId ulid.ULID `json:"user_id"`
}

type RespCreateTeamMembers struct {
	InsertTeamMembers struct {
		AffectedRows int `json:"affected_rows"`
	} `json:"insert_team_members"`
}

func createTeamMembers(objects []InputTeamMember) (*RespCreateTeamMembers, error) {
	client := graphql.NewClient("http://host.docker.internal:8080/v1/graphql")
	req := graphql.NewRequest(`
		mutation CreateTeamMembers($objects: [team_members_insert_input!]!) {
			insert_team_members(objects: $objects) {
			affected_rows
			}
		}
	`)

	req.Var("objects", objects)

	// TODO: make it secret env
	// Request as Admin Role
	req.Header.Set("x-hasura-admin-secret", "sparring-hasura-secret")

	ctx := context.Background()

	var response RespCreateTeamMembers
	if err := client.Run(ctx, req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
