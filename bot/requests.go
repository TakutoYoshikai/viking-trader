package bot

import (
  "io/ioutil"
  "net/http"
  "encoding/json"
  "strconv"
  "viking-trader/model"
  "viking-trader/config"
)

func BankLoginRequest(bankUsername string, bankPassword string) *model.BankAccount {
  url := config.BankHost + "/users/" + bankUsername + "/" + bankPassword
  res, err := http.Get(url)
  if err != nil {
    return nil
  }
  defer res.Body.Close()
  if res.StatusCode != 200 {
    return nil
  }
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return nil
  }
  jsonBytes := ([]byte)(body)
  bankAccount := new(model.BankAccount)
  err = json.Unmarshal(jsonBytes, bankAccount)
  if err != nil {
    return nil
  }
  return bankAccount
}

func GameLoginRequest(gameUsername string, gamePassword string) *model.GameAccount {
  url := config.GameHost + "/users/" + gameUsername + "/" + gamePassword
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
  gameAccount := new(model.GameAccount)
  err = json.Unmarshal(jsonBytes, gameAccount)
  if err != nil {
    return nil
  }
  return gameAccount
}

func CreateRmtItemRequest(gameUsername string, gameItemId int, bankUsername string, price int) *model.RmtItem {
  url := config.RmtHost + "/item/create/" + bankUsername + "/" + gameUsername + "/" + strconv.Itoa(gameItemId) + "/" + strconv.Itoa(price)
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
  rmtItem := new(model.RmtItem)
  err = json.Unmarshal(jsonBytes, rmtItem)
  if err != nil {
    return nil
  }
  return rmtItem

}

func BuyRequest(itemId int, bankUsername string, gameUsername string) *model.TransferRequest {
  url := config.RmtHost + "/item/buy/" + strconv.Itoa(itemId) + "/" + bankUsername + "/" + gameUsername
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
  transferRequest := new(model.TransferRequest)
  err = json.Unmarshal(jsonBytes, transferRequest)
  if err != nil {
    return nil
  }
  return transferRequest
}


func SendGameItemRequest(gameUsername string, gamePassword string, gameItemId int, to string) bool {
  url := config.GameHost + "/send/" + gameUsername + "/" + gamePassword + "/" + strconv.Itoa(gameItemId) + "/" + to
  res, err := http.Get(url)
  if err != nil || res.StatusCode != 200 {
    return false
  }
  defer res.Body.Close()
  return true
}


func SentRequest(itemId int) bool {
  url := config.RmtHost + "/item/sent/" + strconv.Itoa(itemId)
  res, err := http.Get(url)
  if err != nil || res.StatusCode != 200 {
    return false
  }
  defer res.Body.Close()
  return true
}

func TransferReq(bankUsername string, bankPassword string, requestId int) bool {
  url := config.BankHost + "/requests/transfer/" + bankUsername + "/" + bankPassword + "/" + strconv.Itoa(requestId)
  res, err := http.Get(url)
  if err != nil || res.StatusCode != 200 {
    return false
  }
  defer res.Body.Close()
  return true
}

func TransferedRequest(itemId int) bool{
  url := config.RmtHost + "/item/transfered/" + strconv.Itoa(itemId)
  res, err := http.Get(url)
  if err != nil || res.StatusCode != 200 {
    return false
  }
  defer res.Body.Close()
  return true
}
