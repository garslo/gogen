package gogen_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGogen(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gogen Suite")
}
