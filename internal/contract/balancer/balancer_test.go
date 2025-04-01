package balancer

import (
	"alpha-explorer/common/config"
	"alpha-explorer/eth"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestBalancerFilterer_FilterSwap(t *testing.T) {
	config.Init("", "dev")
	eth.InitEthClient()
	vault, err := NewBalancer(common.HexToAddress("0xBA12222222228d8Ba445958a75a0704d566BF2C8"), eth.EthClient.Client)
	if err != nil {
		t.Fatal(err)
	}
	iterator, err := vault.FilterSwap(&bind.FilterOpts{
		Start:   16194258,
		End:     nil,
		Context: nil,
	}, nil, nil, nil)
	if err != nil {
		return
	}
	for iterator.Next() {
		fmt.Println(hex.EncodeToString(iterator.Event.PoolId[:]), iterator.Event.TokenIn, iterator.Event.TokenOut, iterator.Event.AmountIn, iterator.Event.AmountOut)
	}
}
