package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	SIZE_FIELD_X  = 800
	SIZE_FIELD_Y  = 600
	SIZE_RACKET_X = 10
	SIZE_RACKET_Y = 150
	SIZE_BALL     = 16
	SPEED_PLAYERS = 2
)

type vector2f struct {
	X, Y float64
}

func MakeVector2f(x, y float64) vector2f {
	return vector2f{x, y}
}

type vector2 struct {
	X, Y int
}

func MakeVector2(x, y int) vector2 {
	return vector2{x, y}
}

var direct_ball = vector2f{0.75, 0.25}
var speed_ball = 2.0

var coor_player1 = vector2{10,
	SIZE_FIELD_Y/2 - SIZE_RACKET_Y/2}

var coor_player2 = vector2{SIZE_FIELD_X - SIZE_RACKET_X - 10,
	SIZE_FIELD_Y/2 - SIZE_RACKET_Y/2}

var coor_ball = vector2f{SIZE_FIELD_X / 2, SIZE_FIELD_Y / 2}

var pos_mouse_player1 = vector2{0, 0}
var pos_mouse_player2 = vector2{0, 0}

var player1_wins = 0
var player2_wins = 0

var key_player1 string
var key_player2 string

var server_message = "Welcome server!"

type Setting struct {
	Key         string
	SIZE_FIELD  vector2
	SIZE_RACKET vector2
	SIZE_BALL   int
}

func registration(w http.ResponseWriter, r *http.Request) {
	h := sha1.New()
	io.WriteString(h, r.RemoteAddr)

	hash := fmt.Sprintf("%x", h.Sum(nil))

	setting := Setting{Key: hash,
		SIZE_FIELD:  MakeVector2(SIZE_FIELD_X, SIZE_FIELD_Y),
		SIZE_BALL:   SIZE_BALL,
		SIZE_RACKET: MakeVector2(SIZE_RACKET_X, SIZE_RACKET_Y)}

	b, err := json.Marshal(setting)

	if err != nil {
		log.Fatalf("faild mashaling settings: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%s", b)

	if key_player1 == "" {
		log.Println("Client 1 connected!")
		key_player1 = hash
	} else if key_player2 == "" {
		log.Println("Client 2 connected!")
		log.Println("Game play...")
		key_player2 = hash
		go Game()
	} else {
		w.WriteHeader(400)
		log.Println("Faild trying connect...")
	}
}

type Response struct {
	Player1      vector2
	Player2      vector2
	Ball         vector2
	Player1_wins int
	Player2_wins int
}

func handler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(400)
		log.Println("incorrect params")
		return
	}

	key := r.FormValue("key")

	if key != key_player1 && key != key_player2 {
		w.WriteHeader(400)
		log.Printf("unknown key %s!", key)
		return
	}

	pos_mouse_x := r.FormValue("x")
	pos_mouse_y := r.FormValue("y")

	x, err := strconv.Atoi(pos_mouse_x)

	if err != nil {
		w.WriteHeader(400)
		log.Printf("error parsing coordinates mouse %v", err)
		return
	}

	y, err := strconv.Atoi(pos_mouse_y)
	if err != nil {
		w.WriteHeader(400)
		log.Printf("error parsing coordinates mouse %v", err)
		return
	}

	if key == key_player1 {
		pos_mouse_player1.X = x
		pos_mouse_player1.Y = y
	} else {
		pos_mouse_player2.X = x
		pos_mouse_player2.Y = y
	}

	w.Header().Set("Content-Type", "application/json")

	response := Response{
		Player1:      coor_player1,
		Player2:      coor_player2,
		Ball:         MakeVector2(int(coor_ball.X), int(coor_ball.Y)),
		Player1_wins: player1_wins,
		Player2_wins: player2_wins}

	b, err := json.Marshal(response)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	fmt.Fprintf(w, "%s", b)
}

func main() {
	http.HandleFunc("/handle", handler)
	http.HandleFunc("/registration", registration)
	log.Print("Start server port 8080...")
	log.Print("Waiting clients...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Game() {

	for {
		if coor_player1.Y-pos_mouse_player1.Y < 0 {
			if coor_player1.Y+SIZE_RACKET_Y < 800 {
				coor_player1.Y += SPEED_PLAYERS
			}
		}

		if coor_player1.Y-pos_mouse_player1.Y > -5 {
			if coor_player1.Y > 0 {
				coor_player1.Y -= SPEED_PLAYERS
			}
		}

		if coor_player2.Y-pos_mouse_player2.Y < 0 {
			if coor_player2.Y+SIZE_RACKET_Y < 800 {
				coor_player2.Y += SPEED_PLAYERS
			}
		}

		if coor_player2.Y-pos_mouse_player2.Y > -5 {
			if coor_player2.Y > 0 {
				coor_player2.Y -= SPEED_PLAYERS
			}
		}

		time.Sleep(5 * time.Millisecond)

		IsCollision()

		coor_ball = vector2f{coor_ball.X + direct_ball.X*speed_ball, coor_ball.Y + direct_ball.Y*speed_ball}

		if MoveOut() {
			coor_ball = vector2f{SIZE_FIELD_X / 2, SIZE_FIELD_Y / 2}
			direct_ball = vector2f{-direct_ball.X, direct_ball.Y}
		}
	}
}

func MoveOut() bool {

	if coor_ball.X > SIZE_FIELD_X-SIZE_BALL {
		player1_wins += 1
		return true
	}

	if coor_ball.X < 0 {
		player2_wins += 1
		return true
	}

	return false
}

func IsCollision() {

	if int(coor_ball.Y) >= coor_player1.Y-SIZE_BALL &&
		int(coor_ball.Y) <= coor_player1.Y+SIZE_RACKET_Y &&
		int(coor_ball.X) <= coor_player1.X+SIZE_RACKET_X {

		direct_ball = vector2f{-direct_ball.X, direct_ball.Y}
	}

	if int(coor_ball.Y) >= coor_player2.Y-SIZE_BALL &&
		int(coor_ball.Y) <= coor_player2.Y+SIZE_RACKET_Y &&
		int(coor_ball.X+SIZE_BALL) > coor_player2.X {

		direct_ball = vector2f{-direct_ball.X, direct_ball.Y}
	}

	if coor_ball.Y > SIZE_FIELD_Y-SIZE_BALL || coor_ball.Y < 0 {
		direct_ball = vector2f{direct_ball.X, -direct_ball.Y}
	}

}
