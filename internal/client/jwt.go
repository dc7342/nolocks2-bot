package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"nolocks-bot/internal/entity"
)

const (
	refreshURL = "refresh/"
	verifyURL  = "verify/"
	tokenURL   = "token/"
)

type jsonData map[string]string

type TokenJSON struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type JWTClient struct {
	accessToken  string
	refreshToken string
}

func NewJWTClient(conf entity.NoLocksConfig) (*JWTClient, error) {
	conf.EndpointURL += tokenURL
	// Check provided endPoint for validness.
	_, err := url.ParseRequestURI(conf.EndpointURL)
	if err != nil {
		return &JWTClient{}, err
	}
	return &JWTClient{}, nil
}

func (j *JWTClient) Get(conf entity.NoLocksConfig) (string, error) {
	if j.accessToken == "" {
		return j.getToken(conf)
	}

	valid, err := j.verify(conf)
	if err != nil {
		return "", err
	}

	if !valid {
		return j.refresh(conf)
	}

	return j.accessToken, nil
}

func (j *JWTClient) getToken(conf entity.NoLocksConfig) (string, error) {
	// Send auth query.
	resp, err := j.sendJson(conf.EndpointURL+tokenURL, jsonData{"username": conf.User, "password": conf.Pass})
	if err != nil {
		return "", err
	}

	// Check if request completed successfully.
	if resp.StatusCode != http.StatusOK {
		return "", errorConnection
	}

	// Parse response.
	token := j.parse(resp)
	j.accessToken = token.Access
	j.refreshToken = token.Refresh

	return j.accessToken, err
}

func (j *JWTClient) refresh(conf entity.NoLocksConfig) (string, error) {
	if j.refreshToken == "" {
		return "", errorUnobtainedAccess
	}

	// Refresh request.
	resp, err := j.sendJson(conf.EndpointURL+tokenURL+refreshURL, jsonData{"refresh": j.refreshToken})
	if err != nil {
		return "", err
	}

	// Check if request completed successfully.
	if resp.StatusCode != http.StatusOK {
		return "", errorConnection
	}
	token := j.parse(resp)
	j.accessToken = token.Access

	return j.accessToken, err
}

func (j *JWTClient) verify(conf entity.NoLocksConfig) (bool, error) {
	if j.accessToken == "" {
		return false, errorUnobtainedAccess
	}

	// Refresh request.
	resp, err := j.sendJson(conf.EndpointURL+tokenURL+verifyURL, jsonData{"refresh": j.refreshToken})
	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusOK {
		return false, nil
	}

	return true, nil
}

// Parse response to struct.
func (j *JWTClient) parse(resp *http.Response) *TokenJSON {
	defer resp.Body.Close()

	token := new(TokenJSON)
	json.NewDecoder(resp.Body).Decode(token)

	return token
}

func (j *JWTClient) sendJson(url string, data jsonData) (*http.Response, error) {
	jsonRaw, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return http.Post(url, contentTypeJSON, bytes.NewBuffer(jsonRaw))
}
