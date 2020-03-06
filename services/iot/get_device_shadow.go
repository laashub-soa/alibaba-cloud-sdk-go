package iot

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

// GetDeviceShadow invokes the iot.GetDeviceShadow API synchronously
// api document: https://help.aliyun.com/api/iot/getdeviceshadow.html
func (client *Client) GetDeviceShadow(request *GetDeviceShadowRequest) (response *GetDeviceShadowResponse, err error) {
	response = CreateGetDeviceShadowResponse()
	err = client.DoAction(request, response)
	return
}

// GetDeviceShadowWithChan invokes the iot.GetDeviceShadow API asynchronously
// api document: https://help.aliyun.com/api/iot/getdeviceshadow.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDeviceShadowWithChan(request *GetDeviceShadowRequest) (<-chan *GetDeviceShadowResponse, <-chan error) {
	responseChan := make(chan *GetDeviceShadowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDeviceShadow(request)
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

// GetDeviceShadowWithCallback invokes the iot.GetDeviceShadow API asynchronously
// api document: https://help.aliyun.com/api/iot/getdeviceshadow.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDeviceShadowWithCallback(request *GetDeviceShadowRequest, callback func(response *GetDeviceShadowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDeviceShadowResponse
		var err error
		defer close(result)
		response, err = client.GetDeviceShadow(request)
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

// GetDeviceShadowRequest is the request struct for api GetDeviceShadow
type GetDeviceShadowRequest struct {
	*requests.RpcRequest
	ShadowMessage string `position:"Query" name:"ShadowMessage"`
	ProductKey    string `position:"Query" name:"ProductKey"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	DeviceName    string `position:"Query" name:"DeviceName"`
}

// GetDeviceShadowResponse is the response struct for api GetDeviceShadow
type GetDeviceShadowResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	Success       bool   `json:"Success" xml:"Success"`
	Code          string `json:"Code" xml:"Code"`
	ErrorMessage  string `json:"ErrorMessage" xml:"ErrorMessage"`
	ShadowMessage string `json:"ShadowMessage" xml:"ShadowMessage"`
}

// CreateGetDeviceShadowRequest creates a request to invoke GetDeviceShadow API
func CreateGetDeviceShadowRequest() (request *GetDeviceShadowRequest) {
	request = &GetDeviceShadowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "GetDeviceShadow", "Iot", "openAPI")
	return
}

// CreateGetDeviceShadowResponse creates a response to parse from GetDeviceShadow response
func CreateGetDeviceShadowResponse() (response *GetDeviceShadowResponse) {
	response = &GetDeviceShadowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
