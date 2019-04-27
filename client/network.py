import requests
import json

class Network:
    def __init__(self):
        self.key = ""

    def registration(self):
        r = requests.post("http://127.0.0.1:8080/registration")
        if r.status_code == 200:
            setting = r.json()
            self.key = setting["Key"]
            return setting, True
        else:
            return setting, False

    def send_request(self, mouse_pos_x, mouse_pos_y):
        resp = requests.post("http://127.0.0.1:8080/handle", data={"x" : mouse_pos_x, "y": mouse_pos_y, "key" : self.key})

        if resp.status_code == 200:
            return resp.json(), True

        return [], False

