package airec

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

// Meta is a nested struct in airec response
type Meta struct {
	ExtInfo       map[string]interface{} `json:"extInfo" xml:"extInfo"`
	Type          string                 `json:"type" xml:"type"`
	Timestamp     int64                  `json:"timestamp" xml:"timestamp"`
	BucketName    string                 `json:"bucketName" xml:"bucketName"`
	AccessKeyId   string                 `json:"accessKeyId" xml:"accessKeyId"`
	TableName     string                 `json:"tableName" xml:"tableName"`
	CronEnabled   bool                   `json:"cronEnabled" xml:"cronEnabled"`
	GmtModified   string                 `json:"gmtModified" xml:"gmtModified"`
	Status        string                 `json:"status" xml:"status"`
	Category      string                 `json:"category" xml:"category"`
	AlgorithmName string                 `json:"algorithmName" xml:"algorithmName"`
	Description   string                 `json:"description" xml:"description"`
	GmtCreate     string                 `json:"gmtCreate" xml:"gmtCreate"`
	Partition     string                 `json:"partition" xml:"partition"`
	MetaType      string                 `json:"metaType" xml:"metaType"`
	Cron          string                 `json:"cron" xml:"cron"`
	ProjectName   string                 `json:"projectName" xml:"projectName"`
	Path          string                 `json:"path" xml:"path"`
	Threshold     Threshold              `json:"threshold" xml:"threshold"`
}
