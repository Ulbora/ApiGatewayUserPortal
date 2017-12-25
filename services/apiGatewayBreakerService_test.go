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

var RTID22 int64
var RUID22 int64
var RUID222 int64
var BKID int64

func TestGatewayBreakerService_AddRoute(t *testing.T) {
	var r GatewayRouteService
	r.ClientID = testClientStr
	r.Host = "http://localhost:3011"
	r.Token = tempToken
	r.APIKey = testAPIKey

	var rr GatewayRoute
	//rr.ClientID = GwCid33
	rr.Route = "anewroute"
	res := r.AddRoute(&rr)
	RTID22 = res.ID
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayBreakerService_AddRouteURL(t *testing.T) {
	var r GatewayRouteURLService
	r.ClientID = testClientStr
	r.Host = "http://localhost:3011"
	r.Token = tempToken
	r.APIKey = testAPIKey

	var rr GatewayRouteURL
	//rr.ClientID = GwCid33
	rr.RouteID = RTID22
	rr.Name = "blue"
	rr.URL = "http://www.google.com/test"
	res := r.AddRouteURL(&rr)
	RUID22 = res.ID
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayBreakerService_InsertBreaker(t *testing.T) {
	var b GatewayBreakerService
	b.ClientID = testClientStr
	b.Host = "http://localhost:3011"
	b.Token = tempToken
	b.APIKey = testAPIKey

	var bb GatewayBreaker
	//bb.ClientID = GwCid33
	bb.FailoverRouteName = "red"
	bb.FailureThreshold = 3
	bb.HealthCheckTimeSeconds = 60
	bb.RestRouteID = RTID22
	bb.RouteURIID = RUID22
	res := b.InsertBreaker(&bb)
	fmt.Print("breaker res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayBreakerService_GetBreaker(t *testing.T) {
	var b GatewayBreakerService
	b.ClientID = testClientStr
	b.Host = "http://localhost:3011"
	b.Token = tempToken
	b.APIKey = testAPIKey

	res := b.GetBreaker(strconv.FormatInt(RUID22, 10), strconv.FormatInt(RTID22, 10))
	fmt.Print("get breaker res: ")
	fmt.Println(res)
	if res == nil {
		t.Fail()
	} else {
		BKID = res.ID
	}
}

func TestGatewayBreakerService_UpdateBreaker(t *testing.T) {
	var b GatewayBreakerService
	b.ClientID = testClientStr
	b.Host = "http://localhost:3011"
	b.Token = tempToken
	b.APIKey = testAPIKey

	var bb GatewayBreaker
	bb.ID = BKID
	//bb.ClientID = GwCid33
	bb.FailoverRouteName = "green"
	bb.FailureThreshold = 1
	bb.HealthCheckTimeSeconds = 120
	bb.RestRouteID = RTID22
	bb.RouteURIID = RUID22
	res := b.UpdateBreaker(&bb)
	fmt.Print("update breaker res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayBreakerService_GetBreaker2(t *testing.T) {
	var b GatewayBreakerService
	b.ClientID = testClientStr
	b.Host = "http://localhost:3011"
	b.Token = tempToken
	b.APIKey = testAPIKey

	res := b.GetBreaker(strconv.FormatInt(RUID22, 10), strconv.FormatInt(RTID22, 10))
	fmt.Print("get breaker res 2: ")
	fmt.Println(res)
	if res == nil {
		t.Fail()
	}
}

func TestGatewayBreakerService_GetBreakerStatus(t *testing.T) {
	var b GatewayBreakerService
	b.ClientID = testClientStr
	b.Host = "http://localhost:3011"
	b.Token = tempToken
	b.APIKey = testAPIKey

	res := b.GetBreakerStatus(strconv.FormatInt(RUID22, 10))
	fmt.Print("get breaker status: ")
	fmt.Println(res)
	if res == nil {
		t.Fail()
	}
}

func TestGatewayBreakerService_ResetBreakerStatus(t *testing.T) {
	var b GatewayBreakerService
	b.ClientID = testClientStr
	b.Host = "http://localhost:3011"
	b.Token = tempToken
	b.APIKey = testAPIKey

	var bb GatewayBreaker
	//bb.ClientID = GwCid33
	bb.RestRouteID = RTID22
	bb.RouteURIID = RUID22
	res := b.ResetBreaker(&bb)
	fmt.Print("reset breaker res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayBreakerService_DeleteBreakerStatus(t *testing.T) {
	var b GatewayBreakerService
	b.ClientID = testClientStr
	b.Host = "http://localhost:3011"
	b.Token = tempToken
	b.APIKey = testAPIKey

	res := b.DeleteBreaker(strconv.FormatInt(RUID22, 10), strconv.FormatInt(RTID22, 10))
	fmt.Print("delete breaker res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayBreakerService_DeleteRoute(t *testing.T) {
	var r GatewayRouteService
	r.ClientID = testClientStr
	r.Host = "http://localhost:3011"
	r.Token = tempToken
	r.APIKey = testAPIKey

	res := r.DeleteRoute(strconv.FormatInt(RTID22, 10))
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}
