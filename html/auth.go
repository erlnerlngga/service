package html

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func SignupPage(p PageProps) g.Node {
	p.Title = "Sign up"

	return page(p,
		Div(Class("sm:mx-auto sm:w-full sm:max-w-md"),
			card(
				FormEl(Action("/signup"), Method("post"), Class("space-y-6"),
					Div(Class("text-center"),
						h1(g.Text(`Sign up`)),
						a(Href("/login"), g.Text("or log in instead")),
					),

					Div(
						label("name", "Name"),
						input(Type("text"), ID("name"), Name("name"), AutoComplete("name"), Placeholder("Me"), Required(), AutoFocus()),
					),

					Div(
						label("email", "Email"),
						input(Type("email"), ID("email"), Name("email"), AutoComplete("email"), Placeholder("me@example.com"), Required()),
					),

					Div(Class("flex items-center space-x-2"),
						Input(ID("accept"), Name("accept"), Type("checkbox"), Required(), Class("h-4 w-4 rounded border-gray-300 text-cyan-600 focus:ring-cyan-500")),
						Label(For("accept"), Class("text-gray-900"),
							g.Text(`I accept the `),
							a(Href("/legal/terms-of-service"), Target("_blank"), g.Text(`Terms of Service`)),
							g.Text(` and `),
							a(Href("/legal/privacy-policy"), Target("_blank"), g.Text(`Privacy Policy`)),
							g.Text(`.`),
						),
					),

					button(Type("submit"), g.Text(`Sign up`)),
				),
			),
		),
	)
}

func LoginPage(p PageProps) g.Node {
	p.Title = "Log in"

	return page(p,
		Div(Class("sm:mx-auto sm:w-full sm:max-w-md"),
			card(
				FormEl(Action("/login"), Method("post"), Class("space-y-6"),
					Div(Class("text-center"),
						h1(g.Text(`Log in`)),
						a(Href("/signup"), g.Text("or sign up instead")),
					),

					Div(
						label("email", "Email"),
						input(Type("email"), ID("email"), Name("email"), AutoComplete("email"), Placeholder("me@example.com"), Required(), AutoFocus()),
					),

					button(Type("submit"), g.Text(`Log in`)),
				),
			),
		),
	)
}
