package main_test

import (
	"github.com/cloudfoundry/cli/plugin/fakes"
	io_helpers "github.com/cloudfoundry/cli/testhelpers/io"
	. "github.com/dmikusa-pivotal/support_plugin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
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
			Expect(output).To(ContainElement(ContainSubstring("## Output from `target`")))
			Expect(output).To(ContainElement(ContainSubstring("## Output from `apps`")))
			Expect(output).To(ContainElement(ContainSubstring("## Output from `services`")))
		})
	})
	Describe("PromptForYesNo", func() {
		PromptForInput := func(input string) ([]string, bool) {
			var answer bool
			output := io_helpers.CaptureOutput(func() {
				io_helpers.SimulateStdin(input, func(r io.Reader) {
					answer = PromptForYesNo(r, "Is this a dumb question?")
				})
			})
			return output, answer
		}

		It("returns true when 'y' or 'Y' is pressed", func() {
			output, answer := PromptForInput("y\n")
			Expect(output).To(ContainElement(ContainSubstring("Is this a dumb question? (y/n)")))
			Expect(answer).To(BeTrue())
			output, answer = PromptForInput("Y\n")
			Expect(output).To(ContainElement(ContainSubstring("Is this a dumb question? (y/n)")))
			Expect(answer).To(BeTrue())
		})
		It("returns false when anything else is pressed", func() {
			output, answer := PromptForInput("n\n")
			Expect(output).To(ContainElement(ContainSubstring("Is this a dumb question? (y/n)")))
			Expect(answer).To(BeFalse())
			output, answer = PromptForInput("N\n")
			Expect(output).To(ContainElement(ContainSubstring("Is this a dumb question? (y/n)")))
			Expect(answer).To(BeFalse())
			output, answer = PromptForInput("\n")
			Expect(output).To(ContainElement(ContainSubstring("Is this a dumb question? (y/n)")))
			Expect(answer).To(BeFalse())
			output, answer = PromptForInput("askdjfkdj\n")
			Expect(output).To(ContainElement(ContainSubstring("Is this a dumb question? (y/n)")))
			Expect(answer).To(BeFalse())
			output, answer = PromptForInput("kdjf kdjf\n")
			Expect(output).To(ContainElement(ContainSubstring("Is this a dumb question? (y/n)")))
			Expect(answer).To(BeFalse())
		})
	})
})
