package executor_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/kun-lun/common/configuration"
	clogger "github.com/kun-lun/common/logger"
	"github.com/kun-lun/common/storage"
	"github.com/kun-lun/executor/commands"
	. "github.com/kun-lun/executor/pkg/apis"
)

var _ = Describe("Executor", func() {

	var (
		executor Executor
		config   configuration.Configuration
	)

	Describe("Run", func() {
		Context("Command not supported", func() {

			BeforeEach(func() {

				logger := clogger.NewLogger(os.Stdout, os.Stdin)
				usage := commands.NewUsage(logger)
				config = configuration.Configuration{
					Command: "helpx",
				}
				executor = NewExecutor(config, usage, logger, storage.Store{})
			})
			It("should raise one error", func() {
				err := executor.Run()
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
