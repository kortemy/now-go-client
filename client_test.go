package now_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	log "github.com/sirupsen/logrus"

	. "github.com/kortemy/now-go-client/now"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	var testServer *httptest.Server
	testClient := Client{}

	It("should make a authrozied GET call", func() {
		testClient.Token = "test-token"
		testServer = httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				Expect(r.Header.Get("Authorization")).To(Equal("Bearer test-token"))
			}))
		defer testServer.Close()

		err := testClient.Call("GET", testServer.URL, nil, nil)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should make a unauthorized GET call", func() {
		testClient.Token = ""
		testServer = httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				err := Error{
					Code:    "some-code",
					Message: "some-message",
					URL:     "some-url",
				}
				errResponse := ErrorResponse{err}
				w.WriteHeader(403)
				json.NewEncoder(w).Encode(errResponse)
			}))
		defer testServer.Close()

		err := testClient.Call("GET", testServer.URL, nil, nil)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("some-message"))
	})

	It("should make a POST call with some data", func() {
		testServer = httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// body, _ := ioutil.ReadAll(r.Body)
				testResponse := Secret{}
				json.NewDecoder(r.Body).Decode(&testResponse)
				log.Info(testResponse)
				testResponse.UID = "secret-uid"
				json.NewEncoder(w).Encode(testResponse)
			}))
		defer testServer.Close()

		testData := Secret{
			Name:  "secret-name",
			Value: "secret-value",
		}
		testResult := Secret{}
		err := testClient.Call("POST", testServer.URL, testData, &testResult)
		Expect(err).NotTo(HaveOccurred())
		Expect(testResult.UID).To(Equal("secret-uid"))
		Expect(testResult.Name).To(Equal("secret-name"))
		Expect(testResult.Value).To(Equal("secret-value"))
	})
})
