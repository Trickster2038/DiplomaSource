import utils.settings as settings
import utils.utils as utils
from utils.consts import *

import allure
import copy
import json

def ordered(obj):
    if isinstance(obj, dict):
        return sorted((k, ordered(v)) for k, v in obj.items())
    if isinstance(obj, list):
        return sorted(ordered(x) for x in obj)
    else:
        return obj


@allure.description("Test for valid diagram")
@allure.epic("Unit-testing")
@allure.story("Wavedrom")
def test_wavedrom_correct_positive():
    resp = utils.send_request(settings.WAVEDROM_PORT,
                              "wavedrom", Wavedrom.valid_request)
    assert utils.is_ok_response(resp)
    assert ordered(resp.json()) == ordered(Wavedrom.valid_response)


@allure.description("Test for invalid diagram")
@allure.epic("Unit-testing")
@allure.story("Wavedrom")
def test_wavedrom_error_format():
    payload = copy.deepcopy(Wavedrom.valid_request)
    payload["data"] = "jinnjkln"
    resp = utils.send_request(settings.WAVEDROM_PORT,
                              "wavedrom", payload)
    assert utils.is_error_response(resp)
