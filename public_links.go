package pcloud

import (
	"errors"
	"net/url"
	"strconv"
)

// CreateFolder; https://docs.pcloud.com/methods/folder/createfolder.html
func (c *PCloudClient) GetFilePubLink(path string, fileID int) error {
	values := url.Values{
		"auth": {*c.Auth},
	}

	switch {
	case path != "":
		values.Add("path", path)
	case fileID >= 0:
		values.Add("fileid", strconv.Itoa(fileID))
	default:
		return errors.New("bad params")
	}

	return checkResult(c.Client.Get(urlBuilder("getfilepublink", values)))
}
