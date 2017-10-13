"use strict"


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

	function elementFunction(tag) {
		return (options, ...children) => {
			return l(tag, options, ...children)
		}
	}

	function renderNewTree(component) {
		component.tree = component.renderThis()
		component.components = []
		markTree(component, component.tree)
	}

	function markTree(component, tree) {
		traverseTree(tree, (node, path) => {
			if (node.component) {
				if (!path.length)
					throw new Error("Component must be within some other element")
				component.components.push({ component: node.component, path: path })
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
			node.properties.attributes[component.designation] = ""
		})
	}

	function renderComponents(component) {
		for (var i in component.components) {
			var node = component.tree
			var n = 0
			for (; n < component.components[i].path.length - 1; n++)
				node = node.children[component.components[i].path[n]]
			node.children[component.components[i].path[n]] = component.components[i].component.render()
		}
	}

	function traverseTree(tree, callback, path) {
		if (!(tree instanceof virtualDom.VNode))
			return
		if (!path)
			path = []
		var stop = callback(tree, path)
		if (stop)
			return
		for (var i in tree.children)
			traverseTree(tree.children[i], callback, path.concat([i]))
	}

	function modifySelector(selector, designation) {
		selector = selector.split(" ").join("[" + designation + "] ") + "[" + designation + "] "
		selector = selector.split(",").join("[" + designation + "],")
		return selector
	}

	function insertRule(sheet, key, rule) {
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

	var stylesheets = {}

	class Component {
		constructor() {
			this.designation = getDesignation(this)
		}

		render() {
			if (!this.renderThis)
				throw new Error("Component is missing renderThis()")
			if (!stylesheets[this.designation] || this.constructor.styleHasChanged())
				this.style()
			if (!this.tree || this.renderHasChanged())
				renderNewTree(this)
			renderComponents(this)
			return this.tree
		}

		renderHasChanged() {
			return true
		}

		style() {
			if (!this.constructor.styleThis)
				return
			var styles = this.constructor.styleThis()
			if (!stylesheets[this.designation]) {
				stylesheets[this.designation] = document.createElement("style")
				document.head.appendChild(stylesheets[this.designation])
			}
			else {
				while (stylesheets[this.designation].sheet.cssRules.length)
					stylesheets[this.designation].sheet.removeRule(0)
			}
			for (var i in styles)
				insertRule(stylesheets[this.designation].sheet, modifySelector(i, this.designation), styles[i])
		}

		static styleHasChanged() {
			return false
		}
	}

	function getDesignation(component) {
		return "site-" + component.constructor.name
	}

	return {
		update: () => { update() },
		setRenderFunction: (r) => { render = r },
		l: l,
		elementFunction: elementFunction,
		Component: Component
	}
})()

var l = site.l
var Component = site.Component

