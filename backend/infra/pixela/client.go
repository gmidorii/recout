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

var (
	validate = validator.New()
)

type Client interface {
	Increment(userID, token, graph string) error
	CreateUser(user User) error
	CreateGraph(id, graph, name, token string) error
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

func (c *client) CreateUser(user User) error {
	if err := validate.Struct(&user); err != nil {
		return err
	}
	d, err := json.Marshal(user)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", pixelaURL, bytes.NewReader(d))
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

	decoder := json.NewDecoder(resp.Body)
	var resUser PostResponse
	if err := decoder.Decode(&resUser); err != nil {
		return err
	}
	if !resUser.IsSuccess {
		return errors.New("failed create user")
	}

	return nil
}

func (c *client) CreateGraph(id, graph, name, token string) error {
	g := Graph{
		ID:   id,
		Name: name,
		// fixed value
		Unit:           "commit",
		Type:           "int",
		Color:          "shibafu",
		Timezone:       "Asia/Tokyo",
		SelfSufficient: "increment",
	}
	if err := validate.Struct(&g); err != nil {
		return err
	}
	d, err := json.Marshal(g)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%v/%v/graphs", pixelaURL, name)
	log.Println(url)
	req, err := http.NewRequest("POST", url, bytes.NewReader(d))
	if err != nil {
		return errors.Wrap(err, "failed new http request")
	}
	req.Header.Add(pixelaHeaderToken, token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var postRes PostResponse
	if err := decoder.Decode(&postRes); err != nil {
		return err
	}
	if !postRes.IsSuccess {
		return errors.New("failed create graph")
	}
	return nil
}
