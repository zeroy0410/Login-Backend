package captcha

import (
	"Login-Backend/src/config"
	"bytes"
	"encoding/json"
	// "fmt"
	"net/http"
	"strings"
)

type VerifiyCaptchaResonse struct {
	Success bool `json:"success"`
}

func VerifyCaptcha(token string) (bool, error) {
	url := "https://recaptcha.net/recaptcha/api/siteverify?secret=" +
		config.CaptchaConfig.SecretKey + "&response=" + token
	// fmt.Println("hey: "+config.CaptchaConfig.SecretKey)
	client := &http.Client{}
	request,err:=http.NewRequest("POST",url,&strings.Reader{})
	if err != nil {
		return false,err
	}
	resp, err:= client.Do(request)
	if err != nil {
		return false, err
	}
	var body bytes.Buffer
	_,err=body.ReadFrom(resp.Body)
	if err != nil {
		return false, err
	}
	var result VerifiyCaptchaResonse
	err = json.Unmarshal(body.Bytes(),&result)
	if err != nil {
		return false,err
	}
	return result.Success, nil
}	
