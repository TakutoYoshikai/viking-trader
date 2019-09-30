package model

type RmtItem struct {
  Id int
  GameItemId int
  OwnerBankUsername string
  Rarity int
  Name string
  Status int
  Price uint64
  TransferRequest *TransferRequest
  BuyerGameUsername string
}
