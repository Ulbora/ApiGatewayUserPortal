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

var rID int64

func TestClientRoleService_AddClientRole(t *testing.T) {
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
	rID = res.ID
	if res.Success != true {
		t.Fail()
	}
}

func TestClientRoleService_GetClientRoleList(t *testing.T) {
	var c ClientRoleService
	c.ClientID = testClientStr
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	res := c.GetClientRoleList()
	fmt.Print("client role res list: ")
	fmt.Println(res)
	fmt.Print("len: ")
	fmt.Println(len(*res))

	if res == nil || len(*res) == 0 {
		t.Fail()
	}
}

func TestClientRoleService_DeleteClientRole(t *testing.T) {
	var c ClientRoleService
	c.ClientID = testClientStr
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	res := c.DeleteClientRole(strconv.FormatInt(rID, 10))
	fmt.Print("res deleted client role: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}
