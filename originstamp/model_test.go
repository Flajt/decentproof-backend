package originstamp

import (
	"encoding/json"
	"testing"

	models "github.com/Flajt/decentproof-backend/originstamp/models"
)

func TestMarshaling(t *testing.T) {

	t.Run("TimeStampStatus", func(t *testing.T) {
		testJSON := `{
			"data": {
			  "comment": "my fancy comment",
			  "created": false,
			  "date_created": 0,
			  "hash_string": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			  "timestamps": [
				{
				  "seed_id": "my-seed-id-1",
				  "currency_id": 0,
				  "private_key": "a-very-private-key",
				  "submit_status": 0,
				  "timestamp": 0,
				  "transaction": "some-base64-encoded-string"
				}
			  ]
			},
			"error_code": 0,
			"error_message": "string"
		  }`
		myStruct := models.OriginStampCreateTimestampResponse{}
		err := json.Unmarshal([]byte(testJSON), &myStruct)
		if err != nil {
			t.Log(err)
		}
		t.Log(myStruct)
	})
	t.Run("CreateTimeStampResonse", func(t *testing.T) {
		testJson := `{
			"data": {
			  "comment": "string",
			  "created": false,
			  "date_created": 0,
			  "hash_string": "string",
			  "timestamps": [
				{
				  "seed_id": "1234567890ABCDEF",
				  "currency_id": 0,
				  "private_key": "a-very-private-key",
				  "submit_status": 0,
				  "timestamp": 0,
				  "transaction": "string"
				}
			  ]
			},
			"error_code": 0,
			"error_message": "an fancy error message"
		  }`

		myStruct := models.OriginStampCreateTimestampResponse{}
		err := json.Unmarshal([]byte(testJson), &myStruct)
		if err != nil {
			t.Log(err)
		}
		t.Log(myStruct)
	})

	t.Run("ProofResponse", func(t *testing.T) {
		testJson := `{
			"error_code": 0,
			"error_message": null,
			"data": {
				"download_url": "https://api.originstamp.com/v3/timestamp/proof/download?token=<token>&name=proof.Bitcoin.2c5d36be542f8f0e7345d77753a5d7ea61a443ba6a9a86bb060332ad56dba38e.xml",
				"file_name": "proof.Bitcoin.2c5d36be542f8f0e7345d77753a5d7ea61a443ba6a9a86bb060332ad56dba38e.xml",
				"file_size_bytes": 0
			}
		}`

		myStruct := models.OriginStampProofResponse{}
		err := json.Unmarshal([]byte(testJson), &myStruct)
		if err != nil {
			t.Log(err)
		}
		t.Log(myStruct)
	})

	t.Run("WebhookRequestBody", func(t *testing.T) {
		testJson := `{
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
		  }`

		myStruct := models.OriginStampWebhookRequestBody{}
		err := json.Unmarshal([]byte(testJson), &myStruct)
		if err != nil {
			t.Log(err)
		}
		t.Log(myStruct)
	})

}
