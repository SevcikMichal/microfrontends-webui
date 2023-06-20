package server

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/SevcikMichal/microfrontends-webui/model"
)

func ServeSinglePageApplication(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("../web-ui/www/index.html")
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("index.html does not exist!")
			http.NotFound(w, r)
			return
		}
		log.Panic(err)
	}

	fileContents := string(data)

	tmpl, err := template.New("index.html").Parse(fileContents)
	if err != nil {
		log.Panic(err)
	}

	pageData := &model.TemplateData{
		Language:                   "en",
		AppTitle:                   "Micro Frontends WebUI",
		BaseURL:                    "/microfrontends-ui/",
		Description:                "Micro Frontends WebUI",
		MicroFrontendShellContext:  "application-shell",
		MicroFrontendSelector:      "",
		ProgresiveWebAppMode:       "false",
		ContentSecurityPolicyNonce: "123",
		TouchIcon:                  "",
		FavIcon:                    "",
	}

	err = tmpl.Execute(w, pageData)
	if err != nil {
		log.Panic(err)
	}
}

func ServeFile(w http.ResponseWriter, r *http.Request) {
	fileName := strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1]

	var matches []string
	err := filepath.Walk("../web-ui/www", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if fileName == info.Name() {
				matches = append(matches, path)
			}
		}
		return nil
	})

	if err != nil {
		if os.IsNotExist(err) {
			log.Println("File" + fileName + "does not exist!")
			http.NotFound(w, r)
			return
		}
		log.Panic(err)
	}

	if matches == nil {
		log.Println("File" + fileName + "does not exist!")
		http.NotFound(w, r)
		return
	}

	file, _ := os.Open(matches[0])
	http.ServeContent(w, r, matches[0], time.Now(), file)
}

func ServeManifestJson(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("../web-ui/www/manifest.template.json")

	if err != nil {
		if os.IsNotExist(err) {
			log.Println("manifest.template.json does not exist!")
			http.NotFound(w, r)
			return
		}
		log.Panic(err)
	}

	fileContents := string(data)

	tmpl, err := template.New("manifest.json").Parse(fileContents)
	if err != nil {
		log.Panic(err)
	}

	pageData := &model.TemplateData{
		AppTitle:        "Micro Frontends WebUI",
		AppTitleShort:   "µFE",
		BaseURL:         "/microfrontends-ui/",
		AppIconLarge:    "",
		AppIconSmall:    "",
		TouchIcon:       "",
		BackgroundColor: "#ffffff",
		ThemeColor:      "#ffffff",
	}

	err = tmpl.Execute(w, pageData)
	if err != nil {
		log.Panic(err)
	}
}

func PassThrough(w http.ResponseWriter, r *http.Request) {
	log.Println("Request to get web component started.")
	path := r.URL.Path

	req, err := http.NewRequest("GET", "http://localhost:80"+path, r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Proxying request to the module.", "Resolved URL:", "http://localhost:80"+path)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	copyHeaders(w.Header(), resp.Header)
	w.Header().Set("Content-Type", "application/javascript")

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func copyHeaders(dst, src http.Header) {
	for key, values := range src {
		dst.Set(key, strings.Join(values, ", "))
	}
}
