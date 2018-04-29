package acceptance_test

import (
	"fmt"
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
	config     string
	outputFile string
)

var _ = Describe("Acceptance", func() {
	BeforeEach(func() {
		var err error
		tempDir, err = ioutil.TempDir("", "")
		Expect(err).NotTo(HaveOccurred())

		relInputDir := "some-input-dir"
		relInputFile := filepath.Join(relInputDir, "input.txt")

		err = os.Mkdir(filepath.Join(tempDir, relInputDir), os.ModePerm)
		Expect(err).NotTo(HaveOccurred())

		err = ioutil.WriteFile(filepath.Join(tempDir, relInputDir, "input.txt"), []byte("some-input"), os.ModePerm)
		Expect(err).NotTo(HaveOccurred())

		relOutputDir := "some-output-dir"
		relOutputFile := filepath.Join(relOutputDir, "output.txt")
		outputFile = filepath.Join(tempDir, relOutputDir, "output.txt")

		configFile = filepath.Join(tempDir, "config.yml")
		config = fmt.Sprintf(`---
name: some-job
tasks:
- name: some-task
  command:
    name: cp
    args:
    - %s
    - %s
  inputs:
  - %s
  outputs:
  - %s`,
			relInputFile,
			relOutputFile,
			relInputDir,
			relOutputDir)

		err = ioutil.WriteFile(configFile, []byte(config), os.ModePerm)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		os.RemoveAll(tempDir)
	})

	It("Executes the config", func() {
		cmd := exec.Command(binPath, "-config", configFile)
		session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session, TIMEOUT).Should(gexec.Exit(0))

		output, err := ioutil.ReadFile(outputFile)
		Expect(string(output)).To(Equal("some-input"))
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
