package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(minerCmd)
}

var minerCmd = &cobra.Command{
	Use:   "miner",
	Short: "Start a Ethereum miner",
	Long:  `Start a Ethereum miner`,
	Run: func(cmd *cobra.Command, args []string) {

		startNode("ethereum-miner1", cfg.EtherBase)

	},
}
