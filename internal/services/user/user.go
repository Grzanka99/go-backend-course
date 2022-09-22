package user

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/grzanka99/backend-go/internal/utils"
)

func UserLogin(rw http.ResponseWriter, req *http.Request) {
	var res UserInformation = UserInformation{
		Email:    "aaaa@email.com",
		Username: "admin",
	}

	utils.Success(res, rw)
}

func UserRegister(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		utils.InternalServerError(Message{
			Message: "Internal server error",
		}, rw)
		return
	}

	var payload UserRegistrationPayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		utils.BadRequest(Message{
			Message: "Wrong payload",
		}, rw)
		return
	}

	if payload.Username == "" || payload.Email == "" || payload.Password == "" {
		utils.BadRequest(Message{
			Message: "Wrong payload",
		}, rw)
		return
	}

	utils.Success(UserInformation{
		Username: payload.Username,
		Email:    payload.Email,
	}, rw)

	return
}
