package rds

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

// DeleteBackup invokes the rds.DeleteBackup API synchronously
// api document: https://help.aliyun.com/api/rds/deletebackup.html
func (client *Client) DeleteBackup(request *DeleteBackupRequest) (response *DeleteBackupResponse, err error) {
	response = CreateDeleteBackupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteBackupWithChan invokes the rds.DeleteBackup API asynchronously
// api document: https://help.aliyun.com/api/rds/deletebackup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteBackupWithChan(request *DeleteBackupRequest) (<-chan *DeleteBackupResponse, <-chan error) {
	responseChan := make(chan *DeleteBackupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteBackup(request)
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

// DeleteBackupWithCallback invokes the rds.DeleteBackup API asynchronously
// api document: https://help.aliyun.com/api/rds/deletebackup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteBackupWithCallback(request *DeleteBackupRequest, callback func(response *DeleteBackupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteBackupResponse
		var err error
		defer close(result)
		response, err = client.DeleteBackup(request)
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

// DeleteBackupRequest is the request struct for api DeleteBackup
type DeleteBackupRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	BackupId             string           `position:"Query" name:"BackupId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

// DeleteBackupResponse is the response struct for api DeleteBackup
type DeleteBackupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteBackupRequest creates a request to invoke DeleteBackup API
func CreateDeleteBackupRequest() (request *DeleteBackupRequest) {
	request = &DeleteBackupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DeleteBackup", "rds", "openAPI")
	return
}

// CreateDeleteBackupResponse creates a response to parse from DeleteBackup response
func CreateDeleteBackupResponse() (response *DeleteBackupResponse) {
	response = &DeleteBackupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
