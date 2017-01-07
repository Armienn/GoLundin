var body = document.getElementById('jsbox')
body.innerHTML = ""
var canvas = document.createElement("canvas")
body.appendChild(canvas)
canvas.width = 800
canvas.height = 660
var ctx = canvas.getContext('2d')

function clear() {
    ctx.clearRect(0, 0, canvas.width, canvas.height)
}

function draw() {
    clear()
    var second = new Date().getMilliseconds()/5
    ctx.fillStyle = "rgba(0,100,100,0.5)"
    ctx.fillRect (second, 20, 50, 50)
}

if (updateInterval) {
    clearInterval(updateInterval)
}
var updateInterval = setInterval(draw, 50)