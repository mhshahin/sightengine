package sightengine

import (
	"encoding/json"

	"github.com/mhshahin/sightengine/constants"
)

// NudityByURL helps you determine whether and image passed by
// its publicly available URL contains any king of nudity or not.
func (c *Client) NudityByURL(imageURL string) (*Response, error) {
	params := map[string]string{
		"url": imageURL,
	}

	rsp, err := c.Get(constants.Nudity, params)
	if err != nil {
		return nil, err
	}

	newResponse := new(Response)
	if err = json.Unmarshal(rsp, newResponse); err != nil {
		return nil, err
	}

	return newResponse, nil
}

// NudityByFile ....
func (c *Client) NudityByFile(filePath string) (*Response, error) {
	rsp, err := c.Post(constants.Nudity, filePath)
	if err != nil {
		return nil, err
	}

	newResponse := new(Response)
	if err = json.Unmarshal(rsp, newResponse); err != nil {
		return nil, err
	}

	return newResponse, nil
}
