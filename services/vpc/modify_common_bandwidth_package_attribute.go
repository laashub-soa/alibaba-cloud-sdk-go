package vpc

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

// ModifyCommonBandwidthPackageAttribute invokes the vpc.ModifyCommonBandwidthPackageAttribute API synchronously
// api document: https://help.aliyun.com/api/vpc/modifycommonbandwidthpackageattribute.html
func (client *Client) ModifyCommonBandwidthPackageAttribute(request *ModifyCommonBandwidthPackageAttributeRequest) (response *ModifyCommonBandwidthPackageAttributeResponse, err error) {
	response = CreateModifyCommonBandwidthPackageAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyCommonBandwidthPackageAttributeWithChan invokes the vpc.ModifyCommonBandwidthPackageAttribute API asynchronously
// api document: https://help.aliyun.com/api/vpc/modifycommonbandwidthpackageattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyCommonBandwidthPackageAttributeWithChan(request *ModifyCommonBandwidthPackageAttributeRequest) (<-chan *ModifyCommonBandwidthPackageAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyCommonBandwidthPackageAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyCommonBandwidthPackageAttribute(request)
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

// ModifyCommonBandwidthPackageAttributeWithCallback invokes the vpc.ModifyCommonBandwidthPackageAttribute API asynchronously
// api document: https://help.aliyun.com/api/vpc/modifycommonbandwidthpackageattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyCommonBandwidthPackageAttributeWithCallback(request *ModifyCommonBandwidthPackageAttributeRequest, callback func(response *ModifyCommonBandwidthPackageAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyCommonBandwidthPackageAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyCommonBandwidthPackageAttribute(request)
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

// ModifyCommonBandwidthPackageAttributeRequest is the request struct for api ModifyCommonBandwidthPackageAttribute
type ModifyCommonBandwidthPackageAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Description          string           `position:"Query" name:"Description"`
	BandwidthPackageId   string           `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Name                 string           `position:"Query" name:"Name"`
}

// ModifyCommonBandwidthPackageAttributeResponse is the response struct for api ModifyCommonBandwidthPackageAttribute
type ModifyCommonBandwidthPackageAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyCommonBandwidthPackageAttributeRequest creates a request to invoke ModifyCommonBandwidthPackageAttribute API
func CreateModifyCommonBandwidthPackageAttributeRequest() (request *ModifyCommonBandwidthPackageAttributeRequest) {
	request = &ModifyCommonBandwidthPackageAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyCommonBandwidthPackageAttribute", "vpc", "openAPI")
	return
}

// CreateModifyCommonBandwidthPackageAttributeResponse creates a response to parse from ModifyCommonBandwidthPackageAttribute response
func CreateModifyCommonBandwidthPackageAttributeResponse() (response *ModifyCommonBandwidthPackageAttributeResponse) {
	response = &ModifyCommonBandwidthPackageAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
