package model

type RmtItem struct {
  Id int
  GameItemId int
  OwnerBankUsername string
  Rarity int
  Name string
  IsBought bool
  Price uint64
  TransferRequest *TransferRequest
  BuyerGameUsername string
}
