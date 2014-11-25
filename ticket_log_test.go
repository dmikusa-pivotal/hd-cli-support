package main_test

import (
	"bytes"
	. "github.com/dmikusa-pivotal/support-plugin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TicketLog", func() {
	var tl TicketLog
	var buf *bytes.Buffer

	BeforeEach(func() {
		buf = &bytes.Buffer{}
		tl = TicketLog{
			Name:   "in-memory",
			Writer: buf,
		}
	})

	Describe("Basic Properties", func() {
		It("can be created", func() {
			Expect(tl.Name).To(HavePrefix("in-memory"))
		})
	})

	Describe("File Actions", func() {
		It("can write", func() {
			tl.Append(TicketEntry{
				Description: "Hello World!",
				Body:        nil,
			})
			resp, _ := buf.ReadString('\n')
			Expect("## Hello World!\n").Should(Equal(resp))
			resp, _ = buf.ReadString('\n')
			Expect("```")
			resp, _ = buf.ReadString('\n')
			Expect("")
			resp, _ = buf.ReadString('\n')
			Expect("```")
		})
		It("can write with a body", func() {
			tl.Append(TicketEntry{
				Description: "Hello World!",
				Body:        []byte("Some body text"),
			})
			resp, _ := buf.ReadString('\n')
			Expect("## Hello World!\n").Should(Equal(resp))
			resp, _ = buf.ReadString('\n')
			Expect("```")
			resp, _ = buf.ReadString('\n')
			Expect("Some body text")
			resp, _ = buf.ReadString('\n')
			Expect("```")
		})
	})
})
