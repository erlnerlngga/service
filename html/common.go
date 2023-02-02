package html

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"

	"github.com/maragudk/service/model"
)

type PageProps struct {
	Title       string
	Description string
	User        *model.User
}

var hashOnce sync.Once
var appCSSPath, appJSPath string

func page(p PageProps, body ...g.Node) g.Node {
	hashOnce.Do(func() {
		appCSSPath = getHashedPath("public/styles/app.css")
		appJSPath = getHashedPath("public/scripts/app.js")
	})

	return c.HTML5(c.HTML5Props{
		Title:       p.Title,
		Description: p.Description,
		Language:    "en",
		Head: []g.Node{
			Link(Rel("stylesheet"), Href(appCSSPath)),
			Script(Src(appJSPath), Defer()),
			favIcons(),
			openGraph(p.Title, p.Description, "/images/logo.png", ""),
		},
		Body: []g.Node{Class("h-full bg-gray-50"),
			navbar(p),
			container(true,
				g.Group(body),
			),
		},
	})
}

func navbar(p PageProps) g.Node {
	return Div(
		container(false,
			Div(Class("flex justify-between py-2"),
				A(Href("/"), g.Text(`Home`)),
				g.If(p.User == nil,
					A(Href("/signup"), g.Text(`Sign up`)),
				),
			),
		),
	)
}

func container(padY bool, children ...g.Node) g.Node {
	return Div(
		c.Classes{
			"max-w-7xl mx-auto px-4 sm:px-6 lg:px-8": true,
			"py-4 sm:py-6 lg:py-8":                   padY,
		},
		g.Group(children),
	)
}

func prose(children ...g.Node) g.Node {
	return Div(Class("prose prose-lg lg:prose-xl xl:prose-2xl"), g.Group(children))
}

func card(children ...g.Node) g.Node {
	return Div(Class("bg-white py-8 px-4 shadow rounded-lg sm:px-10"), g.Group(children))
}

func h1(children ...g.Node) g.Node {
	return H1(Class("font-medium text-gray-900 text-xl"), g.Group(children))
}

func a(children ...g.Node) g.Node {
	return A(Class("font-medium text-cyan-600 hover:text-cyan-500"), g.Group(children))
}

func label(id, text string) g.Node {
	return Label(For(id), Class("block text-sm text-gray-700 mb-1"), g.Text(text))
}

func input(children ...g.Node) g.Node {
	return Input(Class("block w-full rounded-md border border-gray-300 focus:border-cyan-500 px-3 py-2 placeholder-gray-400 shadow-sm sm:text-sm text-gray-900 focus:ring-cyan-500"), g.Group(children))
}

func button(children ...g.Node) g.Node {
	return Button(Class("block w-full rounded-md bg-cyan-600 hover:bg-cyan-700 px-4 py-2 font-medium text-white focus:outline-none focus:ring-2 focus:ring-cyan-500 focus:ring-offset-2"), g.Group(children))
}

const themeColor = "#ffffff"

func favIcons() g.Node {
	return g.Group([]g.Node{
		Link(Rel("apple-touch-icon"), g.Attr("sizes", "180x180"), Href("/apple-touch-icon.png")),
		Link(Rel("icon"), Type("image/png"), g.Attr("sizes", "32x32"), Href("/favicon-32x32.png")),
		Link(Rel("icon"), Type("image/png"), g.Attr("sizes", "16x16"), Href("/favicon-16x16.png")),
		Link(Rel("manifest"), Href("/manifest.json")),
		Link(Rel("mask-icon"), Href("/safari-pinned-tab.svg"), g.Attr("color", themeColor)),
		Meta(Name("msapplication-TileColor"), Content(themeColor)),
		Meta(Name("theme-color"), Content(themeColor)),
	})
}

func openGraph(title, description, image, url string) g.Node {
	return g.Group([]g.Node{
		Meta(g.Attr("property", "og:type"), Content("website")),
		Meta(g.Attr("property", "og:title"), Content(title)),
		g.If(description != "", Meta(g.Attr("property", "og:description"), Content(description))),
		g.If(image != "", Meta(g.Attr("property", "og:image"), Content(image))),
		g.If(url != "", Meta(g.Attr("property", "og:url"), Content(url))),
	})
}

func getHashedPath(path string) string {
	externalPath := strings.TrimPrefix(path, "public")
	ext := filepath.Ext(path)
	if ext == "" {
		panic("no extension found")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Sprintf("%v.x%v", strings.TrimSuffix(externalPath, ext), ext)
	}

	return fmt.Sprintf("%v.%x%v", strings.TrimSuffix(externalPath, ext), sha256.Sum256(data), ext)
}
