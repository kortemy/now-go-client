package now_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestNow(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Now Suite")
}

var _ = BeforeSuite(func() {
})
