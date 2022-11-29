package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"shopeeadapterapi/middleware"
	_ "shopeeadapterapi/routers"
	"strconv"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	//createAuthTokenUrl()
	runServer()
}

func createAuthTokenUrl() {
	timest := strconv.FormatInt(time.Now().Unix(), 10)
	host := "https://partner.test-stable.shopeemobile.com"
	path := "/api/v2/shop/auth_partner"
	redirectUrl := "https://www.google.com"
	partnerId := strconv.Itoa(1014113)
	partnerKey := "7169744c79764c647258564f4a6b7065526a4f4679716c594c46727970566276"
	baseString := fmt.Sprintf("%s%s%s", partnerId, path, timest)
	h := hmac.New(sha256.New, []byte(partnerKey))
	h.Write([]byte(baseString))
	sign := hex.EncodeToString(h.Sum(nil))
	url := fmt.Sprintf(host+path+"?partner_id=%s&timestamp=%s&sign=%s&redirect=%s", partnerId, timest, sign, redirectUrl)
	fmt.Println(url)
}

func runServer() {
	beego.RunWithMiddleWares("0.0.0.0:8080", middleware.TransformResponse)
}
