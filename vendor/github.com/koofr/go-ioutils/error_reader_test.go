package ioutils_test

import (
	"errors"
	. "github.com/koofr/go-ioutils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ErrorReader", func() {
	It("should return error", func() {
		e := errors.New("my error")
		r := NewErrorReader(e)

		n, err := r.Read([]byte{0})
		Expect(n).To(Equal(0))
		Expect(err).To(Equal(e))

		n, err = r.Read([]byte{0})
		Expect(n).To(Equal(0))
		Expect(err).To(Equal(e))
	})
})
