package bot

import (
  "viking-trader/model"
  "viking-trader/util"
  "math/rand"
  "math"
)


type Bot struct {
  GameUsername string
  GamePassword string
  BankUsername string
  BankPassword string
  Balance int
  GameItems model.GameItems
  RmtItems model.RmtItems
  Orders []int
}

type Bots []*Bot


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
  rmtItem := CreateRmtItemRequest(bot.GameUsername, gameItemId, bot.BankUsername, price)
  if rmtItem != nil {
    bot.Orders = append(bot.Orders, rmtItem.Id)
  }
  return rmtItem
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

func doIt(p float32) bool {
  n := rand.Intn(100)
  if n < int(p * 100) {
    return true
  }
  return false
}

func (bot *Bot) Transfered (itemId int) bool {
  return TransferedRequest(itemId)
}

func (bot *Bot) OrdersTransfered() model.RmtItems {
  var result model.RmtItems = model.RmtItems{}
  for _, itemId := range bot.Orders {
    item := util.GetItem(itemId)
    if item == nil {
      continue
    }
    if item.Status == model.ItemStatusTransfered {
      result = append(result, *item)
    }
  }
  return result
}



func (bot *Bot) RandomAction() {
  if !bot.FetchGameItem() {
    return
  }
  items := util.GetAllItems()
  for _, gameItem := range bot.GameItems {
    if doIt(0.3) {
      bot.CreateRmtItem(gameItem.Id, int(math.Pow(10, float64(gameItem.Rarity))))
    }
    if items == nil {
      return
    }
  }
  if doIt(0.3) {
    if len(*items) == 0 {
      return
    }
    item := (*items)[rand.Intn(len(*items))]
    if item.OwnerGameUsername == bot.GameUsername {
      return
    }
    if item.Status != model.ItemStatusSale {
      return
    }
    if item.Price > uint64(math.Pow(10, float64(item.Rarity))) {
      return
    }
    transferRequest := bot.Buy(item.Id)
    if transferRequest == nil {
      return
    }
    success := bot.Transfer(transferRequest)
    if !success {
      return
    }
    success = bot.Transfered(item.Id)
    if !success {
      return
    }
  }
}
