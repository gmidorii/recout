package pixela

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

const (
	pixelaURL         = "https://pixe.la/v1/users"
	pixelaHeaderToken = "X-USER-TOKEN"
)

type Client interface {
	Increment(userID, token, graph string) error
	CreateUser(user User) error
}

type client struct {
}

func NewClient() Client {
	return &client{}
}

func (c *client) Increment(userID, token, graph string) error {
	url := fmt.Sprintf("%v/%v/graphs/%v/increment", pixelaURL, userID, graph)
	log.Println(url)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return errors.Wrap(err, "failed new http request")
	}
	req.Header.Add(pixelaHeaderToken, token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

type User struct {
	Token               string   `validate:"required,get=8,lte=24"`
	UserName            string   `validate:"required,get=1,lte=32"`
	AgreeTermsOfService Question `validate:"required"`
	NotMinor            Question `validate:"required"`
}

type UserResponse struct {
	Message   string `json:"message"`
	IsSuccess bool   `json:"isSuccess"`
}

type Question string

const (
	Yes Question = "yes"
	No  Question = "no"
)

func (c *client) CreateUser(user User) error {
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return err
	}
	d, err := json.Marshal(&user)
	if err != nil {
		return err
	}

	req := http.NewRequest("POST", pixelaURL, bytes.NewReader(d))
	if err != nil {
		return errors.Wrap(err, "failed new http request")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code = %v", resp.StatusCode)
	}

	var decoder json.Decoder
	var resUser UserResponse
	if err := decoder.Decode(&resUser); err != nil {
		return err
	}
	if !resUser.IsSuccess {
		return errors.New("failed create user")
	}

	return nil
}
