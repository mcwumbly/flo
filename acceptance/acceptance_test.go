package acceptance_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

const TIMEOUT = "5s"

var _ = Describe("Acceptance", func() {
	It("Executes successfully", func() {
		session, err := gexec.Start(exec.Command(binPath), GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session, TIMEOUT).Should(gexec.Exit(0))
	})

})
