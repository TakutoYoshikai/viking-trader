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
      time.Sleep(300 * time.Millisecond)
      for _, b := range bots {
        ordersTransfered := b.OrdersTransfered()
        for _, item := range ordersTransfered {
          success := b.SendGameItem(item.GameItemId, "rmt")
          if !success {
            continue
          }
          success = b.Sent(item.Id)
          if !success {
            continue
          }
        }
        b.RandomAction()
      }
    }
  fmt.Println("end")
}
