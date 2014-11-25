package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSupportPlugin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SupportPlugin Suite")
}
