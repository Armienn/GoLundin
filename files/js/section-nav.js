class SectionNav extends Component {
	renderThis() {
		return l("nav",
			l("h1", "Familien Lundin"),
			l("ul",
				new ButtonNav("Hjem", data.sectionThreads),
				new ButtonNav("Billeder", data.sectionImages),
				new ButtonNav("Sjov", data.sectionCode)),
			l("a", {}, "Log ud")
		)
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
			}
		}
	}
}

class ButtonNav extends Component {
	constructor(text, destination){
		super()
		this.text = text
		this.destination = destination
	}

	renderThis() {
		return l("li", { onclick: () => { this.selectPage(this.destination) } }, this.text)
	}

	styleThis() {
		return {
			"li": {
				height: "2rem",
				lineHeight: "2rem",
				textAlign: "center",
				width: "10rem",
				background: "green"
			}
		}
	}

	selectPage(page) {
		state.currentPage = page
		site.update()
	}
}

