package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/ddosakura/gklang"
	"github.com/ddosakura/pdf-designer/watcher"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"golang.org/x/net/websocket"
)

var (
	workCmd = &cobra.Command{
		Use:   "work",
		Short: "Design pdf",
		Long:  `Design pdf, and print PDF after exit.`,
		Run: func(cmd *cobra.Command, args []string) {
			port := ":" + strconv.Itoa(workPort)

			// TODO: 优化（考虑第三方库）
			// open the web browser
			run, ok := openCommands[runtime.GOOS]
			if !ok {
				println("WARNING", "don't know how to open things on %s platform", runtime.GOOS)
			}
			c := exec.Command(run, "http://localhost"+port)
			go c.Start()

			http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(afero.NewHttpFs(fs))))

			wp := http.FileServer(http.Dir(workSrc))
			http.Handle("/wp/", http.StripPrefix("/wp/", wp))

			go watcher.Init(workSrc)
			http.Handle("/fresh", websocket.Handler(watcher.WsFreshHandler))

			http.HandleFunc("/", preview)

			http.ListenAndServe(port, nil)
		},
	}

	workPort int
	workSrc  string
)

func init() {
	workCmd.PersistentFlags().IntVarP(&workPort, "port", "p", 8080, "port for preview page")
	workCmd.PersistentFlags().StringVarP(&workSrc, "src", "s", "./src", "the work directory")
}

var openCommands = map[string]string{
	"windows": "cmd /c start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func preview(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(200)
	// content-type: text/html; charset=utf-8
	rw.Write([]byte(getHTML()))
}

func getHTML() string {
	d := readFromBlock("/index.html")
	return loadTemplate(d)
}

func readFromBlock(name string) string {
	f, e := fs.Open(name)
	defer f.Close()
	if e != nil {
		gklang.Er(e)
	}
	b, e := ioutil.ReadAll(f)
	if e != nil {
		gklang.Er(e)
	}
	return string(b)
}

func readFromWp(name string) string {

	f, e := os.Open(filepath.Join(workSrc, name))
	defer f.Close()
	if dataSuffix.MatchString(name) && os.IsNotExist(e) {
		f.Close()
		name = dataSuffix.ReplaceAllString(name, ".default.txt")
		f, e = os.Open(filepath.Join(workSrc, name))
	}
	if e != nil {
		gklang.Er(e)
	}
	b, e := ioutil.ReadAll(f)
	if e != nil {
		gklang.Er(e)
	}

	if dataSuffix.MatchString(name) {
		return string(b)
	}

	if jsSuffix.MatchString(name) {
		return fmt.Sprintf("<script>%s</script>", string(b))
	}

	if cssSuffix.MatchString(name) {
		return fmt.Sprintf("<style>%s</style>", string(b))
	}

	if htmlSuffix.MatchString(name) {
		return loadTemplate(string(b))
	}

	gklang.Log(gklang.LWarn, fmt.Sprintf("%s has an unknow suffix", name))
	return string(b)
}

var (
	jsSuffix          = regexp.MustCompile(`.js$`)
	cssSuffix         = regexp.MustCompile(`.css$`)
	htmlSuffix        = regexp.MustCompile(`.html$`)
	dataSuffix        = regexp.MustCompile(`.txt$`)
	defaultDataSuffix = regexp.MustCompile(`.default.txt$`)
)

const (
	hashText1 = "aksdfsfkbnkevelkmvnkjdnvlmalacvsfv"
	hashText2 = "pboytjoibepdvvnkjdnvlmvsdvfdvvssfv"
)

func loadTemplate(t string) string {
	tx := strings.ReplaceAll(t, "\\\\", hashText1)
	tx = strings.ReplaceAll(tx, "\\$", hashText2)

	reg := regexp.MustCompile(`\${[^}]*}`)
	b := reg.MatchString(tx)
	if !b {
		return t
	}

	list := reg.FindAllString(t, -1)
	gklang.Log(gklang.LDebug, "find", list)
	for _, item := range list {
		name := strings.TrimPrefix(item, "${")
		name = strings.TrimSuffix(name, "}")
		gklang.Log(gklang.LDebug, "loading", name)
		tx = strings.ReplaceAll(tx, item, readFromWp(name))
	}

	tx = strings.ReplaceAll(tx, hashText1, "\\\\")
	tx = strings.ReplaceAll(tx, hashText2, "\\$")
	return tx
}
