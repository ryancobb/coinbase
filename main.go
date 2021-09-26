package main

import (
  "os"
	coinbasepro "github.com/preichenberger/go-coinbasepro/v2"
)

func main() {
  os.Setenv("COINBASE_PRO_SANDBOX", "1")

  client := coinbasepro.NewClient()

  client.UpdateConfig(&coinbasepro.ClientConfig{
    BaseURL: "https://api.pro.coinbase.com",
    Key: os.Getenv("COINBASE_KEY"),
    Passphrase: os.Getenv("COINBASE_PASSPHRASE"),
    Secret: os.Getenv("COINBASE_SECRET"),
  })

  accounts, err := client.GetAccounts()
  if err != nil {
    println(err.Error())
  }

  for _, a := range accounts {
    println(a.Balance)
  }
}
 
