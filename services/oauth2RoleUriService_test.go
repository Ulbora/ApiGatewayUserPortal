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

var rID2 int64
var uID2 int64

// func TestRoleURIService_AddClient(t *testing.T) {
// 	var c ClientService
// 	c.ClientID = "403"
// 	c.Host = "http://localhost:3000"
// 	c.Token = tempToken
// 	var uri RedirectURI
// 	uri.URI = "http://googole.com"
// 	var uris []RedirectURI
// 	uris = append(uris, uri)
// 	var cc Client
// 	cc.Email = "ken@ken.com"
// 	cc.Enabled = true
// 	cc.Name = "A Big Company"
// 	cc.RedirectURIs = uris
// 	res := c.AddClient(&cc)
// 	fmt.Print("add client res: ")
// 	fmt.Println(res)
// 	CID7 = res.ClientID
// 	if res.Success != true {
// 		t.Fail()
// 	}
// }

func TestRoleURIService_AddAllowedURI(t *testing.T) {
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
	uID2 = res.ID
	if res.Success != true {
		t.Fail()
	}
}

func TestRoleURIService_AddClientRole(t *testing.T) {
	var c ClientRoleService
	c.ClientID = testClientStr
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	var cr ClientRole
	cr.Role = "user"
	cr.ClientID = testClient
	res := c.AddClientRole(&cr)

	fmt.Print("add client role res: ")
	fmt.Println(res)
	rID2 = res.ID
	if res.Success != true {
		t.Fail()
	}
}

func TestRoleURIService_AddRoleURI(t *testing.T) {
	var c RoleURIService
	c.ClientID = testClientStr
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	var ru RoleURI
	ru.ClientAllowedURIID = uID2
	ru.ClientRoleID = rID2
	res := c.AddRoleURI(&ru)

	fmt.Print("add roleUri res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestRoleURIService_GetRoleURIList(t *testing.T) {
	var c RoleURIService
	c.ClientID = testClientStr
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	res := c.GetRoleURIList(strconv.FormatInt(rID2, 10))
	fmt.Print("roleuri res list: ")
	fmt.Println(res)
	fmt.Print("len: ")
	fmt.Println(len(*res))
	if res == nil || len(*res) != 1 {
		t.Fail()
	}
}

func TestRoleURIService_DeleteRoleURI(t *testing.T) {
	var c RoleURIService
	c.ClientID = testClientStr
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	var ru RoleURI
	ru.ClientAllowedURIID = uID2
	ru.ClientRoleID = rID2
	res := c.DeleteRoleURI(&ru)
	fmt.Print("res deleted roleuri: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestRoleURIService_DeleteAllowedURI(t *testing.T) {
	var c AllowedURIService
	c.ClientID = testClientStr
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	res := c.DeleteAllowedURI(strconv.FormatInt(uID2, 10))
	fmt.Print("res deleted uri: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestRoleURIService_DeleteClientRole(t *testing.T) {
	var c ClientRoleService
	c.ClientID = testClientStr
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	res := c.DeleteClientRole(strconv.FormatInt(rID2, 10))
	fmt.Print("res deleted client role: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

// func TestRoleURIService_DeleteClient(t *testing.T) {
// 	var c ClientService
// 	c.ClientID = "403"
// 	c.Host = "http://localhost:3000"
// 	c.Token = tempToken
// 	res := c.DeleteClient(strconv.FormatInt(CID7, 10))
// 	fmt.Print("res deleted client: ")
// 	fmt.Println(res)
// 	if res.Success != true {
// 		t.Fail()
// 	}
// }
