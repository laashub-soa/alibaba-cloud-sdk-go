package csb

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

// FindBrokerSLOList invokes the csb.FindBrokerSLOList API synchronously
// api document: https://help.aliyun.com/api/csb/findbrokerslolist.html
func (client *Client) FindBrokerSLOList(request *FindBrokerSLOListRequest) (response *FindBrokerSLOListResponse, err error) {
	response = CreateFindBrokerSLOListResponse()
	err = client.DoAction(request, response)
	return
}

// FindBrokerSLOListWithChan invokes the csb.FindBrokerSLOList API asynchronously
// api document: https://help.aliyun.com/api/csb/findbrokerslolist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FindBrokerSLOListWithChan(request *FindBrokerSLOListRequest) (<-chan *FindBrokerSLOListResponse, <-chan error) {
	responseChan := make(chan *FindBrokerSLOListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FindBrokerSLOList(request)
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

// FindBrokerSLOListWithCallback invokes the csb.FindBrokerSLOList API asynchronously
// api document: https://help.aliyun.com/api/csb/findbrokerslolist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FindBrokerSLOListWithCallback(request *FindBrokerSLOListRequest, callback func(response *FindBrokerSLOListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FindBrokerSLOListResponse
		var err error
		defer close(result)
		response, err = client.FindBrokerSLOList(request)
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

// FindBrokerSLOListRequest is the request struct for api FindBrokerSLOList
type FindBrokerSLOListRequest struct {
	*requests.RpcRequest
	CsbId requests.Integer `position:"Query" name:"CsbId"`
}

// FindBrokerSLOListResponse is the response struct for api FindBrokerSLOList
type FindBrokerSLOListResponse struct {
	*responses.BaseResponse
	Code      int                    `json:"Code" xml:"Code"`
	Message   string                 `json:"Message" xml:"Message"`
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	Data      map[string]interface{} `json:"Data" xml:"Data"`
}

// CreateFindBrokerSLOListRequest creates a request to invoke FindBrokerSLOList API
func CreateFindBrokerSLOListRequest() (request *FindBrokerSLOListRequest) {
	request = &FindBrokerSLOListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CSB", "2017-11-18", "FindBrokerSLOList", "", "")
	return
}

// CreateFindBrokerSLOListResponse creates a response to parse from FindBrokerSLOList response
func CreateFindBrokerSLOListResponse() (response *FindBrokerSLOListResponse) {
	response = &FindBrokerSLOListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}