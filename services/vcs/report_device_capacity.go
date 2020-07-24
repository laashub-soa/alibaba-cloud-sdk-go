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

// ReportDeviceCapacity invokes the vcs.ReportDeviceCapacity API synchronously
// api document: https://help.aliyun.com/api/vcs/reportdevicecapacity.html
func (client *Client) ReportDeviceCapacity(request *ReportDeviceCapacityRequest) (response *ReportDeviceCapacityResponse, err error) {
	response = CreateReportDeviceCapacityResponse()
	err = client.DoAction(request, response)
	return
}

// ReportDeviceCapacityWithChan invokes the vcs.ReportDeviceCapacity API asynchronously
// api document: https://help.aliyun.com/api/vcs/reportdevicecapacity.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReportDeviceCapacityWithChan(request *ReportDeviceCapacityRequest) (<-chan *ReportDeviceCapacityResponse, <-chan error) {
	responseChan := make(chan *ReportDeviceCapacityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReportDeviceCapacity(request)
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

// ReportDeviceCapacityWithCallback invokes the vcs.ReportDeviceCapacity API asynchronously
// api document: https://help.aliyun.com/api/vcs/reportdevicecapacity.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReportDeviceCapacityWithCallback(request *ReportDeviceCapacityRequest, callback func(response *ReportDeviceCapacityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReportDeviceCapacityResponse
		var err error
		defer close(result)
		response, err = client.ReportDeviceCapacity(request)
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

// ReportDeviceCapacityRequest is the request struct for api ReportDeviceCapacity
type ReportDeviceCapacityRequest struct {
	*requests.RpcRequest
	StreamCapacities *[]ReportDeviceCapacityStreamCapacities `position:"Body" name:"StreamCapacities"  type:"Repeated"`
	Latitude         string                                  `position:"Body" name:"Latitude"`
	PresetNum        string                                  `position:"Body" name:"PresetNum"`
	DeviceSn         string                                  `position:"Body" name:"DeviceSn"`
	AudioFormat      string                                  `position:"Body" name:"AudioFormat"`
	PTZCapacity      string                                  `position:"Body" name:"PTZCapacity"`
	Longitude        string                                  `position:"Body" name:"Longitude"`
}

// ReportDeviceCapacityStreamCapacities is a repeated param struct in ReportDeviceCapacityRequest
type ReportDeviceCapacityStreamCapacities struct {
	BitrateRange string `name:"BitrateRange"`
	MaxStream    string `name:"MaxStream"`
	EncodeFormat string `name:"EncodeFormat"`
	MaxFrameRate string `name:"MaxFrameRate"`
	Resolution   string `name:"Resolution"`
}

// ReportDeviceCapacityResponse is the response struct for api ReportDeviceCapacity
type ReportDeviceCapacityResponse struct {
	*responses.BaseResponse
	Code          string `json:"Code" xml:"Code"`
	Message       string `json:"Message" xml:"Message"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
	RetryInterval string `json:"RetryInterval" xml:"RetryInterval"`
}

// CreateReportDeviceCapacityRequest creates a request to invoke ReportDeviceCapacity API
func CreateReportDeviceCapacityRequest() (request *ReportDeviceCapacityRequest) {
	request = &ReportDeviceCapacityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "ReportDeviceCapacity", "vcs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReportDeviceCapacityResponse creates a response to parse from ReportDeviceCapacity response
func CreateReportDeviceCapacityResponse() (response *ReportDeviceCapacityResponse) {
	response = &ReportDeviceCapacityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
