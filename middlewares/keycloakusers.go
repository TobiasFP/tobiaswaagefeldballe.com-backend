package middlewares

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type tokenRes struct {
	Access_token       string `json:"access_token"`
	Expires_in         int    `json:"expires_in"`
	Refresh_expires_in int    `json:"refresh_expires_in"`
	Refresh_token      string `json:"refresh_token"`
	Token_type         string `json:"token_type"`
	Notbeforepolicy    int    `json:"notbefore"`
	Session_state      string `json:"session_state"`
	Scope              string `json:"scope"`
}

type groupRes struct {
	Access_token       string `json:"access_token"`
	Expires_in         int    `json:"expires_in"`
	Refresh_expires_in int    `json:"refresh_expires_in"`
	Refresh_token      string `json:"refresh_token"`
	Token_type         string `json:"token_type"`
	Notbeforepolicy    int    `json:"notbefore"`
	Session_state      string `json:"session_state"`
	Scope              string `json:"scope"`
}

type Group struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Path      string   `json:"path"`
	SubGroups []string `json:"subGroups"`
}

type User struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	Enabled       bool   `json:"enabled"`
	EmailVerified bool   `json:"emailVerified"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
}

func SetUsersData(masterClientSecret string, keycloakUrl string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		KEYCLOAK_WORKER_REALM := "ListenToTheRainbow"
		Access_token := getMasterRealToken(keycloakUrl, masterClientSecret)
		ctx.Set("keycloakWorkerRealm", KEYCLOAK_WORKER_REALM)
		ctx.Set("keycloakUrl", keycloakUrl)
		ctx.Set("mastertoken", Access_token)
		ctx.Set("groups", getGroups(keycloakUrl, Access_token, KEYCLOAK_WORKER_REALM))
		ctx.Set("users", getUsers(keycloakUrl, Access_token, KEYCLOAK_WORKER_REALM))
	}
}

func getGroups(keycloakUrl string, Access_token string, KEYCLOAK_WORKER_REALM string) []Group {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, keycloakUrl+"/auth/admin/realms/"+KEYCLOAK_WORKER_REALM+"/groups", nil)
	req.Header.Add("Authorization", "Bearer "+Access_token)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	var groups []Group
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(bodyBytes, &groups)
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}

	}
	return groups
}

func getUsers(keycloakUrl string, Access_token string, KEYCLOAK_WORKER_REALM string) []User {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, keycloakUrl+"/auth/admin/realms/"+KEYCLOAK_WORKER_REALM+"/users", nil)
	req.Header.Add("Authorization", "Bearer "+Access_token)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	var users []User
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(bodyBytes, &users)
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}

	}
	return users
}

func getMasterRealToken(keycloakUrl string, masterClientSecret string) string {
	KEYCLOAK_REALM := "master"
	KEYCLOAK_CLIENT_ID := "admin-cli"
	KEYCLOAK_CLIENT_SECRET := masterClientSecret
	username := "tobias"
	pw := "What4reUL00k1ng4"

	resp, err := http.PostForm(keycloakUrl+"/auth/realms/"+KEYCLOAK_REALM+"/protocol/openid-connect/token", url.Values{
		"grant_type":    {"password"},
		"client_id":     {KEYCLOAK_CLIENT_ID},
		"client_secret": {KEYCLOAK_CLIENT_SECRET},
		"username":      {username},
		"password":      {pw},
	})
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	var tokenResponse tokenRes
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	return tokenResponse.Access_token
}
