class SectionNav extends Component {
	renderThis() {
		if(!this.components){
			this.components = {
				buttonThreads: new ButtonNav("Hjem", data.sectionThreads),
				buttonImages: new ButtonNav("Billeder", data.sectionImages),
				buttonCode: new ButtonNav("Sjov", data.sectionCode)
			}
		}
		return l("div",
			l("ul",
				this.components.buttonThreads,
				this.components.buttonImages,
				this.components.buttonCode),
			l("a", {}, "Log ud")
		)
	}

	renderHasChanged() {
		return false
	}

	static styleThis() {
		return {
			"div": {
				padding: "0.5rem",
				display: "flex",
				fontWeight: "bold",
				fontSize: "1.2rem",
			},
			"ul": {
				display: "flex",
				flexGrow: 1
			},
			"a": {
				flexGrow: 0,
				cursor: "pointer"
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
				cursor: "pointer",
				transition: "0.3s"
			},
			"li:hover": {
				background: "#4b9a8c",
				transition: "0.3s"
			},
			".selected": {
				background: "#61af4e"
			}
		}
	}

	selectPage(page) {
		state.currentPage = page
		site.update()
	}
}

