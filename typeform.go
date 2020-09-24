package typeform

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joeshaw/envdecode"
	"io/ioutil"
	"net/http"
	"time"
)

func Init() error {

	var conf Configuration

	if err := envdecode.Decode(&conf); err != nil {
		return fmt.Errorf("envdecode.Decode: cannot parse the environment parameters: %v", err)
	}

	return conf.setupTypeformWebhook()
}

type Configuration struct {
	Id      string `env:"TYPEFORM_ID,required"`
	Tag     string `env:"TYPEFORM_TAG,required"`
	Token   string `env:"TYPEFORM_TOKEN,required"`
	FormUrl string `env:"TYPEFORM_URL,required"`
	BaseUrl string `env:"TYPEFORM_BASEURL,default=https://api.typeform.com"`
}

// ParseTypeformData accepts a json and returns a WebhookTypeform struct
func ParseTypeformData(data []byte) (WebhookTypeform, error) {
	var form WebhookTypeform

	err := json.Unmarshal(data, &form)
	if err != nil {
		return form, fmt.Errorf("json.Unmarshal: cannot parse incoming typeform data from request: %v", err)
	}

	if form.FormResponse.Hidden.User == "" {
		return form, errors.New("user query paramater is missing within the hidden field")
	}

	if form.FormResponse.Hidden.T == "" {
		return form, errors.New("t query paramater is missing within the hidden field")
	}

	if len(form.FormResponse.Answers) == 0 || len(form.FormResponse.Definition.Fields) == 0 {
		return form, errors.New("no questions and/or answers are found")
	}

	return form, nil
}

func (conf *Configuration) setupTypeformWebhook() error {

	url := fmt.Sprintf("%s/forms/%s/webhooks/%s", conf.BaseUrl, conf.Id, conf.Tag)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("cannot construct request: %v", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("bearer %s", conf.Token))
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("client.Do: cannot do request to TypeForm: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll: cannot read response from TypeForm: %v", err)
	}

	switch res.StatusCode {

	case http.StatusOK:

		var response WebhookDataResponse
		err := json.Unmarshal(body, &response)
		if err != nil {
			return fmt.Errorf("json.Unmarshal: %v", err)
		}

		if response.URL != conf.FormUrl {
			err := conf.enableTypeformWebhook()
			if err != nil {
				return err // returning the error directly because EnableTypeformWebhook is internal and takes care of the error formatting
			}
			return fmt.Errorf("webhook url from Typeform doesn't match preset: %s | %s (setting up again...)", response.URL, conf.FormUrl)
		}

		if response.Tag != conf.Tag {
			return fmt.Errorf("webhook tag from Typeform does not match preset: %s | %s", response.Tag, conf.Tag)
		}

		return nil

	case http.StatusNotFound:
		err := conf.enableTypeformWebhook()
		if err != nil {
			return err
		}

		return nil

	default:
		return fmt.Errorf("failed request with statuscode (%s) and body: %s", res.Status, body)

	}

}

func (conf *Configuration) enableTypeformWebhook() error {

	url := fmt.Sprintf("%s/forms/%s/webhooks/%s", conf.BaseUrl, conf.Id, conf.Tag)

	requestBody := &WebhookCreateRequest{
		URL:     conf.FormUrl,
		Enabled: true,
	}

	requestByte, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("json.Marshall: cannot marshal request body: %v", err)
	}

	requestReader := bytes.NewReader(requestByte)

	req, err := http.NewRequest("PUT", url, requestReader)
	if err != nil {
		return fmt.Errorf("http.NewRequest: %v", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("bearer %s", conf.Token))
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("client.Do: %v", err)

	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("iotutil.ReadAll: %v", err)
	}

	switch res.StatusCode {

	case http.StatusOK:

		//* All fine here

		var response WebhookCreateResponse
		err := json.Unmarshal(body, &response)
		if err != nil {
			return fmt.Errorf("json.Unmarhal: cannot unmarshal response body: %v", err)
		}

		return nil

	default:
		return fmt.Errorf("received non-ok response code (%d) with body: %s", res.StatusCode, body)
	}

}
