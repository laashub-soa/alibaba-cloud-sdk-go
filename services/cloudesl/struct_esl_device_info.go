package cloudesl

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// EslDeviceInfo is a nested struct in cloudesl response
type EslDeviceInfo struct {
	CompanyId           string `json:"CompanyId" xml:"CompanyId"`
	StoreId             string `json:"StoreId" xml:"StoreId"`
	Mac                 string `json:"Mac" xml:"Mac"`
	EslBarCode          string `json:"EslBarCode" xml:"EslBarCode"`
	ItemBarCode         string `json:"ItemBarCode" xml:"ItemBarCode"`
	Vendor              string `json:"Vendor" xml:"Vendor"`
	ConnectAp           string `json:"ConnectAp" xml:"ConnectAp"`
	Type                string `json:"Type" xml:"Type"`
	Model               string `json:"Model" xml:"Model"`
	BeBind              string `json:"BeBind" xml:"BeBind"`
	ScreenWidth         string `json:"ScreenWidth" xml:"ScreenWidth"`
	ScreenHeight        string `json:"ScreenHeight" xml:"ScreenHeight"`
	LastCommunicateTime string `json:"LastCommunicateTime" xml:"LastCommunicateTime"`
	BatteryLevel        int    `json:"BatteryLevel" xml:"BatteryLevel"`
	EslStatus           string `json:"EslStatus" xml:"EslStatus"`
	ShelfCode           string `json:"ShelfCode" xml:"ShelfCode"`
	PositionCode        string `json:"PositionCode" xml:"PositionCode"`
	ItemId              int    `json:"ItemId" xml:"ItemId"`
	ItemTitle           string `json:"ItemTitle" xml:"ItemTitle"`
	ItemActionPrice     int    `json:"ItemActionPrice" xml:"ItemActionPrice"`
	ItemPriceUnit       string `json:"ItemPriceUnit" xml:"ItemPriceUnit"`
}