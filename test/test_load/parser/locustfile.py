from locust import HttpUser, task
from const import *
import copy
import logging
import random


class BaseUser(HttpUser):
    @task
    def synth_correct(self):
        payload = copy.deepcopy(request_parser_correct)
        payload["user_id"] = random.randint(0, 10000)
        payload["level_id"] = random.randint(0, 10000)
        id = payload["user_id"]
        logging.info(f"parse with user ID: {str(id)}")
        self.client.post(
            "/parse", json=payload)
