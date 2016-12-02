package test

import (
  "bytes"
  "encoding/json"
  "io/ioutil"
  "net/http"
  "testing"
  "time"
)

import (
  . "github.com/jellybean4/gosalt/web"
  "github.com/spf13/viper"
  "fmt"
  "io"
)

const (
  testPort = 5500
)

func initServer(t *testing.T) {
  viper.SetConfigFile("/tmp/gosalt/gosalt.yml")
  viper.SetConfigName("gosalt")
  viper.AddConfigPath("/etc/gosalt")
  viper.AutomaticEnv()

  // If a config file is found, read it in.
  if err := viper.ReadInConfig(); err != nil {
    errMesg := fmt.Sprintf("parse config file failed, reason %s", err.Error())
    t.Error(errMesg)
  } else {
    t.Log("starting server...")
    go Serve(testPort)
    time.Sleep(100 * time.Millisecond)
  }
}

func request(t *testing.T, reqStruct interface{}, path string) []byte {
  client := &http.Client{}
  reqData, err := json.Marshal(reqStruct)
  if err != nil {
    t.Errorf("marshal data %s failed, reason %s", reqStruct, err.Error())
    return nil
  }

  body := bytes.NewBuffer(reqData)
  url := fmt.Sprintf("http://192.168.2.220:%d/%s", testPort, path)
  request, err := http.NewRequest(http.MethodPost, url, body)
  request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
  request.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
  request.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
  request.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
  request.Header.Set("Cache-Control", "max-age=0")
  request.Header.Set("Connection", "keep-alive")

  if resp, err := client.Do(request); err != nil {
    t.Errorf("release request failed, reason %s", err.Error())
  } else if resp.StatusCode == 200 {
    if respData, err := ioutil.ReadAll(resp.Body); err != nil && err != io.EOF {
      t.Errorf("read response data failed, reason %s", err.Error())
    } else {
      return respData
    }
  } else {
    t.Errorf("response status not ok %d", resp.StatusCode)
  }
  return nil
}
