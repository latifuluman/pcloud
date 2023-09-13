package pcloud

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
)

// GetFilePubLink; https://docs.pcloud.com/methods/public_links/getfilepublink.html
func (c *PCloudClient) GetFilePubLink(path string, fileID int, isEU bool) (string, error) {
	values := url.Values{
		"auth": {*c.Auth},
	}

	switch {
	case path != "":
		values.Add("path", path)
	case fileID >= 0:
		values.Add("fileid", strconv.Itoa(fileID))
	default:
		return "", errors.New("bad params")
	}
	resp, err := c.Client.Get(urlBuilder("getfilepublink", values, isEU))

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result := struct {
		Link   string `json:"link"`
		Result int    `json:"result"`
		Error  string `json:"error"`
	}{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	if result.Result > 0 {
		return "", errors.New(result.Error)
	}
	return result.Link, nil
}

// GetFolderPubLink; https://docs.pcloud.com/methods/public_links/getfolderpublink.html
func (c *PCloudClient) GetFolderPubLink(path string, fileID int, isEU bool) (string, error) {
	values := url.Values{
		"auth": {*c.Auth},
	}

	switch {
	case path != "":
		values.Add("path", path)
	case fileID >= 0:
		values.Add("fileid", strconv.Itoa(fileID))
	default:
		return "", errors.New("bad params")
	}
	resp, err := c.Client.Get(urlBuilder("getfolderpublink", values, isEU))

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result := struct {
		Link   string `json:"link"`
		Result int    `json:"result"`
		Error  string `json:"error"`
	}{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	if result.Result > 0 {
		return "", errors.New(result.Error)
	}
	return result.Link, nil
}
