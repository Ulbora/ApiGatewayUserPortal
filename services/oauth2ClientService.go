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
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//ClientService service
type ClientService struct {
	Token    string
	ClientID string
	APIKey   string
	UserID   string
	Hashed   string
	Host     string
}

//Client client
type Client struct {
	ClientID     int64         `json:"clientId"`
	Secret       string        `json:"secret"`
	Name         string        `json:"name"`
	WebSite      string        `json:"webSite"`
	Email        string        `json:"email"`
	Enabled      bool          `json:"enabled"`
	RedirectURIs []RedirectURI `json:"redirectUrls"`
}

//ClientResponse resp
type ClientResponse struct {
	Success  bool  `json:"success"`
	ClientID int64 `json:"clientId"`
	Code     int   `json:"code"`
}

// //AddClient add template
// func (c *ClientService) AddClient(client *Client) *ClientResponse {
// 	var rtn = new(ClientResponse)
// 	var addURL = c.Host + "/rs/client/add"
// 	aJSON, err := json.Marshal(client)

// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		req, rErr := http.NewRequest("POST", addURL, bytes.NewBuffer(aJSON))
// 		if rErr != nil {
// 			fmt.Print("request err: ")
// 			fmt.Println(rErr)
// 		} else {
// 			req.Header.Set("Content-Type", "application/json")
// 			req.Header.Set("Authorization", "Bearer "+c.Token)
// 			req.Header.Set("clientId", c.ClientID)
// 			//req.Header.Set("userId", c.UserID)
// 			//req.Header.Set("hashed", c.Hashed)
// 			req.Header.Set("apiKey", c.APIKey)
// 			client := &http.Client{}
// 			resp, cErr := client.Do(req)
// 			if cErr != nil {
// 				fmt.Print("Client Add err: ")
// 				fmt.Println(cErr)
// 			} else {
// 				defer resp.Body.Close()
// 				//fmt.Print("resp: ")
// 				//fmt.Println(resp)
// 				decoder := json.NewDecoder(resp.Body)
// 				error := decoder.Decode(&rtn)
// 				if error != nil {
// 					log.Println(error.Error())
// 				}
// 				rtn.Code = resp.StatusCode
// 			}
// 		}
// 	}
// 	return rtn
// }

// //UpdateClient update UpdateClient
// func (c *ClientService) UpdateClient(client *Client) *ClientResponse {
// 	var rtn = new(ClientResponse)
// 	var upURL = c.Host + "/rs/client/update"

// 	//fmt.Println(content.Text)
// 	aJSON, err := json.Marshal(client)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		req, rErr := http.NewRequest("PUT", upURL, bytes.NewBuffer(aJSON))
// 		if rErr != nil {
// 			fmt.Print("request err: ")
// 			fmt.Println(rErr)
// 		} else {
// 			req.Header.Set("Content-Type", "application/json")
// 			req.Header.Set("Authorization", "Bearer "+c.Token)
// 			req.Header.Set("clientId", c.ClientID)
// 			//req.Header.Set("userId", c.UserID)
// 			//req.Header.Set("hashed", c.Hashed)
// 			req.Header.Set("apiKey", c.APIKey)
// 			client := &http.Client{}
// 			resp, cErr := client.Do(req)
// 			if cErr != nil {
// 				fmt.Print("Client Service Update err: ")
// 				fmt.Println(cErr)
// 			} else {
// 				defer resp.Body.Close()
// 				decoder := json.NewDecoder(resp.Body)
// 				error := decoder.Decode(&rtn)
// 				if error != nil {
// 					log.Println(error.Error())
// 				}
// 				rtn.Code = resp.StatusCode
// 			}
// 		}
// 	}
// 	return rtn
// }

// GetClient get GetClient
func (c *ClientService) GetClient() *Client {
	var rtn = new(Client)
	var gURL = c.Host + "/rs/client/admin/get"
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
		fmt.Print("Client Service read: ")
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

// // GetClientList get client list
// func (c *ClientService) GetClientList() *[]Client {
// 	var rtn = make([]Client, 0)
// 	var gURL = c.Host + "/rs/client/list"
// 	//fmt.Println(gURL)
// 	req, rErr := http.NewRequest("GET", gURL, nil)
// 	if rErr != nil {
// 		fmt.Print("request err: ")
// 		fmt.Println(rErr)
// 	} else {
// 		req.Header.Set("clientId", c.ClientID)
// 		req.Header.Set("Authorization", "Bearer "+c.Token)
// 		req.Header.Set("apiKey", c.APIKey)
// 		client := &http.Client{}
// 		resp, cErr := client.Do(req)
// 		if cErr != nil {
// 			fmt.Print("client list Service read err: ")
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

// //SearchClient SearchClient
// func (c *ClientService) SearchClient(client *Client) *[]Client {
// 	var rtn = make([]Client, 0)
// 	var addURL = c.Host + "/rs/client/search"
// 	aJSON, err := json.Marshal(client)

// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		req, rErr := http.NewRequest("POST", addURL, bytes.NewBuffer(aJSON))
// 		if rErr != nil {
// 			fmt.Print("request err: ")
// 			fmt.Println(rErr)
// 		} else {
// 			req.Header.Set("Content-Type", "application/json")
// 			req.Header.Set("Authorization", "Bearer "+c.Token)
// 			req.Header.Set("clientId", c.ClientID)
// 			//req.Header.Set("userId", c.UserID)
// 			//req.Header.Set("hashed", c.Hashed)
// 			req.Header.Set("apiKey", c.APIKey)
// 			client := &http.Client{}
// 			resp, cErr := client.Do(req)
// 			if cErr != nil {
// 				fmt.Print("Client search err: ")
// 				fmt.Println(cErr)
// 			} else {
// 				defer resp.Body.Close()
// 				//fmt.Print("search resp: ")
// 				//fmt.Println(resp)
// 				decoder := json.NewDecoder(resp.Body)
// 				error := decoder.Decode(&rtn)
// 				if error != nil {
// 					log.Println(error.Error())
// 				}
// 				//rtn.Code = resp.StatusCode
// 			}
// 		}
// 	}
// 	return &rtn
// }

// // DeleteClient delete DeleteClient
// func (c *ClientService) DeleteClient(id string) *ClientResponse {
// 	var rtn = new(ClientResponse)
// 	var gURL = c.Host + "/rs/client/delete/" + id
// 	//fmt.Println(gURL)
// 	req, rErr := http.NewRequest("DELETE", gURL, nil)
// 	if rErr != nil {
// 		fmt.Print("request err: ")
// 		fmt.Println(rErr)
// 	} else {
// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", "Bearer "+c.Token)
// 		req.Header.Set("clientId", c.ClientID)
// 		//req.Header.Set("userId", r.UserID)
// 		//req.Header.Set("hashed", r.Hashed)
// 		req.Header.Set("apiKey", c.APIKey)
// 		client := &http.Client{}
// 		resp, cErr := client.Do(req)
// 		if cErr != nil {
// 			fmt.Print("redirect uri Service delete err: ")
// 			fmt.Println(cErr)
// 		} else {
// 			defer resp.Body.Close()
// 			decoder := json.NewDecoder(resp.Body)
// 			error := decoder.Decode(&rtn)
// 			if error != nil {
// 				log.Println(error.Error())
// 			}
// 			rtn.Code = resp.StatusCode
// 		}
// 	}
// 	return rtn
// }
