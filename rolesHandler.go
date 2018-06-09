package main

/*
 Copyright (C) 2017 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2017 Ken Williamson
 All rights reserved.

 Certain inventions and disclosures in this file may be claimed within
 patents owned or patent applications filed by Ulbora Labs LLC., or third
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
	"strconv"

	"github.com/gorilla/mux"
)

func handleRoles(w http.ResponseWriter, r *http.Request) {
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

			var r services.ClientRoleService
			r.ClientID = clientID
			r.Host = getOauthHost()
			r.Token = token.AccessToken
			res2 := r.GetClientRoleList()
			page.ClientRoles = res2
			//if getAuthCodeClient() == clientID {
			//page.ClientIsSelf = true
			//}
			var sm secSideMenu
			sm.RolesActive = "active teal"
			page.SecSideMenu = &sm

			//fmt.Println(page)
			templates.ExecuteTemplate(w, "roles.html", &page)
		}
	}
}

func handleRoleAdd(w http.ResponseWriter, r *http.Request) {
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

		clientRole := r.FormValue("clientRole")
		fmt.Println(clientRole)

		clientIDStr := session.Values["clientId"].(string)
		clientID, _ := strconv.ParseInt(clientIDStr, 10, 0)
		fmt.Print("clientId: ")
		fmt.Println(clientID)

		session.Values["userLoggenIn"] = true

		token := getToken(w, r)

		var rs services.ClientRoleService
		rs.ClientID = clientIDStr
		rs.Host = getOauthHost()
		rs.Token = token.AccessToken
		resTest := rs.GetClientRoleList()
		var roleExists = false
		for _, rl := range *resTest {
			if rl.Role == clientRole {
				roleExists = true
				break
			}
		}
		if clientRole != "" && clientRole != "superAdmin" && roleExists != true {
			var rr services.ClientRole
			rr.ClientID = clientID
			rr.Role = clientRole
			rres := rs.AddClientRole(&rr)
			fmt.Println(rres)
			if rres.Success != true {
				fmt.Println(rres)
			}
		}
		http.Redirect(w, r, "/clientRoles", http.StatusFound)
	}
}

func handleRoleDelete(w http.ResponseWriter, r *http.Request) {
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
			var rl services.ClientRoleService
			rl.ClientID = clientID
			rl.Host = getOauthHost()
			rl.Token = token.AccessToken
			rres := rl.DeleteClientRole(ID)
			if rres.Success != true {
				fmt.Println(rres)
			}
		}
		http.Redirect(w, r, "/clientRoles", http.StatusFound)
	}
}
