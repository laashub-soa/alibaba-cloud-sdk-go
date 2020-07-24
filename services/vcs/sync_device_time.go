package vcs

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// SyncDeviceTime invokes the vcs.SyncDeviceTime API synchronously
// api document: https://help.aliyun.com/api/vcs/syncdevicetime.html
func (client *Client) SyncDeviceTime(request *SyncDeviceTimeRequest) (response *SyncDeviceTimeResponse, err error) {
	response = CreateSyncDeviceTimeResponse()
	err = client.DoAction(request, response)
	return
}

// SyncDeviceTimeWithChan invokes the vcs.SyncDeviceTime API asynchronously
// api document: https://help.aliyun.com/api/vcs/syncdevicetime.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SyncDeviceTimeWithChan(request *SyncDeviceTimeRequest) (<-chan *SyncDeviceTimeResponse, <-chan error) {
	responseChan := make(chan *SyncDeviceTimeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SyncDeviceTime(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// SyncDeviceTimeWithCallback invokes the vcs.SyncDeviceTime API asynchronously
// api document: https://help.aliyun.com/api/vcs/syncdevicetime.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SyncDeviceTimeWithCallback(request *SyncDeviceTimeRequest, callback func(response *SyncDeviceTimeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SyncDeviceTimeResponse
		var err error
		defer close(result)
		response, err = client.SyncDeviceTime(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// SyncDeviceTimeRequest is the request struct for api SyncDeviceTime
type SyncDeviceTimeRequest struct {
	*requests.RpcRequest
	DeviceTimeStamp string `position:"Body" name:"DeviceTimeStamp"`
	DeviceSn        string `position:"Body" name:"DeviceSn"`
}

// SyncDeviceTimeResponse is the response struct for api SyncDeviceTime
type SyncDeviceTimeResponse struct {
	*responses.BaseResponse
	Code          string `json:"Code" xml:"Code"`
	Message       string `json:"Message" xml:"Message"`
	NTPServer     string `json:"NTPServer" xml:"NTPServer"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
	RetryInterval string `json:"RetryInterval" xml:"RetryInterval"`
	SyncInterval  string `json:"SyncInterval" xml:"SyncInterval"`
	TimeStamp     string `json:"TimeStamp" xml:"TimeStamp"`
}

// CreateSyncDeviceTimeRequest creates a request to invoke SyncDeviceTime API
func CreateSyncDeviceTimeRequest() (request *SyncDeviceTimeRequest) {
	request = &SyncDeviceTimeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "SyncDeviceTime", "vcs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSyncDeviceTimeResponse creates a response to parse from SyncDeviceTime response
func CreateSyncDeviceTimeResponse() (response *SyncDeviceTimeResponse) {
	response = &SyncDeviceTimeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
