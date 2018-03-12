package main

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

import (
	services "ApiGatewayUserPortal/services"
	"fmt"
	"net/http"
)

type secSideMenu struct {
	RedirectURLActive string
	GrantTypeActive   string
	RolesActive       string
	AllowedURIActive  string
	ClientActive      string
	UlboraURIsActive  string
	UsersActive       string
}

type oauthPage struct {
	ClientActive         string
	OauthActive          string
	GwActive             string
	CanDeleteRedirectURI bool
	ClientIsSelf         bool
	SecSideMenu          *secSideMenu
	ClientList           *[]services.Client
	Client               *services.Client
	RedirectURLs         *[]services.RedirectURI
	GrantTypes           *[]services.GrantType
	ClientRoles          *[]services.ClientRole
	AllowedURIs          *[]allowedURIDisplay
	RoleURIs             *[]services.RoleURI
	UserList             *[]services.User
	User                 *services.User
	UserRoleList         *[]services.Role
	UserAssignedRole     int64
}

func handleOauth2(w http.ResponseWriter, r *http.Request) {
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
		clientID := session.Values["clientId"].(string) //:= getAuthCodeClient() //vars["clientId"]

		if clientID != "" {
			var c services.ClientService
			token := getToken(w, r)
			c.ClientID = clientID //getAuthCodeClient()
			c.Host = getOauthHost()
			c.Token = token.AccessToken

			res := c.GetClient()
			//fmt.Println(res)
			var page oauthPage
			page.OauthActive = "active"
			page.Client = res
			var sm secSideMenu
			sm.ClientActive = "active teal"
			page.SecSideMenu = &sm
			//fmt.Println(page)
			templates.ExecuteTemplate(w, "oauth2.html", &page)
		}
	}
}
