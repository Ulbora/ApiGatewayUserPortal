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

package services

import (
	"fmt"
	"strconv"
	"testing"
)

var RTID2225 int64
var RUID22225 int64
var RUID222225 int64

func TestGatewayErrorsService_AddRoute(t *testing.T) {
	var r GatewayRouteService
	r.ClientID = testClientStr
	r.Host = "http://localhost:3011"
	r.Token = tempToken
	r.APIKey = testAPIKey

	var rr GatewayRoute
	//rr.ClientID = GwCid3335
	rr.Route = "anewroute"
	res := r.AddRoute(&rr)
	RTID2225 = res.ID
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayErrorsService_AddRouteURL(t *testing.T) {
	var r GatewayRouteURLService
	r.ClientID = testClientStr
	r.Host = "http://localhost:3011"
	r.Token = tempToken
	r.APIKey = testAPIKey

	var rr GatewayRouteURL
	//rr.ClientID = GwCid3335
	rr.RouteID = RTID2225
	rr.Name = "blue"
	rr.URL = "http://www.google.com/test"
	res := r.AddRouteURL(&rr)
	RUID22225 = res.ID
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayErrorsService_GetErrors(t *testing.T) {
	var e GatewayErrorsService
	e.ClientID = testClientStr
	e.Host = "http://localhost:3011"
	e.Token = tempToken
	e.APIKey = testAPIKey

	var ee GatewayError
	ee.RouteURIID = RUID22225
	ee.RestRouteID = RTID2225
	//ee.ClientID = GwCid3335
	res := e.GetRouteErrors(&ee)
	fmt.Print("get performance res: ")
	fmt.Println(res)
	if res == nil || len(*res) != 0 {
		t.Fail()
	}
}

func TestGatewayErrorsService_DeleteRoute(t *testing.T) {
	var r GatewayRouteService
	r.ClientID = testClientStr
	r.Host = "http://localhost:3011"
	r.Token = tempToken
	r.APIKey = testAPIKey

	res := r.DeleteRoute(strconv.FormatInt(RTID2225, 10))
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}
