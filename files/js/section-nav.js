class SectionNav extends Component {
	renderThis() {
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

	styleThis() {
		return {
			"nav": {
				padding: "1rem",
				display: "flex",
				fontWeight: "bold",
				fontSize: "1.2rem",
			},

			"h1": {
				fontSize: "2rem",
				flexGrow: 0
			},

			"ul": {
				display: "flex",
				flexGrow: 1
			},

			"a": {
				flexGrow: 0
			},

			"li": {
				height: "2rem",
				lineHeight: "2rem",
				textAlign: "center",
				width: "10rem",
				background: "green"
			}
		}
	}
}

