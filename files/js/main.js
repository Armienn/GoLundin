"use strict"

var data = {
	sectionNav: new SectionNav(),
	sectionThreads: new SectionThreads(),
	sectionImages: new SectionImages(),
	sectionCode: new SectionCode()
}

var state = {
	currentPage: data.sectionThreads
}

class FamilienLundinSite extends Component {
	renderThis() {
		return l("div", {},
			l("header", "Familien Lundin"),
			l("nav", data.sectionNav),
			l("section", state.currentPage)
		)
	}

	static styleThis(){
		return {
			header: {
				fontSize: "2rem",
				fontWeight: "bold",
				height: "3rem",
				lineHeight: "3rem",
				paddingLeft: "1rem"
			},
			nav: {
				height: "3rem"
			},
			section: {
				minHeight: "calc(100vh - 6rem)"
			},
			"nav, header":{
				backgroundColor: "#3f5f7f"
			}
		}
	}
}

window.onload = function () {
	var fmsite = new FamilienLundinSite()

	site.setRenderFunction(() => fmsite.render())
	site.update()
}
