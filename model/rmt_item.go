package model
import (
  "time"
)

type RmtItem struct {
  Id int
  GameItemId int
  OwnerBankUsername string
  OwnerGameUsername string
  Rarity int
  Name string
  Status int
  Price uint64
  TransferRequest *TransferRequest
  BuyerGameUsername string
  CreatedAt time.Time
}

type RmtItems []RmtItem
const (
  ItemStatusSale = iota
  ItemStatusOrdered
  ItemStatusTransfered
  ItemStatusSentItem
  ItemStatusCompleted
)
