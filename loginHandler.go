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
	"net/http"

	oauth2 "github.com/Ulbora/go-oauth2-client"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	//s.InitSessionStore(w, r)
	loginImplicit(w, r)
}

// login handler
func handleLogout(w http.ResponseWriter, r *http.Request) {
	removeToken(w, r)
	cookie := &http.Cookie{
		Name:   "ugw-user-session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	cookie2 := &http.Cookie{
		Name:   "ulbora_oauth2_server",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie2)
	http.Redirect(w, r, "/", http.StatusFound)
}

// func authorize(res http.ResponseWriter, req *http.Request) bool {
// 	fmt.Println("in authorize")
// 	fmt.Println(schemeDefault)
// 	var a oauth2.AuthCodeAuthorize
// 	a.ClientID = getAuthCodeClient()
// 	a.OauthHost = getOauthHost()
// 	a.RedirectURI = getRedirectURI(req, authCodeRedirectURI)
// 	a.Scope = "write"
// 	a.State = authCodeState
// 	a.Res = res
// 	a.Req = req
// 	resp := a.AuthCodeAuthorizeUser()
// 	if resp != true {
// 		fmt.Println("Authorize Failed")
// 	}
// 	fmt.Print("Resp: ")
// 	fmt.Println(resp)
// 	return resp
// }

func loginImplicit(w http.ResponseWriter, r *http.Request) {
	//s.InitSessionStore(w, r)
	templates.ExecuteTemplate(w, "login.html", nil)
}

func handleImplicitLogin(w http.ResponseWriter, r *http.Request) {
	s.InitSessionStore(w, r)
	clientID := r.FormValue("clientId")
	session, err := s.GetSession(r)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		session.Values["clientId"] = clientID
		//session.Values["testBool"] = true
		serr := session.Save(r, w)
		fmt.Println(serr)
		fmt.Print("clientId: ")
		fmt.Println(clientID)

		fmt.Print("clientId from session: ")
		fmt.Println(session.Values["clientId"].(string))

		var a oauth2.ImplicitAuthorize
		a.ClientID = clientID
		a.OauthHost = getOauthRedirectHost()
		a.RedirectURI = getRedirectURI(r, implicitRedirectURI)
		a.Scope = "write"
		a.State = authCodeState
		a.Req = r
		a.Res = w
		resp := a.ImplicitAuthorize()
		//fmt.Print("RedirectURI: ")
		//fmt.Println(a.RedirectURI)
		if resp != true {
			fmt.Println("Authorize Failed")
		}
		//fmt.Print("Resp: ")
		//fmt.Println(resp)
	}
}

func handleImplicitToken(res http.ResponseWriter, req *http.Request) {
	token := req.URL.Query().Get("token")
	state := req.URL.Query().Get("state")
	//fmt.Println("handle token")
	if state == authCodeState {
		if token != "" {
			fmt.Println(token)
			//token = resp
			session, err := s.GetSession(req)
			if err != nil {
				fmt.Println(err)
				http.Error(res, err.Error(), http.StatusInternalServerError)
			} else {
				//session.Values["clientId"] = "616"
				//session.Values["testBool"] = true
				session.Values["userLoggenIn"] = true
				var accKey = generateTokenKey()
				session.Values["accessTokenKey"] = accKey
				var resp oauth2.Token
				resp.AccessToken = token
				tokenMap[accKey] = &resp
				//fmt.Print("session id: ")
				//fmt.Println(session.ID)
				err := session.Save(req, res)
				fmt.Println(err)
				http.Redirect(res, req, "/oauth2", http.StatusFound)
				// decode token and get user id
			}
		}
	}
}

// func handleToken(res http.ResponseWriter, req *http.Request) {
// 	code := req.URL.Query().Get("code")
// 	state := req.URL.Query().Get("state")
// 	//fmt.Println("handle token")
// 	if state == authCodeState {
// 		var tn oauth2.AuthCodeToken
// 		tn.OauthHost = getOauthHost()
// 		tn.ClientID = getAuthCodeClient()
// 		tn.Secret = getAuthCodeSecret()
// 		tn.Code = code
// 		tn.RedirectURI = getRedirectURI(req, authCodeRedirectURI)
// 		//fmt.Println("getting token")
// 		resp := tn.AuthCodeToken()
// 		fmt.Print("token len: ")
// 		fmt.Println(len(resp.AccessToken))
// 		if resp != nil && resp.AccessToken != "" {
// 			fmt.Println(resp.AccessToken)
// 			//token = resp
// 			session, err := s.GetSession(req)
// 			if err != nil {
// 				fmt.Println(err)
// 				http.Error(res, err.Error(), http.StatusInternalServerError)
// 			} else {
// 				session.Values["userLoggenIn"] = true
// 				var accKey = generateTokenKey()
// 				session.Values["accessTokenKey"] = accKey
// 				tokenMap[accKey] = resp
// 				//fmt.Print("session id: ")
// 				//fmt.Println(session.ID)
// 				err := session.Save(req, res)
// 				fmt.Println(err)
// 				http.Redirect(res, req, "/oauth2", http.StatusFound)

// 				// decode token and get user id
// 			}
// 		}
// 	}
// }
