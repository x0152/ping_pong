import pygame.freetype
from vector import vector2
from config import *

class Field:
    def __init__(self):
        middle_possition_y = SIZE_GAME_PLACE_Y / 2
        middle_possition_x = SIZE_GAME_PLACE_X / 2

        half_racket_y = SIZE_RACKET_Y / 2

        padding = 10

        self._player1_wins = 0
        self._player2_wins = 0
        self._player1 = vector2(padding, middle_possition_y - half_racket_y) 
        self._player2 = vector2(SIZE_GAME_PLACE_X - padding - SIZE_RACKET_X, middle_possition_y - half_racket_y)
        self._ball = vector2(middle_possition_x, middle_possition_y)

    def draw(self, screen):
        color = (255, 255, 255)
        color_red = (255, 0, 0)
        color_gray = (100, 100, 100)

        rect_line = pygame.Rect(SIZE_GAME_PLACE_X / 2 - 5, 0, 10, SIZE_GAME_PLACE_Y)
        pygame.draw.rect(screen, color_gray, rect_line) 

        rect_ball = pygame.Rect(self._ball.GetX(), self._ball.GetY(), SIZE_BALL, SIZE_BALL)
        pygame.draw.rect(screen, color_red, rect_ball) 

        rect_player1 = pygame.Rect(self._player1.GetX(), self._player1.GetY(), SIZE_RACKET_X, SIZE_RACKET_Y)
        rect_player2 = pygame.Rect(self._player2.GetX(), self._player2.GetY(), SIZE_RACKET_X, SIZE_RACKET_Y)
        pygame.draw.rect(screen, color,  rect_player1)
        pygame.draw.rect(screen, color, rect_player2) 


        font = pygame.font.SysFont("comicsansms", 72)
        text = font.render(str(self._player1_wins), True, color_red)
        screen.blit(text, (SIZE_GAME_PLACE_X / 4, SIZE_GAME_PLACE_Y / 10 + 10))

        font = pygame.font.SysFont("comicsansms", 72)
        text = font.render(str(self._player2_wins), True, color_red)
        screen.blit(text, (SIZE_GAME_PLACE_X / 4 * 3, SIZE_GAME_PLACE_Y / 10 + 10))

        pass
    def update(self, data):
        c_pl1 = vector2(int(data["Player1"]["X"]), int(data["Player1"]["Y"]))
        c_pl2 = vector2(int(data["Player2"]["X"]), int(data["Player2"]["Y"]))
        c_ball = vector2(int(data["Ball"]["X"]), int(data["Ball"]["Y"]))
        pl1_w = int(data["Player1_wins"])
        pl2_w = int(data["Player2_wins"])

        self.r_update(c_pl1, c_pl2, c_ball, pl1_w, pl2_w)

    #Coordinates player1, player2 and ball
    def r_update(self, player1, player2, ball, player1_wins, player2_wins):
        self._player1 = vector2(player1.GetX() * D_X, player1.GetY() * D_Y)
        self._player2 = vector2(player2.GetX() * D_X, player2.GetY() * D_Y)
        self._ball = vector2(ball.GetX() * D_X, ball.GetY() / D_Y)

        self._player1_wins = player1_wins
        self._player2_wins = player2_wins
        
        pass
