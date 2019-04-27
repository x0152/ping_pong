import math
import pygame
import config 
import field
import network

class Game:

    def __init__(self):
        self.network = network.Network()
        pass

    def Start(self):
        pygame.init()

        screen = pygame.display.set_mode([config.SIZE_GAME_PLACE_X, config.SIZE_GAME_PLACE_Y])
        screen.fill(config.COLOR_BACKGROUND)

        pygame.display.set_caption(config.CAPTION_WINDOW)
        clock = pygame.time.Clock()

        f = field.Field()

        setting, ok = self.network.registration() 
        if ok == False:
            print("faild registration!")
            return

        config.configuration(setting)

        done = False
        while done == False:
            pygame.display.flip()

            screen.fill(config.COLOR_BACKGROUND)

            mouse = pygame.mouse.get_pos()
            data, ok = self.network.send_request(int(config.D_X * mouse[0]), int(config.D_Y * mouse[1]))

            f.update(data)
            
            if ok == False:
                print("faild send request!")
                return

            f.draw(screen)

            clock.tick(60)

            for event in pygame.event.get():
                if event.type == pygame.QUIT:
                    done = True


        pygame.quit()

