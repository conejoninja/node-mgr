package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(buildCmd)
}

const sourceTmpFolder = "docker-tmp/go-ethereum"

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build and update the docker image with the latest version",
	Long:  `Build and update the docker image with the latest version`,
	Run: func(cmd *cobra.Command, args []string) {

		os.RemoveAll("docker-tmp/go-ethereum/")

		if cfg.SourceCode != "<nil>" && cfg.SourceCode != "" {
			// Copy from locla folder
			copy_folder(cfg.SourceCode, sourceTmpFolder)
		} else if cfg.GitRepo != "<nil>" && cfg.GitRepo != "" {
			// Clone from private repository
			exe_cmd("git", "clone", "--depth", "1", cfg.GitRepo, sourceTmpFolder)
		} else {
			// Clone from official repository
			exe_cmd("git", "clone", "--depth", "1", "https://github.com/ethereum/go-ethereum.git", sourceTmpFolder)
		}

		exe_cmd("docker", "build", "--tag", cfg.BaseName+":dev", "-f", "docker-tmp/go-ethereum/Dockerfile", "docker-tmp/go-ethereum/")
		exe_cmd("docker", "build", "--tag", cfg.BaseName+":alltools-dev", "-f", "docker-tmp/go-ethereum/Dockerfile.alltools", "docker-tmp/go-ethereum/")

	},
}
