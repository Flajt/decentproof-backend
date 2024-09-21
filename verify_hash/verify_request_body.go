package verify_hash

type VerifyRequestBody struct {
	Hash string `json:"hash"`
	// Currently not used, consider passing as int
	BlockChain string `json:"blockChain"`
}
