package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/bin16/dino/colors"
	"github.com/bin16/dino/dino"
)

var config = parseArgs()

func main() {
	http.HandleFunc("/", appEntryFunc)
	http.HandleFunc("/dino/", dinoImageFunc)
	http.HandleFunc("/img2/", dinoImageV2Func)
	http.HandleFunc("/rand/", dinoEntryV2Func)
	http.HandleFunc("/rand/v1/", dinoEntryFunc)
	http.ListenAndServe(":"+strconv.Itoa(config.port), nil)
}

func appEntryFunc(w http.ResponseWriter, r *http.Request) {
	maxAge := int(7 * 24 * 60 * 60)
	cacheControl := fmt.Sprint("public; max-age=", int(maxAge))
	w.Header().Set("Cache-Control", cacheControl)
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, filepath.Join(config.workDir, "index.html"))
}

func dinoEntryFunc(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.RequestURI, r.RemoteAddr, r.UserAgent())

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	dd, err := dino.NewDino(config.tplPath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	dd.UseTag("xmas")
	dd.UseRandomColors()
	dd.Draw()
	dd.SetWorkDir(config.workDir)
	targetPath, err := dd.AutoExport()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	filename := filepath.Base(targetPath)
	url := "/dino/" + filename[5:5+20]
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func dinoEntryV2Func(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.RequestURI, r.RemoteAddr, r.UserAgent())

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	query := r.URL.Query()
	q := query.Get("q")
	if len(q) == 0 {
		t := time.Now()
		q = strconv.Itoa(int(t.Unix()))
	}

	h := sha256.New()
	h.Write([]byte(q))
	keys := h.Sum(nil)
	pal := colors.MakePal(keys)

	hashed := hex.EncodeToString(keys)
	filename := hashed + ".png"
	targetPath := path.Join(config.workDir, filename)

	dino := colors.NewPixel(config.tplPath, "")
	// dino.SetTag("xmas")
	dino.SetPalette(pal)
	dino.LoadTplImg()
	dino.ExportAs(targetPath)

	url := "/img2/" + filename
	url += "?q=" + q
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

/*
	/img/fullpath
*/
func dinoImageV2Func(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.RequestURI, r.RemoteAddr, r.UserAgent())

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	filename := r.URL.Path[6:]
	fullpath := filepath.Join(config.workDir, filename)

	maxAge := int(365 * 100 * 24 * 60 * 60)
	cacheControl := fmt.Sprint("public; max-age=", int(maxAge))
	w.Header().Set("Cache-Control", cacheControl)
	w.Header().Set("Content-Type", "image/png")
	http.ServeFile(w, r, fullpath)
}

func dinoImageFunc(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.RequestURI, r.RemoteAddr, r.UserAgent())

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	name := r.RequestURI[6:]
	filename := filepath.Join(config.workDir, "dino-"+name+".png")

	maxAge := int(365 * 100 * 24 * 60 * 60)
	cacheControl := fmt.Sprint("public; max-age=", int(maxAge))
	w.Header().Set("Cache-Control", cacheControl)
	w.Header().Set("Content-Type", "image/png")
	http.ServeFile(w, r, filename)
}

type settings struct {
	workDir string
	tplPath string
	port    int
}

func parseArgs() settings {
	var port int
	var workDir, tplPath string
	flag.IntVar(&port, "port", 5000, "--port=xxxx")
	flag.StringVar(&workDir, "dir", "dino_images", "--dir=/path/to/output")
	flag.StringVar(&tplPath, "tpl", "dino.tpl.png", "--tpl=/path/to/tpl.png")
	flag.Parse()

	return settings{
		workDir,
		tplPath,
		port,
	}
}
