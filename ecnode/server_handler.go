// Copyright 2020 The Chubao Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package ecnode

import (
	"encoding/json"
	"net/http"

	"github.com/chubaofs/chubaofs/util/log"
)

func (e *EcNode) getDiskAPI(w http.ResponseWriter, r *http.Request) {
	return
}

func (e *EcNode) getPartitionsAPI(w http.ResponseWriter, r *http.Request) {
	log.LogDebugf("action[getPartitionsAPI]")

	partitions := make([]interface{}, 0)
	e.space.RangePartitions(func(ep *EcPartition) bool {
		partition := &struct {
			ID            uint64   `json:"partition_id"`
			Size          int      `json:"size"`
			Used          int      `json:"used"`
			Status        int      `json:"status"`
			Path          string   `json:"path"`
			DataNodeNum   int      `json:"data_node_num"`
			ParityNodeNum int      `json:"parity_node_num"`
			NodeIndex     int      `json:node_index`
			DataNodes     []string `json:data_nodes`
			ParityNodes   []string `json:parity_nodes`
		}{
			ID:            ep.partitionID,
			Size:          ep.Size(),
			Used:          ep.Used(),
			Status:        ep.Status(),
			Path:          ep.Path(),
			DataNodeNum:   ep.DataNodeNum(),
			ParityNodeNum: ep.ParityNodeNum(),
			NodeIndex:     ep.NodeIndex(),
			DataNodes:     ep.DataNodes(),
			ParityNodes:   ep.ParityNodes(),
		}
		partitions = append(partitions, partition)
		return true
	})
	result := &struct {
		Partitions     []interface{} `json:"partitions"`
		PartitionCount int           `json:"partitionCount"`
	}{
		Partitions:     partitions,
		PartitionCount: len(partitions),
	}
	e.buildSuccessResp(w, result)
	return
}

func (e *EcNode) getPartitionAPI(w http.ResponseWriter, r *http.Request) {
	return
}

func (e *EcNode) getExtentAPI(w http.ResponseWriter, r *http.Request) {
	return
}

func (e *EcNode) getBlockCrcAPI(w http.ResponseWriter, r *http.Request) {
	return
}

func (e *EcNode) getStatAPI(w http.ResponseWriter, r *http.Request) {
	return
}

func (e *EcNode) getRaftStatusAPI(w http.ResponseWriter, r *http.Request) {
	return
}

func (e *EcNode) buildSuccessResp(w http.ResponseWriter, data interface{}) {
	e.buildJSONResp(w, http.StatusOK, data, "")
}

func (e *EcNode) buildFailureResp(w http.ResponseWriter, code int, msg string) {
	e.buildJSONResp(w, code, nil, msg)
}

func (e *EcNode) buildJSONResp(w http.ResponseWriter, code int, data interface{}, msg string) {
	var (
		jsonBody []byte
		err      error
	)
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	body := struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
		Msg  string      `json:"msg"`
	}{
		Code: code,
		Data: data,
		Msg:  msg,
	}
	if jsonBody, err = json.Marshal(body); err != nil {
		return
	}
	w.Write(jsonBody)
}
