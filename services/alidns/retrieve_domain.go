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

// RetrieveDomain invokes the alidns.RetrieveDomain API synchronously
func (client *Client) RetrieveDomain(request *RetrieveDomainRequest) (response *RetrieveDomainResponse, err error) {
	response = CreateRetrieveDomainResponse()
	err = client.DoAction(request, response)
	return
}

// RetrieveDomainWithChan invokes the alidns.RetrieveDomain API asynchronously
func (client *Client) RetrieveDomainWithChan(request *RetrieveDomainRequest) (<-chan *RetrieveDomainResponse, <-chan error) {
	responseChan := make(chan *RetrieveDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RetrieveDomain(request)
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

// RetrieveDomainWithCallback invokes the alidns.RetrieveDomain API asynchronously
func (client *Client) RetrieveDomainWithCallback(request *RetrieveDomainRequest, callback func(response *RetrieveDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RetrieveDomainResponse
		var err error
		defer close(result)
		response, err = client.RetrieveDomain(request)
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

// RetrieveDomainRequest is the request struct for api RetrieveDomain
type RetrieveDomainRequest struct {
	*requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// RetrieveDomainResponse is the response struct for api RetrieveDomain
type RetrieveDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRetrieveDomainRequest creates a request to invoke RetrieveDomain API
func CreateRetrieveDomainRequest() (request *RetrieveDomainRequest) {
	request = &RetrieveDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "RetrieveDomain", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRetrieveDomainResponse creates a response to parse from RetrieveDomain response
func CreateRetrieveDomainResponse() (response *RetrieveDomainResponse) {
	response = &RetrieveDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
