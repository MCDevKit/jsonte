package main

import (
	"encoding/json"
	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/fatih/color"
	"github.com/gammazero/deque"
	ipc "github.com/james-barrow/golang-ipc"
)

type EvaluateExpressionRequest struct {
	ExtraScope map[string]interface{} `json:"extraScope"`
	Expression string                 `json:"expression"`
	Path       string                 `json:"path"`
}

type EvaluateExpressionResponse struct {
	Result    interface{} `json:"result"`
	Action    string      `json:"action"`
	ValueName string      `json:"valueName"`
	IndexName string      `json:"indexName"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

const (
	CloseConnection    int = 1
	EvaluateExpression int = 2
	Error              int = 3
)

func StartIPC(ipcName string, scope utils.NavigableMap[string, interface{}]) error {
	color.NoColor = true
	sc, err := ipc.StartServer(ipcName, nil)
	if err != nil {
		return utils.WrapError(err, "An error occurred while starting the server")
	}
	utils.Logger.Infof("Started IPC server %s", ipcName)
	for {
		message, err := sc.Read()
		if err == nil {
			if message.MsgType <= 0 {
				continue
			}
			if message.MsgType == CloseConnection {
				utils.Logger.Infof("Received close connection message")
				sc.Close()
				break
			} else if message.MsgType == EvaluateExpression {
				utils.Logger.Infof("Received evaluate expression message")
				var request EvaluateExpressionRequest
				err := json.Unmarshal(message.Data, &request)
				if err != nil {
					sendError(sc, utils.WrapErrorf(err, "An error occurred while unmarshalling the request"))
					continue
				}
				s := deque.Deque[interface{}]{}
				s.PushBack(scope)
				s.PushBack(request.ExtraScope)
				utils.Logger.Infof("Evaluating expression %s", request.Expression)
				result, err := jsonte.Eval(request.Expression, s, request.Path)
				if err != nil {
					sendError(sc, utils.WrapErrorf(err, "An error occurred while evaluating the expression %s", request.Expression))
					continue
				}
				response := EvaluateExpressionResponse{
					Result:    utils.UnwrapContainers(result.Value),
					Action:    result.Action.String(),
					ValueName: result.Name,
					IndexName: result.IndexName,
				}
				data, err := json.Marshal(response)
				err = sc.Write(EvaluateExpression, data)
				if err != nil {
					utils.Logger.Errorf("An error occurred while writing to the client: %s", err)
					return err
				}
			}
		} else {
			return utils.WrapErrorf(err, "An error occurred while reading from the client")
		}
	}
	return nil
}

func sendError(sc *ipc.Server, err error) {
	response := ErrorResponse{
		Error: err.Error(),
	}
	data, err := json.Marshal(response)
	if err != nil {
		utils.Logger.Errorf("An error occurred while marshalling the response: %s", err)
		return
	}
	err = sc.Write(Error, data)
	if err != nil {
		utils.Logger.Errorf("An error occurred while writing to the client: %s", err)
		return
	}
}
