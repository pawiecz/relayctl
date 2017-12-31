package main

import (
	"net/http"
	"net/http/httptest"

	"github.com/julienschmidt/httprouter"

	. "github.com/Benjamintf1/ExpandedUnmarshalledMatchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func call(method, path, route string, fnHandler func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	return newRequestRecorder(req, method, route, fnHandler)
}

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
		expected := "relayctl reporting for duty!\n"
		actual := call("GET", "/", "/", Index)
		Expect(actual.Code).To(Equal(http.StatusOK))
		Expect(actual.Body.String()).To(Equal(expected))
	})

	It("should list available relays", func() {
		expected := "{\"meta\":null,\"data\":[{\"id\":\"I1\",\"pin\":17,\"state\":false},{\"id\":\"I2\",\"pin\":27,\"state\":true}]}\n"
		actual := call("GET", "/relays", "/relays", RelayIndex)
		Expect(actual.Code).To(Equal(http.StatusOK))
		Expect(actual.Body.String()).To(MatchUnorderedJSON(expected))
	})

	It("should return an error on show request for nonexistent relay", func() {
		expected := "{\"error\":{\"status\":404,\"title\":\"relay not found\"}}\n"
		actual := call("GET", "/relays/I0", "/relays/:id", RelayShow)
		Expect(actual.Code).To(Equal(http.StatusNotFound))
		Expect(actual.Body.String()).To(MatchJSON(expected))
	})

	It("should return proper data on show request for existing relays", func() {
		expected := "{\"meta\":null,\"data\":{\"id\":\"I1\",\"pin\":17,\"state\":false}}\n"
		actual := call("GET", "/relays/I1", "/relays/:id", RelayShow)
		Expect(actual.Code).To(Equal(http.StatusOK))
		Expect(actual.Body.String()).To(MatchJSON(expected))

		expected = "{\"meta\":null,\"data\":{\"id\":\"I2\",\"pin\":27,\"state\":true}}\n"
		actual = call("GET", "/relays/I2", "/relays/:id", RelayShow)
		Expect(actual.Code).To(Equal(http.StatusOK))
		Expect(actual.Body.String()).To(MatchJSON(expected))
	})
})
