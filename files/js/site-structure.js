"use strict"

function l(tag, options, ...children) {
	if (typeof options == "string" ||
		options instanceof virtualDom.VNode ||
		options instanceof virtualDom.VText ||
		options instanceof Component
	) {
		children.unshift(options)
		options = {}
	}
	for (var i in children)
		if (children[i] instanceof Array)
			children.splice(i, 1, ...children[i])
	for (var i in children)
		if (children[i] instanceof Component) {
			var placeholder = virtualDom.h("placeholder", {})
			placeholder.component = children[i]
			children[i] = placeholder
		}
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
		if (!this.styleElement || this.styleHasChanged())
			this.style()
		if (!this.tree || this.renderHasChanged())
			this.renderNewTree()
		this.renderComponents()
		return this.tree
	}

	renderHasChanged() {
		return true
	}

	renderNewTree() {
		this.tree = this.renderThis()
		for (var i in this.components)
			this.components[i].component.deleteStylesheet()
		this.components = []
		this.markTree(this.tree)
	}

	markTree(tree) {
		this.traverseTree(tree, (node, path) => {
			if (node.component) {
				if (!path.length)
					throw new Error("Component must be within some other element")
				this.components.push({ component: node.component, path: path })
				return true
			}

			if (node.properties.attributes) {
				for (var i in node.properties.attributes)
					if (i.startsWith("site-"))
						return true
			}
			else {
				node.properties.attributes = {}
			}
			node.properties.attributes[this.designation] = ""
		})
	}

	traverseTree(tree, callback, path) {
		if (!(tree instanceof virtualDom.VNode))
			return
		if (!path)
			path = []
		var stop = callback(tree, path)
		if (stop)
			return
		for (var i in tree.children)
			this.traverseTree(tree.children[i], callback, path.concat([i]))
	}

	renderComponents() {
		for (var i in this.components) {
			var node = this.tree
			var n = 0
			for (; n < this.components[i].path.length - 1; n++)
				node = node.children[this.components[i].path[n]]
			node.children[this.components[i].path[n]] = this.components[i].component.render()
		}
	}

	style() {
		if (!this.styleThis)
			return
		var styles = this.styleThis()
		if (!this.styleElement) {
			this.styleElement = document.createElement("style")
			document.head.appendChild(this.styleElement)
		}
		else {
			while (this.styleElement.sheet.cssRules.length)
				this.styleElement.sheet.removeRule(0)
		}
		for (var i in styles)
			this.insertRule(this.styleElement.sheet, this.modifySelector(i), styles[i])
	}

	styleHasChanged() {
		return false
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

	deleteStylesheet() {
		document.head.removeChild(this.styleElement)
		this.styleElement = undefined
	}
}

