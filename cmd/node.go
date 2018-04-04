package cmd

import (
	"log"
	"os"
	"regexp"
	"strings"

	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(nodeCmd)
}

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Start a Ethereum node",
	Long:  `Start a Ethereum node`,
	Run: func(cmd *cobra.Command, args []string) {

		startNode(cfg.BaseName+"-node1", "")

	},
}

func startNode(nodeName string, etherbase string) {
	exe_cmd("docker", "stop", nodeName)
	exe_cmd("docker", "rm", nodeName)

	enodeLine := exe_cmd("docker", "logs", cfg.BaseName+"-bootnode")

	myip := exe_cmd("docker", "exec", cfg.BaseName+"-bootnode", "ifconfig", "eth0")

	rg := regexp.MustCompile("inet addr:([^\\ ]+)\\s+Bcas")
	ip := rg.FindStringSubmatch(myip)

	if len(ip) > 1 {
		enodeLine = strings.Replace(enodeLine, "127.0.0.1", ip[1], -1)
		enodeLine = strings.Replace(enodeLine, "[::]", ip[1], -1)
	}

	tmp := strings.Split(enodeLine, "enode://")
	if len(tmp) > 1 {
		enodeLine = "enode://" + tmp[1]
	}

	enodeLine = strings.Trim(enodeLine, " \n")

	_, err := os.Open(cfg.DataPath + "genesis.json")
	if err != nil && strings.Contains(err.Error(), "no such file or directory") {
		log.Fatal(red("ERROR: No genesis.json found"))
	}

	_, err = os.Open(cfg.DataPath + "." + nodeName + "/keystore")
	if err != nil && strings.Contains(err.Error(), "no such file or directory") {
		fmt.Println(blue("ERROR: No keystore found, running 'geth init'"))
		exe_cmd("docker", "run", "--rm",
			"-v", cfg.DataPath+"."+nodeName+":/root/.ethereum",
			"-v", cfg.DataPath+"genesis.json:/opt/genesis.json",
			cfg.BaseName+":dev",
			"init", "/opt/genesis.json",
		)
	}

	var nodeArgs []string
	if etherbase == "" {
		nodeArgs = []string{
			"run", "-d",
			"--name", nodeName,
			"--network", cfg.BaseName,
			"-v", cfg.DataPath + "." + nodeName + ":/root/.ethereum",
			"-v", cfg.HashPath + ":/root/.ethash",
			"-v", cfg.DataPath + "genesis.json:/opt/genesis.json",
			//"-p", "8545:8545",
			cfg.BaseName+":dev",
			"--bootnodes=" + enodeLine,
			//"--rpc", "--rpcaddr=0.0.0.0", "--rpcapi=db,eth,net,web3,personal", "--rpccorsdomain", "\"*\"",
			"--cache=512", "--verbosity=4", "--maxpeers=3",
		}
	} else {
		nodeArgs = []string{
			"run", "-d",
			"--name", nodeName,
			"--network", cfg.BaseName,
			"-v", cfg.DataPath + "." + nodeName + ":/root/.ethereum",
			"-v", cfg.HashPath + ":/root/.ethash",
			"-v", cfg.DataPath + "genesis.json:/opt/genesis.json",
			"-p", "8545:8545",
			cfg.BaseName+":dev",
			"--bootnodes=" + enodeLine,
			"--rpc", "--rpcaddr=0.0.0.0", "--rpcapi=db,eth,net,web3,personal", "--rpccorsdomain", "\"*\"",
			"--cache=512", "--verbosity=4", "--maxpeers=3",
			"--mine", "--minerthreads=1", "--etherbase=" + etherbase,
		}
	}

	/*if etherbase != "" {
		nodeArgs = append(nodeArgs, "--mine", "--minerthreads=1", "--etherbase="+etherbase)
	}*/

	exe_cmd("docker", nodeArgs...)

}
