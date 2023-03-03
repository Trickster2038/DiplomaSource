import utils.settings as settings
import utils.utils as utils
from utils.consts import *

import allure
import copy
import pytest


@allure.description("Test for proxy-request to CRUD")
@allure.epic("Integrational testing")
@allure.story("Gateway")
def test_correct_proxy_crud():
    resp = utils.send_request(settings.GATEWAY_PORT,
                              "user", CRUD.request_read_user)
    assert utils.is_ok_response(resp)
    assert resp.json()["data"]["nickname"] == "Deni"

@allure.description("Test for proxy-request to Stats")
@allure.epic("Integrational testing")
@allure.story("Gateway")
def test_correct_proxy_stats():
    resp = utils.send_request(settings.GATEWAY_PORT,
                              "stats", Gateway.request_stats_each_level_passed)
    assert utils.is_ok_response(resp)
    assert utils.ordered_json(resp.json()) == utils.ordered_json(StatsGeneral.response_each_level_passed)

# TODO: more tests: no user, not admin, no level, fail code creation/check, ok creation/check