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
	"strconv"

	"github.com/gorilla/mux"
)

// type userPage struct {
// 	ClientActive string
// 	OauthActive  string
// 	GwActive     string
// 	UserList     *[]services.User
// 	User         *services.User
// 	Client       *services.Client
// }

// user handlers-----------------------------------------------------
func handleUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Print("url: ")
	fmt.Println(r.URL)
	fmt.Println(r.Host)
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

			var u services.UserService
			u.ClientID = clientID
			u.Host = getUserHost()
			u.Token = token.AccessToken
			res2 := u.SearchUserList()
			fmt.Println(res2)

			res3 := u.GetRoleList()
			fmt.Println(res3)

			var page oauthPage
			page.OauthActive = "active"
			page.Client = res
			page.UserList = res2
			page.UserRoleList = res3
			var sm secSideMenu
			sm.UsersActive = "active teal"
			page.SecSideMenu = &sm
			templates.ExecuteTemplate(w, "users.html", &page)
		}
	}

}

// func handleAddUser(w http.ResponseWriter, r *http.Request) {
// 	s.InitSessionStore(w, r)
// 	session, err := s.GetSession(r)
// 	if err != nil {
// 		fmt.Println(err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// 	loggedIn := session.Values["userLoggenIn"]
// 	token := getToken(w, r)
// 	fmt.Print("Logged in: ")
// 	fmt.Println(loggedIn)
// 	//fmt.Println(token.AccessToken)
// 	//var res *[]services.Client
// 	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
// 		authorize(w, r)
// 	} else {
// 		templates.ExecuteTemplate(w, "addClient.html", nil)
// 	}

// }

func handleNewUser(w http.ResponseWriter, r *http.Request) {
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
		var uu services.User
		clientID := session.Values["clientId"].(string)
		fmt.Print("clientID: ")
		fmt.Println(clientID)
		uu.ClientID = clientID

		username := r.FormValue("username")
		fmt.Print("username: ")
		fmt.Println(username)
		uu.Username = username

		userRoleIDStr := r.FormValue("userRoleId")
		userRoleID, _ := strconv.ParseInt(userRoleIDStr, 10, 0)
		fmt.Print("userRoleID: ")
		fmt.Println(userRoleID)
		uu.RoleID = userRoleID

		firstName := r.FormValue("firstName")
		fmt.Print("firstName: ")
		fmt.Println(firstName)
		uu.FirstName = firstName

		lastName := r.FormValue("lastName")
		fmt.Print("lastName: ")
		fmt.Println(lastName)
		uu.LastName = lastName

		emailAddress := r.FormValue("emailAddress")
		fmt.Print("emailAddress: ")
		fmt.Println(emailAddress)
		uu.EmailAddress = emailAddress

		password := r.FormValue("password")
		fmt.Print("password: ")
		fmt.Println(password)
		uu.Password = password

		enabled := r.FormValue("enabled")
		fmt.Print("enabled: ")
		fmt.Println(enabled)
		if enabled == "yes" {
			uu.Enabled = true
		}

		var u services.UserService
		token := getToken(w, r)
		u.ClientID = clientID
		u.Host = getUserHost()
		u.Token = token.AccessToken

		//cc.Secret = generateClientSecret()
		res := u.AddUser(&uu)
		if res.Success != true {
			fmt.Print("Add User Failed: ")
			fmt.Println(res)
		}
		http.Redirect(w, r, "/users", http.StatusFound)
	}
}

func handleEditUser(w http.ResponseWriter, r *http.Request) {
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
		vars := mux.Vars(r)
		username := vars["username"]
		clientID := session.Values["clientId"].(string)

		if clientID != "" {
			var c services.ClientService
			token := getToken(w, r)
			c.ClientID = clientID
			c.Host = getOauthHost()
			c.Token = token.AccessToken

			res := c.GetClient()

			var u services.UserService
			u.ClientID = clientID
			u.Host = getUserHost()
			u.Token = token.AccessToken
			res2 := u.GetUser(username)
			fmt.Println(res2)

			res3 := u.GetRoleList()
			fmt.Println(res3)

			var page oauthPage
			page.OauthActive = "active"
			page.Client = res
			page.User = res2
			page.UserRoleList = res3
			page.UserAssignedRole = res2.RoleID

			var sm secSideMenu
			//sm.UsersActive = "active teal"
			page.SecSideMenu = &sm
			templates.ExecuteTemplate(w, "editUser.html", &page)
		}
	}
}

