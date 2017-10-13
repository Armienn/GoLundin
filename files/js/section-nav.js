class SectionNav extends Component {
	renderThis() {
		if(!this.components){
			this.components = {
				buttonThreads: new ButtonNav("Hjem", data.sectionThreads),
				buttonImages: new ButtonNav("Billeder", data.sectionImages),
				buttonCode: new ButtonNav("Sjov", data.sectionCode)
			}
		}
		return l("nav",
			l("h1", "Familien Lundin"),
			l("ul",
				this.components.buttonThreads,
				this.components.buttonImages,
				this.components.buttonCode),
			l("a", {}, l("div", "Log ud"))
		)
	}

	renderHasChanged() {
		return false
	}

	static styleThis() {
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
	constructor(text, destination) {
		super()
		this.text = text
		this.destination = destination
	}

	renderThis() {
		return l(state.currentPage == this.destination ? "li.selected" : "li", {
			onclick: () => { this.selectPage(this.destination) }
		}, this.text)
	}

	static styleThis() {
		return {
			"li": {
				height: "2rem",
				lineHeight: "2rem",
				textAlign: "center",
				width: "10rem",
				background: "green"
			},
			".selected": {
				background: "yellow"
			}
		}
	}

	selectPage(page) {
		state.currentPage = page
		site.update()
	}
}

