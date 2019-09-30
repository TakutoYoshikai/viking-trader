package bot

import (
  "viking-trader/model"
)


type Bot struct {
  GameUsername string
  GamePassword string
  BankUsername string
  BankPassword string
  Balance int
  GameItems model.GameItems
}


func (bot *Bot) FetchBalance() bool {
  bankAccount := BankLoginRequest(bot.BankUsername, bot.BankPassword)
  if bankAccount == nil {
    return false
  }
  bot.Balance = bankAccount.Balance
  return true
}

func (bot *Bot) FetchGameItem() bool {
  gameAccount := GameLoginRequest(bot.GameUsername, bot.GamePassword)
  if gameAccount == nil {
    return false
  }
  bot.GameItems = gameAccount.Items
  return true
}

func (bot *Bot) CreateRmtItem(gameItemId int, price int) *model.RmtItem {
  return CreateRmtItemRequest(bot.GameUsername, gameItemId, bot.BankUsername, price)
}

func (bot *Bot) Buy(itemId int) *model.TransferRequest {
  return BuyRequest(itemId, bot.BankUsername, bot.GameUsername)
}

func (bot *Bot) Transfer(transferRequest *model.TransferRequest) bool {
  return TransferReq(bot.BankUsername, bot.BankPassword, transferRequest.Id)
}

func (bot *Bot) SendGameItem(gameItemId int, to string) bool {
  success := SendGameItemRequest(bot.GameUsername, bot.GamePassword, gameItemId, to)
  if success {
    for i, gameItem := range bot.GameItems {
      if gameItem.Id == gameItemId {
        bot.GameItems = append(bot.GameItems[:i], bot.GameItems[i+1:]...)
        return success
      }
    }
  }
  return success
}

func (bot *Bot) Sent(itemId int) bool {
  return SentRequest(itemId)
}


func Login(gameUsername string, gamePassword string, bankUsername string, bankPassword string) *Bot {
  bankAccount := BankLoginRequest(bankUsername, bankPassword)
  if bankAccount == nil {
    return nil
  }
  gameAccount := GameLoginRequest(gameUsername, gamePassword)
  if gameAccount == nil {
    return nil
  }
  return &Bot {
    GameUsername: gameUsername,
    GamePassword: gamePassword,
    BankUsername: bankUsername,
    BankPassword: bankPassword,
    Balance: bankAccount.Balance,
    GameItems: gameAccount.Items,
  }
}
