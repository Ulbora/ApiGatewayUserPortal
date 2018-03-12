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

type gwSideMenu struct {
	GWActive        string
	GWClientActive  string
	RouteActive     string
	RouteURLsActive string
	EditRoute       string
	EditURL         string
}

type gwPage struct {
	ClientActive          string
	OauthActive           string
	GwActive              string
	ClientIsSelf          bool
	GwSideMenu            *gwSideMenu
	Client                *services.Client
	User                  *services.User
	GatewayClient         *services.GatewayClient
	GatewayRoutes         *[]services.GatewayRoute
	GatewayRoute          *services.GatewayRoute
	GatewayRouteURIs      *[]services.GatewayRouteURL
	GatewayRouteURI       *services.GatewayRouteURL
	CircuitBreaker        *services.GatewayBreaker
	CircuitBreakerEnabled bool
	GatewayRouteURLsDisp  *[]gatewayRouteURLDisp
	GatewayRouteURLDisp   *gatewayRouteURLDisp
	ChartData             string
	ChartDate             *chartDate
	Errors                *[]services.GatewayError
	ErrorPages            int
	ErrorPageRange        *[]errorPageRange
	ErrorPageCurrent      int
}

type gatewayRouteURLDisp struct {
	ID             int64   `json:"id"`
	RouteID        int64   `json:"routeId"`
	ClientID       int64   `json:"clientId"`
	Name           string  `json:"name"`
	URL            string  `json:"url"`
	Active         bool    `json:"active"`
	BreakerStatus  string  `json:"status"`
	Healthy        bool    `json:"healthy"`
	AverageLatency float64 `json:"averageLatency"`
	ErrorCount     int     `json:"errorCount"`
}

type chartDate struct {
	Month int
	Year  int
}

type chartCol struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Type  string `json:"type"`
}

type chartRowVal struct {
	V interface{} `json:"v"`
}

type chartRow struct {
	C []chartRowVal `json:"c"`
}

type chartData struct {
	Cols []chartCol `json:"cols"`
	Rows []chartRow `json:"rows"`
}

type errorPageRange struct {
	Pg int
}

func handleGateway(w http.ResponseWriter, r *http.Request) {
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
			//fmt.Println(res)
			var page gwPage
			page.GwActive = "active"
			page.Client = res
			page.GatewayClient = gres
			var sm gwSideMenu
			sm.GWActive = "active teal"
			page.GwSideMenu = &sm
			//fmt.Println(page)
			templates.ExecuteTemplate(w, "gateway.html", &page)
		}
	}
}
