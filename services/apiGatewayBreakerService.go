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

//GatewayBreakerService service
type GatewayBreakerService struct {
	Token    string
	ClientID string
	APIKey   string
	UserID   string
	Hashed   string
	Host     string
}

//GatewayBreaker GatewayClient
type GatewayBreaker struct {
	ID                     int64     `json:"id"`
	FailureThreshold       int       `json:"failureThreshold"`
	FailureCount           int       `json:"failureCount"`
	LastFailureTime        time.Time `json:"lastFailureTime"`
	HealthCheckTimeSeconds int       `json:"healthCheckTimeSeconds"`
	FailoverRouteName      string    `json:"failoverRouteName"`
	OpenFailCode           int       `json:"openFailCode"`
	RouteURIID             int64     `json:"routeUriId"`
	RestRouteID            int64     `json:"routeId"`
	ClientID               int64     `json:"clientId"`
}

//Status status
type Status struct {
	Warning           bool   `json:"warning"`
	Open              bool   `json:"open"`
	PartialOpen       bool   `json:"partialOpen"`
	FailoverRouteName string `json:"failoverRouteName"`
	OpenFailCode      int    `json:"openFailCode"`
}

//InsertBreaker add
func (r *GatewayBreakerService) InsertBreaker(bk *GatewayBreaker) *GatewayResponse {
	var rtn = new(GatewayResponse)
	var addURL = r.Host + "/rs/gwBreaker/add"
	aJSON, err := json.Marshal(bk)

	if err != nil {
		fmt.Println(err)
	} else {
		req, rErr := http.NewRequest("POST", addURL, bytes.NewBuffer(aJSON))
		if rErr != nil {
			fmt.Print("request err: ")
			fmt.Println(rErr)
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+r.Token)
			req.Header.Set("clientId", r.ClientID)
			//req.Header.Set("userId", c.UserID)
			//req.Header.Set("hashed", c.Hashed)
			req.Header.Set("apiKey", r.APIKey)
			client := &http.Client{}
			resp, cErr := client.Do(req)
			if cErr != nil {
				fmt.Print("Breaker Add err: ")
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
				rtn.Code = resp.StatusCode
			}
		}
	}
	return rtn
}

//UpdateBreaker update
func (r *GatewayBreakerService) UpdateBreaker(bk *GatewayBreaker) *GatewayResponse {
	var rtn = new(GatewayResponse)
	var upURL = r.Host + "/rs/gwBreaker/update"

	//fmt.Println(content.Text)
	aJSON, err := json.Marshal(bk)
	if err != nil {
		fmt.Println(err)
	} else {
		req, rErr := http.NewRequest("PUT", upURL, bytes.NewBuffer(aJSON))
		if rErr != nil {
			fmt.Print("request err: ")
			fmt.Println(rErr)
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+r.Token)
			req.Header.Set("clientId", r.ClientID)
			//req.Header.Set("userId", c.UserID)
			//req.Header.Set("hashed", c.Hashed)
			req.Header.Set("apiKey", r.APIKey)
			client := &http.Client{}
			resp, cErr := client.Do(req)
			if cErr != nil {
				fmt.Print("Breaker Update err: ")
				fmt.Println(cErr)
			} else {
				defer resp.Body.Close()
				decoder := json.NewDecoder(resp.Body)
				error := decoder.Decode(&rtn)
				if error != nil {
					log.Println(error.Error())
				}
				rtn.Code = resp.StatusCode
			}
		}
	}
	return rtn
}

