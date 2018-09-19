# getho-go
[getho.io](https://getho.io) api client written in golang

## Installation
To install `getho-go`, simply execute the following command in a terminal from your `$GOPATH`:

```
go get github.com/popshootjapan/getho-go
```

## How to use

```
gethoClient := getho.New("<YOUR DOMAIN>", nil)

input := &getho.EthGetBalanceInput{
    Address:        "0x5c66b0d82df26e8FE165Be6628F5f5e1f1bccD5C",
    BlockParameter: getho.NewBlockParameter("latest"),
}

req, resp := gethoClient.EthGetBalanceRequest(input)
_ = req.Call()

fmt.Println(res.Result)
// => 0x7759d50972e9800
```

## Test
**NOTE**: Please DO NOT call a lot of requests.

```
$ go test *_test.go -v
```

## TODO
- [ ] Error handling
- [ ] Decode all result parameters

## Supported JSON-RPC methods

- `web3_clientVersion`
- `net_version`
- `net_listening`
- `net_peerCount`
- `eth_gasPrice`
- `eth_blockNumber`
- `eth_getBalance`
- `eth_getTransactionCount`
- `eth_getBlockTransactionCountByHash`
- `eth_getBlockTransactionCountByNumber`
- `eth_sendRawTransaction`
- `eth_call`
- `eth_estimateGas`
- `eth_getBlockByHash`
- `eth_getBlockByNumber`
- `eth_getTransactionByHash`
- `eth_getTransactionReceipt`

## License
`getho-go` is available under the MIT license. See the LICENSE file for more info.
