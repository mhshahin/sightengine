package sightengine

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/mhshahin/sightengine/constants"
	"github.com/mhshahin/sightengine/utility"
)

// Client ...
type Client struct {
	client    *http.Client
	APIUser   string
	APISecret string
	Workflow  string
}

// NewClient returns a new instance of Client
func NewClient(settings *Client) *Client {
	var httpClient = http.DefaultClient

	if settings.client != nil {
		httpClient = settings.client
	}

	client := &Client{
		client:    httpClient,
		APIUser:   settings.APIUser,
		APISecret: settings.APISecret,
		Workflow:  settings.Workflow,
	}

	return client
}

// Get ...
func (c *Client) Get(path string, params map[string]string) ([]byte, error) {
	body, err := c.doRequest(http.MethodGet, path, params, nil, "")
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Post ...
func (c *Client) Post(reqPath, filePath string) ([]byte, error) {
	var b bytes.Buffer

	w := multipart.NewWriter(&b)

	if len(c.Workflow) != 0 {
		fw, err := w.CreateFormField("workflow")
		if err != nil {
			return nil, err
		}

		if len(c.Workflow) != 0 {
			_, err = io.Copy(fw, strings.NewReader(c.Workflow))
			if err != nil {
				return nil, err
			}
		}
	}

	fw, err := w.CreateFormField("api_user")
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fw, strings.NewReader(c.APIUser))
	if err != nil {
		return nil, err
	}

	fw, err = w.CreateFormField("api_secret")
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fw, strings.NewReader(c.APISecret))
	if err != nil {
		return nil, err
	}

	var mediaName []string

	goos := runtime.GOOS
	switch goos {
	case "darwin":
		mediaName = strings.Split(filePath, "/")
	case "linux":
		mediaName = strings.Split(filePath, "/")
	default:
		return nil, fmt.Errorf("unexpected operating system %s", goos)
	}

	fw, err = w.CreateFormFile("media", mediaName[len(mediaName)-1])
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fw, file)
	if err != nil {
		log.Println(err)
	}
	w.Close()

	body, err := c.doRequest(http.MethodPost, reqPath, nil, bytes.NewReader(b.Bytes()), w.FormDataContentType())
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) doRequest(httpMethod, path string, params map[string]string, reqBody io.Reader, contentType string) ([]byte, error) {
	authParams := map[string]string{
		"api_user":   c.APIUser,
		"api_secret": c.APISecret,
	}

	req, err := http.NewRequest(httpMethod, createURL(path, authParams, params), reqBody)
	if err != nil {
		return nil, err
	}

	if httpMethod == http.MethodPost {
		req.Header.Set("Content-Type", contentType)
	}

	rsp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode != http.StatusOK {
		err = fmt.Errorf("unexpected status code %d, returned: %s", rsp.StatusCode, string(rspBody))
	}

	return rspBody, nil
}

func createURL(path string, params ...map[string]string) string {
	var url string

	qp := utility.MapToURLValues(params...)
	url = fmt.Sprintf("%s/%s.json?%s", constants.Endpoint, path, qp.Encode())

	return url
}
