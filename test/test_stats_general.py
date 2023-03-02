import utils.settings as settings
import utils.utils as utils
from utils.consts import *

import allure
import copy
import pytest


@allure.description("Test for invalid VCD file")
@allure.epic("Unit-testing")
@allure.story("General stats")
@pytest.mark.parametrize("payload, response", [
    (StatsGeneral.request_each_level_passed_correct,
     StatsGeneral.response_each_level_passed_correct),
    (StatsGeneral.request_each_avg_efforts_correct,
     StatsGeneral.response_each_avg_efforts_correct)])
def test_each_level_passed_correct(payload, response):
    resp = utils.send_request(settings.STATS_PORT,
                              "generalstats", payload)
    assert utils.is_ok_response(resp)
    assert utils.ordered_json(resp.json()) == utils.ordered_json(response)
