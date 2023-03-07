from locust import HttpUser, task
from const import *
import copy
import random


class BaseUser(HttpUser):
    @task
    def stats_general_each_level_passed(self):
        self.client.post(
            "/stats", json=Stats.request_general_each_level_passed)

    @task
    def stats_personal_levels_statuses(self):
        self.client.post(
            "/stats", json=Stats.request_personal_each_level_passed)

    @task
    def stats_general_solutions_dist(self):
        self.client.post("/stats", json=Stats.request_general_solutions_dist)

    @task
    def crud_read_levelsdata(self):
        self.client.post("/levels", json=CRUD.request_read_levels_data)

    # FIXME: failing on connections overflow
    # @task
    # def analyzer_check_code_ok(self):
    #     payload = copy.deepcopy(Analyzer.request_check_program_ok)
    #     payload["user_id"] = random.randint(0, 10000)
    #     self.client.post("/check", json=payload)
