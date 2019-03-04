package cmd

import (
	"net/http"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
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
