package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"

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

		for _, merge := range config.Merges {
			for _, target := range merge.Targets {
				basePath := helper.CleanJoinPath(config.Base, target)
				if !filepath.IsAbs(config.Base) {
					basePath = helper.CleanJoinPath(filepath.Dir(configFile), basePath)
				}

				overlayPath := helper.CleanJoinPath(filepath.Dir(configFile), target)

				// read base file
				base, err := helper.ReadFile(basePath)
				if err != nil {
					return err
				}
				baseMap := make(map[string]interface{})
				if err = json.Unmarshal([]byte(base), &baseMap); err != nil {
					return err
				}

				// read overlay file
				overlay, err := helper.ReadFile(overlayPath)
				if err != nil {
					return err
				}
				overlayMap := make(map[string]interface{})
				if err = json.Unmarshal([]byte(overlay), &overlayMap); err != nil {
					return err
				}

				resultMap := helper.MergeMap(baseMap, overlayMap, helper.MergeMode(merge.Mode))
				result, err := json.Marshal(resultMap)
				if err != nil {
					return err
				}

				var out bytes.Buffer
				if config.Format {
					err = json.Indent(&out, result, "", "  ")
					if err != nil {
						return err
					}
				} else {
					if _, err = out.Write(result); err != nil {
						return err
					}
				}

				outputPath := helper.CleanJoinPath(config.Output, config.Namespace, target)
				if !filepath.IsAbs(outputPath) {
					outputPath = helper.CleanJoinPath(filepath.Dir(configFile), outputPath)
				}
				if err := helper.WriteFile(filepath.Dir(outputPath), filepath.Base(outputPath), out.Bytes()); err != nil {
					return err
				}
			}
		}

		return nil
	},
}

type Config struct {
	// Namespace
	Namespace string

	// Base
	Base string

	// Output
	Output string

	// Format
	Format bool

	// Merges
	Merges []MergeConfig
}

type MergeConfig struct {
	// Mode is merge mode.
	// value are add or noAdd
	Mode string

	// Targets are targets of merge
	Targets []string
}

var config = Config{}

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

	if err := viper.Unmarshal(&config); err != nil {
		cobra.CheckErr(err)
	}
}
