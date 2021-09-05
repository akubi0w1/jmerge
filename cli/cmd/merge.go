package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/akubi0w1/jmerge"
	"github.com/akubi0w1/jmerge/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var MergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "merge multiple json file into one",
	Long:  "merge multiple json files into one.",
	RunE: func(cmd *cobra.Command, args []string) error {
		readConfig()

		for _, merge := range cfg.Merges {
			for _, target := range merge.Targets {
				basePath := helper.CleanJoinPath(cfg.Base, target)
				if !filepath.IsAbs(cfg.Base) {
					basePath = helper.CleanJoinPath(filepath.Dir(configFile), basePath)
				}

				overlayPath := helper.CleanJoinPath(filepath.Dir(configFile), target)

				fmt.Printf("merge: base=%s => overlay=%s\n", basePath, overlayPath)

				out, err := jmerge.MergeJSONByFile(basePath, overlayPath, merge.Mode, cfg.Format)
				if err != nil {
					return err
				}

				outputPath := helper.CleanJoinPath(cfg.Output, cfg.Namespace, target)
				if !filepath.IsAbs(outputPath) {
					outputPath = helper.CleanJoinPath(filepath.Dir(configFile), outputPath)
				}
				if err := helper.WriteFile(filepath.Dir(outputPath), filepath.Base(outputPath), out); err != nil {
					return err
				}
			}
		}
		fmt.Printf("merge complete!\n")

		return nil
	},
}

type config struct {
	// Namespace
	Namespace string

	// Base
	Base string

	// Output
	Output string

	// Format
	Format bool

	// Merges
	Merges []mergeConfig
}

// Validate validates config
func (c *config) Validate() error {
	if c.Namespace == "" {
		return fmt.Errorf("config error: namespace is required")
	}

	if c.Base == "" {
		return fmt.Errorf("config error: output is required")
	}

	if c.Output == "" {
		return fmt.Errorf("config error: output is required")
	}
	return nil
}

type mergeConfig struct {
	// Mode is merge mode.
	// value are add or noAdd
	Mode jmerge.MergeMode

	// Targets are targets of merge
	Targets []string
}

var cfg = config{}

var configFile string

func init() {
	cobra.OnInitialize(initConfig)

	MergeCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "jmerge.yaml", "config file path")
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName("jmerge")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		showPath, err := filepath.Abs(viper.ConfigFileUsed())
		if err != nil {
			showPath = viper.ConfigFileUsed()
		}
		fmt.Println("Using config file:", showPath)
	}
}

func readConfig() {
	if err := viper.ReadInConfig(); err != nil {
		cobra.CheckErr(err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		cobra.CheckErr(err)
	}

	if err := cfg.Validate(); err != nil {
		cobra.CheckErr(err)
	}
}
