module code.cloudfoundry.org/eirini-loggregator-bridge

require (
	code.cloudfoundry.org/eirinix v0.3.1-0.20200908072226-2c03042398ea
	code.cloudfoundry.org/go-loggregator/v8 v8.0.3
	github.com/mitchellh/mapstructure v1.2.2 // indirect
	github.com/onsi/ginkgo v1.12.1
	github.com/onsi/gomega v1.10.1
	github.com/spf13/cobra v0.0.7
	github.com/spf13/viper v1.6.3
	go.uber.org/zap v1.15.0
	k8s.io/api v0.18.6
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v0.18.6
	sigs.k8s.io/controller-runtime v0.6.2
)

go 1.13
