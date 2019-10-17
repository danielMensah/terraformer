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

// DescribeBackupDatabase invokes the rds.DescribeBackupDatabase API synchronously
// api document: https://help.aliyun.com/api/rds/describebackupdatabase.html
func (client *Client) DescribeBackupDatabase(request *DescribeBackupDatabaseRequest) (response *DescribeBackupDatabaseResponse, err error) {
	response = CreateDescribeBackupDatabaseResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBackupDatabaseWithChan invokes the rds.DescribeBackupDatabase API asynchronously
// api document: https://help.aliyun.com/api/rds/describebackupdatabase.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeBackupDatabaseWithChan(request *DescribeBackupDatabaseRequest) (<-chan *DescribeBackupDatabaseResponse, <-chan error) {
	responseChan := make(chan *DescribeBackupDatabaseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBackupDatabase(request)
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

// DescribeBackupDatabaseWithCallback invokes the rds.DescribeBackupDatabase API asynchronously
// api document: https://help.aliyun.com/api/rds/describebackupdatabase.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeBackupDatabaseWithCallback(request *DescribeBackupDatabaseRequest, callback func(response *DescribeBackupDatabaseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBackupDatabaseResponse
		var err error
		defer close(result)
		response, err = client.DescribeBackupDatabase(request)
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

// DescribeBackupDatabaseRequest is the request struct for api DescribeBackupDatabase
type DescribeBackupDatabaseRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	BackupId             string           `position:"Query" name:"BackupId"`
}

// DescribeBackupDatabaseResponse is the response struct for api DescribeBackupDatabase
type DescribeBackupDatabaseResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	DatabaseNames string `json:"DatabaseNames" xml:"DatabaseNames"`
}

// CreateDescribeBackupDatabaseRequest creates a request to invoke DescribeBackupDatabase API
func CreateDescribeBackupDatabaseRequest() (request *DescribeBackupDatabaseRequest) {
	request = &DescribeBackupDatabaseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeBackupDatabase", "rds", "openAPI")
	return
}

// CreateDescribeBackupDatabaseResponse creates a response to parse from DescribeBackupDatabase response
func CreateDescribeBackupDatabaseResponse() (response *DescribeBackupDatabaseResponse) {
	response = &DescribeBackupDatabaseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
