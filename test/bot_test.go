package test

import (
  "testing"
  "viking-trader/bot"
)

func TestBot(t *testing.T) {
  botA := bot.Login (
    "player1",
    "password1",
    "person1",
    "password1",
  )

  botB := bot.Login (
    "player2",
    "password2",
    "person2",
    "password2",
  )

  botC := bot.Login (
    "dummy",
    "dummy",
    "dummy",
    "dummy",
  )

  if botA == nil {
    t.Error("ログインできるはずのアカウントにログインできなかった")
  }
  if botB == nil {
    t.Error("ログインできるはずのアカウントにログインできなかった")
  }
  if botC != nil {
    t.Error("存在しないアカウントにログインできてしまった")
  }
  botABalance := botA.Balance
  botBBalance := botB.Balance
  if len(botA.GameItems) == 0 {
    t.Error("アイテムを取得できていない")
  }
  if len(botB.GameItems) == 0 {
    t.Error("アイテムを取得できていない")
  }
  item := botA.CreateRmtItem(botA.GameItems[0].Id, 100)
  if item == nil {
    t.Error("rmtのitemを作成できなかった")
  }
  itemId := item.Id
  transferRequest := botB.Buy(itemId)
  if transferRequest == nil {
    t.Error("商品を買えなかった")
  }
  success := botB.Transfer(transferRequest)
  if !success {
    t.Error("振込に失敗した")
  }
  botB.FetchBalance()
  if botB.Balance != botBBalance - transferRequest.Amount {
    t.Error("振込んでも残高が適切にかわらなかった")
  }
  success = botA.SendGameItem(item.GameItemId, "rmt")
  if !success {
    t.Error("アイテムを送信できなかった")
  }
  success = botA.Sent(itemId)
  if !success {
    t.Error("送信したことの報告に失敗した")
  }
  success = botB.FetchGameItem()
  if !success {
    t.Error("game itemのfetchに失敗した")
  }
  success = false
  for _, gameItem := range botB.GameItems {
    if gameItem.Id == item.GameItemId {
      success = true
    }
  }
  if !success {
    t.Error("取引が完了してもゲームアイテムが送られてない")
  }
  botA.FetchBalance()
  if botABalance >= botA.Balance {
    t.Error("出品者が報酬をうけとってない")
  }
  t.Log("Bot終了")
}
