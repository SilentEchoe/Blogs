package main

import (
	"fmt"
	"time"
)

const token = `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWRtaW4iLCJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwidWlkIjoiZjkzN2JiMzctYmFmNC0xMWVjLWJkYWYtY2E0ZWM0MTdkNTBlIiwicHJlZmVycmVkX3VzZXJuYW1lIjoiYWRtaW4iLCJmZWRlcmF0ZWRfY2xhaW1zIjp7ImNvbm5lY3Rvcl9pZCI6InN5c3RlbSIsInVzZXJfaWQiOiJhZG1pbiJ9LCJzaWQiOiIiLCJsb2dpbl90eXBlIjoiIiwiYXVkIjpudWxsLCJleHAiOjE2NzUzMDM1MDd9.f4_acDc3daMofLh1Rp0o1DsEzIUwN6R63lGRvsQhCqo`

func main() {
	url := "http://10.3.70.149:30001/"
	//startDate := time.Now().ParseDuration("-24h") - time.Hour
	d, _ := time.ParseDuration("-24h")
	startDate := time.Now().Add(d).Unix()
	endDate := time.Now().Unix()

	//url := "/api/aslan/stat/dashboard/build?startDate=%s&endDate=%s"
	//fmt.Sprintf(url,)
	url = fmt.Sprintf(url+"/api/aslan/stat/dashboard/build?startDate=%d&endDate=%d", startDate, endDate)
	fmt.Println("%s", url)
}
