package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/j4ng5y/go-lumber"
	"github.com/j4ng5y/tfstrap/tfstrap"
	"github.com/spf13/cobra"
)

var (
	// Init Logging
	log = lumber.New()

	// Setup CLI variables
	remoteStateProvider string

	// Setup CLI
	rootCmd = &cobra.Command{
		Use:     "tfstrap",
		Version: "0.1.0",
		Short:   "A simple bootstrapper for Terraform that uses good defaults",
		Long:    "",
		Args:    cobra.ExactArgs(1),
		Run:     rootFunc,
	}
)

func rootFunc(ccmd *cobra.Command, args []string) {
	D := &tfstrap.DirectoryStructure{
		RootDirPath:          args[0],
		DirectoryPermissions: 0660,
		ConfigFile:           "config.tf",
		VariablesFile:        "variables.tf",
		ModulesDirectoryName: "_mods",
	}

	s, err := os.Stat(args[0])
	if err != nil {
		if os.IsNotExist(err) {
			if err := D.Write(); err != nil {
				log.Fatal(err.Error())
			}
		}
		log.Fatal(err.Error())
	}
	if !s.IsDir() {
		log.Fatal(fmt.Sprintf("'%s' is a regular file, please remove and try again", args[0]))
	}
	fmt.Printf("'%s' already exists. Would you like to overwrite it? (y/N): ", args[0])
	reader := bufio.NewReader(os.Stdin)
	c, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal(err.Error())
	}
	switch c {
	case 'Y':
		if err := os.RemoveAll(args[0]); err != nil {
			log.Fatal(err.Error())
		}
		if err := D.Write(); err != nil {
			log.Fatal(err.Error())
		}
	case 'y':
		if err := os.RemoveAll(args[0]); err != nil {
			log.Fatal(err.Error())
		}
		if err := D.Write(); err != nil {
			log.Fatal(err.Error())
		}
	case 'N':
		os.Exit(1)
	case 'n':
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&remoteStateProvider, "remote-state-provider", "r", "", "The remote state provider to bootstrap")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}
