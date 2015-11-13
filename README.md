# golang-twilio-sms
SMS Service for Twilio written in Golang

https://godoc.org/github.com/pevika/golang-twilio-sms

# Example

```

service := twiliosms.NewService(accountSid, authToken, phonenumber)
resp, err := service.Send(userPhonenumber, "Your verification code for MySuperApplication is: 987654)

```
