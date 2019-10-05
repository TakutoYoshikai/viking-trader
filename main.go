package main

import (
  "viking-trader/bot"
  "strconv"
  "fmt"
  "time"
)

func main() {
  fmt.Println("start")
  bots := bot.Bots{}
    for i := 0; i < 100; i++ {
      b := bot.Login("player" + strconv.Itoa(i), "password" + strconv.Itoa(i), "person" + strconv.Itoa(i), "password" + strconv.Itoa(i))
      if b == nil {
        fmt.Println("ログインできなかった")
        return
      }
      bots = append(bots, b)
    }
    for i := 0; i < 50; i++ {
      for _, b := range bots {
        time.Sleep(300 * time.Millisecond)
        ordersTransfered := b.OrdersTransfered()
        remainOrders := []int{}
        for _, item := range ordersTransfered {
          success := b.SendGameItem(item.GameItemId, "rmt")
          if !success {
            remainOrders = append(remainOrders, item.Id)
            continue
          }
          success = b.Sent(item.Id)
          if !success {
            remainOrders = append(remainOrders, item.Id)
            continue
          }
        }
        b.Orders = remainOrders
        b.RandomAction()
      }
    }
  fmt.Println("end")
}
