package game

import "../vector"

func (game *Game) SetCoorMousePlayer1(coor vector.Vector2) {
	game.pos_mouse_player1 = coor
}

func (game *Game) SetCoorMousePlayer2(coor vector.Vector2) {
	game.pos_mouse_player2 = coor
}

func (game *Game) SetKeyPlayer1(key string) {
	game.key_player1 = key
}

func (game *Game) SetKeyPlayer2(key string) {
	game.key_player2 = key
}
