package cmd

import (
	"compress/flate"
	"crypto/md5"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/ddosakura/gklang"
	"github.com/mholt/archiver"
	"github.com/spf13/cobra"
)

var (
	saveCmd = &cobra.Command{
		Use:   "save",
		Short: "Save a template file from project",
		Long:  `Save a template file from project.`,
		Run: func(cmd *cobra.Command, args []string) {
			tmp := pkgS()
			rename(tmp, savePath, saveTemplate+".pdt")
		},
	}

	savePath     string
	saveTemplate string
	saveSrc      string
)

func init() {
	saveCmd.PersistentFlags().StringVarP(&savePath, "path", "p", "./", "the path to template file")
	saveCmd.PersistentFlags().StringVarP(&saveTemplate, "template", "t", "tempate", "the name of template file")
	saveCmd.PersistentFlags().StringVarP(&saveSrc, "src", "s", "./src", "the path to project")
}

func pkgS() string {
	tg := archiver.NewTarGz()
	tg.CompressionLevel = flate.BestCompression
	tg.OverwriteExisting = true
	// tg.Tar.MkdirAll = true
	// e := tg.Archive(allFile, tgName)

	rand := MD5(saveSrc) + ".tar.gz"
	list, err := filepath.Glob(saveSrc + "/*")
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

/*
func listcheckDefault(list []string) []string {
	l := make([]string, 0)
	m := make(map[string]bool)
	for _, item := range list {
		d := dataSuffix.MatchString(item)
		dd := defaultDataSuffix.MatchString(item)
		if d {
			if dd {
				k := defaultDataSuffix.ReplaceAllString(item, ".txt")
				m[k] = false
				l = append(l, item)
			} else {
				m[item] = true
			}
		} else {
			l = append(l, item)
		}
	}
	for k, v := range m {
		if v {
			l = append(l, k)
		}
	}
	return l
}
*/

// MD5 util
func MD5(i string) string {
	h := md5.New()
	h.Write([]byte(i))
	return string(h.Sum(nil))
}

/*
func walkArchive(path string, eDirs *[]string, all []string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		gklang.Er(err)
	}
	dirs := make([]string, 0)
	for _, v := range files {
		p := filepath.Join(path, v.Name())
		if v.IsDir() {
			if !inStringArray(v.Name(), eDirs) {
				dirs = append(dirs, p)
			}
		} else {
			gklang.Log(gklang.LDebug, "archive", p)
			// e := tg.Archive([]string{p}, tgName)
			// if e != nil {
			// 	er(e)
			// }
			all = append(all, p)
		}
	}
	for _, v := range dirs {
		all = walkArchive(v, eDirs, all)
	}
	return all
}

func inStringArray(value string, arr *[]string) bool {
	for _, v := range *arr {
		if value == v {
			return true
		}
	}
	return false
}
*/

func rename(src, dest, filename string) {
	err := os.MkdirAll(dest, 0755)
	if err != nil {
		gklang.Er(err)
	}

	targetFile := path.Join(dest, filename)

	// Try to rename generated source.
	if err := os.Rename(src, targetFile); err == nil {
		return
	}
	// If the rename failed (might do so due to temporary file residing on a
	// different device), try to copy byte by byte.
	rc, err := os.Open(src)
	if err != nil {
		gklang.Er(err)
	}
	defer func() {
		rc.Close()
		os.Remove(src) // ignore the error, source is in tmp.
	}()

	wc, err := os.Create(targetFile)
	if err != nil {
		gklang.Er(err)
	}
	defer wc.Close()

	if _, err = io.Copy(wc, rc); err != nil {
		// Delete remains of failed copy attempt.
		os.RemoveAll(dest)
	}
}
