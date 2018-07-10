package bitcointx

// TxIn Transaction input
type TxIn struct {
	Outpoint struct {
		hash  HashCode
		index uint32
	}
	ScriptBytes     uint64
	SignaturaScript string
	sequence        uint32
}

// TxOut Trasaction output
type TxOut struct {
	Value             int64
	PubKeyScriptBytes uint64
	PubKeyScript      string
}

// Transaction in or out
type Transaction struct {
	Version    uint32
	TxInBytes  uint64
	TxInList   []TxIn
	TxOutBytes uint64
	TxOutList  []TxOut
	LockTime   uint32
}
