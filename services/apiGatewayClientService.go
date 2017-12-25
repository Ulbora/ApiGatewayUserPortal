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
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//GatewayClientService service
type GatewayClientService struct {
	Token    string
	ClientID string
	APIKey   string
	UserID   string
	Hashed   string
	Host     string
}

//GatewayClient GatewayClient
type GatewayClient struct {
	ClientID int64  `json:"clientId"`
	APIKey   string `json:"apiKey"`
	Enabled  bool   `json:"enabled"`
	Level    string `json:"level"`
}

//GatewayResponse resp
type GatewayResponse struct {
	Success bool  `json:"success"`
	ID      int64 `json:"id"`
	Code    int   `json:"code"`
}

// GetClient get GetClient
func (c *GatewayClientService) GetClient() *GatewayClient {
	var rtn = new(GatewayClient)
	var gURL = c.Host + "/rs/gwClientUser/get"
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("GET", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("clientId", c.ClientID)
		req.Header.Set("Authorization", "Bearer "+c.Token)
		req.Header.Set("apiKey", c.APIKey)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		fmt.Print("resp: ")
		fmt.Println(resp)
		if cErr != nil {
			fmt.Print("Client Service read err: ")
			fmt.Println(cErr)
		} else {
			defer resp.Body.Close()
			decoder := json.NewDecoder(resp.Body)
			error := decoder.Decode(&rtn)
			if error != nil {
				log.Println(error.Error())
			}
		}
	}
	return rtn
}
