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
	constructor() {
		super()
		this.threadComponents = []
		this.currentThreads = []
	}
	renderThis() {
		return l("section",
			...this.threads()
		)
	}

	threads() {
		if (this.currentThreads !== core.currentThreads || this.threadComponents.length != core.currentThreads.length)
			this.threadComponents = core.currentThreads.map(e => new ThreadComponent(e))
		this.currentThreads = core.currentThreads
		return this.threadComponents
	}
}

class ThreadComponent extends Component {
	constructor(thread, open) {
		super()
		this.open = open
		this.thread = thread
		this.threadComponents = []
	}

	renderThis() {
		return l("section",
			l("div.top" + (this.open ? ".open" : ""), {
				onclick: () => {
					this.open = !this.open
					site.update()
				}
			},
				l("div.left",
					l("h1", this.thread.title),
					l("pre" + (this.open ? ".open" : ""), this.thread.mainMessage),
					this.editLine()
				),
				l("div.right",
					l("div", this.thread.time),
					l("div", this.thread.author)
				)
			),
			l("div.expand" + (this.open ? ".open" : ""),
				...this.threads()
			)
		)
	}

	editLine() {
		return l("div",
			l("span", {}, "reply")
		)
	}

	static styleThis() {
		return {
			".top:hover": {
				backgroundColor: "#e0e0e0",
				transition: "0.3s"
			},
			".top": {
				cursor: "pointer",
				height: "3rem",
				overflow: "hidden",
				transition: "0.3s"
			},
			".top.open": {
				height: "initial",
			},
			".left": {
				padding: "0.5rem",
				float: "left"
			},
			".left h1": {
				fontWeight: "bold",
				fontSize: "1rem"
			},
			"pre": {
				padding: "0.5rem",
				color: "darkgray"
			},
			"pre.open": {
				color: "black"
			},
			".right": {
				padding: "0.5rem",
				float: "right",
				textAlign: "right"
			},
			".expand": {
				padding: "0 0 0 1rem",
				display: "none"
			},
			".expand.open": {
				display: "block"
			},
			span: {
				fontSize: "0.7rem"
			}
		}
	}

	threads() {
		if (this.currentThreads !== this.thread.responses || this.threadComponents.length != this.thread.responses.length)
			this.threadComponents = this.thread.responses.map(e => new ThreadComponent(e, true))
		this.currentThreads = this.thread.responses
		return this.threadComponents
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
