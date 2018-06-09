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

var RTID2 int64
var RUID int64
var RUID2 int64

func TestGatewayRouteURLService_AddRoute(t *testing.T) {
	var r GatewayRouteService
	r.ClientID = testClientStr
	r.Host = "http://localhost:3011"
	r.Token = tempToken
	r.APIKey = testAPIKey

	var rr GatewayRoute
	//rr.ClientID = GwCid3
	rr.Route = "anewroute"
	res := r.AddRoute(&rr)
	RTID2 = res.ID
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayRouteURLService_AddRouteURL(t *testing.T) {
	var r GatewayRouteURLService
	r.ClientID = testClientStr
	r.Host = "http://localhost:3011"
	r.Token = tempToken
	r.APIKey = testAPIKey

	var rr GatewayRouteURL
	//rr.ClientID = GwCid3
	rr.RouteID = RTID2
	rr.Name = "blue"
	rr.URL = "http://www.google.com/test"
	res := r.AddRouteURL(&rr)
	RUID = res.ID
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayRouteURLService_UpdateRouteURL(t *testing.T) {
	var r GatewayRouteURLService
	r.ClientID = testClientStr
	r.Host = "http://localhost:3011"
	r.Token = tempToken
	r.APIKey = testAPIKey

	var rr GatewayRouteURL
	rr.ID = RUID
	//rr.ClientID = GwCid3
	rr.RouteID = RTID2
	rr.Name = "green"
	rr.URL = "http://www.google.com/green/test"
	res := r.UpdateRouteURL(&rr)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayRouteURLService_GetRouteURL(t *testing.T) {
	var r GatewayRouteURLService
	r.ClientID = testClientStr
	r.Host = "http://localhost:3011"
	r.Token = tempToken
	r.APIKey = testAPIKey

	res := r.GetRouteURL(strconv.FormatInt(RUID, 10), strconv.FormatInt(RTID2, 10))
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Name != "green" || res.URL != "http://www.google.com/green/test" {
		t.Fail()
	}
}

func TestGatewayRouteURLService_AddRouteURL2(t *testing.T) {
	var r GatewayRouteURLService
	r.ClientID = testClientStr
	r.Host = "http://localhost:3011"
	r.Token = tempToken
	r.APIKey = testAPIKey

	var rr GatewayRouteURL
	//rr.ClientID = GwCid3
	rr.RouteID = RTID2
	rr.Name = "blue"
	rr.URL = "http://www.google.com/test"
	res := r.AddRouteURL(&rr)
	RUID2 = res.ID
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayRouteURLService_ActivateRouteURL(t *testing.T) {
	var r GatewayRouteURLService
	r.ClientID = testClientStr
	r.Host = "http://localhost:3011"
	r.Token = tempToken
	r.APIKey = testAPIKey

	var rr GatewayRouteURL
	rr.ID = RUID2
	//rr.ClientID = GwCid3
	rr.RouteID = RTID2

	res := r.ActivateRouteURL(&rr)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayRouteURLService_GetRouteURL2(t *testing.T) {
	var r GatewayRouteURLService
	r.ClientID = testClientStr
	r.Host = "http://localhost:3011"
	r.Token = tempToken
	r.APIKey = testAPIKey

	res := r.GetRouteURL(strconv.FormatInt(RUID2, 10), strconv.FormatInt(RTID2, 10))
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Name != "blue" || res.URL != "http://www.google.com/test" || res.Active != true {
		t.Fail()
	}
}

func TestGatewayRouteURLService_ActivateRouteURL2(t *testing.T) {
	var r GatewayRouteURLService
	r.ClientID = testClientStr
	r.Host = "http://localhost:3011"
	r.Token = tempToken
	r.APIKey = testAPIKey

	var rr GatewayRouteURL
	rr.ID = RUID
	//rr.ClientID = GwCid3
	rr.RouteID = RTID2

	res := r.ActivateRouteURL(&rr)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayRouteURLService_GetRouteURL3(t *testing.T) {
	var r GatewayRouteURLService
	r.ClientID = testClientStr
	r.Host = "http://localhost:3011"
	r.Token = tempToken
	r.APIKey = testAPIKey

	res := r.GetRouteURL(strconv.FormatInt(RUID2, 10), strconv.FormatInt(RTID2, 10))
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Name != "blue" || res.URL != "http://www.google.com/test" || res.Active != false {
		t.Fail()
	}
}

func TestGatewayRouteURLService_GetRouteURLList(t *testing.T) {
	var r GatewayRouteURLService
	r.ClientID = testClientStr
	r.Host = "http://localhost:3011"
	r.Token = tempToken
	r.APIKey = testAPIKey

	res := r.GetRouteURLList(strconv.FormatInt(RTID2, 10))
	fmt.Print("res: ")
	fmt.Println(res)
	if len(*res) < 2 {
		t.Fail()
	}
}

func TestGatewayRouteURLService_DeleteRouteURL(t *testing.T) {
	var r GatewayRouteURLService
	r.ClientID = testClientStr
	r.Host = "http://localhost:3011"
	r.Token = tempToken
	r.APIKey = testAPIKey

	res := r.DeleteRouteURL(strconv.FormatInt(RUID, 10), strconv.FormatInt(RTID2, 10))
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != false {
		t.Fail()
	}
}

func TestGatewayRouteURLService_DeleteRouteURL2(t *testing.T) {
	var r GatewayRouteURLService
	r.ClientID = testClientStr
	r.Host = "http://localhost:3011"
	r.Token = tempToken
	r.APIKey = testAPIKey

	res := r.DeleteRouteURL(strconv.FormatInt(RUID2, 10), strconv.FormatInt(RTID2, 10))
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayRouteURLService_GetRouteURLList2(t *testing.T) {
	var r GatewayRouteURLService
	r.ClientID = testClientStr
	r.Host = "http://localhost:3011"
	r.Token = tempToken
	r.APIKey = testAPIKey

	res := r.GetRouteURLList(strconv.FormatInt(RTID2, 10))
	fmt.Print("res: ")
	fmt.Println(res)
	if len(*res) != 1 {
		t.Fail()
	}
}

func TestGatewayRouteURLService_DeleteRoute(t *testing.T) {
	var r GatewayRouteService
	r.ClientID = testClientStr
	r.Host = "http://localhost:3011"
	r.Token = tempToken
	r.APIKey = testAPIKey

	res := r.DeleteRoute(strconv.FormatInt(RTID2, 10))
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}
