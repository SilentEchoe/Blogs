package million

type PayloadReq struct {
	WindowsVersion string    `json:"version"`
	Token          string    `json:"token"`
	Payloads       []Payload `json:"data"`
}

type Payload struct {
	BucketName string `json:"backetname"`
}
