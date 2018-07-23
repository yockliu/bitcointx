package bitcointx

import (
	"crypto/sha256"

	. "github.com/yockliu/bitcoinlib"
)

// Coinbase coinbase tx
type Coinbase struct {
	Cell
	Version  uint32
	HashID   *HashCode
	Height   uint32
	in       *TXIn
	Out      *TXOut
	lockTime uint32
}

// NewCoinbase generate coinbase instance
func NewCoinbase(address string, height uint32, value uint64) *Coinbase {
	coinbase := Coinbase{}

	coinbase.Version = 2
	coinbase.Height = height

	outpoint := Outpoint{}
	outpoint.TxHash = HashCode{}
	outpoint.N = 0xFFFFFFFF
	in := TXIn{}
	in.outpoint = &outpoint
	coinbase.in = &in

	out := NewTXOut(value, address)
	coinbase.Out = out

	coinbase.lockTime = 0

	// fmt.Printf("New Coinbase: %x\n", coinbase.Hash())

	return &coinbase
}

// Serialize Serializable to byte array
func (coinbase *Coinbase) Serialize() []byte {
	data := []byte{}

	version := Uint32ToBytes(coinbase.Version)
	inputCount := CompactSizeUint{Value: 1}
	inputOutpointHash := coinbase.in.outpoint.TxHash
	inputOutpointN := Uint32ToBytes(coinbase.in.outpoint.N)

	coinbaseData := coinbase.serializeCoinbaseData()
	coinbaseDataLen := CompactSizeUint{Value: uint64(len(coinbaseData))}

	sequence := Uint32ToBytes(0xFFFFFFFF)

	outputCount := CompactSizeUint{Value: 1}
	outputValue := Uint64ToBytes(coinbase.Out.Value)

	pubKeyScript := coinbase.Out.PubKeyScript
	pubKeyScriptSize := CompactSizeUint{Value: uint64(len(pubKeyScript))}

	lockTime := Uint32ToBytes(coinbase.lockTime)

	data = append(data, version...)
	data = append(data, inputCount.Bytes()...)
	data = append(data, inputOutpointHash[:]...)
	data = append(data, inputOutpointN...)
	data = append(data, coinbaseDataLen.Bytes()...)
	data = append(data, coinbaseData...)
	data = append(data, sequence...)
	data = append(data, outputCount.Bytes()...)
	data = append(data, outputValue...)
	data = append(data, pubKeyScriptSize.Bytes()...)
	data = append(data, pubKeyScript...)
	data = append(data, lockTime...)

	return data
}

func (coinbase *Coinbase) serializeCoinbaseData() []byte {
	data := []byte{}

	height := Uint32ToBytes(coinbase.Height)

	data = append(data, height...)

	return data
}

// Deserialize Deserializable from byte array
func (coinbase *Coinbase) Deserialize([]byte) {

}

// Hash Hashable
func (coinbase *Coinbase) Hash() *HashCode {
	if coinbase.HashID != nil {
		return coinbase.HashID
	}

	bytes := coinbase.Serialize()
	hash := HashCode(sha256.Sum256(bytes))
	coinbase.HashID = &hash

	return coinbase.HashID
}
