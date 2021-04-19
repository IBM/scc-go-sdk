package examples

import "os"

var apiKey = os.Getenv("apiKey")
var url = os.Getenv("URL") //Iam endpoint url. Required only when using dev or preprod environment
var accountID = os.Getenv("accountID")
