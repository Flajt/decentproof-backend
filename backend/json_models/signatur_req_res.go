package decentproof_functions

type SignatureRequestBody struct {
	Data string `json:"data"`
}

type SignatureResponseBody struct {
	Signature string `json:"signature"`
}
