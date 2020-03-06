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

// QueryDeviceServiceData invokes the iot.QueryDeviceServiceData API synchronously
// api document: https://help.aliyun.com/api/iot/querydeviceservicedata.html
func (client *Client) QueryDeviceServiceData(request *QueryDeviceServiceDataRequest) (response *QueryDeviceServiceDataResponse, err error) {
	response = CreateQueryDeviceServiceDataResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDeviceServiceDataWithChan invokes the iot.QueryDeviceServiceData API asynchronously
// api document: https://help.aliyun.com/api/iot/querydeviceservicedata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDeviceServiceDataWithChan(request *QueryDeviceServiceDataRequest) (<-chan *QueryDeviceServiceDataResponse, <-chan error) {
	responseChan := make(chan *QueryDeviceServiceDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDeviceServiceData(request)
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

// QueryDeviceServiceDataWithCallback invokes the iot.QueryDeviceServiceData API asynchronously
// api document: https://help.aliyun.com/api/iot/querydeviceservicedata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDeviceServiceDataWithCallback(request *QueryDeviceServiceDataRequest, callback func(response *QueryDeviceServiceDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDeviceServiceDataResponse
		var err error
		defer close(result)
		response, err = client.QueryDeviceServiceData(request)
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

// QueryDeviceServiceDataRequest is the request struct for api QueryDeviceServiceData
type QueryDeviceServiceDataRequest struct {
	*requests.RpcRequest
	Identifier    string           `position:"Query" name:"Identifier"`
	EndTime       requests.Integer `position:"Query" name:"EndTime"`
	StartTime     requests.Integer `position:"Query" name:"StartTime"`
	ProductKey    string           `position:"Query" name:"ProductKey"`
	Asc           requests.Integer `position:"Query" name:"Asc"`
	IotId         string           `position:"Query" name:"IotId"`
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	DeviceName    string           `position:"Query" name:"DeviceName"`
}

// QueryDeviceServiceDataResponse is the response struct for api QueryDeviceServiceData
type QueryDeviceServiceDataResponse struct {
	*responses.BaseResponse
	RequestId    string                       `json:"RequestId" xml:"RequestId"`
	Success      bool                         `json:"Success" xml:"Success"`
	Code         string                       `json:"Code" xml:"Code"`
	ErrorMessage string                       `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryDeviceServiceData `json:"Data" xml:"Data"`
}

// CreateQueryDeviceServiceDataRequest creates a request to invoke QueryDeviceServiceData API
func CreateQueryDeviceServiceDataRequest() (request *QueryDeviceServiceDataRequest) {
	request = &QueryDeviceServiceDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryDeviceServiceData", "Iot", "openAPI")
	return
}

// CreateQueryDeviceServiceDataResponse creates a response to parse from QueryDeviceServiceData response
func CreateQueryDeviceServiceDataResponse() (response *QueryDeviceServiceDataResponse) {
	response = &QueryDeviceServiceDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
