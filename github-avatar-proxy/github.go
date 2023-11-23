package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type GithubResponse struct {
	Login     string `json:"login"`
	Id        int    `json:"id"`
	AvatarUrl string `json:"avatar_url"`
}

var avatarMemo = NewMemo(func(requestParam string) (interface{}, error) {
	log.Println("requesting avatar for: ", requestParam)
	return getAvatar(requestParam)
})

func getAvatar(requestParam string) ([]byte, error) {
	param := strings.Split(requestParam, "&")
	githubName, size := param[0], param[1]

	avatar, err := getGithubUserAvatar(fmt.Sprintf(`https://github.com/%s.png?size=%s`, githubName, size))
	if err != nil {
		return nil, err
	}

	return avatar, nil
}

func getGithubUserAvatar(url string) ([]byte, error) {
	avatarResp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("avatar url: '%s', request github avatar api err: %v", url, err)
	}
	defer avatarResp.Body.Close()

	data, err := ioutil.ReadAll(avatarResp.Body)
	if err != nil {
		return nil, fmt.Errorf("avatar url: '%s', read github avatar api response body err: %v", url, err)
	}

	return data, nil
}