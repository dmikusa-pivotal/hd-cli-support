package main_test

import (
	"github.com/cloudfoundry/cli/plugin/fakes"
	io_helpers "github.com/cloudfoundry/cli/testhelpers/io"
	. "github.com/dmikusa-pivotal/support_plugin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SupportPlugin", func() {
	Describe(".Run", func() {
		var fakeCliConnection *fakes.FakeCliConnection
		var supportPlugin *SupportPlugin

		BeforeEach(func() {
			fakeCliConnection = &fakes.FakeCliConnection{}
			supportPlugin = &SupportPlugin{}
		})

		It("run three tasks", func() {
			io_helpers.CaptureOutput(func() {
				supportPlugin.Run(fakeCliConnection, []string{})
			})
			Expect(fakeCliConnection.CliCommandCallCount()).To(Equal(3))
			Expect(fakeCliConnection.CliCommandArgsForCall(0)[0]).To(Equal("target"))
			Expect(fakeCliConnection.CliCommandArgsForCall(1)[0]).To(Equal("apps"))
			Expect(fakeCliConnection.CliCommandArgsForCall(2)[0]).To(Equal("services"))
		})

		It("creates a ticket log", func() {
			Expect(supportPlugin.TicketLog).ShouldNot(BeNil())
		})
	})
})
