package patching_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kun-lun/common/helpers"
)

var _ = Describe("PasswordGenerator", func() {

	var (
		generator PasswordGenerator
	)

	BeforeEach(func() {
		generator = NewPasswordGenerator()
	})
	Describe("Generate", func() {
		Context("Everything OK", func() {
			It("should succeed", func() {
				password, err := generator.Generate(0)
				Expect(err).To(BeNil())
				Expect(password).NotTo(BeNil())
				Expect(len(password)).To(Equal(20))
			})
		})
	})
})
