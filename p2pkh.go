package bitcointx

// scriptPubKey: OP_DUP OP_HASH160 <pubKeyHash> OP_EQUALVERIFY OP_CHECKSIG
// scriptSig: <sig> <pubKey>

import (
	. "github.com/yockliu/bitcoinlib"
)

// P2PHKGenPubKeyScript generator scriptPubKey from public key
func P2PHKGenPubKeyScript(address string) []byte {
	pubKeyHash := GetPubKeyHashFromAddress(address)

	scriptPubKey := []byte{opDup, opHash160}
	scriptPubKey = append(scriptPubKey, pubKeyHash...)
	scriptPubKey = append(scriptPubKey, opEqualVerify, opCheckSig)

	return scriptPubKey
}

// P2PHKParsePubKeyScript parse scriptPubKey to code statement
func P2PHKParsePubKeyScript(scriptPubKey []byte) []interface{} {
	return []interface{}{}
}

// P2PHKGenSigScript generator scriptSig from private key
func P2PHKGenSigScript(signature []byte, keyPair *KeyPair) []byte {
	return append(signature, keyPair.PublicKey...)
}

// P2PHKParseSigScript parse scriptSig to code statement
func P2PHKParseSigScript(scriptSig []byte) []interface{} {
	return []interface{}{}
}
