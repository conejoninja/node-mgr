package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(wipeCmd)
}

var wipeCmd = &cobra.Command{
	Use:   "wipe",
	Short: "Wipe all data to start a new private ethereum network",
	Long:  `Wipe all data to start a new private ethereum network`,
	Run: func(cmd *cobra.Command, args []string) {


		exe_cmd("docker", "stop", cfg.BaseName+"-bootnode")
		exe_cmd("docker", "stop", cfg.BaseName+"-node1")
		exe_cmd("docker", "stop", cfg.BaseName+"-miner1")

		exe_cmd("docker", "rm", cfg.BaseName+"-bootnode")
		exe_cmd("docker", "rm", cfg.BaseName+"-node1")
		exe_cmd("docker", "rm", cfg.BaseName+"-miner1")

		os.RemoveAll(cfg.BootnodePath)
		os.RemoveAll(cfg.HashPath)
		// TODO, take from configuration how many to remove
		os.RemoveAll(cfg.DataPath+"."+cfg.BaseName+"-node1")
		os.RemoveAll(cfg.DataPath+"."+cfg.BaseName+"-miner1")

	},
}
