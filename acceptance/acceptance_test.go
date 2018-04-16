package acceptance_test

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

const TIMEOUT = "5s"

var (
	tempDir    string
	configFile string
)

var _ = Describe("Acceptance", func() {
	BeforeEach(func() {
		var err error
		tempDir, err = ioutil.TempDir("", "")
		Expect(err).NotTo(HaveOccurred())
		configFile = filepath.Join(tempDir, "config.yml")
	})

	AfterEach(func() {
		os.RemoveAll(tempDir)
	})

	It("Reads the config file", func() {
		config := []byte("some-yaml")
		err := ioutil.WriteFile(configFile, config, os.ModePerm)
		Expect(err).NotTo(HaveOccurred())
		cmd := exec.Command(binPath, "-config", configFile)
		session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session.Out).Should(gbytes.Say(string(config)))
		Eventually(session, TIMEOUT).Should(gexec.Exit(0))
	})

	Context("When the -config flag is not provided", func() {
		It("Exits 1 and outputs usage", func() {
			cmd := exec.Command(binPath)
			session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())

			Eventually(session.Err).Should(gbytes.Say("Usage: flo -config FILE"))
			Eventually(session, TIMEOUT).Should(gexec.Exit(1))
		})
	})

	Context("When the config file cannot be read", func() {
		It("Exits 1 and outputs the error", func() {
			cmd := exec.Command(binPath, "-config", "/path/not/found/config.yml")
			session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())

			Eventually(session.Err).Should(gbytes.Say("Error reading config:"))
			Eventually(session, TIMEOUT).Should(gexec.Exit(1))
		})
	})
})
