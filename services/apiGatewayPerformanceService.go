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
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//GatewayPerformanceService service
type GatewayPerformanceService struct {
	Token    string
	ClientID string
	APIKey   string
	UserID   string
	Hashed   string
	Host     string
}

//GatewayPerformance GatewayPerformance
type GatewayPerformance struct {
	ID             int64     `json:"id"`
	Calls          int64     `json:"calls"`
	LatencyMsTotal int64     `json:"latencyMsTotal"`
	Entered        time.Time `json:"entered"`
	RouteURIID     int64     `json:"routeUriId"`
	RestRouteID    int64     `json:"routeId"`
	ClientID       int64     `json:"clientId"`
}

//GetRoutePerformance GetRoutePerformance
func (p *GatewayPerformanceService) GetRoutePerformance(pp *GatewayPerformance) *[]GatewayPerformance {
	var rtn = make([]GatewayPerformance, 0)
	var pURL = p.Host + "/rs/gwPerformance"
	aJSON, err := json.Marshal(pp)

	if err != nil {
		fmt.Println(err)
	} else {
		req, rErr := http.NewRequest("POST", pURL, bytes.NewBuffer(aJSON))
		if rErr != nil {
			fmt.Print("request err: ")
			fmt.Println(rErr)
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+p.Token)
			req.Header.Set("clientId", p.ClientID)
			//req.Header.Set("userId", c.UserID)
			//req.Header.Set("hashed", c.Hashed)
			req.Header.Set("apiKey", p.APIKey)
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
