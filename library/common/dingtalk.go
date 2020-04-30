package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"
)

type DingTalkMsg struct {
	Msgtype string       `json:"msgtype"`
	Text    DingTalkText `json:"text" `
	At      AtMobile     `json:"at" `
}

type DingTalkText struct {
	Content string `json:"content"`
}

type AtMobile struct {
	AtMobiles []string `json:"atMobiles"`
}

func SendDingTalk(content string) {
	if Cfg.Env == ENV_DEVELOP_NET || Cfg.DingTalkUrl == "" || Cfg.DingTalkSecret == "" {
		return
	}

	// Current timestamp
	timestamp := Now()
	// Signature string format
	stringToSign := strconv.FormatInt(timestamp, 10) + "\n" + Cfg.DingTalkSecret
	// Sign
	hmacCode := getHmacCode(stringToSign)
	// URL Encode
	sign := url.QueryEscape(hmacCode)
	// Request Body
	param := DingTalkMsg{
		Msgtype: "text",
		Text: DingTalkText{
			Content: Cfg.Env + " : " + content,
		},
		At: AtMobile{
			AtMobiles: Cfg.AtMobiles,
		},
	}
	result, ok := PostDingTalk(fmt.Sprintf(Cfg.DingTalkUrl, timestamp, sign), param)
	if !ok {
		Logger.Errorf("send ding talk error, err: %v", result)
	}
}

// HmacSHA256 && Base64 Encode
func getHmacCode(s string) string {
	h := hmac.New(sha256.New, []byte(Cfg.DingTalkSecret))
	h.Write([]byte(s))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
