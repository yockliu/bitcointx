package bitcointx

// scriptPubKey: OP_DUP OP_HASH160 <pubKeyHash> OP_EQUALVERIFY OP_CHECKSIG
// scriptSig: <sig> <pubKey>

import (
	. "github.com/yockliu/bitcoinlib"
)

// P2PHKGenScriptPubKey generator scriptPubKey from public key
func P2PHKGenScriptPubKey(address string) []byte {
	pubKeyHash := GetPubKeyHashFromAddress(address)

	scriptPubKey := []byte{opDup, opHash160}
	scriptPubKey = append(scriptPubKey, pubKeyHash...)
	scriptPubKey = append(scriptPubKey, opEqualVerify, opCheckSig)

	return scriptPubKey
}

// P2PHKParseScriptPubKey parse scriptPubKey to code statement
func P2PHKParseScriptPubKey(scriptPubKey []byte) []interface{} {
	return []interface{}{}
}

// P2PHKGenScriptSig generator scriptSig from private key
func P2PHKGenScriptSig(signature []byte, keyPair *KeyPair) []byte {
	return append(signature, keyPair.PublicKey...)
}

// P2PHKParseScriptSig parse scriptSig to code statement
func P2PHKParseScriptSig(scriptSig []byte) []interface{} {
	return []interface{}{}
}
