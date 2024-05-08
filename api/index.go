package handler

// package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

type StackData struct {
	Path   template.HTMLAttr
	Width  int
	Height int
	Gap    int
	Title  string
}

type TagData struct {
	Title string
}

var techStack = map[string][]StackData{
	"tech": {
		{
			Path:   template.HTMLAttr("src/public/go.svg"),
			Width:  40,
			Height: 40,
			Gap:    2,
			Title:  "Go",
		},
		{
			Path:   template.HTMLAttr("src/public/java.svg"),
			Width:  30,
			Height: 30,
			Gap:    2,
			Title:  "Java",
		},
		{
			Path:   template.HTMLAttr("src/public/python.svg"),
			Width:  10,
			Height: 10,
			Gap:    2,
			Title:  "Python",
		},
		{
			Path:   template.HTMLAttr("src/public/django.svg"),
			Width:  30,
			Height: 30,
			Gap:    2,
			Title:  "Django",
		},
		{
			Path:   template.HTMLAttr("src/public/postgres.svg"),
			Width:  30,
			Height: 30,
			Gap:    2,
			Title:  "PostgreSQL",
		},
		{
			Path:   template.HTMLAttr("src/public/node.svg"),
			Width:  30,
			Height: 30,
			Gap:    0,
			Title:  "Node.js",
		},
		{
			Path:   template.HTMLAttr("src/public/next.svg"),
			Width:  30,
			Height: 30,
			Gap:    2,
			Title:  "Next.js",
		},
		{
			Path: template.HTMLAttr("src/public/supabase.svg"),

			Gap:   1,
			Title: "Supabase",
		},
		{
			Path:   template.HTMLAttr("src/public/swift.svg"),
			Width:  30,
			Height: 30,
			Gap:    2,
			Title:  "Swift",
		},
		{
			Path:   template.HTMLAttr("src/public/git.svg"),
			Width:  30,
			Height: 30,
			Gap:    2,
			Title:  "Git",
		},
		{
			Path:   template.HTMLAttr("src/public/react.svg"),
			Width:  30,
			Height: 30,
			Gap:    2,
			Title:  "React.js",
		},
		{
			Path:   template.HTMLAttr("src/public/tailwind.svg"),
			Width:  30,
			Height: 30,
			Gap:    0,
			Title:  "TailwindCSS",
		},
		{
			Path:   template.HTMLAttr("src/public/html.svg"),
			Width:  30,
			Height: 30,
			Title:  "HTML5",
		},
		{
			Path:   template.HTMLAttr("src/public/htmx.svg"),
			Width:  40,
			Height: 40,
			Title:  "HTMX",
		},
		{
			Path:   template.HTMLAttr("src/public/css.svg"),
			Width:  40,
			Height: 40,
			Title:  "CSS3",
		},
		{
			Path:   template.HTMLAttr("src/public/javascript.svg"),
			Width:  40,
			Height: 40,
			Title:  "Javascript",
		},
		{
			Path:   template.HTMLAttr("src/public/typescript.svg"),
			Width:  40,
			Height: 40,
			Title:  "Typescript",
		},
		{
			Path:   template.HTMLAttr("src/public/astro.svg"),
			Width:  40,
			Height: 40,
			Title:  "Astro.js",
		},
		{
			Path:   template.HTMLAttr("src/public/bootstrap.svg"),
			Width:  40,
			Height: 40,
			Title:  "Bootstrap",
		},
		{
			Path:   template.HTMLAttr("src/public/chakra.svg"),
			Width:  40,
			Height: 40,
			Title:  "ChakraUI",
		},
	},
}

var tags = map[string][]TagData{
	"bbyHtmx": {
		{Title: "Astro"},
		{Title: "HTMX"},
		{Title: "Alpine.js"},
		{Title: "Javascript"},
		{Title: "Tailwind"},
		{Title: "OpenAI"},
		{Title: "Supabase"},
	},
	"nxtStore": {
		{Title: "Next.js"},
		{Title: "React.js"},
		{Title: "Zustand"},
		{Title: "Tailwind"},
		{Title: "Axios"},
		{Title: "Shadcn"},
	},
	"gamePlus": {
		{Title: "React.js"},
		{Title: "ChakraUI"},
		{Title: "Vite"},
		{Title: "Axios"},
		{Title: "Zustand"},
	},
}

//go:embed all:src
var staticFiles embed.FS
var templates = template.Must(template.ParseFS(staticFiles, "src/templates/*.html", "src/templates/components/*.html"))

func Main() {
	router := http.NewServeMux()
	fs := http.FileServer(http.Dir("src"))
	router.Handle("GET /src/", http.StripPrefix("/src/", fs))

	router.HandleFunc("GET /", Handler)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path == "/output.css" {

	// 	w.Header().Set("Content-Type", "text/css")
	// 	outputCSS, _ := staticFiles.ReadFile("src/output.css")
	// 	w.Write(outputCSS)
	// 	return
	// }

	// if strings.HasPrefix(r.URL.Path, "/src/public/") && (strings.HasSuffix(r.URL.Path, ".svg") || strings.HasSuffix(r.URL.Path, ".png")) {
	// 	contentType := "image/svg+xml"
	// 	if strings.HasSuffix(r.URL.Path, ".png") {
	// 		contentType = "image/png"
	// 	}
	// 	w.Header().Set("Content-Type", contentType)
	// 	imageData, err := staticFiles.ReadFile(strings.TrimPrefix(r.URL.Path, "/"))
	// 	if err != nil {
	// 		http.NotFound(w, r)
	// 		return
	// 	}
	// 	w.Write(imageData)
	// 	return
	// }
	// switch r.URL.Path {
	// case "/":
	// 	// fmt.Fprintf(w, "<h1>Success</h1>")
	// 	err := templates.ExecuteTemplate(w, "index.html", techStack)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// case "/projects":
	// 	if r.Header.Get("HX-Request") == "true" {
	// 		err := templates.ExecuteTemplate(w, "projects.html", tags)
	// 		if err != nil {
	// 			http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		}
	// 	} else {
	// 		http.Redirect(w, r, "/", http.StatusFound)
	// 	}
	// case "/blog":
	// 	if r.Header.Get("HX-Request") == "true" {
	// 		err := templates.ExecuteTemplate(w, "blog.html", tags)
	// 		if err != nil {
	// 			http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		}
	// 	} else {
	// 		http.Redirect(w, r, "/", http.StatusFound)
	// 	}
	// case "/main":
	// 	if r.Header.Get("HX-Request") == "true" {
	// 		err := templates.ExecuteTemplate(w, "main-htmx.html", techStack)
	// 		if err != nil {
	// 			http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		}
	// 	} else {
	// 		http.Redirect(w, r, "/", http.StatusFound)
	// 	}
	// default:
	// 	http.NotFound(w, r)
	// }
	templates.ExecuteTemplate(w, "index.html", techStack)
}
