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
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//GatewayErrorsService service
type GatewayErrorsService struct {
	Token    string
	ClientID string
	APIKey   string
	UserID   string
	Hashed   string
	Host     string
}

//GatewayError GatewayError
type GatewayError struct {
	ID          int64     `json:"id"`
	Code        int       `json:"code"`
	Message     string    `json:"message"`
	Entered     time.Time `json:"entered"`
	RouteURIID  int64     `json:"routeUriId"`
	RestRouteID int64     `json:"routeId"`
	ClientID    int64     `json:"clientId"`
}

//GetRouteErrors GetRouteErrors
func (e *GatewayErrorsService) GetRouteErrors(ee *GatewayError) *[]GatewayError {
	var rtn = make([]GatewayError, 0)
	var pURL = e.Host + "/rs/gwErrors"
	aJSON, err := json.Marshal(ee)

	if err != nil {
		fmt.Println(err)
	} else {
		req, rErr := http.NewRequest("POST", pURL, bytes.NewBuffer(aJSON))
		if rErr != nil {
			fmt.Print("request err: ")
			fmt.Println(rErr)
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+e.Token)
			req.Header.Set("clientId", e.ClientID)
			//req.Header.Set("userId", c.UserID)
			//req.Header.Set("hashed", c.Hashed)
			req.Header.Set("apiKey", e.APIKey)
			client := &http.Client{}
			resp, cErr := client.Do(req)
			if cErr != nil {
				fmt.Print("performance err: ")
				fmt.Println(cErr)
			} else {
				defer resp.Body.Close()
				//fmt.Print("resp: ")
				//fmt.Println(resp)
				decoder := json.NewDecoder(resp.Body)
				error := decoder.Decode(&rtn)
				if error != nil {
					log.Println(error.Error())
				}
			}
		}
	}
	return &rtn
}
