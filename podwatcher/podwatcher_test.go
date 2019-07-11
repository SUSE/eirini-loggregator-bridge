package podwatcher_test

import (
	config "github.com/SUSE/eirini-loggregator-bridge/config"
	. "github.com/SUSE/eirini-loggregator-bridge/podwatcher"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PodWatcher", func() {

	Describe("Config", func() {
		Context("when initializing", func() {
			It("sets the config", func() {
				pw := NewPodWatcher(&config.ConfigType{Namespace: "test"})
				Expect(pw.Config).ToNot(BeNil())
				Expect(pw.Config.Namespace).To(Equal("test"))
			})
		})
	})

})