package github

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// GetAuthorizeUrl 获取用户身份认证url
// https://docs.github.com/cn/developers/apps/building-oauth-apps/authorizing-oauth-apps
func GetAuthorizeUrl(clientId, redirectUri, scope, state string) string {
	u, _ := url.Parse("https://github.com/login/oauth/authorize")
	query := u.Query()
	query.Add("client_id", clientId)
	query.Add("redirect_uri", redirectUri)
	query.Add("scope", scope)
	query.Add("state", state)
	// 解析RawQuery并返回"值，您得到的只是URL查询值的副本，而不是"实时引用"，
	// 因此修改该副本不会对原始查询产生任何影响。
	// 为了修改原始查询，您必须分配给原始RawQuery
	u.RawQuery = query.Encode()
	return u.String()
}

// GetAccessToken 获取AccessToken
func GetAccessToken(clientId, clientSecret, code, redirectUri string) (*GetAccessTokenResp, error) {
	data, _ := json.Marshal(map[string]string{
		"client_id":     clientId,
		"client_secret": clientSecret,
		"code":          code,
		"redirect_uri":  redirectUri,
	})

	method := "POST"

	client := &http.Client{}
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

/*
{
    "access_token": "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",
    "token_type": "bearer",
    "scope": ""
}
*/
