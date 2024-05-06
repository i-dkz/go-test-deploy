package handler

import (
	"fmt"
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

var templates = template.Must(template.ParseFiles(
	"home.html",
	"projects.html",
	"blog.html",
	"main.html",
	"header.html",
	"footer.html",
	"modal.html",
	"main-htmx.html",
	"tag.html",
))

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

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Success</h1>")
}

func Main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", Handler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
