package test

import (
  "encoding/json"
  "testing"
)

import (
  . "github.com/jellybean4/gosalt/mesg"
)

func TestDeploy(t *testing.T) {
  startServer(t)

  reqStruct := &DeployReq{
    Version: "trunk",
    Server:  "*",
  }

  data, repStruct := request(t, reqStruct, "deploy/do"), &DeployRep{}
  if err := json.Unmarshal(data, repStruct); err != nil {
    t.Errorf("parse deploy response failed, reason %s", err.Error())
  }
}
