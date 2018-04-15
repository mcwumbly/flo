package acceptance_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

var binPath string

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Acceptance Test Suite")
}

var _ = BeforeSuite(func() {
	var err error
	srcPath := os.Getenv("GOPATH") + "/src/github.com/mcwumbly/flo/main.go"
	binPath, err = gexec.Build(srcPath, "-race")
	Expect(err).NotTo(HaveOccurred())
})
