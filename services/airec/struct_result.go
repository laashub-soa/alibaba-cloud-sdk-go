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

// Result is a nested struct in airec response
type Result struct {
	OnlineTime     string                 `json:"onlineTime" xml:"onlineTime"`
	RuleId         string                 `json:"ruleId" xml:"ruleId"`
	State          string                 `json:"state" xml:"state"`
	Num            int                    `json:"num" xml:"num"`
	Detail         map[string]interface{} `json:"detail" xml:"detail"`
	GmtModified    int64                  `json:"gmtModified" xml:"gmtModified"`
	VersionId      string                 `json:"versionId" xml:"versionId"`
	FinishRate     int                    `json:"finishRate" xml:"finishRate"`
	OfflineTime    string                 `json:"offlineTime" xml:"offlineTime"`
	Name           string                 `json:"name" xml:"name"`
	RankingModelId string                 `json:"rankingModelId" xml:"rankingModelId"`
	Base           bool                   `json:"base" xml:"base"`
	GmtCreate      int64                  `json:"gmtCreate" xml:"gmtCreate"`
	SceneId        string                 `json:"sceneId" xml:"sceneId"`
	DivideType     string                 `json:"divideType" xml:"divideType"`
	Progress       int                    `json:"progress" xml:"progress"`
	InstanceId     string                 `json:"instanceId" xml:"instanceId"`
	Status         string                 `json:"status" xml:"status"`
	Description    string                 `json:"description" xml:"description"`
	FinishTime     int                    `json:"finishTime" xml:"finishTime"`
	MetricType     string                 `json:"metricType" xml:"metricType"`
	ExperimentId   string                 `json:"experimentId" xml:"experimentId"`
	BucketCount    int                    `json:"bucketCount" xml:"bucketCount"`
	InUse          string                 `json:"inUse" xml:"inUse"`
	Meta           map[string]interface{} `json:"meta" xml:"meta"`
	PvCount        int64                  `json:"pvCount" xml:"pvCount"`
	Buckets        []string               `json:"buckets" xml:"buckets"`
	Algorithms     []Algorithm            `json:"algorithms" xml:"algorithms"`
	MetricData     []MetricData           `json:"metricData" xml:"metricData"`
	DataPoints     []DataPoints           `json:"dataPoints" xml:"dataPoints"`
}
