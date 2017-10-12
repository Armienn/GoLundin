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

	var designationCounter = 0

	return {
		update: () => { update() },
		setRenderFunction: (r) => { render = r },
		elementFunction: elementFunction,
		nextDesignation: () => { designationCounter++; return "site-" + designationCounter }
	}
})()

class Component {
	constructor() {
		this.designation = site.nextDesignation()
	}

	render() {
		if (!this.renderThis)
			throw new Error("Component is missing renderThis()")
		var tree = this.renderThis()
		this.markTree(tree)
		if (!this.element)
			this.style()
		return tree
	}

	markTree(tree) {
		if (!(tree instanceof virtualDom.VNode))
			return
		if (tree.properties.attributes) {
			for (var i in tree.properties.attributes)
				if (i.startsWith("site-"))
					return
		}
		else {
			tree.properties.attributes = {}
		}
		tree.properties.attributes[this.designation] = ""
		for (var i in tree.children)
			this.markTree(tree.children[i])
	}

	style() {
		if (!this.styleThis)
			return
		var styles = this.styleThis()
		if (!this.element) {
			this.element = document.createElement("style")
			document.head.appendChild(this.element)
		}
		else {
			while (this.element.sheet.cssRules.length)
				this.element.sheet.removeRule(0)
		}
		for (var i in styles)
			this.insertRule(this.element.sheet, this.modifySelector(i), styles[i])
	}

	insertRule(sheet, key, rule) {
		if (typeof rule == "string") {
			if (rule.trim()[0] != "{")
				rule = "{" + rule + "}"
			sheet.insertRule(key + rule)
			return
		}
		var index = sheet.insertRule(key + "{}")
		for (var i in rule)
			sheet.cssRules[index].style[i] = rule[i]
	}

	modifySelector(selector) {
		selector = selector.split(" ").join("[" + this.designation + "] ") + "[" + this.designation + "] "
		selector = selector.split(",").join("[" + this.designation + "],")
		return selector
	}
}

