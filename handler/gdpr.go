package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"nico-chan/encrypt"
	"nico-chan/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Gdpr(ctx *gin.Context) {
	gdprResp := model.GdprResp{
		ResponseData: model.GdprRes{
			EnableGdpr:      true,
			IsEea:           false,
			ServerTimestamp: time.Now().Unix(),
		},
		ReleaseInfo: []any{},
		StatusCode:  200,
	}
	resp, err := json.Marshal(gdprResp)
	CheckErr(err)

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")))

	ctx.String(http.StatusOK, string(resp))
}
