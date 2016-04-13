package main_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	. "github.com/crowdriff/go-to-sample"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Context("DefaultHandler", func() {
		It("should return some nice JSON", func() {
			server := httptest.NewServer(http.HandlerFunc(DefaultHandler))
			defer server.Close()

			res, err := http.Get(server.URL)
			if err != nil {
				log.Fatal(err)
			}

			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}

			Ω(string(body)).Should(ContainSubstring("go-to-sample"))
		})
	})

})
