package github

import (
	"fmt"
	"testing"
)

func TestGetAuthorizeUrl(t *testing.T) {
	url := GetAuthorizeUrl("07c67ec26f545c24a9d7", "", "", "123")
	fmt.Println(url)
}

func TestGetAccessToken(t *testing.T) {
	token, err := GetAccessToken("07c67ec26f545c24a9d7", "7a97745370c590b40ef72f56f24dfcb65a3f6bc8", "e7bfc43b47eb8c981eb5", "")
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}
