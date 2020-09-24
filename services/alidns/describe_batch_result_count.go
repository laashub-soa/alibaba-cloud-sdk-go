package alidns

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

// DescribeBatchResultCount invokes the alidns.DescribeBatchResultCount API synchronously
func (client *Client) DescribeBatchResultCount(request *DescribeBatchResultCountRequest) (response *DescribeBatchResultCountResponse, err error) {
	response = CreateDescribeBatchResultCountResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBatchResultCountWithChan invokes the alidns.DescribeBatchResultCount API asynchronously
func (client *Client) DescribeBatchResultCountWithChan(request *DescribeBatchResultCountRequest) (<-chan *DescribeBatchResultCountResponse, <-chan error) {
	responseChan := make(chan *DescribeBatchResultCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBatchResultCount(request)
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

// DescribeBatchResultCountWithCallback invokes the alidns.DescribeBatchResultCount API asynchronously
func (client *Client) DescribeBatchResultCountWithCallback(request *DescribeBatchResultCountRequest, callback func(response *DescribeBatchResultCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBatchResultCountResponse
		var err error
		defer close(result)
		response, err = client.DescribeBatchResultCount(request)
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

// DescribeBatchResultCountRequest is the request struct for api DescribeBatchResultCount
type DescribeBatchResultCountRequest struct {
	*requests.RpcRequest
	BatchType    string           `position:"Query" name:"BatchType"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	Lang         string           `position:"Query" name:"Lang"`
	TaskId       requests.Integer `position:"Query" name:"TaskId"`
}

// DescribeBatchResultCountResponse is the response struct for api DescribeBatchResultCount
type DescribeBatchResultCountResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Status       int    `json:"Status" xml:"Status"`
	TotalCount   int    `json:"TotalCount" xml:"TotalCount"`
	SuccessCount int    `json:"SuccessCount" xml:"SuccessCount"`
	FailedCount  int    `json:"FailedCount" xml:"FailedCount"`
	Reason       string `json:"Reason" xml:"Reason"`
	BatchType    string `json:"BatchType" xml:"BatchType"`
	TaskId       int64  `json:"TaskId" xml:"TaskId"`
}

// CreateDescribeBatchResultCountRequest creates a request to invoke DescribeBatchResultCount API
func CreateDescribeBatchResultCountRequest() (request *DescribeBatchResultCountRequest) {
	request = &DescribeBatchResultCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeBatchResultCount", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeBatchResultCountResponse creates a response to parse from DescribeBatchResultCount response
func CreateDescribeBatchResultCountResponse() (response *DescribeBatchResultCountResponse) {
	response = &DescribeBatchResultCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
