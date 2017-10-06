"use strict"

function l(tag, options, ...children) {
	if (typeof options == "string" ||
		options instanceof virtualDom.VNode ||
		options instanceof virtualDom.VText
	) {
		children.unshift(options)
		options = {}
	}
	for (var i in children)
		if (children[i] instanceof Array)
			children.splice(i, 1, ...children[i])
	return virtualDom.h(tag, options, children)
}

var site = (() => {
	var vtree
	var rootNode
	var render
	function update() {
		if (!vtree) {
			vtree = render()
			rootNode = virtualDom.create(vtree)
			document.body.innerHTML = ""
			document.body.appendChild(rootNode)
			return
		}
		var newTree = render();
		var patches = virtualDom.diff(vtree, newTree)
		rootNode = virtualDom.patch(rootNode, patches)
		vtree = newTree
	}

	function elementFunction(tag) {
		return (options, ...children) => {
			return l(tag, options, ...children)
		}
	}

	return {
		update: () => { update() },
		setRenderFunction: (r) => { render = r },
		elementFunction: elementFunction
	}
})()

/*
class Section {
	set render(value) {
		this.renderFunction = value
	}

	get render() {
		return this.renderSection
	}

	renderSection() {
		if (!this.isChanged)
			throw new Error("isChanged function needs to be set in a Section")
		if (!this.renderFunction)
			throw new Error("render function needs to be set in a Section")
		var changes = this.isChanged()
		if (changes)
			this.tree = this.renderFunction()
		return this.tree
	}
}
*/
