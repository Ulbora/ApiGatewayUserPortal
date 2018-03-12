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
	"math"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

var pageSize = 20

func handleRouteURLsErrors(w http.ResponseWriter, r *http.Request) {
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

			//fmt.Println(gres)

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

			var es services.GatewayErrorsService
			es.ClientID = clientID
			es.Host = getGatewayHost()
			es.Token = token.AccessToken
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

				var ess services.GatewayError
				ess.ClientID = u.ClientID
				ess.RouteURIID = u.ID
				ess.RestRouteID = u.RouteID
				eRes := es.GetRouteErrors(&ess)
				gudisp.ErrorCount = len(*eRes)
				// var lat int64
				// var cnt int64
				//var dv float64
				//dv = 1000
				// for _, p := range *pRes {
				// 	lat += p.LatencyMsTotal
				// 	cnt += p.Calls
				// }
				// if lat > 0 && cnt > 0 {
				// 	aveLat := (lat / cnt)
				// 	al := float64(aveLat)
				// 	al = al / 1000
				// 	gudisp.AverageLatency = al
				// 	//gudisp.AverageLatency = 1.000
				// 	//fmt.Print("latency: ")
				// 	//fmt.Println(al)
				// }
				//fmt.Print("latency val: ")
				//fmt.Println(gudisp.AverageLatency)

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
			templates.ExecuteTemplate(w, "gatewayRouteUrlsErrors.html", &page)
		}
	}
}

func handleRouteURLError(w http.ResponseWriter, r *http.Request) {
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
		urlID := vars["urlId"]
		routeID := vars["routeId"]
		clientID := session.Values["clientId"].(string)
		pageStr := vars["page"]
		var pageNum int
		var err error
		if pageStr != "" {
			if pageNum, err = strconv.Atoi(pageStr); err != nil {
				pageNum = 1
			}
		} else {
			pageNum = 1
		}
		//fmt.Print("urlId: ")
		//fmt.Println(urlID)
		//fmt.Print("routeId: ")
		//fmt.Println(routeID)

		if clientID != "" && routeID != "" && urlID != "" {
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

			//fmt.Println(gres)

			var gr services.GatewayRouteService
			gr.ClientID = clientID
			gr.Host = getGatewayHost()
			gr.Token = token.AccessToken
			var grr *services.GatewayRoute
			wg.Add(1)
			go func(routeID string) {
				grr = gr.GetRoute(routeID)
				defer wg.Done()
			}(routeID)

			var gu services.GatewayRouteURLService
			gu.ClientID = clientID
			gu.Host = getGatewayHost()
			gu.Token = token.AccessToken
			wg.Add(1)
			var u *services.GatewayRouteURL
			go func(urlID string, routeID string) {
				u = gu.GetRouteURL(urlID, routeID)
				defer wg.Done()
			}(urlID, routeID)

			var es services.GatewayErrorsService
			es.ClientID = clientID
			es.Host = getGatewayHost()
			es.Token = token.AccessToken
			wg.Add(1)
			var eRes *[]services.GatewayError
			go func(urlID string, routeID string, clientID string) {
				var ess services.GatewayError
				ess.ClientID, _ = strconv.ParseInt(clientID, 10, 0)
				ess.RouteURIID, _ = strconv.ParseInt(urlID, 10, 0)
				ess.RestRouteID, _ = strconv.ParseInt(routeID, 10, 0)
				eRes = es.GetRouteErrors(&ess)
				defer wg.Done()
			}(urlID, routeID, clientID)

			wg.Wait()

			//fmt.Print("route: ")
			//fmt.Println(grr)

			//fmt.Print("errors: ")
			//fmt.Println(eRes)

			var gudisp gatewayRouteURLDisp
			gudisp.ID = u.ID
			gudisp.Name = u.Name
			gudisp.URL = u.URL
			gudisp.RouteID = u.RouteID
			gudisp.ClientID = u.ClientID
			gudisp.Active = u.Active

			tnow := time.Now()
			var cdate chartDate
			m := tnow.Month()
			y := tnow.Year()
			cdate.Month = int(m) - 1
			cdate.Year = int(y)

			var page gwPage
			page.GwActive = "active"
			page.Client = res
			page.GatewayClient = gres
			page.GatewayRoute = grr
			page.GatewayRouteURLDisp = &gudisp
			ep := make([]services.GatewayError, 0)
			epr := make([]errorPageRange, 0)
			////page.Errors = eRes
			//for (pageSize * page)
			pcf := float64(len(*eRes)) / float64(pageSize)
			pc := math.Ceil(pcf)
			pcint := int(pc)
			page.ErrorPages = pcint
			page.ErrorPageCurrent = pageNum
			//var pcrng int
			//pcrng = 1

			//fmt.Print("error page number: ")
			//fmt.Println(page.ErrorPages)

			for pcrng := 1; pcrng <= pcint; pcrng++ {
				var eprr errorPageRange
				eprr.Pg = pcrng
				epr = append(epr, eprr)
			}
			//fmt.Print("error page range: ")
			//fmt.Println(epr)

			for i, e := range *eRes {
				if (i < (pageSize * pageNum)) && (i >= (pageSize * (pageNum - 1))) {
					ep = append(ep, e)
				}
			}
			page.Errors = &ep
			page.ErrorPageRange = &epr

			var sm gwSideMenu
			sm.EditURL = "active teal"
			page.GwSideMenu = &sm
			//fmt.Println(page)
			templates.ExecuteTemplate(w, "gatewayRouteUrlViewErrors.html", &page)
		}
	}
}
