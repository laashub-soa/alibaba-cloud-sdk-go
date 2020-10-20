package opensearch

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

// DescribeSecondRank invokes the opensearch.DescribeSecondRank API synchronously
func (client *Client) DescribeSecondRank(request *DescribeSecondRankRequest) (response *DescribeSecondRankResponse, err error) {
	response = CreateDescribeSecondRankResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSecondRankWithChan invokes the opensearch.DescribeSecondRank API asynchronously
func (client *Client) DescribeSecondRankWithChan(request *DescribeSecondRankRequest) (<-chan *DescribeSecondRankResponse, <-chan error) {
	responseChan := make(chan *DescribeSecondRankResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSecondRank(request)
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

// DescribeSecondRankWithCallback invokes the opensearch.DescribeSecondRank API asynchronously
func (client *Client) DescribeSecondRankWithCallback(request *DescribeSecondRankRequest, callback func(response *DescribeSecondRankResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSecondRankResponse
		var err error
		defer close(result)
		response, err = client.DescribeSecondRank(request)
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

// DescribeSecondRankRequest is the request struct for api DescribeSecondRank
type DescribeSecondRankRequest struct {
	*requests.RoaRequest
	AppId            requests.Integer `position:"Path" name:"appId"`
	Name             string           `position:"Path" name:"name"`
	AppGroupIdentity string           `position:"Path" name:"appGroupIdentity"`
}

// DescribeSecondRankResponse is the response struct for api DescribeSecondRank
type DescribeSecondRankResponse struct {
	*responses.BaseResponse
	RequestId string                     `json:"requestId" xml:"requestId"`
	Result    ResultInDescribeSecondRank `json:"result" xml:"result"`
}

// CreateDescribeSecondRankRequest creates a request to invoke DescribeSecondRank API
func CreateDescribeSecondRankRequest() (request *DescribeSecondRankRequest) {
	request = &DescribeSecondRankRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "DescribeSecondRank", "/v4/openapi/app-groups/[appGroupIdentity]/apps/[appId]/second-ranks/[name]", "opensearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeSecondRankResponse creates a response to parse from DescribeSecondRank response
func CreateDescribeSecondRankResponse() (response *DescribeSecondRankResponse) {
	response = &DescribeSecondRankResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
