class SectionThreads extends Component {
	renderThis() {
		if (!this.components) {
			this.components = {
				overview: new ThreadOverview(),
				new: new ThreadNew()
			}
		}
		return l("div",
			this.renderNav(),
			l("div.main", this.components.overview)
		)
	}

	renderNav() {
		return l("nav",
			l("h1", "Sektioner:"),
			...this.sections()
		)
	}

	static styleThis() {
		return {
			"div.main": {
				height: "100%",
				color: "black",
				marginLeft: "10rem"
			},
			"nav": {
				position: "fixed",
				width: "10rem",
				height: "100%",
				backgroundColor: "#88bbcc"
			},
			"nav h1": {
				width: "10rem",
				height: "1.5rem",
				lineHeight: "1.5rem",
				fontSize: "1rem",
				fontWeight: "bold",
				textAlign: "center",
				color: "#3f5f7f"
			},
			"nav div": {
				width: "10rem",
				height: "1.5rem",
				lineHeight: "1.5rem",
				fontWeight: "bold",
				textAlign: "center",
				cursor: "pointer",
				backgroundColor: "#88bbcc",
				transition: "0.3s"
			},
			"nav div:hover": {
				backgroundColor: "#679aab",
				transition: "0.3s"
			}
		}
	}

	sections() {
		return core.sections.map(e => {
			return l("div", {
				onclick: () => {
					core.currentSection = e
					site.update()
				}
			}, e)
		})
	}
}

class ThreadOverview extends Component {
	renderThis() {
		return l("section",
			...this.threads()
		)
	}

	threads() {
		return core.currentThreads().map(e => new ThreadComponent(e))
	}
}
class ThreadComponent extends Component {
	constructor(thread) {
		super()
		this.thread = thread
	}

	renderThis() {
		return l("section",
			l("div.top",
				l("div.left",
					l("h1", this.thread.title),
					l("pre", this.thread.mainMessage)),
				l("div.right",
					l("div", this.thread.time),
					l("div", this.thread.author)
				)
			),
			l("div.expand",
				l("div")
			)
		)
	}

	static styleThis() {
		return {
			".top": {
				height: "3rem",
				overflow: "hidden"
			},
			".left": {
				padding: "0.5rem",
				float: "left"
			},
			".left h1": {
				fontWeight: "bold",
				fontSize: "1rem"
			},
			".right": {
				padding: "0.5rem",
				float: "right",
				textAlign: "right"
			}
		}
	}
}

class ThreadNew extends Component {
	renderThis() {
		return l("section",
			l("div", "blub"),
			l("div", "asdf")
		)
	}
}
