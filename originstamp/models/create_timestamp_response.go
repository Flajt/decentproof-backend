package originstamp

/*
{
  "data": {
    "comment": "string",
    "created": false,
    "date_created": 0,
    "hash_string": "string",
    "timestamps": [
      {
        "seed_id": "string",
        "currency_id": 0,
        "private_key": "string",
        "submit_status": 0,
        "timestamp": 0,
        "transaction": "string"
      }
    ]
  },
  "error_code": 0,
  "error_message": "string"
}*/

// Contains Originstamp create timestamp response payload

type OriginStampCreateTimestampResponse struct {
	Data         OriginStampCreateTimestampData `json:"data"`
	ErrorCode    int                            `json:"error_code"`
	ErrorMessage string                         `json:"error_message"`
}

type OriginStampCreateTimestampData struct {
	Comment     string                 `json:"comment"`
	Created     bool                   `json:"created"`
	DateCreated int64                  `json:"date_created"`
	HashString  string                 `json:"hash_string"`
	Timestamps  []OriginStampTimeStamp `json:"timestamps"`
}
