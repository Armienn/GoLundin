var body = document.getElementById('jsbox')
body.innerHTML = ""
var canvas = document.createElement("canvas")
body.appendChild(canvas)
var ctx = canvas.getContext('2d')

function clear() {
    ctx.clearRect(0, 0, canvas.width, canvas.height)
}

function mulle() {
    clear()
    var second = new Date().getSeconds()
    ctx.fillStyle = "rgba(0,100,100,0.5)"
    ctx.fillRect (second, 20, 50, 50)
}

if (updateInterval) {
    clearInterval(updateInterval)
}
var updateInterval = setInterval(mulle, 1000)