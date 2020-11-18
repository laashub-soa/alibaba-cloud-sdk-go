package green

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

// VideoAsyncScan invokes the green.VideoAsyncScan API synchronously
func (client *Client) VideoAsyncScan(request *VideoAsyncScanRequest) (response *VideoAsyncScanResponse, err error) {
	response = CreateVideoAsyncScanResponse()
	err = client.DoAction(request, response)
	return
}

// VideoAsyncScanWithChan invokes the green.VideoAsyncScan API asynchronously
func (client *Client) VideoAsyncScanWithChan(request *VideoAsyncScanRequest) (<-chan *VideoAsyncScanResponse, <-chan error) {
	responseChan := make(chan *VideoAsyncScanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.VideoAsyncScan(request)
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

// VideoAsyncScanWithCallback invokes the green.VideoAsyncScan API asynchronously
func (client *Client) VideoAsyncScanWithCallback(request *VideoAsyncScanRequest, callback func(response *VideoAsyncScanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *VideoAsyncScanResponse
		var err error
		defer close(result)
		response, err = client.VideoAsyncScan(request)
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

// VideoAsyncScanRequest is the request struct for api VideoAsyncScan
type VideoAsyncScanRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// VideoAsyncScanResponse is the response struct for api VideoAsyncScan
type VideoAsyncScanResponse struct {
	*responses.BaseResponse
}

// CreateVideoAsyncScanRequest creates a request to invoke VideoAsyncScan API
func CreateVideoAsyncScanRequest() (request *VideoAsyncScanRequest) {
	request = &VideoAsyncScanRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2018-05-09", "VideoAsyncScan", "/green/video/asyncscan", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateVideoAsyncScanResponse creates a response to parse from VideoAsyncScan response
func CreateVideoAsyncScanResponse() (response *VideoAsyncScanResponse) {
	response = &VideoAsyncScanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
