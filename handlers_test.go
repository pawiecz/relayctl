package main

import (
	"net/http"
	"net/http/httptest"

	"github.com/julienschmidt/httprouter"

	. "github.com/Benjamintf1/ExpandedUnmarshalledMatchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func newRequestRecorder(req *http.Request, method string, strPath string, fnHandler func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) *httptest.ResponseRecorder {
	router := httprouter.New()
	router.Handle(method, strPath, fnHandler)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}

var _ = Describe("Main", func() {
	BeforeSuite(func() {
		module["I1"] = &Relay{ID: "I1", Pin: 17, State: false}
		module["I2"] = &Relay{ID: "I2", Pin: 27, State: true}
	})

	It("should report for duty on the default route", func() {
		expectedResponse := "relayctl reporting for duty!\n"
		req := httptest.NewRequest("GET", "/", nil)
		rr := newRequestRecorder(req, "GET", "/", Index)
		Expect(rr.Code).To(Equal(http.StatusOK))
		Expect(rr.Body.String()).To(Equal(expectedResponse))
	})

	It("should list available relays", func() {
		expectedResponse := "{\"meta\":null,\"data\":[{\"id\":\"I1\",\"pin\":17,\"state\":false},{\"id\":\"I2\",\"pin\":27,\"state\":true}]}\n"
		req := httptest.NewRequest("GET", "/relays", nil)
		rr := newRequestRecorder(req, "GET", "/relays", RelayIndex)
		Expect(rr.Code).To(Equal(http.StatusOK))
		Expect(rr.Body.String()).To(MatchUnorderedJSON(expectedResponse))
	})

	It("should return an error on show request for nonexistent relay", func() {
		expectedResponse := "{\"error\":{\"status\":404,\"title\":\"Relay not found\"}}\n"
		req := httptest.NewRequest("GET", "/relays/I0", nil)
		rr := newRequestRecorder(req, "GET", "/relays/:id", RelayShow)
		Expect(rr.Code).To(Equal(http.StatusNotFound))
		Expect(rr.Body.String()).To(MatchJSON(expectedResponse))
	})

	It("should return proper data on show request for existing relays", func() {
		expectedResponse := "{\"meta\":null,\"data\":{\"id\":\"I1\",\"pin\":17,\"state\":false}}\n"
		req := httptest.NewRequest("GET", "/relays/I1", nil)
		rr := newRequestRecorder(req, "GET", "/relays/:id", RelayShow)
		Expect(rr.Code).To(Equal(http.StatusOK))
		Expect(rr.Body.String()).To(MatchJSON(expectedResponse))

		expectedResponse = "{\"meta\":null,\"data\":{\"id\":\"I2\",\"pin\":27,\"state\":true}}\n"
		req = httptest.NewRequest("GET", "/relays/I2", nil)
		rr = newRequestRecorder(req, "GET", "/relays/:id", RelayShow)
		Expect(rr.Code).To(Equal(http.StatusOK))
		Expect(rr.Body.String()).To(MatchJSON(expectedResponse))
	})
})
