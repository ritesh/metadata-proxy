package main

//IamCreds contains credentials
type IamCreds struct {
	Code            string `json:"Code"`
	LastUpdated     string `json:"LastUpdated"`
	AccessKeyID     string `json:"AccessKeyId"`
	SecretAccessKey string `json:"SecretAccessKey"`
	Token           string `json:"Token"`
	Expiration      string `json:"Expiration"`
}
