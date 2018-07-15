package bitcointx

import (
	"crypto/sha256"

	. "github.com/yockliu/bitcoinlib"
)

// Transaction in or out
type Transaction struct {
	Hash     HashCode
	version  uint32
	in       []TXIn
	out      []TXOut
	lockTime uint32
}

// TXIn Transaction input
type TXIn struct {
	outpoint  Outpoint
	keyPair   KeyPair
	signature []byte
	sequence  uint32
	ScriptSig []byte // generate when serialize
}

// TXOut Trasaction output
type TXOut struct {
	value        uint64 // Satoshis
	address      []byte
	ScriptPubKey []byte // greerate when serialize
}

// Outpoint of the tx
type Outpoint struct {
	txHash HashCode
	n      uint32
	utxo   TXOut
}

// NewTXIn New TXIn struct
func NewTXIn(hashOfTX HashCode, indexInTX uint32, utxo TXOut, senderKey KeyPair) TXIn {
	txIn := TXIn{}
	txIn.outpoint = Outpoint{txHash: hashOfTX, n: indexInTX, utxo: utxo}
	txIn.keyPair = senderKey
	txIn.sequence = 0xFFFFFFFF
	return txIn
}

// NewTXOut new TXOut struct
func NewTXOut(value uint64, address []byte) TXOut {
	txOut := TXOut{}
	txOut.value = value
	txOut.address = address
	return txOut
}

// NewTransaction new Transaction,
func NewTransaction(txIns []TXIn, txOuts []TXOut) Transaction {
	tx := Transaction{}

	tx.version = 1
	tx.in = txIns
	tx.out = txOuts
	tx.lockTime = 0

	// txIn should be signed every one single by
	for _, txIn := range txIns {
		txIn.signature = tx.signContent(txIn)
	}

	txBytes := tx.Serialize()
	tx.Hash = sha256.Sum256(txBytes)

	return tx
}

// sign content for the tx in
func (tx *Transaction) signContent(txIn TXIn) []byte {
	signature := Sign(txIn.keyPair.PrivateKey, tx.hashContent(txIn))
	return signature
}

func (tx *Transaction) hashContent(txIn TXIn) []byte {
	content := Uint32ToBytes(tx.version)
	content = append(content, Uint32ToBytes(tx.lockTime)...)

	for _, in := range tx.in {
		value := Uint64ToBytes(in.outpoint.utxo.value)
		content = append(content, value...)
		content = append(content, in.outpoint.utxo.address...)
	}
	for _, out := range tx.out {
		value := Uint64ToBytes(out.value)
		content = append(content, value...)
		content = append(content, out.address...)
	}

	content = append(content, txIn.outpoint.utxo.address...)

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

	utxoHash := txIn.outpoint.txHash[:]
	utxoIndex := Uint32ToBytes(txIn.outpoint.n)
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

	value := Uint64ToBytes(txOut.value)
	pubKeyScript := P2PHKGenScriptPubKey(txOut.address)
	pubKeyScriptSize := CompactSizeUint{Value: uint64(len(pubKeyScript))}

	data = append(data, value...)
	data = append(data, pubKeyScriptSize.Bytes()...)
	data = append(data, pubKeyScript...)

	return data
}
