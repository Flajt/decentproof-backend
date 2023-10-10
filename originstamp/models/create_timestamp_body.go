package originstamp

/*
{
  "comment": "test",
  "hash": "2c5d36be542f8f0e7345d77753a5d7ea61a443ba6a9a86bb060332ad56dba38e",
  "notifications": [
    {
      "currency": 0,
      "notification_type": 0,
      "target": "mail@notification.de"
    },
    {
      "currency": 0,
      "notification_type": 1,
      "target": "https://webhook-notification.url"
    }
  ],
  "url": "string"
}
*/

// Contains Originstamp timestamp request payload
type OriginStampTimestampRequestBody struct {
	Comment       string                          `json:"comment"`
	Hash          string                          `json:"hash"`
	Notifications []OriginStampNotificationTarget `json:"notifications"`
}

type OriginStampNotificationTarget struct {
	// 0 for BTC, 1 for ETH, 2 for OAN, 100 for Süddeutsche Zeitung
	Currency int `json:"currency"`
	// 0 for email, 1 for webhook
	NotificationType int `json:"notification_type"`
	// email address or webhook url
	Target string `json:"target"`
}
