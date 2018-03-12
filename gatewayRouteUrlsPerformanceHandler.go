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
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

func handleRouteURLsPerformance(w http.ResponseWriter, r *http.Request) {
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

			var ps services.GatewayPerformanceService
			ps.ClientID = clientID
			ps.Host = getGatewayHost()
			ps.Token = token.AccessToken
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

				var pss services.GatewayPerformance
				pss.ClientID = u.ClientID
				pss.RouteURIID = u.ID
				pss.RestRouteID = u.RouteID
				pRes := ps.GetRoutePerformance(&pss)
				var lat int64
				var cnt int64
				//var dv float64
				//dv = 1000
				for _, p := range *pRes {
					lat += p.LatencyMsTotal
					cnt += p.Calls
				}
				if lat > 0 && cnt > 0 {
					aveLat := (lat / cnt)
					al := float64(aveLat)
					al = al / 1000
					gudisp.AverageLatency = al
					//gudisp.AverageLatency = 1.000
					//fmt.Print("latency: ")
					//fmt.Println(al)
				}
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
			templates.ExecuteTemplate(w, "gatewayRouteUrlsPerformance.html", &page)
		}
	}
}

func handleRouteURLPerformanceByDate(w http.ResponseWriter, r *http.Request) {
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

			var ps services.GatewayPerformanceService
			ps.ClientID = clientID
			ps.Host = getGatewayHost()
			ps.Token = token.AccessToken
			wg.Add(1)
			var pRes *[]services.GatewayPerformance
			go func(urlID string, routeID string, clientID string) {
				var pss services.GatewayPerformance
				pss.ClientID, _ = strconv.ParseInt(clientID, 10, 0)
				pss.RouteURIID, _ = strconv.ParseInt(urlID, 10, 0)
				pss.RestRouteID, _ = strconv.ParseInt(routeID, 10, 0)
				pRes = ps.GetRoutePerformance(&pss)
				defer wg.Done()
			}(urlID, routeID, clientID)

			wg.Wait()

			//fmt.Print("route: ")
			//fmt.Println(grr)

			//fmt.Print("performance: ")
			//fmt.Println(pRes)

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

			page.ChartDate = &cdate
			chdata := buildChartData(pRes)
			aJSON, _ := json.Marshal(chdata)
			page.ChartData = string(aJSON)
			//fmt.Println("chart data JSON: ")
			//fmt.Println(page.ChartData)

			var sm gwSideMenu
			sm.EditURL = "active teal"
			page.GwSideMenu = &sm
			//fmt.Println(page)
			templates.ExecuteTemplate(w, "gatewayRouteUrlPerformanceByDate.html", &page)
		}
	}
}

func buildChartData(pl *[]services.GatewayPerformance) *chartData {
	var rtn chartData
	var c = make([]chartCol, 0)

	var c1 chartCol
	c1.ID = "month"
	c1.Label = "Month"
	c1.Type = "string"
	c = append(c, c1)

	var c2 chartCol
	c2.ID = "lat"
	c2.Label = "Average Latency (Milliseconds)"
	c2.Type = "number"
	c = append(c, c2)

	var c3 chartCol
	c3.ID = "calls"
	c3.Label = "Call Count"
	c3.Type = "number"
	c = append(c, c3)

	rtn.Cols = c

	var rl = make([]chartRow, 0)

	for _, p := range *pl {
		var rvl = make([]chartRowVal, 0)

		var rvd chartRowVal

		rvd.V = p.Entered
		rvl = append(rvl, rvd)

		var rvla chartRowVal
		if p.Calls > 0 {
			var lat = p.LatencyMsTotal / p.Calls
			var latf = float64(lat)
			latf = latf / 1000
			rvla.V = latf
		} else {
			rvla.V = 0
		}
		rvl = append(rvl, rvla)

		var rvc chartRowVal
		rvc.V = p.Calls
		rvl = append(rvl, rvc)

		var r chartRow
		r.C = rvl
		rl = append(rl, r)
	}
	rtn.Rows = rl
	//fmt.Print("chart data: ")
	//fmt.Println(rtn)

	return &rtn
}
