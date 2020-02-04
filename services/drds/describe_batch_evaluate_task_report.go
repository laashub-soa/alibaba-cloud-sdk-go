package drds

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

// DescribeBatchEvaluateTaskReport invokes the drds.DescribeBatchEvaluateTaskReport API synchronously
// api document: https://help.aliyun.com/api/drds/describebatchevaluatetaskreport.html
func (client *Client) DescribeBatchEvaluateTaskReport(request *DescribeBatchEvaluateTaskReportRequest) (response *DescribeBatchEvaluateTaskReportResponse, err error) {
	response = CreateDescribeBatchEvaluateTaskReportResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBatchEvaluateTaskReportWithChan invokes the drds.DescribeBatchEvaluateTaskReport API asynchronously
// api document: https://help.aliyun.com/api/drds/describebatchevaluatetaskreport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeBatchEvaluateTaskReportWithChan(request *DescribeBatchEvaluateTaskReportRequest) (<-chan *DescribeBatchEvaluateTaskReportResponse, <-chan error) {
	responseChan := make(chan *DescribeBatchEvaluateTaskReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBatchEvaluateTaskReport(request)
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

// DescribeBatchEvaluateTaskReportWithCallback invokes the drds.DescribeBatchEvaluateTaskReport API asynchronously
// api document: https://help.aliyun.com/api/drds/describebatchevaluatetaskreport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeBatchEvaluateTaskReportWithCallback(request *DescribeBatchEvaluateTaskReportRequest, callback func(response *DescribeBatchEvaluateTaskReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBatchEvaluateTaskReportResponse
		var err error
		defer close(result)
		response, err = client.DescribeBatchEvaluateTaskReport(request)
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

// DescribeBatchEvaluateTaskReportRequest is the request struct for api DescribeBatchEvaluateTaskReport
type DescribeBatchEvaluateTaskReportRequest struct {
	*requests.RpcRequest
	BatchEvaluateTaskId requests.Integer `position:"Query" name:"BatchEvaluateTaskId"`
}

// DescribeBatchEvaluateTaskReportResponse is the response struct for api DescribeBatchEvaluateTaskReport
type DescribeBatchEvaluateTaskReportResponse struct {
	*responses.BaseResponse
	RequestId string                                `json:"RequestId" xml:"RequestId"`
	Success   bool                                  `json:"Success" xml:"Success"`
	Data      DataInDescribeBatchEvaluateTaskReport `json:"Data" xml:"Data"`
}

// CreateDescribeBatchEvaluateTaskReportRequest creates a request to invoke DescribeBatchEvaluateTaskReport API
func CreateDescribeBatchEvaluateTaskReportRequest() (request *DescribeBatchEvaluateTaskReportRequest) {
	request = &DescribeBatchEvaluateTaskReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeBatchEvaluateTaskReport", "Drds", "openAPI")
	return
}

// CreateDescribeBatchEvaluateTaskReportResponse creates a response to parse from DescribeBatchEvaluateTaskReport response
func CreateDescribeBatchEvaluateTaskReportResponse() (response *DescribeBatchEvaluateTaskReportResponse) {
	response = &DescribeBatchEvaluateTaskReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}