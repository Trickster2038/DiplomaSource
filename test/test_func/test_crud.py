import utils.settings as settings
import utils.utils as utils
from utils.consts import *

import allure
import copy
import pytest


@allure.description("Test for valid read requests")
@allure.epic("Unit-testing")
@allure.story("CRUD")
@pytest.mark.parametrize("payload, response", [
    (CRUD.request_read_user,
     CRUD.response_read_user),
    (CRUD.request_read_level_brief,
     CRUD.response_read_level_brief),
    (CRUD.request_read_level_data,
     CRUD.response_read_level_data),
    (CRUD.request_read_all_level_brief,
        CRUD.response_read_all_level_brief),
    (CRUD.request_check_succesful,
     CRUD.response_check_succesful)])
def test_correct_read_requests(payload, response):
    resp = utils.send_request(settings.CRUD_PORT,
                              "crud", payload)
    assert utils.is_ok_response(resp)
    assert utils.ordered_json(resp.json()) == utils.ordered_json(response)

@allure.description("Test for valid read requests")
@allure.epic("Unit-testing")
@allure.story("CRUD")
@pytest.mark.parametrize("payload, response", [
    (CRUD.request_read_user,
     CRUD.response_read_user),
    (CRUD.request_read_level_brief,
     CRUD.response_read_level_brief),
    (CRUD.request_read_level_data,
     CRUD.response_read_level_data)])
def test_correct_create_requests(payload, response):
    assert True

@allure.description("Test for valid create requests")
@allure.epic("Unit-testing")
@allure.story("CRUD")
@pytest.mark.parametrize("payload, response", [
    (CRUD.request_read_user,
     CRUD.response_read_user),
    (CRUD.request_read_level_brief,
     CRUD.response_read_level_brief)])
def test_correct_update_requests(payload, response):
    assert True


@allure.description("Test for valid update requests")
@allure.epic("Unit-testing")
@allure.story("CRUD")
@pytest.mark.parametrize("payload, response", [
    (CRUD.request_read_user,
     CRUD.response_read_user),
    (CRUD.request_read_level_brief,
     CRUD.response_read_level_brief)])
def test_correct_update_requests(payload, response):
    assert True

@allure.description("Test for valid delete requests")
@allure.epic("Unit-testing")
@allure.story("CRUD")
@pytest.mark.parametrize("payload, response", [
    (CRUD.request_read_user,
     CRUD.response_read_user),
    (CRUD.request_read_level_brief,
     CRUD.response_read_level_brief)])
def test_correct_update_requests(payload, response):
    assert True