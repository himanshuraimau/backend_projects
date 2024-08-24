package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey returns the API key from the Authorization header

// GetAPIKey returns the API key from the Authorization header
// AuthHeader: ApiKey{insert-api-key-here}

func GetAPIKey(headers http.Header) (string,error){
       val:= headers.Get("Authorization")
       if val == "" {
			   return "", errors.New("no API key provided")
	   }

	 vals := strings.Split(val, " ")
	 if len(vals) != 2 {
		 return "", errors.New("invalid API key format")
	 }
	 if vals[0] != "ApiKey" {
		 return "", errors.New("invalid API key type")
	 }
	 return vals[1], nil
}