func handleUpdateUserInfo(w http.ResponseWriter, r *http.Request) {
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
		var uu services.UserInfo
		clientID := session.Values["clientId"].(string)
		fmt.Print("clientID: ")
		fmt.Println(clientID)
		uu.ClientID = clientID

		username := r.FormValue("username")
		fmt.Print("username: ")
		fmt.Println(username)
		uu.Username = username

		userRoleIDStr := r.FormValue("userRoleId")
		userRoleID, _ := strconv.ParseInt(userRoleIDStr, 10, 0)
		fmt.Print("userRoleID: ")
		fmt.Println(userRoleID)
		uu.RoleID = userRoleID

		firstName := r.FormValue("firstName")
		fmt.Print("firstName: ")
		fmt.Println(firstName)
		uu.FirstName = firstName

		lastName := r.FormValue("lastName")
		fmt.Print("lastName: ")
		fmt.Println(lastName)
		uu.LastName = lastName

		emailAddress := r.FormValue("emailAddress")
		fmt.Print("emailAddress: ")
		fmt.Println(emailAddress)
		uu.EmailAddress = emailAddress

		var u services.UserService
		token := getToken(w, r)
		u.ClientID = clientID
		u.Host = getUserHost()
		u.Token = token.AccessToken

		//cc.Secret = generateClientSecret()
		res := u.UpdateUser(&uu)
		if res.Success != true {
			fmt.Print("Update User Info Failed: ")
			fmt.Println(res)
		}
		http.Redirect(w, r, "/editUser/"+username, http.StatusFound)
	}
}

func handleUpdateUserEnable(w http.ResponseWriter, r *http.Request) {
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
		var uu services.UserDis
		clientID := session.Values["clientId"].(string)
		fmt.Print("clientID: ")
		fmt.Println(clientID)
		uu.ClientID = clientID

		username := r.FormValue("username")
		fmt.Print("username: ")
		fmt.Println(username)
		uu.Username = username

		enabled := r.FormValue("enabled")
		fmt.Print("enabled: ")
		fmt.Println(enabled)
		if enabled == "yes" {
			uu.Enabled = true
		}

		var u services.UserService
		token := getToken(w, r)
		u.ClientID = clientID
		u.Host = getUserHost()
		u.Token = token.AccessToken

		//cc.Secret = generateClientSecret()
		res := u.UpdateUser(&uu)
		if res.Success != true {
			fmt.Print("Update User Enabled Failed: ")
			fmt.Println(res)
		}
		http.Redirect(w, r, "/editUser/"+username, http.StatusFound)
	}
}

func handleUpdateUserPw(w http.ResponseWriter, r *http.Request) {
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
		var uu services.UserPW
		clientID := session.Values["clientId"].(string)
		fmt.Print("clientID: ")
		fmt.Println(clientID)
		uu.ClientID = clientID

		username := r.FormValue("username")
		fmt.Print("username: ")
		fmt.Println(username)
		uu.Username = username

		password := r.FormValue("password")
		fmt.Print("password: ")
		fmt.Println(password)
		uu.Password = password

		var u services.UserService
		token := getToken(w, r)
		u.ClientID = clientID
		u.Host = getUserHost()
		u.Token = token.AccessToken

		//cc.Secret = generateClientSecret()
		res := u.UpdateUser(&uu)
		if res.Success != true {
			fmt.Print("Update User Password Failed: ")
			fmt.Println(res)
		}
		http.Redirect(w, r, "/editUser/"+username, http.StatusFound)
	}
}
