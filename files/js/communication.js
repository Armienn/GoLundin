"use strict"

function getRequest(url, callback) {
	var xmlHttp = new XMLHttpRequest();
	xmlHttp.onreadystatechange = function () {
		if (xmlHttp.readyState == 4) {
			if (xmlHttp.status == 200)
				callback(xmlHttp.responseText)
			else {
				// error(ish?)
			}
		}
	}
	xmlHttp.onerror = function () {
		// error
	}
	xmlHttp.open("GET", url, true)
	xmlHttp.send()
}

function postRequest(url, parameters, callback) {
	var xmlHttp = new XMLHttpRequest();
	xmlHttp.onreadystatechange = function () {
		if (xmlHttp.readyState == 4) {
			if (xmlHttp.status == 200)
				callback(xmlHttp.responseText)
			else {
				// error(ish?)
			}
		}
	}
	xmlHttp.onerror = function () {
		// error
	}
	xmlHttp.open("POST", url, true)
	xmlHttp.send(JSON.stringify(parameters))
}

class Core {
	constructor() {
		this.sections = ["Generelt", "Diverse"]
		this.currentSection = "Generelt"
		this.threads = []
	}

	currentThreads() {
		return this.threads.filter(e => e.section == this.currentSection)
	}
}
var core = new Core()

class Thread {
	constructor(source) {
		this.id = 0
		this.title = ""
		this.mainMessage = ""
		this.author = ""
		this.section = ""
		this.time = new Date()
		this.responses = []
		for (var i in source) this[i] = source[i]
	}
}

core.threads = [
	new Thread({
		id: 1,
		title: "Test besked",
		mainMessage: "Her er en lille test besked. Blablabla.\nAzg nazg thrakatul.",
		author: "Kristjan",
		section: "Diverse",
		time: "blub",
		responses: []
	}),
	new Thread({
		id: 2,
		title: "Test besked 2",
		mainMessage: "Her er endnu en lille test besked. Blablabla.\nAzg nazg thrakatul.",
		author: "Kristjan",
		section: "Diverse",
		time: "blub",
		responses: []
	}),
	new Thread({
		id: 3,
		title: "Test besked 3",
		mainMessage: "Her er endnu en lille test besked. Blablabla.\nAzg nazg thrakatul.",
		author: "Kristjan",
		section: "Generelt",
		time: "blub",
		responses: []
	})
]
