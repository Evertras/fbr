package server

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Evertras/fbr/lib/static"
)

// Config contains all configuration to run the server
type Config struct {
	// ReadStaticFilesPerRequest determines if the server will read from disk on each request
	// or use the precompiled static files.  Useful for development, should not
	// be on otherwise.
	ReadStaticFilesPerRequest bool
}

// Server is an HTTP server that will serve static content and handle web socket connections
type Server struct {
	ctx context.Context
	cfg Config
}

// New creates a new server that's ready to listen but hasn't started yet
func New(ctx context.Context, cfg Config) *Server {
	return &Server{
		ctx: ctx,
		cfg: cfg,
	}
}

// Listen will start listening and block until the server closes
func (s *Server) Listen(addr string) error {
	mux := http.NewServeMux()

	// <jank>
	// Note: the following is jank for prototype purposes, this should
	// be an in-memory file system in a for-reals app... but this is easier
	if s.cfg.ReadStaticFilesPerRequest {
		log.Println("Reading files from disk for every request, ONLY USE THIS FOR DEV MODE!")

		fileReaderFactory := func(f string, contentType string) func(w http.ResponseWriter, req *http.Request) {
			return func(w http.ResponseWriter, req *http.Request) {
				data, err := ioutil.ReadFile(f)

				if err != nil {
					log.Printf("Error reading %s: %v", f, err)
					w.WriteHeader(500)
					return
				}

				w.Header().Set("Content-Type", contentType)
				w.Write(data)
			}
		}

		mux.HandleFunc("/", fileReaderFactory("./front/index.html", "text/html"))
		mux.HandleFunc("/wasm_exec.js", fileReaderFactory("./front/wasm_exec.js", "script/javascript"))
		mux.HandleFunc("/style.css", fileReaderFactory("./front/style.css", "text/css"))
		mux.HandleFunc("/lib.wasm", fileReaderFactory("./front/lib.wasm", "application/wasm"))
	} else {
		mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			io.WriteString(w, static.StaticHtmlIndex)
		})
		mux.HandleFunc("/wasm_exec.js", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "script/javascript")
			io.WriteString(w, static.StaticJsWasmExec)
		})
		mux.HandleFunc("/style.css", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "text/css")
			io.WriteString(w, static.StaticCssStyle)
		})
		mux.HandleFunc("/lib.wasm", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/wasm")
			w.Header().Set("Content-Encoding", "gzip")
			w.Write(static.StaticLibWasm)
		})
	}
	// </jank>

	//mux.HandleFunc("/join", join(s))

	httpServer := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return httpServer.ListenAndServe()
}
