package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"strings"
	"fmt"
)

func init() {
	rootCmd.AddCommand(bootnodeCmd)
}


var bootnodeCmd = &cobra.Command{
	Use:   "bootnode",
	Short: "Start a bootnode, so your network is easily discoverable",
	Long:  `Start a bootnode, so your network is easily discoverable`,
	Run: func(cmd *cobra.Command, args []string) {

			startBootnode(cfg.BaseName+"-bootnode")

	},
}

func startBootnode(bootnode string) {
	exe_cmd("docker", "stop", bootnode)
	exe_cmd("docker", "rm", bootnode)

	os.MkdirAll(cfg.BootnodePath, 0777)
	file, err := os.Open(cfg.BootnodePath + "/boot.key")

	if err != nil && strings.Contains(err.Error(), "no such file or directory") {
		fmt.Printf("%s\n", blue(fmt.Sprintf("%s/boot.key not found, generating...", cfg.BootnodePath)))

		exe_cmd("docker", "run", "--rm",
			"-v", cfg.BootnodePath+":/opt/bootnode",
			cfg.BaseName+":alltools-dev",
			"bootnode", "--genkey", "/opt/bootnode/boot.key",
		)

		fmt.Printf("%s\n", green("DONE"))
	}
	defer file.Close()

	out := exe_cmd("docker", "network", "ls")
	if !strings.Contains(out, cfg.BaseName) {
		exe_cmd("docker", "network", "create", cfg.BaseName)
	}

	exe_cmd("docker", "run", "-d",
		"--name", bootnode,
		"-v", cfg.BootnodePath+":/opt/bootnode",
		"--network", cfg.BaseName,
		cfg.BaseName+":alltools-dev", "bootnode",
		"--nodekey", "/opt/bootnode/boot.key", "--verbosity=3",
	)

}