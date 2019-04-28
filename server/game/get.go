package game

import "../vector"

func (game *Game) GetCoorPlayer1() vector.Vector2 {
	return game.coor_player1
}

func (game *Game) GetCoorBall() vector.Vector2f {
	return game.coor_ball
}

func (game *Game) GetCoorPlayer2() vector.Vector2 {
	return game.coor_player2
}

func (game *Game) GetPlayer1Wins() int {
	return game.player1_wins
}

func (game *Game) GetPlayer2Wins() int {
	return game.player2_wins
}

func (game *Game) GetKeyPlayer1() string {
	return game.key_player1
}

func (game *Game) GetKeyPlayer2() string {
	return game.key_player2
}
