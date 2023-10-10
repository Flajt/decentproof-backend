package originstamp

/*
	{
	  "created": false,
	  "date_created": 1541203188245,
	  "comment": null,
	  "hash_string": "2c5d36be542f8f0e7345d77753a5d7ea61a443ba6a9a86bb060332ad56dba38e",
	  "timestamps": [
	    {
	      "seed_id": "f98a61e8-22bd-4ea5-ab9e-02723ff78c40",
	      "currency_id": 0,
	      "transaction": "aed3db9ef94953f65e93d56a4e5bcf234d43e27a1b3e7ce0f274cc7ed750d0e2",
	      "private_key": "5e92ec09501a5d39e251a151f84b5e2228312c445eb23b4e1de6360e27bad54b",
	      "timestamp": 1541203656000,
	      "submit_status": 3
	    }
	  ]
	}
*/
// Contains Originstamp webhook payload
type OriginStampWebhookRequestBody struct {
	Created     bool                   `json:"created"`
	DateCreated int64                  `json:"date_created"`
	Comment     string                 `json:"comment"`
	HashString  string                 `json:"hash_string"`
	Timestamps  []OriginStampTimeStamp `json:"timestamps"`
}
