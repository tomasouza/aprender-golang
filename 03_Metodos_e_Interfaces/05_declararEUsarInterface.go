package main

import "fmt"

type BlockchainAdapter interface {
	gasPrice() GasPriceResponse
}

type GasPriceResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  string `json:"result"`
}

func ConsultarGAS(blockchainAdapter BlockchainAdapter) GasPriceResponse {
	return blockchainAdapter.gasPrice()
}

type BlockchainAdapterMock struct {
}

func (mock *BlockchainAdapterMock) gasPrice() GasPriceResponse {
	return GasPriceResponse{
		Jsonrpc: "2.0",
		Id:      53,
		Result:  "0x3e8",
	}
}

func main() {
	mock := &BlockchainAdapterMock{}
	result := ConsultarGAS(mock)
	fmt.Println("Gas price:", result)
}
