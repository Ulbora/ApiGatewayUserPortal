/*
 Copyright (C) 2017 Ulbora Labs Inc. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2017 Ken Williamson
 All rights reserved.

 Certain inventions and disclosures in this file may be claimed within
 patents owned or patent applications filed by Ulbora Labs Inc., or third
 parties.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Affero General Public License as published
 by the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Affero General Public License for more details.

 You should have received a copy of the GNU Affero General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	oauth2 "github.com/Ulbora/go-oauth2-client"
	"github.com/dgrijalva/jwt-go"
)

func getTemplate() string {
	var rtn = ""
	if os.Getenv("TEMPLATE_LOC") != "" {
		rtn = os.Getenv("TEMPLATE_LOC")
	} else {
		rtn = "default"
	}
	return rtn
}

// func getAuthCodeClient() string {
// 	var rtn = ""
// 	if os.Getenv("AUTH_CODE_CLIENT_ID") != "" {
// 		rtn = os.Getenv("AUTH_CODE_CLIENT_ID")
// 	} else {
// 		rtn = authCodeClient
// 	}
// 	return rtn
// }

func getGatewayAPIKey() string {
	var rtn = ""
	if os.Getenv("GATEWAY_API_KEY") != "" {
		rtn = os.Getenv("GATEWAY_API_KEY")
	} else {
		rtn = gwAPIKey // fix this
	}
	return rtn
}

func getAuthCodeSecret() string {
	var rtn = ""
	if os.Getenv("AUTH_CODE_CLIENT_SECRET") != "" {
		rtn = os.Getenv("AUTH_CODE_CLIENT_SECRET")
	} else {
		rtn = authCodeSecret
	}
	return rtn
}
func getOauthHost() string {
	var rtn = ""
	if os.Getenv("AUTH_HOST") != "" {
		rtn = os.Getenv("AUTH_HOST")
	} else {
		rtn = "http://localhost:3000"
	}
	return rtn
}

func getGatewayHost() string {
	var rtn = ""
	if os.Getenv("GATEWAY_HOST") != "" {
		rtn = os.Getenv("GATEWAY_HOST")
	} else {
		rtn = "http://localhost:3011"
	}
	return rtn
}

func getUserHost() string {
	var rtn = ""
	if os.Getenv("USER_HOST") != "" {
		rtn = os.Getenv("USER_HOST")
	} else {
		rtn = "http://localhost:3001"
	}
	return rtn
}
func getRedirectURI(req *http.Request, path string) string {
	var scheme = req.URL.Scheme
	var serverHost string
	if scheme != "" {
		serverHost = req.URL.String()
	} else {
		serverHost = schemeDefault + req.Host + path
	}
	return serverHost
}

func getContentHost() string {
	var rtn = ""
	if os.Getenv("API_GATEWAY_HOST") != "" {
		rtn = os.Getenv("API_GATEWAY_HOST")
	} else if os.Getenv("CONTENT_HOST") != "" {
		rtn = os.Getenv("CONTENT_HOST")
	} else {
		rtn = "http://localhost:3011/content"
	}
	return rtn
}

func getMailHost() string {
	var rtn = ""
	if os.Getenv("API_GATEWAY_HOST") != "" {
		rtn = os.Getenv("API_GATEWAY_HOST")
	} else if os.Getenv("MAIL_HOST") != "" {
		rtn = os.Getenv("MAIL_HOST")
	} else {
		rtn = "http://localhost:3011/mail"
	}
	return rtn
}

func getImageHost() string {
	var rtn = ""
	if os.Getenv("API_GATEWAY_HOST") != "" {
		rtn = os.Getenv("API_GATEWAY_HOST")
	} else if os.Getenv("IMAGE_HOST") != "" {
		rtn = os.Getenv("IMAGE_HOST")
	} else {
		rtn = "http://localhost:3011/image"
	}
	return rtn
}

func getTemplateHost() string {
	var rtn = ""
	if os.Getenv("API_GATEWAY_HOST") != "" {
		rtn = os.Getenv("API_GATEWAY_HOST")
	} else if os.Getenv("TEMPLATE_HOST") != "" {
		rtn = os.Getenv("TEMPLATE_HOST")
	} else {
		rtn = "http://localhost:3011/template"
	}
	return rtn
}

func getChallengeHost() string {
	var rtn = ""
	if os.Getenv("API_GATEWAY_HOST") != "" {
		rtn = os.Getenv("API_GATEWAY_HOST")
	} else if os.Getenv("CHALLENGE_HOST") != "" {
		rtn = os.Getenv("CHALLENGE_HOST")
	} else {
		rtn = "http://localhost:3011/challenge"
	}
	return rtn
}

func getHashedUser(w http.ResponseWriter, r *http.Request) string {
	var rtn string
	//fmt.Println(token.AccessToken)
	token := getToken(w, r)
	tk, err := jwt.Parse(token.AccessToken, func(parsedToken *jwt.Token) (interface{}, error) {
		return parsedToken, nil
	})
	if err != nil {
		fmt.Println(err)
	}
	if tk != nil {
		if claims, ok := tk.Claims.(jwt.MapClaims); ok {
			uid := claims["userId"]
			//fmt.Println(uid)
			if uid != nil {
				rtn = uid.(string)
			}
		}
	} else {
		rtn = ""
	}
	//fmt.Println(rtn)
	return rtn
}

// func getRefreshToken(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("getting refresh token")
// 	token := getToken(w, r)
// 	var tn oauth2.AuthCodeToken
// 	tn.OauthHost = getOauthHost()
// 	tn.ClientID = getAuthCodeClient()
// 	tn.Secret = getAuthCodeSecret()
// 	tn.RefreshToken = token.RefreshToken
// 	fmt.Print("refresh token being sent: ")
// 	fmt.Println(tn.RefreshToken)
// 	resp := tn.AuthCodeRefreshToken()
// 	fmt.Print("refresh token response: ")
// 	fmt.Println(resp)
// 	if resp != nil && resp.AccessToken != "" {
// 		fmt.Print("new token: ")
// 		fmt.Println(resp.AccessToken)
// 		//token = resp
// 		session, err := s.GetSession(r)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		} else {
// 			session.Values["userLoggenIn"] = true
// 			session.Save(r, w)
// 			tokenKey := session.Values["accessTokenKey"]
// 			tokenMap[tokenKey.(string)] = resp
// 			//http.Redirect(res, req, "/admin/main", http.StatusFound)

// 			// decode token and get user id
// 		}
// 	}

// }

// func getCredentialsToken() {
// 	fmt.Println("getting Client Credentials token")
// 	var tn oauth2.ClientCredentialsToken
// 	tn.OauthHost = getOauthHost()
// 	tn.ClientID = getAuthCodeClient()
// 	tn.Secret = getAuthCodeSecret()
// 	resp := tn.ClientCredentialsToken()
// 	//fmt.Print("credentils token response: ")
// 	//fmt.Println(resp)
// 	if resp != nil && resp.AccessToken != "" {
// 		//fmt.Print("new credentials token: ")
// 		//fmt.Println(resp.AccessToken)
// 		credentialToken = resp
// 	}
// }

func generateTokenKey() string {
	return RandStringRunes(9)
}

func generateAPIKey() string {
	return RandStringRunes(35)
}

// func generateClientSecret() string {
// 	return RandStringRunes(50)
// }

//**************Random Generator******************************
func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//RandStringRunes RandStringRunes
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func getToken(w http.ResponseWriter, r *http.Request) *oauth2.Token {
	session, err := s.GetSession(r)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var token *oauth2.Token
	if tokenKey := session.Values["accessTokenKey"]; tokenKey != nil {
		token = tokenMap[tokenKey.(string)]
	}
	return token
}

func removeToken(w http.ResponseWriter, r *http.Request) {
	session, err := s.GetSession(r)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tokenKey := session.Values["accessTokenKey"]
	delete(tokenMap, tokenKey.(string))
}
