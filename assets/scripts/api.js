const API = {}
API.print = function () {
    window.print()
}
API.__blackBody = false
API.blackBody = function () {
    API.__blackBody = !API.__blackBody
    document.body.className = API.__blackBody ? 'highlight' : ''
}
API.pages = function (n) {
    if (!API.__splitLines) {
        API.__splitLines = document.getElementById('split-lines')
    }
    if (n > 0) {
        API.__splitLines.className = ""
        API.__splitLines.innerHTML = new Array(n + 1).join('<hr />')
    } else {
        API.__splitLines.className = "hide"
    }
}