//ResetBreaker update
func (r *GatewayBreakerService) ResetBreaker(bk *GatewayBreaker) *GatewayResponse {
	var rtn = new(GatewayResponse)
	var upURL = r.Host + "/rs/gwBreaker/reset"

	//fmt.Println(content.Text)
	aJSON, err := json.Marshal(bk)
	if err != nil {
		fmt.Println(err)
	} else {
		req, rErr := http.NewRequest("POST", upURL, bytes.NewBuffer(aJSON))
		if rErr != nil {
			fmt.Print("request err: ")
			fmt.Println(rErr)
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+r.Token)
			req.Header.Set("clientId", r.ClientID)
			//req.Header.Set("userId", c.UserID)
			//req.Header.Set("hashed", c.Hashed)
			req.Header.Set("apiKey", r.APIKey)
			client := &http.Client{}
			resp, cErr := client.Do(req)
			if cErr != nil {
				fmt.Print("Breaker reset err: ")
				fmt.Println(cErr)
			} else {
				fmt.Print("reset resp: ")
				fmt.Println(resp)
				defer resp.Body.Close()
				decoder := json.NewDecoder(resp.Body)
				error := decoder.Decode(&rtn)
				if error != nil {
					log.Println(error.Error())
				}
				rtn.Code = resp.StatusCode
			}
		}
	}
	return rtn
}

// GetBreaker get
func (r *GatewayBreakerService) GetBreaker(urlID string, routeID string) *GatewayBreaker {
	var rtn = new(GatewayBreaker)
	var gURL = r.Host + "/rs/gwBreaker/get/" + urlID + "/" + routeID
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("GET", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("clientId", r.ClientID)
		req.Header.Set("Authorization", "Bearer "+r.Token)
		req.Header.Set("apiKey", r.APIKey)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("Breaker read err: ")
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

// GetBreakerStatus get
func (r *GatewayBreakerService) GetBreakerStatus(routeURLID string) *Status {
	var rtn = new(Status)
	var gURL = r.Host + "/rs/gwBreaker/status/" + routeURLID
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("GET", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("clientId", r.ClientID)
		req.Header.Set("Authorization", "Bearer "+r.Token)
		req.Header.Set("apiKey", r.APIKey)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("Breaker status read err: ")
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

// // GetRouteURLList get
// func (r *GatewayRouteURLService) GetRouteURLList(routeID string, clientID string) *[]GatewayRouteURL {
// 	var rtn = make([]GatewayRouteURL, 0)
// 	var gURL = r.Host + "/rs/gwRouteUrlSuper/list/" + routeID + "/" + clientID
// 	//fmt.Println(gURL)
// 	req, rErr := http.NewRequest("GET", gURL, nil)
// 	if rErr != nil {
// 		fmt.Print("request err: ")
// 		fmt.Println(rErr)
// 	} else {
// 		req.Header.Set("clientId", r.ClientID)
// 		req.Header.Set("Authorization", "Bearer "+r.Token)
// 		req.Header.Set("apiKey", r.APIKey)
// 		client := &http.Client{}
// 		resp, cErr := client.Do(req)
// 		if cErr != nil {
// 			fmt.Print("route list Service read err: ")
// 			fmt.Println(cErr)
// 		} else {
// 			defer resp.Body.Close()
// 			decoder := json.NewDecoder(resp.Body)
// 			error := decoder.Decode(&rtn)
// 			if error != nil {
// 				log.Println(error.Error())
// 			}
// 		}
// 	}
// 	return &rtn
// }

// DeleteBreaker delete
func (r *GatewayBreakerService) DeleteBreaker(urlID string, routeID string) *GatewayResponse {
	var rtn = new(GatewayResponse)
	var gURL = r.Host + "/rs/gwBreaker/delete/" + urlID + "/" + routeID
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("DELETE", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("Authorization", "Bearer "+r.Token)
		req.Header.Set("clientId", r.ClientID)
		//req.Header.Set("userId", r.UserID)
		//req.Header.Set("hashed", r.Hashed)
		req.Header.Set("apiKey", r.APIKey)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("Breaker delete err: ")
			fmt.Println(cErr)
		} else {
			defer resp.Body.Close()
			decoder := json.NewDecoder(resp.Body)
			error := decoder.Decode(&rtn)
			if error != nil {
				log.Println(error.Error())
			}
			rtn.Code = resp.StatusCode
		}
	}
	return rtn
}
