"use strict"

window.onload = function () {

	function render() {
		return l("div", {},
			renderNav(),
			renderPage()
		)
	}

	var state = 0

	site.setRenderFunction(render)
	site.update()

	setInterval(function () {
		state++;
		site.update()
	}, 1000);
}

function renderNav() {
	return l("nav",
		l("h1", "Familien Lundin"),
		l("ul",
			l("li", {}, "Hjem"),
			l("li", {}, "Billeder"),
			l("li", {}, "Sjov")),
		l("a", {}, "Log ud")
	)
}

function renderPage() {
	return "Hej du"
}