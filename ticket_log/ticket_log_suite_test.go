package ticket_log_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTicketLog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TicketLog Suite")
}
