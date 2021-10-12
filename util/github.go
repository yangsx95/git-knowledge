package util

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"time"
)

// GetGithubAccessToken 获取AccessToken
func GetGithubAccessToken(clientId, clientSecret, code, redirectUri string) (*GetAccessTokenResp, error) {
	data, _ := json.Marshal(map[string]string{
		"client_id":     clientId,
		"client_secret": clientSecret,
		"code":          code,
		"redirect_uri":  redirectUri,
	})

	method := "POST"

	client := &http.Client{Timeout: 20 * time.Second}
	req, err := http.NewRequest(method, "https://github.com/login/oauth/access_token", bytes.NewReader(data))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	rspObj := &GetAccessTokenResp{}
	err = json.Unmarshal(respBody, rspObj)
	if err != nil {
		return nil, err
	}
	return rspObj, nil
}

type GetAccessTokenResp struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorUri         string `json:"error_uri"`
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	Scope            string `json:"scope"`
}

func GetGithubClient(accessToken string) (*github.Client, context.Context) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client, ctx
}
