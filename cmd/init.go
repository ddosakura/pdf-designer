package cmd

import (
	"compress/flate"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/ddosakura/gklang"
	"github.com/mholt/archiver"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Init project from a template file",
		Long:  `Init project from a template file.`,
		Run: func(cmd *cobra.Command, args []string) {
			if !builtInTemplate(initTemplate) {
				unpkg()
			}
		},
	}

	initTemplate string
	initSrc      string
)

func init() {
	initCmd.PersistentFlags().StringVarP(&initTemplate, "template", "t", "origin", "the path to template file")
	initCmd.PersistentFlags().StringVarP(&initSrc, "src", "s", "./src", "the path to project")
}

func unpkg() {
	tg := archiver.NewTarGz()
	tg.CompressionLevel = flate.BestCompression
	tg.OverwriteExisting = true

	err := tg.Unarchive(initTemplate, initSrc)
	if err != nil {
		gklang.Er(err)
	}
}

func builtInTemplate(t string) (has bool) {
	root := "/template/" + t
	has, e := afero.Exists(fs, root)
	if e != nil {
		gklang.Er(e)
	}
	if has {
		e = afero.Walk(fs, root, func(p string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				b, e := afero.ReadFile(fs, p)
				if e != nil {
					gklang.Er(e)
				}

				rp := filepath.Join(initSrc, strings.Replace(p, root, "", 1))
				gklang.Log(gklang.LInfo, rp)

				e = os.MkdirAll(path.Dir(rp), 0755)
				if e != nil {
					gklang.Er(e)
				}

				f, e := os.Create(rp)
				defer f.Close()
				if e != nil {
					gklang.Er(e)
				}

				_, e = f.Write(b)
				if e != nil {
					gklang.Er(e)
				}
			}
			return nil
		})
		if e != nil {
			gklang.Er(e)
		}
	}
	return
}
