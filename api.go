package getho

type input interface {
	method() string
	params() interface{}
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

// Web3ClientVersionRequest generates a request which calls JSON-RPC method of "web3_clientVersion" for go-ethereum's RPC API.
//
// 	Example call a request using the Web3ClientVersionRequest method.
//
// 	req, resp := cli.Web3ClientVersionRequest(&getho.Web3ClientVersionInput{})
//  err := req.Call()
//  if err != nil {
//    fmt.Println(err)
//	}
//  fmt.Println(resp)
//
// See also, https://github.com/ethereum/wiki/wiki/JSON-RPC#web3_clientversion
func (g *Getho) Web3ClientVersionRequest(input *Web3ClientVersionInput) (req *Request, output *Web3ClientVersionOutput) {
	if input != nil {
		input = &Web3ClientVersionInput{}
	}

	output = &Web3ClientVersionOutput{}
	req = g.newRequest(input, output)
	return
}

type Web3ClientVersionInput struct {
}

func (g *Web3ClientVersionInput) method() string {
	return "web3_clientVersion"
}

func (g *Web3ClientVersionInput) params() interface{} {
	return []string{}
}

type Web3ClientVersionOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

// NetVersionRequest generates a request which calls JSON-RPC method of "net_version" for go-ethereum's RPC API.
//
// 	Example call a request using the NetVersionRequest method.
//
// 	req, resp := cli.NetVersionRequest(&getho.NetVersionInput{})
//  err := req.Call()
//  if err != nil {
//    fmt.Println(err)
//	}
//  fmt.Println(resp)
//
// See also, https://github.com/ethereum/wiki/wiki/JSON-RPC#net_version
func (g *Getho) NetVersionRequest(input *NetVersionInput) (req *Request, output *NetVersionOutput) {
	if input == nil {
		input = &NetVersionInput{}
	}

	output = &NetVersionOutput{}
	req = g.newRequest(input, output)
	return
}

type NetVersionInput struct {
}

func (g *NetVersionInput) method() string {
	return "net_version"
}

func (g *NetVersionInput) params() interface{} {
	return []string{}
}

type NetVersionOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

// NetListeningRequest generates a request which calls JSON-RPC method of "net_listening" for go-ethereum's RPC API.
//
// 	Example call a request using the NetListeningRequest method.
//
// 	req, resp := cli.NetListeningRequest(&getho.NetListeningInput{})
//  err := req.Call()
//  if err != nil {
//    fmt.Println(err)
//	}
//  fmt.Println(resp)
//
// See also, https://github.com/ethereum/wiki/wiki/JSON-RPC#net_listening
func (g *Getho) NetListeningRequest(input *NetListeningInput) (req *Request, output *NetListeningOutput) {
	if input == nil {
		input = &NetListeningInput{}
	}

	output = &NetListeningOutput{}
	req = g.newRequest(input, output)
	return
}

type NetListeningInput struct {
}

func (g *NetListeningInput) method() string {
	return "net_listening"
}

func (g *NetListeningInput) params() interface{} {
	return []string{}
}

type NetListeningOutput struct {
	Error  *Error `json:"error"`
	Result bool   `json:"result"`
}

// NetPeerCountRequest generates a request which calls JSON-RPC method of "net_peerCount" for go-ethereum's RPC API.
//
// 	Example call a request using the NetPeerCountRequest method.
//
// 	req, resp := cli.NetPeerCountRequest(&getho.NetPeerCountInput{})
//  err := req.Call()
//  if err != nil {
//    fmt.Println(err)
//	}
//  fmt.Println(resp)
//
// See also, https://github.com/ethereum/wiki/wiki/JSON-RPC#net_peercount
func (g *Getho) NetPeerCountRequest(input *NetPeerCountInput) (req *Request, output *NetPeerCountOutput) {
	if input == nil {
		input = &NetPeerCountInput{}
	}

	output = &NetPeerCountOutput{}
	req = g.newRequest(input, output)
	return
}

type NetPeerCountInput struct {
}

func (g *NetPeerCountInput) method() string {
	return "net_peerCount"
}

func (g *NetPeerCountInput) params() interface{} {
	return []string{}
}

type NetPeerCountOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

// EthGasPriceRequest generates a request which calls JSON-RPC method of "eth_gasPrice" for go-ethereum's RPC API.
//
// 	Example call a request using the EthGasPriceRequest method.
//
// 	req, resp := cli.EthGasPriceRequest(&getho.EthGasPriceInput{})
//  err := req.Call()
//  if err != nil {
//    fmt.Println(err)
//	}
//  fmt.Println(resp)
//
// See also, https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_gasprice
func (g *Getho) EthGasPriceRequest(input *EthGasPriceInput) (req *Request, output *EthGasPriceOutput) {
	if input != nil {
		input = &EthGasPriceInput{}
	}

	output = &EthGasPriceOutput{}
	req = g.newRequest(input, output)
	return
}

type EthGasPriceInput struct {
}

func (i *EthGasPriceInput) method() string {
	return "eth_gasPrice"
}

func (i *EthGasPriceInput) params() interface{} {
	return []string{}
}

type EthGasPriceOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

// EthBlockNumberRequest generates a request which calls JSON-RPC method of "eth_blockNumber" for go-ethereum's RPC API.
//
// 	Example call a request using the EthBlockNumberRequest method.
//
// 	req, resp := cli.EthBlockNumberRequest(&getho.EthBlockNumberInput{})
//  err := req.Call()
//  if err != nil {
//    fmt.Println(err)
//	}
//  fmt.Println(resp)
//
// See also, https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_blocknumber
func (g *Getho) EthBlockNumberRequest(input *EthBlockNumberInput) (req *Request, output *EthBlockNumberOutput) {
	if input != nil {
		input = &EthBlockNumberInput{}
	}

	output = &EthBlockNumberOutput{}
	req = g.newRequest(input, output)
	return
}

type EthBlockNumberInput struct {
}

func (g *EthBlockNumberInput) method() string {
	return "eth_blockNumber"
}

func (g *EthBlockNumberInput) params() interface{} {
	return []string{}
}

type EthBlockNumberOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

// EthGetBalanceRequest generates a request which calls JSON-RPC method of "eth_getBalance" for go-ethereum's RPC API.
//
// 	Example call a request using the EthGetBalanceRequest method.
//
// 	req, resp := cli.EthGetBalanceRequest(&getho.EthGetBalanceInput{Address: "0x5c66b0d82df26e8FE165Be6628F5f5e1f1bccD5C", BlockParameter: getho.NewBlockParameter("latest")})
//  err := req.Call()
//  if err != nil {
//    fmt.Println(err)
//	}
//  fmt.Println(resp)
//
// See also, https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getbalance
func (g *Getho) EthGetBalanceRequest(input *EthGetBalanceInput) (req *Request, output *EthGetBalanceOutput) {
	if input == nil {
		input = &EthGetBalanceInput{}
	}

	output = &EthGetBalanceOutput{}
	req = g.newRequest(input, output)
	return
}

type EthGetBalanceInput struct {
	Address        string
	BlockParameter BlockParameter
}

func (g *EthGetBalanceInput) method() string {
	return "eth_getBalance"
}

func (g *EthGetBalanceInput) params() interface{} {
	return []string{g.Address, g.BlockParameter.Value}
}

type EthGetBalanceOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

// EthGetTransactionCountRequest generates a request which calls JSON-RPC method of "eth_getTransactionCount" for go-ethereum's RPC API.
//
// 	Example call a request using the EthGetTransactionCountRequest method.
//
// 	req, resp := cli.EthGetTransactionCountRequest(&getho.EthGetTransactionCountInput{Address: "0x5c66b0d82df26e8FE165Be6628F5f5e1f1bccD5C", BlockParameter: getho.NewBlockParameter("latest")})
//  err := req.Call()
//  if err != nil {
//    fmt.Println(err)
//	}
//  fmt.Println(resp)
//
// See also, https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_gettransactioncount
func (g *Getho) EthGetTransactionCountRequest(input *EthGetTransactionCountInput) (req *Request, output *EthGetTransactionCountOutput) {
	if input == nil {
		input = &EthGetTransactionCountInput{}
	}

	output = &EthGetTransactionCountOutput{}
	req = g.newRequest(input, output)
	return
}

type EthGetTransactionCountInput struct {
	Address        string
	BlockParameter BlockParameter
}

func (g *EthGetTransactionCountInput) method() string {
	return "eth_getTransactionCount"
}

func (g *EthGetTransactionCountInput) params() interface{} {
	return []string{g.Address, g.BlockParameter.Value}
}

type EthGetTransactionCountOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

// EthGetBlockTransactionCountByHashRequest generates a request which calls JSON-RPC method of "eth_getBlockTransactionCountByHash" for go-ethereum's RPC API.
//
// 	Example call a request using the  method.
//
//	req, resp := cli.EthGetBlockTransactionCountByHashRequest(&getho.EthGetBlockTransactionCountByHashInput{Hash: "0x873d463378e59d8d1e08d5dda12624784a535e96103a9cda04de21735d53ab7b"})
//  err := req.Call()
//  if err != nil {
//    fmt.Println(err)
//	}
//  fmt.Println(resp)
//
// See also, https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getblocktransactioncountbyhash
func (g *Getho) EthGetBlockTransactionCountByHashRequest(input *EthGetBlockTransactionCountByHashInput) (req *Request, output *EthGetBlockTransactionCountByHashOutput) {
	if input == nil {
		input = &EthGetBlockTransactionCountByHashInput{}
	}

	output = &EthGetBlockTransactionCountByHashOutput{}
	req = g.newRequest(input, output)
	return
}

type EthGetBlockTransactionCountByHashInput struct {
	Hash string
}

func (i *EthGetBlockTransactionCountByHashInput) method() string {
	return "eth_getBlockTransactionCountByHash"
}

func (i *EthGetBlockTransactionCountByHashInput) params() interface{} {
	return []string{i.Hash}
}

type EthGetBlockTransactionCountByHashOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

//  EthGetBlockTransactionCountByNumberRequest generates a request which calls JSON-RPC method of "eth_getBlockTransactionCountByNumber" for go-ethereum's RPC API.
//
// 	Example call a request using the  method.
//
// 	req, resp := cli.EthGetBlockTransactionCountByNumberRequest(&getho.EthGetBlockTransactionCountByNumberInput{BlockParameter: getho.NewBlockParameterWithNumber(3622371)})
//  err := req.Call()
//  if err != nil {
//    fmt.Println(err)
//	}
//  fmt.Println(resp)
//
// See also, https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getblocktransactioncountbynumber
func (g *Getho) EthGetBlockTransactionCountByNumberRequest(input *EthGetBlockTransactionCountByNumberInput) (req *Request, output *EthGetBlockTransactionCountByNumberOutput) {
	if input == nil {
		input = &EthGetBlockTransactionCountByNumberInput{}
	}

	output = &EthGetBlockTransactionCountByNumberOutput{}
	req = g.newRequest(input, output)
	return
}

type EthGetBlockTransactionCountByNumberInput struct {
	BlockParameter BlockParameter
}

func (g *EthGetBlockTransactionCountByNumberInput) method() string {
	return "eth_getBlockTransactionCountByNumber"
}

func (g *EthGetBlockTransactionCountByNumberInput) params() interface{} {
	return []string{g.BlockParameter.Value}
}

type EthGetBlockTransactionCountByNumberOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

//  EthSendTransactionRequest generates a request which calls JSON-RPC method of "eth_sendTransaction" for go-ethereum's RPC API.
//
// 	Example call a request using the  method.
//
// 	req, resp := cli.EthSendTransactionRequest(&getho.EthSendTransactionInput{
//		Transaction: getho.Transaction{
//  		From: "0xb60e8dd61c5d32be8058bb8eb970870f07233155",
//  		To: "0xd46e8dd67c5d32be8058bb8eb970870f07244567",
//  		Gas: "0x76c0", // 30400
//  		GasPrice: "0x9184e72a000", // 10000000000000
//  		Value: "0x9184e72a", // 2441406250
//  		Data: "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"})
//		}
//	}
//  err := req.Call()
//  if err != nil {
//    fmt.Println(err)
//	}
//  fmt.Println(resp)
//
// See also, https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_sendtransaction
func (g *Getho) EthSendTransactionRequest(input *EthSendTransactionInput) (req *Request, output *EthSendTransactionOutput) {
	if input == nil {
		input = &EthSendTransactionInput{}
	}

	output = &EthSendTransactionOutput{}
	req = g.newRequest(input, output)
	return
}

type EthSendTransactionInput struct {
	Transaction Transaction
}

func (g *EthSendTransactionInput) method() string {
	return "eth_sendTransaction"
}

func (g *EthSendTransactionInput) params() interface{} {
	return []interface{}{g.Transaction.toMap()}
}

type EthSendTransactionOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

//  EthSendRawTransactionRequest generates a request which calls JSON-RPC method of "eth_sendRawTransaction" for go-ethereum's RPC API.
//
// 	Example call a request using the  method.
//
// 	req, resp := cli.EthSendRawTransactionRequest(&getho.EthSendRawTransactionInput{Data: "0xf86b0185037e11d600825208945c66b0d82df26e8fe165be6628f5f5e1f1bccd5c87037c36dc865000802aa009fa4bca67b16d301eab882acff84d8025d97175d646289fb94f14765fade2f0a00abedcbb522da03e89124394bf56e15a6c246cc046c50a1b5a6e21b04bdeca63"})
//  err := req.Call()
//  if err != nil {
//    fmt.Println(err)
//	}
//  fmt.Println(resp)
//
// See also, https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_sendrawtransaction
func (g *Getho) EthSendRawTransactionRequest(input *EthSendRawTransactionInput) (req *Request, output *EthSendRawTransactionOutput) {
	if input == nil {
		input = &EthSendRawTransactionInput{}
	}

	output = &EthSendRawTransactionOutput{}
	req = g.newRequest(input, output)
	return
}

type EthSendRawTransactionInput struct {
	Data string
}

func (g *EthSendRawTransactionInput) method() string {
	return "eth_sendRawTransaction"
}

func (g *EthSendRawTransactionInput) params() interface{} {
	return []string{g.Data}
}

type EthSendRawTransactionOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

//  EthCall generates a request which calls JSON-RPC method of "eth_Call" for go-ethereum's RPC API.
//
// 	Example call a request using the  method.
//
// 	req, resp := cli.EthCallRequest(&getho.EthCallInput{Transaction: &getho.Transaction{}, BlockParamter: getho.NewBlockParameter("latest")})
//  err := req.Call()
//  if err != nil {
//    fmt.Println(err)
//	}
//  fmt.Println(resp)
//
// See also, https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_call
func (g *Getho) EthCallRequest(input *EthCallInput) (req *Request, output *EthCallOutput) {
	if input == nil {
		input = &EthCallInput{}
	}

	output = &EthCallOutput{}
	req = g.newRequest(input, output)
	return
}

type EthCallInput struct {
	Transaction    Transaction
	BlockParameter BlockParameter
}

func (g *EthCallInput) method() string {
	return "eth_call"
}

func (g *EthCallInput) params() interface{} {
	return []interface{}{g.Transaction.toMap(), g.BlockParameter.Value}
}

type EthCallOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

//  EthEstimateGas generates a request which calls JSON-RPC method of "eth_estimateGas" for go-ethereum's RPC API.
//
// 	Example call a request using the  method.
//
// 	req, resp := cli.EthEstimateGasRequest(&getho.EthEstimateGasInput{Transaction: &getho.Transaction{}})
//  err := req.Call()
//  if err != nil {
//    fmt.Println(err)
//	}
//  fmt.Println(resp)
//
// See also, https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_estimategas
func (g *Getho) EthEstimateGasRequest(input *EthEstimateGasInput) (req *Request, output *EthEstimateGasOutput) {
	if input == nil {
		input = &EthEstimateGasInput{}
	}

	output = &EthEstimateGasOutput{}
	req = g.newRequest(input, output)
	return
}

type EthEstimateGasInput struct {
	Transaction Transaction
}

func (g *EthEstimateGasInput) method() string {
	return "eth_estimateGas"
}

func (g *EthEstimateGasInput) params() interface{} {
	return []interface{}{g.Transaction.toMap()}
}

type EthEstimateGasOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

//  EthGetBlockByHash generates a request which calls JSON-RPC method of "eth_getBlockByHash" for go-ethereum's RPC API.
//
// 	Example call a request using the  method.
//
// 	req, resp := cli.EthGetBlockByHashRequest(&getho.EthGetBlockByHashInput{Hash: "0x873d463378e59d8d1e08d5dda12624784a535e96103a9cda04de21735d53ab7b", IsFull: true})
//  err := req.Call()
//  if err != nil {
//    fmt.Println(err)
//	}
//  fmt.Println(resp)
//
// See also, https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getblockbyhash
func (g *Getho) EthGetBlockByHashRequest(input *EthGetBlockByHashInput) (req *Request, output *EthGetBlockByHashOutput) {
	if input == nil {
		input = &EthGetBlockByHashInput{}
	}

	output = &EthGetBlockByHashOutput{}
	req = g.newRequest(input, output)
	return
}

type EthGetBlockByHashInput struct {
	Hash   string
	IsFull bool
}

func (g *EthGetBlockByHashInput) method() string {
	return "eth_getBlockByHash"
}

func (g *EthGetBlockByHashInput) params() interface{} {
	return []interface{}{g.Hash, g.IsFull}
}

type EthGetBlockByHashOutput struct {
	Error  *Error `json:"error"`
	Result *Block `json:"result"`
}

//  EthGetBlockByNumber generates a request which calls JSON-RPC method of "eth_ethGetBlockByNumber" for go-ethereum's RPC API.
//
// 	Example call a request using the  method.
//
//	req, resp := cli.EthGetBlockByNumberRequest(&getho.EthGetBlockByNumberInput{BlockParameter: getho.NewBlockParameterWithNumber(3622371), IsFull: true})
//  err := req.Call()
//  if err != nil {
//    fmt.Println(err)
//	}
//  fmt.Println(resp)
//
// See also, https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getblockbynumber
func (g *Getho) EthGetBlockByNumberRequest(input *EthGetBlockByNumberInput) (req *Request, output *EthGetBlockByNumberOutput) {
	if input == nil {
		input = &EthGetBlockByNumberInput{}
	}

	output = &EthGetBlockByNumberOutput{}
	req = g.newRequest(input, output)
	return
}

type EthGetBlockByNumberInput struct {
	BlockParameter BlockParameter
	IsFull         bool
}

func (g *EthGetBlockByNumberInput) method() string {
	return "eth_getBlockByNumber"
}

func (g *EthGetBlockByNumberInput) params() interface{} {
	return []interface{}{g.BlockParameter.Value, g.IsFull}
}

type EthGetBlockByNumberOutput struct {
	Error  *Error `json:"error"`
	Result *Block `json:"result"`
}

//  EthGetTransactionByHash generates a request which calls JSON-RPC method of "eth_getTransactionByHash" for go-ethereum's RPC API.
//
// 	Example call a request using the  method.
//
// 	req, resp := cli.EthGetTransactionByHashRequest(&getho.EthGetTransactionByHashInput{Hash: "0x5a774f4da8a20a53e67b69565e2cb50b62feb5d0731a6f29f158e38e24ae3ed7"})
//  err := req.Call()
//  if err != nil {
//    fmt.Println(err)
//	}
//  fmt.Println(resp)
//
// See also, https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_gettransactionbyhash
func (g *Getho) EthGetTransactionByHashRequest(input *EthGetTransactionByHashInput) (req *Request, output *EthGetTransactionByHashOutput) {
	if input == nil {
		input = &EthGetTransactionByHashInput{}
	}

	output = &EthGetTransactionByHashOutput{}
	req = g.newRequest(input, output)
	return
}

type EthGetTransactionByHashInput struct {
	Hash string
}

func (g *EthGetTransactionByHashInput) method() string {
	return "eth_getTransactionByHash"
}

func (g *EthGetTransactionByHashInput) params() interface{} {
	return []interface{}{g.Hash}
}

type EthGetTransactionByHashOutput struct {
	Error  *Error          `json:"error"`
	Result *RawTransaction `json:"result"`
}

//  EthGetTransactionReceipt generates a request which calls JSON-RPC method of "eth_getTransactionReceipt" for go-ethereum's RPC API.
//
// 	Example call a request using the  method.
//
// 	req, resp := cli.EthGetTransactionReceiptRequest(&getho.EthGetTransactionReceiptInput{Hash: "0x5a774f4da8a20a53e67b69565e2cb50b62feb5d0731a6f29f158e38e24ae3ed7"})
//  err := req.Call()
//  if err != nil {
//    fmt.Println(err)
//	}
//  fmt.Println(resp)
//
// See also, https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_gettransactionreceipt
func (g *Getho) EthGetTransactionReceiptRequest(input *EthGetTransactionReceiptInput) (req *Request, output *EthGetTransactionReceiptOutput) {
	if input == nil {
		input = &EthGetTransactionReceiptInput{}
	}

	output = &EthGetTransactionReceiptOutput{}
	req = g.newRequest(input, output)
	return
}

type EthGetTransactionReceiptInput struct {
	Hash string
}

func (g *EthGetTransactionReceiptInput) method() string {
	return "eth_getTransactionReceipt"
}

func (g *EthGetTransactionReceiptInput) params() interface{} {
	return []string{g.Hash}
}

type EthGetTransactionReceiptOutput struct {
	Error  *Error              `json:"error"`
	Result *TransactionReceipt `json:"result"`
}
