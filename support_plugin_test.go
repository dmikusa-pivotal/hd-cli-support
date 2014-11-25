package main_test

import (
	"github.com/cloudfoundry/cli/plugin/fakes"
	io_helpers "github.com/cloudfoundry/cli/testhelpers/io"
	. "github.com/dmikusa-pivotal/support-plugin"
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

		It("does nothing", func() {
			io_helpers.CaptureOutput(func() {
				supportPlugin.Run(fakeCliConnection, []string{})
			})
			Expect(fakeCliConnection.CliCommandCallCount()).To(BeZero())
		})
	})
})
