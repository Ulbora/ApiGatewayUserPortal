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

var gtID int64

func TestGrantTypeService_AddGrantType(t *testing.T) {
	var c GrantTypeService
	c.ClientID = testClientStr
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	var gt GrantType
	gt.GrantType = "code"
	gt.ClientID = testClient
	res := c.AddGrantType(&gt)

	fmt.Print("add grant type res: ")
	fmt.Println(res)
	gtID = res.ID
	if res.Success != true {
		t.Fail()
	}
}

func TestGrantTypeService_GetGrantTypeList(t *testing.T) {
	var c GrantTypeService
	c.ClientID = testClientStr
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	res := c.GetGrantTypeList()
	fmt.Print("grant type res list: ")
	fmt.Println(res)
	fmt.Print("len: ")
	fmt.Println(len(*res))

	if res == nil || len(*res) == 0 {
		t.Fail()
	}
}

func TestGrantTypeService_DeleteGrantType(t *testing.T) {
	var c GrantTypeService
	c.ClientID = testClientStr
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	res := c.DeleteGrantType(strconv.FormatInt(gtID, 10))
	fmt.Print("res deleted uri: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}
