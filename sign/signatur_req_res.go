package sign

type BlockChain uint8

const (
	Bitcoin  BlockChain = 0
	Ethereum BlockChain = 1
)

type SignatureRequestBody struct {
	Data       string     `json:"data"`
	Email      string     `json:"email"`
	BlockChain BlockChain `json:"blockchain"`
}

type SignatureResponseBody struct {
	Signature string `json:"signature"`
}
