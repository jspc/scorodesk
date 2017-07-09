package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/zeebox/go-http-middleware"
	"github.com/zeebox/goose4"
)

var (
	FilePath = flag.String("c", "testdata/config.json", "config file")

	C Config

	ErrorLogger *log.Logger
)

// goose4
var (
	buildNumber  string
	buildMachine string
	builtBy      string
	builtWhen    string
	compiler     string
	sha          string
)

func main() {
	flag.Parse()

	var err error
	ErrorLogger = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)

	C, err = LoadConfig(*FilePath)
	if err != nil {
		ErrorLogger.Panic(err)
	}

	t, err := strconv.Atoi(builtWhen)
	if err != nil {
		ErrorLogger.Panic(err)
	}

	c := goose4.Config{
		ArtifactID:      "scorodesk",
		BuildNumber:     buildNumber,
		BuildMachine:    buildMachine,
		BuiltBy:         builtBy,
		BuiltWhen:       time.Unix(int64(t), 0),
		CompilerVersion: compiler,
		GitSha:          sha,
		RunbookURI:      ``,
		Version:         buildNumber,
	}
	se4, err := goose4.NewGoose4(c)

	router := httprouter.New()
	router.POST("/incoming", Incoming)
	router.NotFound = middleware.NewMiddleware(se4)

	ErrorLogger.Fatal(http.ListenAndServe(C.Listen, middleware.NewMiddleware(router)))
}
