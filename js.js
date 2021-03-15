canv = querySelector("#canvas")
var ctx = canv.getContext('2d')

var boardImage = new Image()
boardImage.src = 'board.png'

window.onload = () => {
   ctx.drawImage(boardImage, 12, 12)
}
