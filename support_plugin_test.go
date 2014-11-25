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
			Expect(fakeCliConnection.CliCommandWithoutTerminalOutputCallCount()).To(Equal(3))
			Expect(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(0)[0]).To(Equal("target"))
			Expect(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(1)[0]).To(Equal("apps"))
			Expect(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(2)[0]).To(Equal("services"))
		})

		It("creates a ticket log", func() {
			Expect(supportPlugin.TicketLog).ShouldNot(BeNil())
		})

		It("prints a greeting", func() {
			output := io_helpers.CaptureOutput(func() {
				supportPlugin.Run(fakeCliConnection, []string{})
			})
			Expect(output[0]).To(ContainSubstring("Don't Panic!"))
		})

		It("prints ticket log", func() {
			output := io_helpers.CaptureOutput(func() {
				supportPlugin.Run(fakeCliConnection, []string{})
			})
			Expect(output[4]).To(ContainSubstring("## Output from `target`"))
			Expect(output[8]).To(ContainSubstring("## Output from `apps`"))
			Expect(output[12]).To(ContainSubstring("## Output from `services`"))
		})
	})
})
