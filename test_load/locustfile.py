from locust import HttpUser, task
from const import *


class BaseUser(HttpUser):
    @task
    def analyzer_code_ok(self):
        # host = "https://10.247.123.172" 
        self.client.post("/check", json=Analyzer.request_code_ok
                         )
