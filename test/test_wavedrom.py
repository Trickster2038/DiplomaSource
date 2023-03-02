import utils.settings as settings
import utils.utils as utils
from utils.consts import *

import allure
import copy

@allure.description("Test for valid diagram")
@allure.epic("Unit-testing")
@allure.story("Wavedrom")
def test_correct_positive():
    resp = utils.send_request(settings.WAVEDROM_PORT,
                              "wavedrom", Wavedrom.valid_request)
    assert utils.is_ok_response(resp)
    assert utils.ordered_json(resp.json()) == utils.ordered_json(Wavedrom.valid_response)


@allure.description("Test for invalid diagram")
@allure.epic("Unit-testing")
@allure.story("Wavedrom")
def test_error_format():
    payload = copy.deepcopy(Wavedrom.valid_request)
    payload["data"] = "jinnjkln"
    resp = utils.send_request(settings.WAVEDROM_PORT,
                              "wavedrom", payload)
    assert utils.is_error_response(resp)
