package cmd

import "github.com/spf13/cobra"

const (
	Version = "1.0.0"
	IsoDate = "2006-01-02"
)

var PostRunMsg = func(cmd *cobra.Command, args []string) {
	log.Infof("\n\n \t\t  Follow @samit_gh in twitter\n")
}
