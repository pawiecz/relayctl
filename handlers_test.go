package main

import (
	"net/http"
	"net/http/httptest"

	"github.com/julienschmidt/httprouter"

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
	It("should report for duty on the default route", func() {
		expectedResponse := "relayctl reporting for duty!\n"
		req := httptest.NewRequest("GET", "/", nil)
		rr := newRequestRecorder(req, "GET", "/", Index)
		Expect(rr.Code).To(Equal(http.StatusOK))
		Expect(rr.Body.String()).To(Equal(expectedResponse))
	})
})
