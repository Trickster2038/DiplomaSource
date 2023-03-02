import utils.settings as settings
import utils.utils as utils
from utils.consts import *

import allure
import copy

@allure.description("Test for valid VCD file")
@allure.epic("Unit-testing")
@allure.story("Parser")
def test_correct_positive():
    resp = utils.send_request(settings.PARSER_PORT,
                              "parse", Parser.valid_request)
    assert utils.is_ok_response(resp)
    assert utils.ordered_json_safe_list(resp.json()) == utils.ordered_json_safe_list(Parser.valid_response)

@allure.description("Test for invalid VCD file")
@allure.epic("Unit-testing")
@allure.story("Parser")
def test_error_in_vcd():
    payload = copy.deepcopy(Parser.valid_request)
    payload["data"] = "jinnjkln"
    resp = utils.send_request(settings.PARSER_PORT,
                              "parse", payload)
    assert utils.is_error_response(resp)
    assert "VCD parsing error" in resp.json()["message"]