package cmd

import (
	"fmt"
	"os"

	"os/exec"

	"log"
	"path/filepath"
	"strings"

	"io"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	BaseName     string
	DataPath     string
	BootnodePath string
	HashPath     string
	SourceCode   string
	GitRepo      string
	EtherBase    string
}

var cfg Config
var red func(a ...interface{}) string
var yellow func(a ...interface{}) string
var green func(a ...interface{}) string
var blue func(a ...interface{}) string

var rootCmd = &cobra.Command{
	Use:   "nodemgr",
	Short: "NodeMgr is a manager for deploying Ethereum nodes",
	Long:  `NodeMgr is a manager for deploying Ethereum nodes`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	red = color.New(color.FgRed).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	green = color.New(color.FgGreen).SprintFunc()
	blue = color.New(color.FgBlue).SprintFunc()


}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func exe_cmd(cmd string, args ...string) string {
	outStr := ""

	fmt.Printf("%s %s %s\n", yellow("EXEC:"), cmd, strings.Join(args, " "))
	out, err := exec.Command(cmd, args...).CombinedOutput()
	outStr = string(out)
	if err != nil {
		if outStr != "" {
			fmt.Printf("%s %s: %s\n", red("OUTPUT:"), err, outStr)
		}
	}
	return outStr
}

func copy_folder(source string, dest string) (err error) {

	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			err = copy_folder(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err = copy_file(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

func copy_file(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		err = os.Chmod(dest, 0777)
	}

	return
}

func initConfig() {
	if _, err := os.Stat("./config.yml"); err != nil {
		fmt.Println("Error: config.yml file does not exist")
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	cfg.BaseName = fmt.Sprint(viper.Get("base_name"))
	if cfg.BaseName == "<nil>" {
		cfg.BaseName = "ethereum"
	}

	cfg.DataPath = fmt.Sprint(viper.Get("data_path"))
	if cfg.DataPath == "<nil>" || cfg.DataPath == "." {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		cfg.DataPath = dir
	}
	cfg.DataPath = strings.Replace(cfg.DataPath+"/", "//", "/", -1)
	cfg.BootnodePath = cfg.DataPath + ".bootnode"
	cfg.HashPath = cfg.DataPath + ".ethash"

	cfg.SourceCode = fmt.Sprint(viper.Get("source_code"))
	if cfg.SourceCode[len(cfg.SourceCode)-1:] == "/" {
		cfg.SourceCode = cfg.SourceCode[:len(cfg.SourceCode)-1]
	}
	cfg.GitRepo = fmt.Sprint(viper.Get("git_repo"))

	cfg.EtherBase = fmt.Sprint(viper.Get("etherbase"))

	if cfg.EtherBase == "<nil>" || cfg.EtherBase == "" {
		cfg.EtherBase = "0x0000000000000000000000000000000000000001"
	}

}
