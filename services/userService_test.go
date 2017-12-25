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
	"testing"
)

var UID = "bob123456789"

func TestUserService_AddUser(t *testing.T) {
	var u UserService
	u.ClientID = testClientStr
	u.Host = "http://localhost:3001"
	u.Token = tempToken
	var user User
	user.ClientID = testClientStr
	user.EmailAddress = "bob@bob.com"
	user.Enabled = true
	user.FirstName = "bob"
	user.LastName = "bob"
	user.Password = "bob"
	user.RoleID = 1
	user.Username = UID

	res := u.AddUser(&user)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestUserService_UpdateUserPassword(t *testing.T) {
	var u UserService
	u.ClientID = testClientStr
	u.Host = "http://localhost:3001"
	u.Token = tempToken
	var user UserPW
	user.Username = UID
	user.ClientID = testClientStr
	user.Password = "bobbby"

	res := u.UpdateUser(&user)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestUserService_UpdateUserDisable(t *testing.T) {
	var u UserService
	u.ClientID = testClientStr
	u.Host = "http://localhost:3001"
	u.Token = tempToken
	var user UserDis
	user.Username = UID
	user.ClientID = testClientStr
	user.Enabled = false

	res := u.UpdateUser(&user)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestUserService_UpdateUserInfo(t *testing.T) {
	var u UserService
	u.ClientID = testClientStr
	u.Host = "http://localhost:3001"
	u.Token = tempToken
	var user UserInfo
	user.Username = UID
	user.ClientID = testClientStr
	user.EmailAddress = "bobbby@bob.com"
	user.FirstName = "bobby"
	user.RoleID = 1
	user.LastName = "williams"

	res := u.UpdateUser(&user)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestUserService_UpdateUserDisable2(t *testing.T) {
	var u UserService
	u.ClientID = testClientStr
	u.Host = "http://localhost:3001"
	u.Token = tempToken
	var user UserDis
	user.Username = UID
	user.ClientID = testClientStr
	user.Enabled = true

	res := u.UpdateUser(&user)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestUserService_GetUser(t *testing.T) {
	var u UserService
	u.ClientID = testClientStr
	u.Host = "http://localhost:3001"
	u.Token = tempToken

	res := u.GetUser(UID)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Username != UID || res.Enabled == false {
		t.Fail()
	}
}

// func TestUserService_GetUserList(t *testing.T) {
// 	var u UserService
// 	u.ClientID = testClientStr
// 	u.Host = "http://localhost:3001"
// 	u.Token = tempToken

// 	res := u.GetUserList()
// 	fmt.Print("user list res: ")
// 	fmt.Println(res)
// 	if len(*res) == 0 {
// 		t.Fail()
// 	}
// }

func TestUserService_SearchUserList(t *testing.T) {
	var u UserService
	u.ClientID = testClientStr
	u.Host = "http://localhost:3001"
	u.Token = tempToken

	res := u.SearchUserList()
	fmt.Print("res: ")
	fmt.Println(res)
	if len(*res) == 0 {
		t.Fail()
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	var u UserService
	u.ClientID = testClientStr
	u.Host = "http://localhost:3001"
	u.Token = tempToken

	res := u.DeleteUser(UID)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestUserService_GetRoleList(t *testing.T) {
	var u UserService
	u.ClientID = testClientStr
	u.Host = "http://localhost:3001"
	u.Token = tempToken

	res := u.GetRoleList()
	fmt.Print("role res: ")
	fmt.Println(res)
	if len(*res) == 0 {
		t.Fail()
	}
}
