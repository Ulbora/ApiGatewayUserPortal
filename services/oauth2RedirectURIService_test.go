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

var rdID int64

func TestRedirectURIService_AddRedirectURI(t *testing.T) {
	var c RedirectURIService
	c.ClientID = testClientStr
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	var uri RedirectURI
	uri.URI = "http://yahoooo.com"
	uri.ClientID = testClient
	res := c.AddRedirectURI(&uri)

	fmt.Print("add uri res: ")
	fmt.Println(res)
	rdID = res.ID
	if res.Success != true {
		t.Fail()
	}
}

func TestRedirectURIService_GetRedirectURIList(t *testing.T) {
	var c RedirectURIService
	c.ClientID = testClientStr
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	res := c.GetRedirectURIList()
	fmt.Print("uri res list: ")
	fmt.Println(res)
	fmt.Print("len: ")
	fmt.Println(len(*res))
	if res == nil || len(*res) == 0 {
		t.Fail()
	}
}

func TestRedirectURIService_DeleteRedirectURI(t *testing.T) {
	var c RedirectURIService
	c.ClientID = testClientStr
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	res := c.DeleteRedirectURI(strconv.FormatInt(rdID, 10))
	fmt.Print("res deleted uri: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}
