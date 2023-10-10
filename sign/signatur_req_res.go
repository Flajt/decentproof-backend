package sign

type SignatureRequestBody struct {
	Data  string `json:"data"`
	Email string `json:"email"`
}

type SignatureResponseBody struct {
	Signature string `json:"signature"`
}
