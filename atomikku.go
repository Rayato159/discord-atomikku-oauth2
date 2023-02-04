package atomikku

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Rayato159/discord-atomikku-oauth2/models"
	kawaiihttp "github.com/Rayato159/kawaii-sender"
)

type atomikkuApp struct {
	ClientId     string
	ClientSecret string
	RedirectUrl  string
	Scopes       string
}

const (
	baseTokenUrl string = "https://discord.com/api/oauth2/token"
	version      string = "v10"
)

func scopesConcator(scopes ...string) (*string, error) {
	if len(scopes) == 0 {
		return nil, fmt.Errorf("error, scopes are empty")
	}
	var concated string
	for i := range scopes {
		if i != len(scopes)-1 {
			concated += fmt.Sprintf("%s ", scopes[i])
		} else {
			concated += scopes[i]
		}
	}
	return &concated, nil
}

func SetAtomikkuConfig(clientId, clientSecret, redirect string, scopes ...string) (*atomikkuApp, error) {
	switch {
	case clientId == "":
		return nil, fmt.Errorf("error, client_id is missing")
	case redirect == "":
		return nil, fmt.Errorf("error, redirect_url is missing")
	}
	scopesStr, err := scopesConcator(scopes...)
	if err != nil {
		return nil, err
	}
	return &atomikkuApp{
		ClientId:    clientId,
		RedirectUrl: redirect,
		Scopes:      *scopesStr,
	}, nil
}

func (a *atomikkuApp) UrlGenerator(state string) string {
	return fmt.Sprintf(
		"https://discord.com/api/oauth2/authorize?client_id=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s",
		a.ClientId,
		a.RedirectUrl,
		a.Scopes,
		state,
	)
}

func (a *atomikkuApp) GetAccessToken(code string) (*models.Token, error) {
	url := baseTokenUrl

	type body struct {
		ClientId     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		GrantType    string `json:"grant_type"`
		Code         string `json:"code"`
		RedirectUrl  string `json:"redirect_url"`
	}
	resBytes, err := kawaiihttp.FireHttpRequest(
		kawaiihttp.Post,
		url,
		&body{
			ClientId:     a.ClientId,
			ClientSecret: a.ClientSecret,
			GrantType:    "authorization_code",
			Code:         code,
			RedirectUrl:  a.RedirectUrl,
		},
		time.Second*120,
	)
	if err != nil {
		return nil, err
	}

	accessToken := new(models.Token)
	if err := json.Unmarshal(resBytes, &accessToken); err != nil {
		return nil, fmt.Errorf("error, response is invalid")
	}
	return accessToken, nil
}

func (a *atomikkuApp) RefreshToken(refreshToken string) (*models.Token, error) {
	url := baseTokenUrl

	type body struct {
		ClientId     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		GrantType    string `json:"grant_type"`
		RefreshToken string `json:"refresh_token"`
	}
	resBytes, err := kawaiihttp.FireHttpRequest(
		kawaiihttp.Post,
		url,
		&body{
			ClientId:     a.ClientId,
			ClientSecret: a.ClientSecret,
			GrantType:    "refresh_token",
			RefreshToken: refreshToken,
		},
		time.Second*120,
	)
	if err != nil {
		return nil, err
	}

	accessToken := new(models.Token)
	if err := json.Unmarshal(resBytes, &accessToken); err != nil {
		return nil, fmt.Errorf("error, response is invalid")
	}
	return accessToken, nil
}
