package model

type TransferRequest struct {
  Id int
  From string
  To string
  Amount int
  Transfered bool
}
