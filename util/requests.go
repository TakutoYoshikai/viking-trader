package util

import (
  "net/http"
  "encoding/json"
  "io/ioutil"
  "viking-trader/config"
  "viking-trader/model"
)

func GetAllItems() *[]model.RmtItem {
  url := config.RmtHost + "/items"
  res, err := http.Get(url)
  if err != nil || res.StatusCode != 200 {
    return nil
  }
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return nil
  }
  jsonBytes := ([]byte)(body)
  items := new([]model.RmtItem)
  err = json.Unmarshal(jsonBytes, items)
  if err != nil {
    return nil
  }
  return items
}
