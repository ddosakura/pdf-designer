package cmd

import (
	"fmt"

	"github.com/ddosakura/gklang"
	"github.com/ddosakura/pdf-designer/sblock"
	sbgo "github.com/ddosakura/sblock/libs/golang"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "pdf-designer",
		Short: "a designer for PDF, useful for writing resumes.",
		Long: `A designer for PDF, useful for writing resumes.
When I started writing my resume, I saw many websites that provided templates.
These websites provide online resume design.
So I came up with the idea,
"As a front-end worker, why don't I create an exclusive design tool for my resume?"`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)

var (
	fs afero.Fs
)

// Execute the CLI
func Execute() {
	var err error
	fs, err = sbgo.New(sblock.Raw, "default")
	if err != nil {
		gklang.Er(err)
	}

	fmt.Println(logo)

	rootCmd.Version = version
	rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(saveCmd)
	rootCmd.AddCommand(workCmd)
}
