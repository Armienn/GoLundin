class SectionNav {
	render() {
		return l("nav",
			l("h1", "Familien Lundin"),
			l("ul",
				l("li", { onclick: () => { this.selectPage(data.sectionThreads) } }, "Hjem"),
				l("li", { onclick: () => { this.selectPage(data.sectionImages) } }, "Billeder"),
				l("li", { onclick: () => { this.selectPage(data.sectionCode) } }, "Sjov")),
			l("a", {}, "Log ud")
		)
	}

	selectPage(page) {
		state.currentPage = page
		site.update()
	}
}

