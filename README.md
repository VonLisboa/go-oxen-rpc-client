Go Oxen RPC Client
====================

<p align="center">
<img src="https://github.com/vonlisboa/go-oxen-rpc-client/raw/master/media/img/monero_gopher.png" alt="Oxen Gopher" width="200" />
</p>

A client implementation for the Oxen wallet and daemon RPC written in go.
This package is inspired by https://github.com/gabstv/go-monero, https://monero-ecosystem/go-oxen-rpc-client
and https://github.com/monero-ecosystem/go-monero-rpc-client

## Wallet RPC Client

[![GoDoc](https://godoc.org/github.com/vonlisboa/go-oxen-rpc-client/wallet?status.svg)](https://godoc.org/github.com/vonlisboa/go-oxen-rpc-client/wallet)

### Oxen RPC Version
The ```go-oxen-rpc-client/wallet``` package is the RPC client for [Oxen Wallet RPC](https://docs.oxen.io/oxen-docs/using-the-oxen-blockchain/advanced/wallet-rpc-calls).

### Installation

```sh
go get -u github.com/vonlisboa/go-oxen-rpc-client
```

#### Spawn the oxen-wallet-rpc daemon (without rpc login):

```sh
./oxen-wallet-rpc --wallet-file /home/$user/stagenetwallet/stagenetwallet --daemon-address pool.cloudissh.com:38081 --stagenet --rpc-bind-port 6061 --password 'mystagenetwalletpassword' --disable-rpc-login
```
You can use our remote node for the stagenet running at pool.cloudissh.com port `38081`.

#### Go code:

```Go
package main

import (
  "encoding/json"
  "fmt"
  "log"

  "github.com/vonlisboa/go-oxen-rpc-client/wallet"
)

func checkerr(err error) {
  if err != nil {
    log.Panic(err)
  }
}

func main() {
  // Start a wallet client instance
  client := wallet.New(wallet.Config{
    Address: "http://127.0.0.1:6061/json_rpc",
  })

  // check wallet balance
  resp, err := client.GetBalance(&wallet.RequestGetBalance{AccountIndex: 0})
  checkerr(err)
  res, _ := json.MarshalIndent(resp, "", "\t")
  fmt.Print(string(res))

  // get incoming transfers
  resp1, err := client.GetTransfers(&wallet.RequestGetTransfers{
    AccountIndex: 0,
    In:           true,
  })
  checkerr(err)
  for _, in := range resp1.In {
    res, _ := json.MarshalIndent(in, "", "\t")
    fmt.Print(string(res))
  }
}
```

### Spawn the oxen-wallet-rpc daemon (with rpc login):

```sh
./oxen-wallet-rpc --wallet-file /home/$user/stagenetwallet/stagenetwallet --daemon-address pool.cloudissh.com:38081 --stagenet --rpc-bind-port 6061 --password 'mystagenetwalletpassword' --rpc-login test:testpass
```

#### Go code:

```Go
package main

import (
  "encoding/json"
  "fmt"
  "log"

  "github.com/vonlisboa/go-oxen-rpc-client/wallet"
)

func checkerr(err error) {
  if err != nil {
    log.Panic(err)
  }
}

func main() {
  t := httpdigest.New("test", "testpass")

  // Start a wallet client instance
  client := wallet.New(wallet.Config{
    Address: "http://127.0.0.1:6061/json_rpc",
    Transport: t,
  })

  // check wallet balance
  resp, err := client.GetBalance(&wallet.RequestGetBalance{AccountIndex: 0})
  checkerr(err)
  res, _ := json.MarshalIndent(resp, "", "\t")
  fmt.Print(string(res))

  // get incoming transfers
  resp1, err := client.GetTransfers(&wallet.RequestGetTransfers{
    AccountIndex: 0,
    In:           true,
  })
  checkerr(err)
  for _, in := range resp1.In {
    res, _ := json.MarshalIndent(in, "", "\t")
    fmt.Print(string(res))
  }
}
```

# Daemon RPC Client

As of now, only the wallet RPC has been implemented. The daemon RPC will follow very soon.

# Contribution
* You can fork this, extend it and contribute back.
* You can contribute with pull requests.

# Donations
I love OXEN (Loki) and building applications for and on top of Oxen.

You can make me happy by donating Oxen to the following address:

```
LQN2bEMzX4kPJZHhFheboD6gbTnLhzCDQZGrHkuw4JsaEtVWKmD631bLnwoZsppsSSYV9YCy3hjLyURGYLF2cwn3Cn8vgYh
```

# LICENSE
MIT License