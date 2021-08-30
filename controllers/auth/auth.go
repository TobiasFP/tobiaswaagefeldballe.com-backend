package auth

import (
	"backend/config"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

// Auth is A simple struct to handle authentication configuration
type Auth struct {
	Provider *oidc.Provider
	Config   oauth2.Config
}

// Login is the login funktion, that redirects the user to keycloak
func (auth Auth) Login(ctx *gin.Context) {
	state, err := randString(16)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.SetCookie("state", state, int(time.Hour.Seconds()), "", "", false, true)
	ctx.Redirect(http.StatusFound, auth.Config.AuthCodeURL(state))
}

// Callback handles the callback from keycloak/OIDC
func (auth Auth) Callback(ctx *gin.Context) {
	config := config.GetConfig()

	state, err := ctx.Cookie("state")
	if err != nil {
		ctx.Error(err)
		return
	}
	if ctx.Query("state") != state {
		ctx.Error(errors.New("state did not match"))
		return
	}
	oauth2Token, err := auth.Config.Exchange(ctx, ctx.Query("code"))
	if err != nil {
		ctx.Error(errors.New("Failed to exchange token: " + err.Error()))
		return
	}

	// Get IDToken info
	// Extract the ID Token from OAuth2 token.
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		ctx.Error(errors.New("No id token in claims" + err.Error()))
		return
	}

	ctx.Redirect(
		http.StatusPermanentRedirect,
		config.GetString("appUrl")+"/"+
			"callback/"+
			oauth2Token.AccessToken+"/"+
			oauth2Token.RefreshToken+"/"+
			rawIDToken+"/"+
			strconv.FormatInt(oauth2Token.Expiry.Unix(), 10),
	)
}

type tokenInput struct {
	AccessToken string `json:"AccessToken"`
	Idtoken     string `json:"idtoken"`
}

// GetUserInfo retrieves the user info Via a http post
// (since the user info is too big to be received in the callback)
func (auth Auth) GetUserInfo(ctx *gin.Context, oauth2Token oauth2.Token) (*oidc.UserInfo, error) {
	return auth.Provider.UserInfo(ctx, oauth2.StaticTokenSource(&oauth2Token))
}

type idInfo struct {
	Email      string            `json:"email"`
	Verified   bool              `json:"email_verified"`
	Aud        string            `json:"aud"`
	Sub        string            `json:"sub"`
	RealmRoles lttrResourceRealm `json:"realm_access"`
	Groups     []string          `json:"groups"`
}
type lttrResourceRealm struct {
	Roles []string `json:"roles"`
}

// GetFormattedUserInfo retrieves the user info Via a http post, and formats it nicely in json
func (auth Auth) GetFormattedUserInfo(ctx *gin.Context) {
	var token tokenInput
	err := ctx.Bind(&token)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error ": err.Error(),
		})
	}
	var verifier = auth.Provider.Verifier(&oidc.Config{ClientID: auth.Config.ClientID})
	idToken, err := verifier.Verify(ctx, token.Idtoken)
	if err != nil {
		ctx.Error(errors.New("Could not verify id token " + err.Error()))
		return
	}

	// Extract custom claims
	var idInfo idInfo
	if err := idToken.Claims(&idInfo); err != nil {
		ctx.Error(errors.New("Could not parse claims " + err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, idInfo)
}

func randString(nByte int) (string, error) {
	b := make([]byte, nByte)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}
