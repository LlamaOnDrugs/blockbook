package kittycoin

import (
	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
	"github.com/trezor/blockbook/bchain"
	"github.com/trezor/blockbook/bchain/coins/btc"
)

const (
	// MainnetMagic is mainnet network constant
	MainnetMagic wire.BitcoinNet = 0x1516171c
	// TestnetMagic is testnet network constant
	TestnetMagic wire.BitcoinNet = 0x1213141f
	// RegtestMagic is regtest network constant
	RegtestMagic wire.BitcoinNet = 0xdcb7c1fc
)

var (
	// MainNetParams are parser parameters for mainnet
	MainNetParams chaincfg.Params
	// TestNetParams are parser parameters for testnet
	TestNetParams chaincfg.Params
	// RegtestParams are parser parameters for regtest
	RegtestParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic

	// Address encoding magics
	MainNetParams.PubKeyHashAddrID = []byte{46} // base58 prefix: K
	MainNetParams.ScriptHashAddrID = []byte{85} // base58 prefix: b

	TestNetParams = chaincfg.TestNet3Params
	TestNetParams.Net = TestnetMagic

	// Address encoding magics
	TestNetParams.PubKeyHashAddrID = []byte{153} // base58 prefix: y
	TestNetParams.ScriptHashAddrID = []byte{196}  // base58 prefix: 8 or 9

	RegtestParams = chaincfg.RegressionNetParams
	RegtestParams.Net = RegtestMagic

	// Address encoding magics
	RegtestParams.PubKeyHashAddrID = []byte{140} // base58 prefix: y
	RegtestParams.ScriptHashAddrID = []byte{19}  // base58 prefix: 8 or 9
}

// KittyCoinParser handle
type KittyCoinParser struct {
	*btc.BitcoinParser
	baseparser *bchain.BaseParser
}

// NewKittyCoinParser returns new KittyCoinParser instance
func NewKittyCoinParser(params *chaincfg.Params, c *btc.Configuration) *KittyCoinParser {
	return &KittyCoinParser{
		BitcoinParser: btc.NewBitcoinParser(params, c),
		baseparser:    &bchain.BaseParser{},
	}
}

// GetChainParams contains network parameters for the main KittyCoin network,
// the regression test KittyCoin network, the test KittyCoin network and
// the simulation test KittyCoin network, in this order
func GetChainParams(chain string) *chaincfg.Params {
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err == nil {
			err = chaincfg.Register(&TestNetParams)
		}
		if err == nil {
			err = chaincfg.Register(&RegtestParams)
		}
		if err != nil {
			panic(err)
		}
	}
	switch chain {
	case "test":
		return &TestNetParams
	case "regtest":
		return &RegtestParams
	default:
		return &MainNetParams
	}
}

// PackTx packs transaction to byte array using protobuf
func (p *KittyCoinParser) PackTx(tx *bchain.Tx, height uint32, blockTime int64) ([]byte, error) {
	return p.baseparser.PackTx(tx, height, blockTime)
}

// UnpackTx unpacks transaction from protobuf byte array
func (p *KittyCoinParser) UnpackTx(buf []byte) (*bchain.Tx, uint32, error) {
	return p.baseparser.UnpackTx(buf)
}
