package bitcointx

import (
	"crypto/sha256"

	. "github.com/yockliu/bitcoinlib"
)

// Transaction in or out
type Transaction struct {
	Hashable
	Serializable
	HashID   *HashCode
	version  uint32
	in       []*TXIn
	out      []*TXOut
	lockTime uint32
}

// TXIn Transaction input
type TXIn struct {
	outpoint  *Outpoint
	keyPair   *KeyPair
	signature []byte
	sequence  uint32
	ScriptSig []byte // generate when serialize
}

// TXOut Trasaction output
type TXOut struct {
	Value        uint64 // Satoshis
	Address      string
	ScriptPubKey []byte // greerate when serialize
}

// Outpoint of the tx
type Outpoint struct {
	TxHash HashCode
	N      uint32
	Utxo   *TXOut
	Lock   bool
}

// NewTXIn New TXIn struct
func NewTXIn(outpoint *Outpoint, senderKey *KeyPair) *TXIn {
	txIn := TXIn{}
	txIn.outpoint = outpoint
	txIn.keyPair = senderKey
	txIn.sequence = 0xFFFFFFFF
	return &txIn
}

// NewTXOut new TXOut struct
func NewTXOut(value uint64, address string) *TXOut {
	txOut := TXOut{}
	txOut.Value = value
	txOut.Address = address
	return &txOut
}

// NewTransaction new Transaction,
func NewTransaction(txIns []*TXIn, txOuts []*TXOut) *Transaction {
	tx := Transaction{}

	tx.version = 1
	tx.in = txIns
	tx.out = txOuts
	tx.lockTime = 0

	// txIn should be signed every one single by
	for _, txIn := range txIns {
		txIn.signature = tx.signContent(txIn)
	}

	tx.Hash()

	return &tx
}

// sign content for the tx in
func (tx *Transaction) signContent(txIn *TXIn) []byte {
	signature := Sign(txIn.keyPair.PrivateKey, tx.hashContent(txIn))
	return signature
}

func (tx *Transaction) hashContent(txIn *TXIn) []byte {
	content := Uint32ToBytes(tx.version)
	content = append(content, Uint32ToBytes(tx.lockTime)...)

	for _, in := range tx.in {
		value := Uint64ToBytes(in.outpoint.Utxo.Value)
		content = append(content, value...)
		content = append(content, in.outpoint.Utxo.Address...)
	}
	for _, out := range tx.out {
		value := Uint64ToBytes(out.Value)
		content = append(content, value...)
		content = append(content, out.Address...)
	}

	content = append(content, txIn.outpoint.Utxo.Address...)

	hash := sha256.Sum256(content)

	return hash[:]
}

// Serialize serialize TX to byte array
func (tx *Transaction) Serialize() []byte {
	data := []byte{}

	version := Uint32ToBytes(tx.version)
	data = append(data, version...)

	inSize := CompactSizeUint{Value: uint64(len(tx.in))}
	data = append(data, inSize.Bytes()...)
	for _, input := range tx.in {
		data = append(data, input.Serialize()...)
	}

	outSize := CompactSizeUint{Value: uint64(len(tx.out))}
	data = append(data, outSize.Bytes()...)
	for _, output := range tx.out {
		data = append(data, output.Serialize()...)
	}

	lockTime := Uint32ToBytes(tx.lockTime)
	data = append(data, lockTime...)

	return data
}

// Serialize serialize TXIn to byte array
func (txIn *TXIn) Serialize() []byte {
	data := []byte{}

	utxoHash := txIn.outpoint.TxHash[:]
	utxoIndex := Uint32ToBytes(txIn.outpoint.N)
	sigScript := P2PHKGenScriptSig(txIn.signature, txIn.keyPair)
	sigScriptSize := CompactSizeUint{Value: uint64(len(sigScript))}
	sequenceBytes := Uint32ToBytes(txIn.sequence)

	data = append(data, utxoHash...)
	data = append(data, utxoIndex...)
	data = append(data, sigScriptSize.Bytes()...)
	data = append(data, sigScript...)
	data = append(data, sequenceBytes...)

	return data
}

// Serialize serialize TXOut to byte array
func (txOut *TXOut) Serialize() []byte {
	data := []byte{}

	value := Uint64ToBytes(txOut.Value)
	pubKeyScript := P2PHKGenScriptPubKey(txOut.Address)
	pubKeyScriptSize := CompactSizeUint{Value: uint64(len(pubKeyScript))}

	data = append(data, value...)
	data = append(data, pubKeyScriptSize.Bytes()...)
	data = append(data, pubKeyScript...)

	return data
}

// Deserialize Serializable
func (tx *Transaction) Deserialize([]byte) {
	// TODO
}

// Hash Hashable
func (tx *Transaction) Hash() *HashCode {
	if tx.HashID != nil {
		return tx.HashID
	}
	bytes := tx.Serialize()
	hash := HashCode(sha256.Sum256(bytes))
	tx.HashID = &hash
	return tx.HashID
}
