package test

import (
  "viking-trader/util"
  "viking-trader/bot"
  "testing"
)

func TestUtil(t *testing.T) {
  botA := bot.Login (
    "player4",
    "password4",
    "person4",
    "password4",
  )
  if botA == nil {
    t.Error("ログインできるはずのアカウントにログインできなかった")
  }
  item := botA.CreateRmtItem(botA.GameItems[0].Id, 100)
  if item == nil {
    t.Error("rmtのitemを作成できなかった")
  }
  items := util.GetAllItems()
  if items == nil {
    t.Error("アイテムが取得できなかった")
  }
  if len(*items) == 0 {
    t.Error("itemsの取得が適切でない（個数がゼロだった）")
  }
  success := false
  for _, itm := range *items {
    if itm.Id == item.Id {
      success = true
    }
  }
  if !success {
    t.Error("作成したはずのitemがitemsにない")
  }
  t.Log("util終了")
}
