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

var aID int64

func TestAllowedURIService_AddAllowedURI(t *testing.T) {
	var c AllowedURIService
	c.ClientID = testClientStr
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	var uri AllowedURI
	uri.URI = "/rs/mail/send"
	uri.ClientID = testClient
	res := c.AddAllowedURI(&uri)

	fmt.Print("add uri res: ")
	fmt.Println(res)
	aID = res.ID
	if res.Success != true {
		t.Fail()
	}
}

func TestAllowedURIService_UpdateAllowedURI(t *testing.T) {
	var c AllowedURIService
	c.ClientID = testClientStr
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	var uri AllowedURI
	uri.ID = aID
	uri.URI = "/rs/mail/get"
	uri.ClientID = testClient
	res := c.UpdateAllowedURI(&uri)
	fmt.Print("update uri res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestAllowedURIService_GetAllowedURI(t *testing.T) {
	var c AllowedURIService
	c.ClientID = testClientStr
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	res := c.GetAllowedURI(strconv.FormatInt(aID, 10))
	fmt.Print("get uri res: ")
	fmt.Println(res)
	if res.URI != "/rs/mail/get" {
		t.Fail()
	}
}

func TestAllowedURIService_GetAllowedURIList(t *testing.T) {
	var c AllowedURIService
	c.ClientID = testClientStr
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	res := c.GetAllowedURIList()
	fmt.Print("uri res list: ")
	fmt.Println(res)
	fmt.Print("len: ")
	fmt.Println(len(*res))
	if res == nil || len(*res) == 0 {
		t.Fail()
	}
}

func TestAllowedURIService_DeleteAllowedURI(t *testing.T) {
	var c AllowedURIService
	c.ClientID = testClientStr
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	res := c.DeleteAllowedURI(strconv.FormatInt(aID, 10))
	fmt.Print("res deleted uri: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}
