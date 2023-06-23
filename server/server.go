package server

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/SevcikMichal/microfrontends-webui/configuration"
	"github.com/SevcikMichal/microfrontends-webui/model"
	"golang.org/x/net/html"
)

func ServeSinglePageApplication(w http.ResponseWriter, r *http.Request) {
	language, _ := requestMatchLanguage(r, configuration.GetAcceptsLanguages())

	htmlToFind := "index." + language + ".html"
	matches := getAllPossibleFiles(htmlToFind)
	bestFit := getFirstMatchingFile("web-ui/www", matches)

	data, err := os.ReadFile(bestFit)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("index.html does not exist!")
			http.NotFound(w, r)
			return
		}
		log.Panic(err)
	}

	fileContents := string(data)

	node, err := html.Parse(strings.NewReader(fileContents))
	if err != nil {
		log.Panic(err)
	}

	nonce, _ := generateNonce()

	var findAllScriptsAddNonce func(*html.Node)
	findAllScriptsAddNonce = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "script" {
			// Add nonce attribute to script tag
			attr := html.Attribute{Key: "nonce", Val: nonce}
			n.Attr = append(n.Attr, attr)
		}

		// Recursively process child nodes
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findAllScriptsAddNonce(c)
		}
	}

	findAllScriptsAddNonce(node)

	var buf bytes.Buffer
	if err := html.Render(&buf, node); err != nil {
		log.Panic(err)
	}
	fileContents = buf.String()

	tmpl, err := template.New("index.html").Parse(fileContents)
	if err != nil {
		log.Panic(err)
	}

	pageData := &model.TemplateData{
		Language:                   language,
		AppTitle:                   configuration.GetApplicationTitle(language),
		BaseURL:                    configuration.GetBaseURL(),
		Description:                configuration.GetApplicationDescription(language),
		MicroFrontendShellContext:  configuration.GetApplicationShellContext(),
		MicroFrontendSelector:      strings.Join(configuration.GetWebcomponentsSelector(), ","),
		ProgresiveWebAppMode:       configuration.GetPwaMode(),
		ContentSecurityPolicyNonce: nonce,
		TouchIcon:                  configuration.GetTouchIcon(),
		FavIcon:                    configuration.GetFaviconIco(),
	}

	cspHeader := configuration.GetHttpCspHeader()
	cspHeader = strings.ReplaceAll(cspHeader, "{NONCE_VALUE}", nonce)

	w.Header().Set("Content-Security-Policy", cspHeader)

	err = tmpl.Execute(w, pageData)
	if err != nil {
		log.Panic(err)
	}
}

func ServeFile(w http.ResponseWriter, r *http.Request) {
	fileName := strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1]

	if fileName == "sw.mjs" {
		fileName = configuration.GetServiceWorker()
	}

	matches := getAllPossibleFiles(fileName)
	bestFit := getFirstMatchingFile("web-ui/www", matches)

	file, err := os.Open(bestFit)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("index.html does not exist!")
			http.NotFound(w, r)
			return
		}
		log.Panic(err)
	}

	http.ServeContent(w, r, matches[0], time.Now(), file)
}

func ServeManifestJson(w http.ResponseWriter, r *http.Request) {
	language, _ := requestMatchLanguage(r, configuration.GetAcceptsLanguages())
	manifest := configuration.GetManifestTemplate()
	parts := strings.Split(manifest, ".")
	ext := "." + parts[len(parts)-1]
	base := strings.Join(parts[:len(parts)-1], ".")

	manifestToFind := base + "." + language + ext

	matches := getAllPossibleFiles(manifestToFind)
	bestFit := getFirstMatchingFile("web-ui/www", matches)

	data, err := os.ReadFile(bestFit)

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
		AppTitle:        configuration.GetApplicationTitle(language),
		AppTitleShort:   configuration.GetApplicationTitleShort(language),
		BaseURL:         configuration.GetBaseURL(),
		AppIconLarge:    configuration.GetAppIconLarge(),
		AppIconSmall:    configuration.GetAppIconSmall(),
		TouchIcon:       configuration.GetTouchIcon(),
		BackgroundColor: configuration.GetManifestBackgroundColor(),
		ThemeColor:      configuration.GetManifestBackgroundColor(),
	}

	err = tmpl.Execute(w, pageData)
	if err != nil {
		log.Panic(err)
	}
}

func generateNonce() (string, error) {
	codes := make([]byte, 128)
	_, err := rand.Read(codes)
	if err != nil {
		return "", err
	}

	text := string(codes)
	nonce := base64.StdEncoding.EncodeToString([]byte(text))

	return nonce, nil
}
