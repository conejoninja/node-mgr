package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(allCmd)
}

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Start a Ethereum bootnode, a node and a miner",
	Long:  `Start a Ethereum bootnode, a node and a miner`,
	Run: func(cmd *cobra.Command, args []string) {

		startBootnode(cfg.BaseName+"-bootnode")
		startNode(cfg.BaseName+"-node1", "")
		startNode(cfg.BaseName+"-miner1", cfg.EtherBase)

	},
}
