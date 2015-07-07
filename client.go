package pipl

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	libraryVersion = "0.1"
	defaultBaseURL = "http://api.pipl.com/search/"
	v4Version      = "v4"
	defaultTimeout = 10 * time.Second
)

type V4Client struct {
	ApiKey    string
	Version   string
	BaseURL   string
	UserAgent string
	Timeout   time.Duration
}

func NewV4Client(key string) *V4Client {
	c := &V4Client{}
	c.ApiKey = key
	baseURL := defaultBaseURL + v4Version + "/"
	c.UserAgent = "go-pipl-client " + libraryVersion
	c.BaseURL = baseURL
	return c
}

func (c *V4Client) request(timeout time.Duration, params map[string]string) (PiplResponse, error) {
	req, _ := url.Parse(c.BaseURL)
	p := url.Values{}

	for k, v := range params {
		p.Add(k, v)
	}
	p.Add("key", c.ApiKey)

	// maybe move http client up into struct
	client := http.Client{
		Timeout: timeout,
	}
	response, err := client.PostForm(req.String(), p)
	person := PiplResponse{}

	if err != nil {
		return person, err
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return person, err
		}
		if err = json.Unmarshal(contents, &person); err != nil {
			return person, err
		}
		if len(person.Error) > 0 {
			return person, errors.New(person.Error)
		}
		return person, nil
	}
}

func (c *V4Client) SearchPerson(p Person, timeout time.Duration, params map[string]string) (PiplResponse, error) {
	person, err := json.Marshal(p)
	if err != nil {
		return PiplResponse{}, err
	}
	params["person"] = string(person)
	return c.request(timeout, params)
}

func (c *V4Client) SearchPointer(search_pointer string, timeout time.Duration, params map[string]string) (PiplResponse, error) {
	params["search_pointer"] = search_pointer
	return c.request(timeout, params)
}

func (c *V4Client) SearchParams(search_pointer string, timeout time.Duration, params map[string]string) (PiplResponse, error) {
	return c.request(timeout, params)
}
