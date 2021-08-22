package SMS

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//This function is to create a nondeterministic random code to be sent through SMS
func CreateCode() string {
	var code string = ""
	for i := 0; i < 6; i++ {
		rand.Seed(time.Now().Unix() + int64(i))
		var j int = rand.Intn(10)
		code += strconv.Itoa(j)
	}
	return code
}

//This function is to send SMS and return the status
func Send() bool {
	var sentStatus bool = false
	var sentMessage string = "The code is\n" + CreateCode()
	accountSid := "ACXXXX"
	authToken := "XXXXXX"
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	msgData := url.Values{}
	msgData.Set("To", "NUMBER_TO")
	msgData.Set("From", "NUMBER_FROM")
	msgData.Set("Body", sentMessage)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			sentStatus = true
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
	}
	return sentStatus
}
