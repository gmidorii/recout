package pixela

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

const (
	pixelaURL         = "https://pixe.la/v1/users"
	pixelaHeaderToken = "X-USER-TOKEN"
)

type Client interface {
	Increment(userID, token, graph string) error
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
