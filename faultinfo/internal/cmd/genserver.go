package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	flags "github.com/jessevdk/go-flags"
	openapi "github.com/nasa9084/go-openapi"
)

// Error type for this package
type Error string

func (e Error) Error() string { return string(e) }

// Error constants
const (
	ErrVersion Error = "OpenAPI Version must be 3.0.0"
)

type options struct {
	SpecFile string `short:"f" long:"file"`
}

func main() { os.Exit(_main()) }

func _main() int {
	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		log.Println(err)
		return 1
	}
	spec, err := openapi.Load(opts.SpecFile)
	if err != nil {
		log.Println(err)
		return 1
	}
	if err := generate(spec); err != nil {
		log.Println(err)
		return 1
	}
	return 0
}

func generate(spec *openapi.Document) error {
	if spec.Version != "3.0.0" {
		return ErrVersion
	}
	if err := generateServer(spec); err != nil {
		return err
	}
	if _, err := os.Stat("handlers_gen.go"); os.IsNotExist(err) {
		if err := generateHandlers(spec); err != nil {
			return err
		}
	}
	return nil
}

func generateServer(spec *openapi.Document) error {
	buf := bytes.Buffer{}
	buf.WriteString(`package faultinfo
// DO NOT EDITS. Automatically generated.

`)
	stdlibs := []string{
		"net/http",
	}
	extlibs := []string{
		"github.com/gorilla/mux",
	}
	writeImports(&buf, stdlibs, extlibs)
	buf.WriteString(`type Server struct {
	*mux.Router
}

func New() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}
	s.bindRoutes()
	return s
}

func (s *Server) Run(l string) error {
	return http.ListenAndServe(l, s.Router)
}

func (s *Server) bindRoutes() {
`)
	writeRoutes(&buf, spec.Paths)
	buf.WriteString("}\n")

	file, err := os.OpenFile("faultinfo_gen.go", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, &buf)
	return err
}

func writeImports(out io.Writer, stdlibs, extlibs []string) {
	if len(stdlibs) == 0 && len(extlibs) == 0 {
		return
	}

	fmt.Fprintln(out, "import (")
	for _, p := range stdlibs {
		fmt.Fprintf(out, "\t\"%s\"\n", p)
	}
	if len(stdlibs) > 0 && len(extlibs) > 0 {
		fmt.Fprint(out, "\n")
	}
	for _, p := range extlibs {
		fmt.Fprintf(out, "\t\"%s\"\n", p)
	}
	fmt.Fprintf(out, ")\n\n")
}

func writeRoutes(out io.Writer, paths openapi.Paths) {
	for path, item := range paths {
		if item.Get != nil {
			writeRoute(out, "GET", path)
		}
		if item.Put != nil {
			writeRoute(out, "PUT", path)
		}
		if item.Post != nil {
			writeRoute(out, "POST", path)
		}
		if item.Delete != nil {
			writeRoute(out, "DELETE", path)
		}
		if item.Options != nil {
			writeRoute(out, "OPTIONS", path)
		}
		if item.Head != nil {
			writeRoute(out, "HEAD", path)
		}
		if item.Patch != nil {
			writeRoute(out, "PATCH", path)
		}
		if item.Trace != nil {
			writeRoute(out, "TRACE", path)
		}
	}
}

func writeRoute(out io.Writer, method, path string) {
	fmt.Fprintf(out,
		"\ts.Router.HandleFunc(\"%s\", %s).Method(\"%s\")\n",
		path,
		strings.ToLower(method)+strings.Join(strings.Split(strings.Replace(strings.Replace(path, "{", "", -1), "}", "", -1), "/"), ""),
		method,
	)
}

func generateHandlers(spec *openapi.Document) error {
	buf := bytes.Buffer{}
	fmt.Fprintf(&buf, `package faultinfo

import (
	"net/http"
)

`)

	writeHandlers(&buf, spec.Paths)

	file, err := os.OpenFile("handlers_gen.go", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, &buf)
	return err
}

func writeHandlers(out io.Writer, paths openapi.Paths) {
	for path, item := range paths {
		if item.Get != nil {
			writeHandlerSkel(out, "GET", path, item)
		}
		if item.Put != nil {
			writeHandlerSkel(out, "PUT", path, item)
		}
		if item.Post != nil {
			writeHandlerSkel(out, "POST", path, item)
		}
		if item.Delete != nil {
			writeHandlerSkel(out, "DELETE", path, item)
		}
		if item.Options != nil {
			writeHandlerSkel(out, "OPTIONS", path, item)
		}
		if item.Head != nil {
			writeHandlerSkel(out, "HEAD", path, item)
		}
		if item.Patch != nil {
			writeHandlerSkel(out, "PATCH", path, item)
		}
		if item.Trace != nil {
			writeHandlerSkel(out, "TRACE", path, item)
		}
	}
}

func writeHandlerSkel(out io.Writer, method, path string, item *openapi.PathItem) {
	fmt.Fprintf(out, "func %s(w http.ResponseWriter, r *http.Request) {\n",
		strings.ToLower(method)+strings.Join(strings.Split(strings.Replace(strings.Replace(path, "{", "", -1), "}", "", -1), "/"), ""),
	)
	fmt.Fprint(out, "}\n\n")
}
