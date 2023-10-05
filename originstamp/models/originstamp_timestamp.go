package originstamp

/*
	{
	     "seed_id": "f98a61e8-22bd-4ea5-ab9e-02723ff78c40",
	     "currency_id": 0,
	     "transaction": "aed3db9ef94953f65e93d56a4e5bcf234d43e27a1b3e7ce0f274cc7ed750d0e2",
	     "private_key": "5e92ec09501a5d39e251a151f84b5e2228312c445eb23b4e1de6360e27bad54b",
	     "timestamp": 1541203656000,
	     "submit_status": 3
	   }
*/
type OriginStampTimeStamp struct {
	SeedID       string `json:"seed_id"`
	CurrencyID   int    `json:"currency_id"`
	Transaction  string `json:"transaction"`
	PrivateKey   string `json:"private_key"`
	Timestamp    int64  `json:"timestamp"`
	SubmitStatus int    `json:"submit_status"`
}
