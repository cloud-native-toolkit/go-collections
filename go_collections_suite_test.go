package go_collections_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoCollections(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoCollections Suite")
}
