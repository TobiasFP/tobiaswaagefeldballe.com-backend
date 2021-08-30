package middlewares

import (
	b64 "encoding/base64"
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthData struct {
	Sub          string   `json:"sub"`
	Realm_access roles    `json:"realm_access"`
	Name         string   `json:"name"`
	Groups       []string `json:"groups"`
}

type roles struct {
	Roles []string `json:"roles"`
}

func SetAuthData() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := strings.Fields(ctx.Request.Header.Get("Authorization"))[1]
		tokenUserInfoPart := strings.Split(token, ".")[1]
		decoded, _ := b64.StdEncoding.DecodeString(tokenUserInfoPart)
		decodedString := string(decoded)
		var authData AuthData
		if err := json.Unmarshal([]byte(decodedString), &authData); err != nil {
			if err := json.Unmarshal([]byte(decodedString+"}"), &authData); err != nil {
				if err := json.Unmarshal([]byte(decodedString+"\"}"), &authData); err != nil {
					panic(err)
				}
			}
		}

		ctx.Set("userdata", authData)
	}
}
