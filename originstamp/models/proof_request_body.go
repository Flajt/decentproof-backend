package originstamp

/*
{
  "currency": 0,
  "hash_string": "2c5d36be542f8f0e7345d77753a5d7ea61a443ba6a9a86bb060332ad56dba38e",
  "proof_type": 0
}
*/

// Contains Originstamp proof request payload
type OriginStampProofRequestBody struct {
	// 0 for BTC, 1 for ETH, 2 for OAN, 100 for Süddeutsche Zeitung
	Currency   int    `json:"currency"`
	HashString string `json:"hash_string"`
	// 0 for Merkle Tree, 1 for PDF
	ProofType int `json:"proof_type"`
}
