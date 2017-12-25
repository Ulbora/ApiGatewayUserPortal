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
	services "ApiGatewayUserPortal/services"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func handleRedirectURLs(w http.ResponseWriter, r *http.Request) {
	s.InitSessionStore(w, r)
	session, err := s.GetSession(r)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	loggedIn := session.Values["userLoggenIn"]
	token := getToken(w, r)
	fmt.Print("in main page. Logged in: ")
	fmt.Println(loggedIn)
	//fmt.Println(token.AccessToken)
	//var res *[]services.Client
	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
		//authorize(w, r)
		loginImplicit(w, r)
	} else {
		session.Values["userLoggenIn"] = true
		//vars := mux.Vars(r)
		clientID := session.Values["clientId"].(string)

		if clientID != "" {
			var c services.ClientService
			token := getToken(w, r)
			c.ClientID = clientID
			c.Host = getOauthHost()
			c.Token = token.AccessToken

			res := c.GetClient()
			//fmt.Println(res)
			var page oauthPage
			page.OauthActive = "active"
			page.Client = res

			var r services.RedirectURIService
			r.ClientID = clientID
			r.Host = getOauthHost()
			r.Token = token.AccessToken
			res2 := r.GetRedirectURIList()
			rul := make([]services.RedirectURI, 0)
			for _, u := range *res2 {
				if strings.Contains(u.URI, "Handler") {
					continue
				} else {
					rul = append(rul, u)
				}
			}
			page.RedirectURLs = &rul
			if len(*res2) > 1 {
				page.CanDeleteRedirectURI = true
			}
			var sm secSideMenu
			sm.RedirectURLActive = "active teal"
			page.SecSideMenu = &sm
			//fmt.Println(page)
			templates.ExecuteTemplate(w, "redirectUrls.html", &page)
		}
	}
}

func handleRedirectURLAdd(w http.ResponseWriter, r *http.Request) {
	s.InitSessionStore(w, r)
	session, err := s.GetSession(r)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	loggedIn := session.Values["userLoggenIn"]
	token := getToken(w, r)
	fmt.Print("in main page. Logged in: ")
	fmt.Println(loggedIn)
	//fmt.Println(token.AccessToken)
	//var res *[]services.Client
	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
		//authorize(w, r)
		loginImplicit(w, r)
	} else {

		redirectURL := r.FormValue("redirectURL")
		fmt.Println(redirectURL)

		clientIDStr := session.Values["clientId"].(string)
		clientID, _ := strconv.ParseInt(clientIDStr, 10, 0)
		fmt.Print("clientId: ")
		fmt.Println(clientID)

		session.Values["userLoggenIn"] = true

		token := getToken(w, r)
		var rl services.RedirectURIService
		rl.ClientID = clientIDStr
		rl.Host = getOauthHost()
		rl.Token = token.AccessToken

		var rr services.RedirectURI
		rr.ClientID = clientID
		rr.URI = redirectURL
		ares := rl.AddRedirectURI(&rr)
		fmt.Println("redirect add: ")
		fmt.Println(ares)
		if ares.Success != true {
			fmt.Println(ares)
		}
		http.Redirect(w, r, "/clientRedirectUrls", http.StatusFound)
	}
}

func handleRedirectURLDelete(w http.ResponseWriter, r *http.Request) {
	s.InitSessionStore(w, r)
	session, err := s.GetSession(r)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	loggedIn := session.Values["userLoggenIn"]
	token := getToken(w, r)
	fmt.Print("in main page. Logged in: ")
	fmt.Println(loggedIn)
	//fmt.Println(token.AccessToken)
	//var res *[]services.Client
	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
		//authorize(w, r)
		loginImplicit(w, r)
	} else {
		session.Values["userLoggenIn"] = true
		vars := mux.Vars(r)
		ID := vars["id"]
		clientID := session.Values["clientId"].(string)

		if ID != "" && clientID != "" {
			token := getToken(w, r)
			var rl services.RedirectURIService
			rl.ClientID = clientID
			rl.Host = getOauthHost()
			rl.Token = token.AccessToken
			dres := rl.DeleteRedirectURI(ID)
			if dres.Success != true {
				fmt.Println(dres)
			}
		}
		http.Redirect(w, r, "/clientRedirectUrls", http.StatusFound)
	}
}
