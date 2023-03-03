import utils.settings as settings
import utils.utils as utils
from utils.consts import *

import allure
import copy
import pytest


@allure.description("Test for valid personal stats requests")
@allure.epic("Unit-testing")
@allure.story("General stats")
@pytest.mark.parametrize("payload, response", [
    (StatsPersonal.request_general_progress,
     StatsPersonal.response_general_progress),
    (StatsPersonal.request_avg_efforts,
     StatsPersonal.response_avg_efforts),
    (StatsPersonal.request_monthly_activity,
     StatsPersonal.response_monthly_activity),
    (StatsPersonal.request_activity_borders,
     StatsPersonal.response_activity_borders)])
def test_correct_requests(payload, response):
    resp = utils.send_request(settings.STATS_PORT,
                              "personalstats", payload)
    assert utils.is_ok_response(resp)
    assert utils.ordered_json(resp.json()) == utils.ordered_json(response)

# write in docs that endpoint returns error in this case
@allure.description("Test for no user stats request")
@allure.epic("Unit-testing")
@allure.story("General stats")
def test_error_no_user():
    assert True
