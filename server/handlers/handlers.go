package handlers

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"../consts"
	"../game"
	"../vector"
)

type Setting struct {
	Key         string
	SIZE_FIELD  vector.Vector2
	SIZE_RACKET vector.Vector2
	SIZE_BALL   int
}

var g = game.Game{}

func Registration(w http.ResponseWriter, r *http.Request) {
	h := sha1.New()
	io.WriteString(h, r.RemoteAddr)

	hash := fmt.Sprintf("%x", h.Sum(nil))

	setting := Setting{Key: hash,
		SIZE_FIELD:  vector.MakeVector2(consts.SIZE_FIELD_X, consts.SIZE_FIELD_Y),
		SIZE_BALL:   consts.SIZE_BALL,
		SIZE_RACKET: vector.MakeVector2(consts.SIZE_RACKET_X, consts.SIZE_RACKET_Y)}

	b, err := json.Marshal(setting)

	if err != nil {
		log.Fatalf("faild mashaling settings: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%s", b)

	if g.GetKeyPlayer1() == "" {
		log.Println("Client 1 connected!")
		g.SetKeyPlayer1(hash)
	} else if g.GetKeyPlayer2() == "" {
		log.Println("Client 2 connected!")
		log.Println("Game play...")
		g.SetKeyPlayer2(hash)

		g.Init()
		go g.Play()
	} else {
		w.WriteHeader(400)
		log.Println("Faild trying connect...")
	}
}

type Response struct {
	Player1      vector.Vector2
	Player2      vector.Vector2
	Ball         vector.Vector2
	Player1_wins int
	Player2_wins int
}

func Handler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(400)
		log.Println("incorrect params")
		return
	}

	key := r.FormValue("key")

	if key != g.GetKeyPlayer1() && key != g.GetKeyPlayer2() {
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

	if key == g.GetKeyPlayer1() {
		g.SetCoorMousePlayer1(vector.Vector2{x, y})
	} else {
		g.SetCoorMousePlayer2(vector.Vector2{x, y})
	}

	w.Header().Set("Content-Type", "application/json")

	response := Response{
		Player1:      g.GetCoorPlayer1(),
		Player2:      g.GetCoorPlayer2(),
		Ball:         vector.MakeVector2(int(g.GetCoorBall().X), int(g.GetCoorBall().Y)),
		Player1_wins: g.GetPlayer1Wins(),
		Player2_wins: g.GetPlayer2Wins()}

	b, err := json.Marshal(response)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	fmt.Fprintf(w, "%s", b)
}
