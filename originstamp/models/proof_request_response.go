package originstamp

/*
{
    "error_code": 0,
    "error_message": null,
    "data": {
        "download_url": "https://api.originstamp.com/v3/timestamp/proof/download?token=<token>&name=proof.Bitcoin.2c5d36be542f8f0e7345d77753a5d7ea61a443ba6a9a86bb060332ad56dba38e.xml",
        "file_name": "proof.Bitcoin.2c5d36be542f8f0e7345d77753a5d7ea61a443ba6a9a86bb060332ad56dba38e.xml",
        "file_size_bytes": 0
    }
}
*/

// Contains Originstamp proof response payload
type OriginStampProofResponse struct {
	ErrorCode    int                 `json:"error_code"`
	ErrorMessage string              `json:"error_message"`
	Data         OrignStampProofData `json:"data"`
}

type OrignStampProofData struct {
	DownloadURL   string `json:"download_url"`
	FileName      string `json:"file_name"`
	FileSizeBytes int    `json:"file_size_bytes"`
}
