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
	"testing"
)

func TestGatewayClientService_GetClient(t *testing.T) {
	var c GatewayClientService
	c.ClientID = testClientStr
	c.Host = "http://localhost:3011"
	c.Token = tempToken
	c.APIKey = testAPIKey

	res := c.GetClient()
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Enabled == false || res.ClientID != testClient {
		t.Fail()
	}
}
