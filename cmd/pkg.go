package cmd

import (
	"compress/flate"
	"path/filepath"

	"github.com/ddosakura/gklang"
	"github.com/mholt/archiver"
	"github.com/spf13/cobra"
)

var (
	pkgCmd = &cobra.Command{
		Use:   "pkg",
		Short: "Pkg a project",
		Long:  `Pkg a project.`,
		Run: func(cmd *cobra.Command, args []string) {
			tmp := pkg()
			rename(tmp, pkgPath, pkgName)
		},
	}

	pkgPath string
	pkgName string
	pkgSrc  string
)

func init() {
	pkgCmd.PersistentFlags().StringVarP(&pkgPath, "path", "p", "./", "the path to pkg file")
	pkgCmd.PersistentFlags().StringVarP(&pkgName, "name", "n", "proj.pdp", "the name of project")
	pkgCmd.PersistentFlags().StringVarP(&pkgSrc, "src", "s", "./src", "the path to project")
}

func pkg() string {
	tg := archiver.NewTarGz()
	tg.CompressionLevel = flate.BestCompression
	tg.OverwriteExisting = true
	// tg.Tar.MkdirAll = true
	// e := tg.Archive(allFile, tgName)

	rand := MD5(pkgSrc) + ".tar.gz"
	list, err := filepath.Glob(pkgSrc + "/*")
	// list = listcheckDefault(list)
	// gklang.Log(gklang.LDebug, list)
	if err != nil {
		gklang.Er(err)
	}
	err = tg.Archive(list, rand)
	if err != nil {
		gklang.Er(err)
	}
	return rand
}
