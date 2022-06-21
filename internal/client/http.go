package client

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"nolocks-bot/internal/entity"
	"path"
	"strconv"
	"strings"
)

const (
	contentTypeJSON      = "application/json"
	contentTypeForm      = "application/x-www-form-urlencoded"
	contentTypeMultipart = "multipart/form-data"

	locationURL = "location/"
)

type HTTPClient struct {
	conf entity.NoLocksConfig
	jwt  JWT
}

func NewHTTPClient(conf entity.NoLocksConfig) (*HTTPClient, error) {
	jwt, err := NewJWTClient(conf)
	if err != nil {
		return nil, err
	}

	return &HTTPClient{conf: conf, jwt: jwt}, nil
}

func (h *HTTPClient) Send(loc *entity.Location) error {
	// Body buffer.
	bBuf := &bytes.Buffer{}
	// Body write.
	bWrite := multipart.NewWriter(bBuf)

	// Craft multipart form.
	// Ignore err handling it should be fine here.
	bWrite.WriteField("longitude", strconv.FormatFloat(loc.Lon, 'f', 6, 64))
	bWrite.WriteField("latitude", strconv.FormatFloat(loc.Lat, 'f', 6, 64))
	bWrite.WriteField("comment", loc.Comment)

	// If user sent an image, it is being downloaded from Telegram
	// servers and copies to the file writer.
	if loc.ImageUrl != "" {
		// Download an image and get io.Reader.
		img, err := h.downloadFile(loc.ImageUrl)
		if err != nil {
			return err
		}

		// Create io.Writer.
		fileWriter, _ := bWrite.CreateFormFile(path.Ext(loc.ImageUrl), path.Base(loc.ImageUrl))
		// Copy io.Reader to io.Writer.
		_, err = io.Copy(fileWriter, img)
		// Don't forget to close img io.ReadCloser!
		img.Close()
		if err != nil {
			return err
		}
	}

	// Get content type we've got.
	contentType := bWrite.FormDataContentType()
	if err := bWrite.Close(); err != nil {
		return err
	}

	// Get JWT token.
	token, err := h.jwt.Get(h.conf)
	if err != nil {
		return err
	}

	// Craft a request.
	client := http.Client{}
	req, err := http.NewRequest("POST", h.conf.EndpointURL+locationURL, strings.NewReader(bBuf.String()))
	req.Header.Add("Authorization", fmt.Sprintf("%s %s", "Bearer", token))
	req.Header.Add("Content-Type", contentType)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf(resp.Status)
	}

	return nil
}

func (h *HTTPClient) Get(loc entity.Location) ([]entity.Location, error) {
	// Isn't implemented yet
	return nil, nil
}

func (h *HTTPClient) GetAll() ([]entity.Location, error) {
	// Isn't implemented yet
	return nil, nil
}
