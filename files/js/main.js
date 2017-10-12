"use strict"

var state
var data

window.onload = function () {
	data = {
		sectionNav: new SectionNav(),
		sectionThreads: new SectionThreads(),
		sectionImages: new SectionImages(),
		sectionCode: new SectionCode()
	}

	state = {
		currentPage: data.sectionThreads
	}

	function render() {
		return l("div", {},
			data.sectionNav.render(),
			state.currentPage.render()
		)
	}

	site.setRenderFunction(render)
	site.update()
}