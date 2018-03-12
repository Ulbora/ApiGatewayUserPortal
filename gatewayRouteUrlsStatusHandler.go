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
	"sync"

	"github.com/gorilla/mux"
)

func handleRouteURLsStatus(w http.ResponseWriter, r *http.Request) {
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
		ID := vars["routeId"]
		clientID := session.Values["clientId"].(string)

		if clientID != "" && ID != "" {
			var wg sync.WaitGroup

			var c services.ClientService
			token := getToken(w, r)
			c.ClientID = clientID
			c.Host = getOauthHost()
			c.Token = token.AccessToken
			wg.Add(1)
			var res *services.Client
			go func() {
				res = c.GetClient()
				defer wg.Done()
			}()

			var g services.GatewayClientService
			//token := getToken(w, r)
			g.ClientID = clientID
			g.Host = getGatewayHost()
			g.Token = token.AccessToken
			var gres *services.GatewayClient
			wg.Add(1)
			go func() {
				gres = g.GetClient()
				defer wg.Done()
			}()

			fmt.Println(gres)

			var gr services.GatewayRouteService
			gr.ClientID = clientID
			gr.Host = getGatewayHost()
			gr.Token = token.AccessToken
			var grr *services.GatewayRoute
			wg.Add(1)
			go func(routeID string) {
				grr = gr.GetRoute(routeID)
				defer wg.Done()
			}(ID)

			var gu services.GatewayRouteURLService
			gu.ClientID = clientID
			gu.Host = getGatewayHost()
			gu.Token = token.AccessToken
			wg.Add(1)
			var grus *[]services.GatewayRouteURL
			go func(routeID string) {
				grus = gu.GetRouteURLList(ID)
				defer wg.Done()
			}(ID)
			wg.Wait()

			var cu services.GatewayBreakerService
			cu.ClientID = clientID
			cu.Host = getGatewayHost()
			cu.Token = token.AccessToken
			//var cRes *services.GatewayBreaker

			var grusDisp = make([]gatewayRouteURLDisp, 0)
			for _, u := range *grus {
				var gudisp gatewayRouteURLDisp
				gudisp.ID = u.ID
				gudisp.Name = u.Name
				gudisp.URL = u.URL
				gudisp.RouteID = u.RouteID
				gudisp.ClientID = u.ClientID
				gudisp.Active = u.Active
				cRes := cu.GetBreakerStatus(strconv.FormatInt(u.ID, 10))
				if cRes.Open == true {
					gudisp.Healthy = false
					gudisp.BreakerStatus = "Open"
				} else if cRes.PartialOpen == true {
					gudisp.Healthy = true
					gudisp.BreakerStatus = "Partially Open"
				} else if cRes.Warning == true {
					gudisp.Healthy = true
					gudisp.BreakerStatus = "Warning"
				} else {
					gudisp.Healthy = true
					gudisp.BreakerStatus = "Normal"
				}

				grusDisp = append(grusDisp, gudisp)
			}

			var page gwPage
			page.GwActive = "active"
			page.Client = res
			page.GatewayClient = gres
			page.GatewayRoute = grr
			page.GatewayRouteURLsDisp = &grusDisp
			var sm gwSideMenu
			sm.EditRoute = "active teal"
			page.GwSideMenu = &sm
			//fmt.Println(page)
			templates.ExecuteTemplate(w, "gatewayRouteUrlsStatus.html", &page)
		}
	}
}

func handleResetBreaker(w http.ResponseWriter, r *http.Request) {
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

		IDStr := vars["urlId"]
		ID, _ := strconv.ParseInt(IDStr, 10, 0)
		fmt.Println(IDStr)

		routeIDStr := vars["routeId"]
		//routeID, _ := strconv.ParseInt(routeIDStr, 10, 0)
		fmt.Println(routeIDStr)

		clientIDStr := session.Values["clientId"].(string)
		clientID, _ := strconv.ParseInt(clientIDStr, 10, 0)

		fmt.Print("clientId: ")
		fmt.Println(clientIDStr)

		if IDStr != "" && clientIDStr != "" {
			token := getToken(w, r)

			var cu services.GatewayBreakerService
			cu.ClientID = clientIDStr
			cu.Host = getGatewayHost()
			cu.Token = token.AccessToken

			var cuu services.GatewayBreaker
			cuu.ClientID = clientID
			cuu.RouteURIID = ID

			//guu.Name = name
			//guu.URL = gwURL

			cuRes := cu.ResetBreaker(&cuu)
			if cuRes.Success != true {
				fmt.Println(cuRes)
			}
			http.Redirect(w, r, "/gatewayRouteUrlsStatus/"+routeIDStr, http.StatusFound)
		}
	}
}
