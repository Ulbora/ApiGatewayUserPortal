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

	"github.com/gorilla/mux"
)

type gwRoutesDisplay struct {
	ID           int64
	URI          string
	ClientID     int64
	AssignedRole int64
}

func handleRoutes(w http.ResponseWriter, r *http.Request) {
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

			var g services.GatewayClientService
			//token := getToken(w, r)
			g.ClientID = clientID
			g.Host = getGatewayHost()
			g.Token = token.AccessToken

			gres := g.GetClient()
			//fmt.Println(gres)

			var gr services.GatewayRouteService
			gr.ClientID = clientID
			gr.Host = getGatewayHost()
			gr.Token = token.AccessToken
			grr := gr.GetRouteList()
			fmt.Print("routes: ")
			fmt.Println(grr)
			var page gwPage
			page.GwActive = "active"
			page.Client = res
			page.GatewayClient = gres
			page.GatewayRoutes = grr
			var sm gwSideMenu
			sm.RouteActive = "active teal"
			page.GwSideMenu = &sm
			//fmt.Println(page)
			templates.ExecuteTemplate(w, "gatewayRoutes.html", &page)
		}
	}
}

func handleRoutesAdd(w http.ResponseWriter, r *http.Request) {
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
		gwRoute := r.FormValue("gwRoute")
		fmt.Println(gwRoute)

		clientIDStr := session.Values["clientId"].(string)
		clientID, _ := strconv.ParseInt(clientIDStr, 10, 0)
		fmt.Print("clientId: ")
		fmt.Println(clientID)
		if gwRoute != "" && clientIDStr != "" {
			var gr services.GatewayRouteService
			gr.ClientID = clientIDStr
			gr.Host = getGatewayHost()
			gr.Token = token.AccessToken

			var grr services.GatewayRoute
			grr.ClientID = clientID
			grr.Route = gwRoute
			grRes := gr.AddRoute(&grr)
			if grRes.Success != true {
				fmt.Println(grRes)
			}
			fmt.Println(grRes)
		}
		http.Redirect(w, r, "/gatewayRoutes", http.StatusFound)
	}
}

func handleRouteEdit(w http.ResponseWriter, r *http.Request) {
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

		routeID := vars["id"]
		clientID := session.Values["clientId"].(string)

		if clientID != "" && routeID != "" {
			var c services.ClientService
			token := getToken(w, r)
			c.ClientID = clientID
			c.Host = getOauthHost()
			c.Token = token.AccessToken

			res := c.GetClient()

			var g services.GatewayClientService
			//token := getToken(w, r)
			g.ClientID = clientID
			g.Host = getGatewayHost()
			g.Token = token.AccessToken

			gres := g.GetClient()
			//fmt.Println(gres)

			var gr services.GatewayRouteService
			gr.ClientID = clientID
			gr.Host = getGatewayHost()
			gr.Token = token.AccessToken
			grr := gr.GetRoute(routeID)
			var page gwPage
			page.GwActive = "active"
			page.Client = res
			page.GatewayClient = gres
			page.GatewayRoute = grr
			var sm gwSideMenu
			//sm.RouteActive = "active teal"
			page.GwSideMenu = &sm
			//fmt.Println(page)
			templates.ExecuteTemplate(w, "editGatewayRoute.html", &page)
		}
	}
}

func handleRouteUpdate(w http.ResponseWriter, r *http.Request) {
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

		IDStr := r.FormValue("id")
		fmt.Println(IDStr)
		ID, _ := strconv.ParseInt(IDStr, 10, 0)
		fmt.Println(ID)

		clientIDStr := session.Values["clientId"].(string)
		clientID, _ := strconv.ParseInt(clientIDStr, 10, 0)
		fmt.Print("clientId: ")
		fmt.Println(clientID)

		gwRoute := r.FormValue("gwRoute")
		fmt.Print("gwRoute: ")
		fmt.Println(gwRoute)

		if IDStr != "" && clientIDStr != "" && gwRoute != "" {
			var gr services.GatewayRouteService
			gr.ClientID = clientIDStr
			gr.Host = getGatewayHost()
			gr.Token = token.AccessToken

			var grr services.GatewayRoute
			grr.ClientID = clientID
			grr.Route = gwRoute
			grr.ID = ID
			grRes := gr.UpdateRoute(&grr)
			if grRes.Success != true {
				fmt.Println(grRes)
			}
		}
		http.Redirect(w, r, "/gatewayRoutes", http.StatusFound)

	}
}

func handleRoutesDelete(w http.ResponseWriter, r *http.Request) {
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

		IDStr := vars["id"]
		fmt.Println(IDStr)

		clientIDStr := session.Values["clientId"].(string)

		fmt.Print("clientId: ")
		fmt.Println(clientIDStr)

		if IDStr != "" && clientIDStr != "" {
			token := getToken(w, r)

			var gr services.GatewayRouteService
			gr.ClientID = clientIDStr
			gr.Host = getGatewayHost()
			gr.Token = token.AccessToken

			gres := gr.DeleteRoute(IDStr)

			fmt.Println(gres)
			http.Redirect(w, r, "/gatewayRoutes", http.StatusFound)
		}
	}
}
