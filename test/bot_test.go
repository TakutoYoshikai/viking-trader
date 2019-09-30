package test

import (
  "testing"
  "viking-trader"
)

func TestBot(t *testing.T) {
  botA := main.Login {
    GameUsername: "player1",
    GamePassword: "password1",
    BankUsername: "person1",
    BankPassword: "password1",
  }

  botB := main.Login {
    GameUsername: "player2",
    GamePassword: "password2",
    BankUsername: "person2",
    BankPassword: "password2",
  }

  botC := main.Login {
    GameUsername: "dummy",
    GamePassword: "dummy",
    BankUsername: "dummy",
    BankPassword: "dummy",
  }

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
  if botABalance == nil {
    t.Error("銀行の残高が取得できなかった")
  }
  botBBalance := botB.Balance
  if botBBalance == nil {
    t.Error("銀行の残高が取得できなかった")
  }

  item := bot.CreateRmtItem(gameItemId string, price int)
  if item == nil {
    t.Error("rmtのitemを作成できなかった")
  }
  itemId := item.Id
  transferRequest := botA.Buy(itemId)
  if transferRequest != nil {
    t.Error("自分の商品を自分で買えてしまった")
  }
  transferRequest = botB.Buy(itemId)
  if transferRequest == nil {
    t.Error("商品を買えなかった")
  }
  bankAccountB := botB.Transfer(transferRequest)
  if bankAccountB.Balance != botBBalance - botBtransferRequest.Amount {
    t.Error("振込んでも残高が適切にかわらなかった")
  }
  if bankAccountB.Balance != botB.Balance {
    t.Error("最新のBalanceが反映されていない")
  }
  success := botA.SendGameItem(item.GameItemId, "rmt")
  if !success {
    t.Error("アイテムを送信できなかった")
  }
  success = botA.Sent(itemId)
  if !success {
    t.Error("送信したことの報告に失敗した")
  }
  gameItems := botB.FetchGameItem()
  success = false
  for i, gameItem := range gameItems {
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
