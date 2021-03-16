canv = Document.querySelector("#canvas")
let ctx = canv.getContext('2d')

let boardImage = new Image()
boardImage.src = 'board.png'

window.onload = () => {
   ctx.drawImage(boardImage, 12, 12)
}
