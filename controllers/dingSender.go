package controllers

import "github.com/wanghuiyt/ding"

func SendDingMsg(msg string) {
	d := ding.Webhook{
		AccessToken: "77f579efbefeefc316b55d3caea1ba1963db2f1319aa7520cbfd9626de073fdc",    // 上面获取的 access_token
		Secret:      "SEC72586e3f7ff6db4b2ad24eac905f308a9ddb0b1b9809af31e5623a14abb424b2", // 上面获取的加签的值
	}
	_ = d.SendMessage(msg)
}
