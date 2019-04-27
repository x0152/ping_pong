import pygame


SIZE_GAME_PLACE_X = 800
SIZE_GAME_PLACE_Y = 600 

D_X = SIZE_GAME_PLACE_X / 800
D_Y = SIZE_GAME_PLACE_Y / 600

SIZE_RACKET_X = 10
SIZE_RACKET_Y = 150

SIZE_BALL = 16

COLOR_BACKGROUND = pygame.Color(0x00, 0x00, 0x00)

CAPTION_WINDOW = "Ping pong"

def configuration(setting):
    SIZE_GAME_PLACE_X = int(setting["SIZE_FIELD"]["X"])
    SIZE_GAME_PLACE_Y = int(setting["SIZE_FIELD"]["Y"])

    SIZE_RACKET_X = int(setting["SIZE_RACKET"]["X"])
    SIZE_RACKET_Y = int(setting["SIZE_RACKET"]["Y"])

    SIZE_BALL = int(setting["SIZE_BALL"])
