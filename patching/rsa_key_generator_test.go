package patching_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kun-lun/common/helpers"
)

var _ = Describe("RsaKeyGenerator", func() {

	var (
		generator RSAKeyGenerator
	)

	BeforeEach(func() {
		generator = NewRSAKeyGenerator()
	})
	Describe("Generate", func() {
		Context("Everything OK", func() {
			It("should succeed", func() {
				rsa_key, err := generator.Generate(nil)
				Expect(err).To(BeNil())
				Expect(rsa_key).NotTo(BeNil())
			})
		})
	})
})
