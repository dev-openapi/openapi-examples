package main

import (
	"context"
	"testing"

	miniapp "github.com/dev-openapi/douyin-miniapp"
)

var (
	appid  = "tt45ed3350dad1429e01"
	secret = "c80c34c6d70ad3fc72d11465c7fb8f6a4f289008"
	addr   = "https://open-sandbox.douyin.com"

	oauthService = miniapp.NewOauthService(miniapp.WithAddr(addr))
)

func TestGetClientToken(t *testing.T) {
	req := &miniapp.GetClientTokenReq{
		ClientKey:    appid,
		ClientSecret: secret,
		GrantType:    "client_credential",
	}
	res, err := oauthService.GetClientToken(context.Background(), req)
	if err != nil || res == nil {
		t.Error(res, err)
		return
	}
	if res == nil || res.Data == nil || res.Data.ErrorCode != 0 {
		t.Error(res, err)
		return
	}
	t.Log(res, err)
}
