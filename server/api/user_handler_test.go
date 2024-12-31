package api

import (
	"bytes"
	"encoding/json"
	"hotel-reservation/types"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
)

func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.User)
	app.Post("/", userHandler.HandlePostUser)

	params := types.CreateUserParams{
		Email:     "some@test.con",
		FirstName: "James",
		LastName:  "Foo",
		Password:  "sdsdsdsdsdsgsv",
	}

	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	const timeout = time.Second * 2
	resp, err := app.Test(req, int(timeout))
	if err != nil {
		t.Error(err)
	}
	var user types.User
	json.NewDecoder(resp.Body).Decode(&user)

	if len(user.ID) == 0 {
		t.Error("expecting a user id to be set")
	}
	if len(user.EncryptedPassword) > 0 {
		t.Error("expecting the EncryptedPassword not to be included in the json response")
	}
	if user.FirstName != params.FirstName {
		t.Errorf("expected firstname %s got %s", params.FirstName, user.FirstName)
	}
	if user.LastName != params.LastName {
		t.Errorf("expected lastname %s got %s", params.FirstName, user.FirstName)
	}
	if user.Email != params.Email {
		t.Errorf("expected email %s got %s", params.FirstName, user.FirstName)
	}
}
