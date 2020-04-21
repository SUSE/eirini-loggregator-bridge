package cmd

import (
	"fmt"
	"os"

	eirinix "github.com/SUSE/eirinix"

	configpkg "github.com/SUSE/eirini-loggregator-bridge/config"
	. "github.com/SUSE/eirini-loggregator-bridge/logger"
	podwatcher "github.com/SUSE/eirini-loggregator-bridge/podwatcher"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var kubeconfig string

var config configpkg.ConfigType

var rootCmd = &cobra.Command{
	Use:   "eirini-loggregator-bridge",
	Short: "eirini-loggregator-bridge streams Eirini application logs to CloudFoundry loggregator",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		err = config.Validate()
		if err != nil {
			LogError(err.Error())
			os.Exit(1)
		}

		filter := false
		x := eirinix.NewManager(
			eirinix.ManagerOptions{
				Namespace:           config.Namespace,
				KubeConfig:          kubeconfig,
				OperatorFingerprint: "eirini-loggregator-bridge", // Not really used for now, but setting it up for future
				FilterEiriniApps:    &filter,
			})

		x.AddWatcher(podwatcher.NewPodWatcher(config))

		err = x.Watch()
		if err != nil {
			LogError(err.Error())
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		LogError(err.Error())
		os.Exit(1)
	}
}

// Loggregator TLS:
// https://github.com/cloudfoundry/go-loggregator/blob/master/tls.go
// https://docs.cloudfoundry.org/loggregator/architecture.html
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	rootCmd.PersistentFlags().StringVar(&kubeconfig, "kubeconfig", "", "kubeconfig file path. This is optional, in cluster config will be used if not set")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			LogError("Can't read config:", err.Error())
			os.Exit(1)
		}
	}

	viper.Unmarshal(&config)
	viper.AutomaticEnv() // read in environment variables that match

	LogDebug("Namespace: ", fmt.Sprintf("%s", viper.GetString("namespace")))
	LogDebug("Loggregator-endpoint: ", fmt.Sprintf("%s", viper.GetString("loggregator-endpoint")))
	LogDebug("Loggregator-ca-path: ", fmt.Sprintf("%s", viper.GetString("loggregator-ca-path")))
	LogDebug("Loggregator-cert-path: ", fmt.Sprintf("%s", viper.GetString("loggregator-cert-path")))
	LogDebug("Loggregator-key-path: ", fmt.Sprintf("%s", viper.GetString("loggregator-key-path")))
}
