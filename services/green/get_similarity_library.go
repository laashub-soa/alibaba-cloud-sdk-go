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

// GetSimilarityLibrary invokes the green.GetSimilarityLibrary API synchronously
func (client *Client) GetSimilarityLibrary(request *GetSimilarityLibraryRequest) (response *GetSimilarityLibraryResponse, err error) {
	response = CreateGetSimilarityLibraryResponse()
	err = client.DoAction(request, response)
	return
}

// GetSimilarityLibraryWithChan invokes the green.GetSimilarityLibrary API asynchronously
func (client *Client) GetSimilarityLibraryWithChan(request *GetSimilarityLibraryRequest) (<-chan *GetSimilarityLibraryResponse, <-chan error) {
	responseChan := make(chan *GetSimilarityLibraryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSimilarityLibrary(request)
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

// GetSimilarityLibraryWithCallback invokes the green.GetSimilarityLibrary API asynchronously
func (client *Client) GetSimilarityLibraryWithCallback(request *GetSimilarityLibraryRequest, callback func(response *GetSimilarityLibraryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSimilarityLibraryResponse
		var err error
		defer close(result)
		response, err = client.GetSimilarityLibrary(request)
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

// GetSimilarityLibraryRequest is the request struct for api GetSimilarityLibrary
type GetSimilarityLibraryRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// GetSimilarityLibraryResponse is the response struct for api GetSimilarityLibrary
type GetSimilarityLibraryResponse struct {
	*responses.BaseResponse
}

// CreateGetSimilarityLibraryRequest creates a request to invoke GetSimilarityLibrary API
func CreateGetSimilarityLibraryRequest() (request *GetSimilarityLibraryRequest) {
	request = &GetSimilarityLibraryRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2018-05-09", "GetSimilarityLibrary", "/green/similarity/library/get", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetSimilarityLibraryResponse creates a response to parse from GetSimilarityLibrary response
func CreateGetSimilarityLibraryResponse() (response *GetSimilarityLibraryResponse) {
	response = &GetSimilarityLibraryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
