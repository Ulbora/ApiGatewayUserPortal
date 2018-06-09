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

package services

import (
	"fmt"
	"strconv"
	"testing"
)

var RTID222 int64
var RUID2222 int64
var RUID22222 int64

func TestGatewayPerformanceService_AddRoute(t *testing.T) {
	var r GatewayRouteService
	r.ClientID = testClientStr
	r.Host = "http://localhost:3011"
	r.Token = tempToken
	r.APIKey = testAPIKey

	var rr GatewayRoute
	//rr.ClientID = GwCid3334
	rr.Route = "anewroute"
	res := r.AddRoute(&rr)
	RTID222 = res.ID
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayPerformanceService_AddRouteURL(t *testing.T) {
	var r GatewayRouteURLService
	r.ClientID = testClientStr
	r.Host = "http://localhost:3011"
	r.Token = tempToken
	r.APIKey = testAPIKey

	var rr GatewayRouteURL
	//rr.ClientID = GwCid3334
	rr.RouteID = RTID222
	rr.Name = "blue"
	rr.URL = "http://www.google.com/test"
	res := r.AddRouteURL(&rr)
	RUID2222 = res.ID
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayPerformanceService_GetPerformance(t *testing.T) {
	var p GatewayPerformanceService
	p.ClientID = testClientStr
	p.Host = "http://localhost:3011"
	p.Token = tempToken
	p.APIKey = testAPIKey

	var pp GatewayPerformance
	pp.RouteURIID = RUID2222
	pp.RestRouteID = RTID222
	//pp.ClientID = GwCid3334
	res := p.GetRoutePerformance(&pp)
	fmt.Print("get performance res: ")
	fmt.Println(res)
	if res == nil || len(*res) != 0 {
		t.Fail()
	}
}

func TestGatewayPerformanceService_DeleteRoute(t *testing.T) {
	var r GatewayRouteService
	r.ClientID = testClientStr
	r.Host = "http://localhost:3011"
	r.Token = tempToken
	r.APIKey = testAPIKey

	res := r.DeleteRoute(strconv.FormatInt(RTID222, 10))
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}
