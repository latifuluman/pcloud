package pcloud

import (
	"encoding/json"
	"errors"
	"net/url"
)

// Login client; https://docs.pcloud.com/methods/intro/authentication.html
func (c *PCloudClient) Login(username string, password string, isEU bool) error {
	values := url.Values{
		"getauth":  {"1"},
		"username": {username},
		"password": {password},
	}

	buf, err := convertToBuffer(c.Client.Get(urlBuilder("userinfo", values, isEU)))
	if err != nil {
		return err
	}

	result := struct {
		Auth   string `json:"auth"`
		Result int    `json:"result"`
		Error  string `json:"error"`
	}{}

	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		return err
	}

	if result.Result != 0 {
		return errors.New(result.Error)
	}

	c.Auth = &result.Auth
	return nil
}

// Logout client; https://docs.pcloud.com/methods/auth/logout.html
func (c *PCloudClient) Logout(isEU bool) error {
	values := url.Values{
		"auth": {*c.Auth},
	}

	if err := checkResult(c.Client.Get(urlBuilder("logout", values, isEU))); err != nil {
		return err
	}

	c.Auth = nil
	return nil
}
