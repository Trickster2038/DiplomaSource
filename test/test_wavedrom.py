import utils.settings as settings
import utils.utils as utils
from utils.consts import *

import allure
import copy
import json

@allure.description("Test for valid diagram")
@allure.epic("Unit-testing")
@allure.story("Wavedrom")
def test_wavedrom_correct_positive():
    resp = utils.send_request(settings.WAVEDROM_PORT,
                              "wavedrom", Wavedrom.valid_request)
    assert utils.is_ok_response(resp)
    assert json.dumps(resp.json()) == json.dumps(Wavedrom.valid_response)

@allure.description("Test for invalid diagram")
@allure.epic("Unit-testing")
@allure.story("Wavedrom")
def test_wavedrom_error_format():
    payload = copy.deepcopy(Wavedrom.valid_request)
    payload["data"] = "jinnjkln"
    resp = utils.send_request(settings.WAVEDROM_PORT,
                              "wavedrom", payload)
    assert utils.is_error_response(resp)