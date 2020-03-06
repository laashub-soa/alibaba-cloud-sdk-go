package retailcloud

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

// ListPersistentVolumeClaim invokes the retailcloud.ListPersistentVolumeClaim API synchronously
// api document: https://help.aliyun.com/api/retailcloud/listpersistentvolumeclaim.html
func (client *Client) ListPersistentVolumeClaim(request *ListPersistentVolumeClaimRequest) (response *ListPersistentVolumeClaimResponse, err error) {
	response = CreateListPersistentVolumeClaimResponse()
	err = client.DoAction(request, response)
	return
}

// ListPersistentVolumeClaimWithChan invokes the retailcloud.ListPersistentVolumeClaim API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/listpersistentvolumeclaim.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPersistentVolumeClaimWithChan(request *ListPersistentVolumeClaimRequest) (<-chan *ListPersistentVolumeClaimResponse, <-chan error) {
	responseChan := make(chan *ListPersistentVolumeClaimResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPersistentVolumeClaim(request)
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

// ListPersistentVolumeClaimWithCallback invokes the retailcloud.ListPersistentVolumeClaim API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/listpersistentvolumeclaim.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPersistentVolumeClaimWithCallback(request *ListPersistentVolumeClaimRequest, callback func(response *ListPersistentVolumeClaimResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPersistentVolumeClaimResponse
		var err error
		defer close(result)
		response, err = client.ListPersistentVolumeClaim(request)
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

// ListPersistentVolumeClaimRequest is the request struct for api ListPersistentVolumeClaim
type ListPersistentVolumeClaimRequest struct {
	*requests.RpcRequest
	AppId      requests.Integer `position:"Query" name:"AppId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	EnvId      requests.Integer `position:"Query" name:"EnvId"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// ListPersistentVolumeClaimResponse is the response struct for api ListPersistentVolumeClaim
type ListPersistentVolumeClaimResponse struct {
	*responses.BaseResponse
	Code       int                           `json:"Code" xml:"Code"`
	PageNumber int                           `json:"PageNumber" xml:"PageNumber"`
	RequestId  string                        `json:"RequestId" xml:"RequestId"`
	PageSize   int                           `json:"PageSize" xml:"PageSize"`
	TotalCount int64                         `json:"TotalCount" xml:"TotalCount"`
	ErrorMsg   string                        `json:"ErrorMsg" xml:"ErrorMsg"`
	Data       []PersistentVolumeClaimDetail `json:"Data" xml:"Data"`
}

// CreateListPersistentVolumeClaimRequest creates a request to invoke ListPersistentVolumeClaim API
func CreateListPersistentVolumeClaimRequest() (request *ListPersistentVolumeClaimRequest) {
	request = &ListPersistentVolumeClaimRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "ListPersistentVolumeClaim", "retailcloud", "openAPI")
	return
}

// CreateListPersistentVolumeClaimResponse creates a response to parse from ListPersistentVolumeClaim response
func CreateListPersistentVolumeClaimResponse() (response *ListPersistentVolumeClaimResponse) {
	response = &ListPersistentVolumeClaimResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}