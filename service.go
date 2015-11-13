//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package twiliosms

import (
	"net/http"
	"net/url"
	"strings"
)

type Service struct {
	accountID		string
	authToken		string
	phonenumber		string
}

func NewService (accountID string, authToken string, phonenumber string) *Service {
	return &Service{
		accountID: accountID,
		authToken: authToken,
		phonenumber: phonenumber,
	}
}

func (this *Service) Send (phonenumber string, content string) (*http.Response, error) {
	values := url.Values{}
	values.Set("To", phonenumber)
	values.Set("From", this.phonenumber)
	values.Set("Body", content)
	urlValues := strings.NewReader(values.Encode())
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.twilio.com/2010-04-01/Accounts/" + this.accountID + "/Messages.json", urlValues)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(this.accountID, this.authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return client.Do(req)
}