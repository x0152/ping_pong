package game

import (
	"time"

	"../consts"
	"../vector"
)

type Game struct {
	direct_ball  vector.Vector2f
	speed_ball   float64
	coor_player1 vector.Vector2
	coor_player2 vector.Vector2

	coor_ball vector.Vector2f

	pos_mouse_player1 vector.Vector2
	pos_mouse_player2 vector.Vector2
	player1_wins      int
	player2_wins      int

	key_player1 string
	key_player2 string
}

func (game *Game) Init() {
	game.direct_ball = vector.Vector2f{0.75, 0.25}
	game.speed_ball = 2.0
	game.coor_player1 = vector.Vector2{10,
		consts.SIZE_FIELD_Y/2 - consts.SIZE_RACKET_Y/2}

	game.coor_player2 = vector.Vector2{consts.SIZE_FIELD_X - consts.SIZE_RACKET_X - 10,
		consts.SIZE_FIELD_Y/2 - consts.SIZE_RACKET_Y/2}

	game.coor_ball = vector.Vector2f{consts.SIZE_FIELD_X / 2, consts.SIZE_FIELD_Y / 2}

	game.pos_mouse_player1 = vector.Vector2{0, 0}
	game.pos_mouse_player2 = vector.Vector2{0, 0}
	game.player1_wins = 0
	game.player2_wins = 0

}

func (game *Game) Play() {

	for {
		if game.coor_player1.Y-game.pos_mouse_player1.Y < 0 {
			if game.coor_player1.Y+consts.SIZE_RACKET_Y < 800 {
				game.coor_player1.Y += consts.SPEED_PLAYERS
			}
		}

		if game.coor_player1.Y-game.pos_mouse_player1.Y > -5 {
			if game.coor_player1.Y > 0 {
				game.coor_player1.Y -= consts.SPEED_PLAYERS
			}
		}

		if game.coor_player2.Y-game.pos_mouse_player2.Y < 0 {
			if game.coor_player2.Y+consts.SIZE_RACKET_Y < 800 {
				game.coor_player2.Y += consts.SPEED_PLAYERS
			}
		}

		if game.coor_player2.Y-game.pos_mouse_player2.Y > -5 {
			if game.coor_player2.Y > 0 {
				game.coor_player2.Y -= consts.SPEED_PLAYERS
			}
		}

		time.Sleep(5 * time.Millisecond)

		game.HandleCollision()

		game.coor_ball = vector.Vector2f{game.coor_ball.X + game.direct_ball.X*game.speed_ball, game.coor_ball.Y + game.direct_ball.Y*game.speed_ball}

		game.HandleMoveOut()
	}
}

func (game *Game) HandleMoveOut() {

	if game.coor_ball.X > consts.SIZE_FIELD_X-consts.SIZE_BALL {
		game.player1_wins += 1
		game.coor_ball = vector.Vector2f{consts.SIZE_FIELD_X / 2, consts.SIZE_FIELD_Y / 2}
		game.direct_ball = vector.Vector2f{-game.direct_ball.X, game.direct_ball.Y}
	}

	if game.coor_ball.X < 0 {
		game.player2_wins += 1
		game.coor_ball = vector.Vector2f{consts.SIZE_FIELD_X / 2, consts.SIZE_FIELD_Y / 2}
		game.direct_ball = vector.Vector2f{-game.direct_ball.X, game.direct_ball.Y}
	}
}

func (game *Game) HandleCollision() {

	if int(game.coor_ball.Y) >= game.coor_player1.Y-consts.SIZE_BALL &&
		int(game.coor_ball.Y) <= game.coor_player1.Y+consts.SIZE_RACKET_Y &&
		int(game.coor_ball.X) <= game.coor_player1.X+consts.SIZE_RACKET_X {

		game.direct_ball = vector.Vector2f{-game.direct_ball.X, game.direct_ball.Y}
	}

	if int(game.coor_ball.Y) >= game.coor_player2.Y-consts.SIZE_BALL &&
		int(game.coor_ball.Y) <= game.coor_player2.Y+consts.SIZE_RACKET_Y &&
		int(game.coor_ball.X+consts.SIZE_BALL) > game.coor_player2.X {

		game.direct_ball = vector.Vector2f{-game.direct_ball.X, game.direct_ball.Y}
	}

	if game.coor_ball.Y > consts.SIZE_FIELD_Y-consts.SIZE_BALL || game.coor_ball.Y < 0 {
		game.direct_ball = vector.Vector2f{game.direct_ball.X, -game.direct_ball.Y}
	}

}
