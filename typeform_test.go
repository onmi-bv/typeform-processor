package typeform

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"
)

func TestInit(t *testing.T) {

	tag := "123456"
	formID := "abcdef"

	os.Setenv("TYPEFORM_ID", formID)
	os.Setenv("TYPEFORM_TAG", tag)
	os.Setenv("TYPEFORM_TOKEN", "secrettoken")

	mux := http.NewServeMux()
	// the default http package does not support the parsing of Url parameters so this will hit the endpoint
	mux.HandleFunc(fmt.Sprintf("/forms/%s/webhooks/%s", formID, tag), handleFormsWebhook(t))

	ts := httptest.NewServer(mux)
	defer ts.Close()

	os.Setenv("TYPEFORM_URL", ts.URL+"/webhook")
	os.Setenv("TYPEFORM_BASEURL", ts.URL)

	testCases := []struct {
		desc   string
		mutate func()
		want   string
	}{
		{
			desc: "Correct authorization",
			mutate: func() {
				return
			},
			want: "",
		},
		{
			desc: "Incorrect authorization",
			mutate: func() {
				os.Setenv("TYPEFORM_TOKEN", "adifferentsecrettoken")
				return
			},
			want: "SetupTypeformWebhook: failed request with statuscode (401 Unauthorized) and body: Incorrect Authorization token\n",
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Init(%q)", tc.desc), func(t *testing.T) {

			// run whatever we want to mutate for this case
			tc.mutate()

			got := Init()

			// if we got an error, check if the error is as expected
			if got != nil {
				if got.Error() != tc.want {
					t.Errorf("Init(%q) = %q; want %q", tc.desc, got, tc.want)
				}
			}

			// if we did did not get an error, but are expecting one, that is also an error
			if got == nil && tc.want != "" {
				t.Errorf("Init(%q) is empty; want %q", tc.desc, tc.want)
			}
		})
	}

}

func handleFormsWebhook(t *testing.T) http.HandlerFunc {

	type Webhook struct {
		ID        string    `json:"id"`
		FormID    string    `json:"form_id"`
		Tag       string    `json:"tag"`
		URL       string    `json:"url"`
		Enabled   bool      `json:"enabled"`
		VerifySsl bool      `json:"verify_ssl"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	var webhook Webhook

	// first time token is correct, subtests will change this for bad authentication
	token := os.Getenv("TYPEFORM_TOKEN")

	return func(w http.ResponseWriter, r *http.Request) {

		tokenFromRequest := r.Header.Get("Authorization")

		if fmt.Sprintf("bearer %s", token) != tokenFromRequest {
			http.Error(w, "Incorrect Authorization token", http.StatusUnauthorized)
			return
		}

		formID, tag := parsePathParameters(r.URL)

		// retreives a single webhook (https://developer.typeform.com/webhooks/reference/retrieve-single-webhook/)
		if r.Method == "GET" {

			// if the webhook is already defined
			if webhook.Tag == tag && webhook.FormID == formID {

				res, err := json.Marshal(webhook)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

				w.Header().Set("Content-Type", "application/json")
				w.Write(res)
				return
			}

			http.Error(w, "No webhook found", http.StatusNotFound)
			return
		}

		// update or create a webhook (https://developer.typeform.com/webhooks/reference/create-or-update-webhook/)
		if r.Method == "PUT" {

			type RequestBody struct {
				URL     string `json:"url"`
				Enabled bool   `json:"enabled"`
			}

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			var requestBody RequestBody

			err = json.Unmarshal(body, &requestBody)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			webhook.Enabled = requestBody.Enabled
			webhook.URL = requestBody.URL
			webhook.Tag = tag
			webhook.FormID = formID

			res, err := json.Marshal(webhook)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(res)
			return

		}

		http.Error(w, "", http.StatusMethodNotAllowed)

	}
}

func parsePathParameters(u *url.URL) (string, string) {

	var formID string
	var tag string

	path := u.Path

	// \/forms\/([a-zA-Z0-9]+)\/webhooks\/([a-zA-Z0-9]+)

	paths := strings.Split(path, "/")
	formID = paths[1]
	tag = paths[3]

	// some regex magic here

	return formID, tag
}

func TestParseTypeformData(t *testing.T) {

	testCases := []struct {
		desc string
		arg  []byte
		want string
	}{
		{desc: "Correct submission", arg: correctSubmission, want: ""},
		{desc: "Empty submission", arg: emptySubmission, want: "ParseTypeformData::json.Unmarshal: cannot parse incoming typeform data from request: unexpected end of JSON input"},
		{desc: "Faulty submission", arg: faultySubmission, want: "ParseTypeformData: no questions and/or answers are found"},
		{desc: "Missing user submission", arg: missingUserSubmission, want: "ParseTypeformData: user query paramater is missing within the hidden field"},
		{desc: "Missing t submission", arg: missingTSubmission, want: "ParseTypeformData: t query paramater is missing within the hidden field"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("ParseTypeformData(%q)", tc.desc), func(t *testing.T) {

			_, got := ParseTypeformData(tc.arg)

			// if we got an error, check if the error is as expected
			if got != nil {
				if got.Error() != tc.want {
					t.Errorf("ParseTypeformData(%q) = %q; want %q", tc.desc, got, tc.want)
				}
			}

			// if we did did not get an error, but are expecting one, that is also an error
			if got == nil && tc.want != "" {
				t.Errorf("ParseTypeformData(%q) is empty; want %q", tc.desc, tc.want)
			}
		})
	}

}
