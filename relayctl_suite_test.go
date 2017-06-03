package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRelayctl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Relayctl Suite")
}
