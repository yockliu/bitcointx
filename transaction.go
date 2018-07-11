package bitcointx

import (
	. "github.com/yockliu/bitcoinlib"
)

// TxIn Transaction input
type TxIn struct {
	Outpoint struct {
		TxID  HashCode
		Index uint32
	}
	ScriptSig []byte
	Sequence  uint32
}

// TxOut Trasaction output
type TxOut struct {
	TxID         HashCode
	Value        uint64 // Satoshis
	ScriptPubKey []byte
}

// Transaction in or out
type Transaction struct {
	Version   uint32
	TxInList  []TxIn
	TxOutList []TxOut
	LockTime  uint32
}

// GenTxIn gen Transaction input that can unlock the TxOut, should check if the TXO is Unspend first
func GenTxIn(txOut TxOut, senderKeyPair KeyPair) TxIn {
	content := []byte{}
	txIn := TxIn{}
	txIn.Outpoint.TxID = txOut.TxID
	txIn.ScriptSig = P2PHKGenScriptSig(content, senderKeyPair)
	txIn.Sequence = 0xFFFFFFFF
	return txIn
}

// GenTxOut generator T
func GenTxOut(value uint64, address []byte) TxOut {
	txOut := TxOut{}
	txOut.Value = value
	txOut.ScriptPubKey = P2PHKGenScriptPubKey(address)
	return txOut
}

// Unlock use stack to do the unlock
func Unlock(txOut TxOut, txIn TxIn) bool {
	return false
}
