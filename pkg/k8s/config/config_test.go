package config_test

import (
	"fmt"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/grafana/xk6-chaos/pkg/k8s/config"
	"k8s.io/client-go/util/homedir"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "K8s Config Suite")
}

var _ = Describe("Configuration", func() {
	When("resolving the config file", func() {
		It("should default to the homedir", func() {
			err := os.Unsetenv("K6_CHAOS_KUBECONFIG")
			path := config.GetConfigPath()

			Expect(err).NotTo(HaveOccurred())
			Expect(path).To(Equal(fmt.Sprintf("%s/.kube/config", homedir.HomeDir())))
		})

		It("should pick the value from $KUBECONFIG when available", func() {
			expected := "/some/test/path"
			err := os.Setenv("K6_CHAOS_KUBECONFIG", expected)
			path := config.GetConfigPath()

			Expect(err).NotTo(HaveOccurred())
			Expect(path).To(Equal(expected))
		})
	})
})
