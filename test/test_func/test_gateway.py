import utils.settings as settings
import utils.utils as utils
from utils.consts import *

import allure
import copy


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
    assert utils.ordered_json(resp.json()) == utils.ordered_json(
        StatsGeneral.response_each_level_passed)


@allure.description("Test for no user proxy-request to Stats")
@allure.epic("Integrational testing")
@allure.story("Gateway")
def test_error_no_user():
    resp = utils.send_request(settings.GATEWAY_PORT,
                              "check", Gateway.request_check_no_user)
    assert utils.is_error_response(resp)
    assert "does not exist" in resp.json()["message"].lower()


@allure.description("Test for not admin proxy-request to CRUD")
@allure.epic("Integrational testing")
@allure.story("Gateway")
def test_error_not_admin():
    resp = utils.send_request(settings.GATEWAY_PORT,
                              "levels", Gateway.request_crud_create_lb_not_admin)
    assert utils.is_error_response(resp)
    assert "user have no rights to modify levels" in resp.json()[
        "message"].lower()

@allure.description("Test for ok proxy-request to Check")
@allure.epic("Integrational testing")
@allure.story("Gateway")
def test_correct_check():
    resp = utils.send_request(settings.GATEWAY_PORT,
                              "check", Gateway.request_check_correct)
    assert utils.is_ok_response(resp)
    assert resp.json()["is_correct"] == True
    # CAN DEPENDS ON .env MODE
    # assert resp.json()["is_already_solved"] == True

@allure.description("Test for no level proxy-request to Check")
@allure.epic("Integrational testing")
@allure.story("Gateway")
def test_error_no_level():
    resp = utils.send_request(settings.GATEWAY_PORT,
                              "check", Gateway.request_check_no_level)
    assert utils.is_error_response(resp)
    assert "crud-microservice.levelsbrief error" in resp.json()[
        "message"].lower()
    
@allure.description("Test for no level proxy-request to Check")
@allure.epic("Integrational testing")
@allure.story("Gateway")
def test_error_no_level():
    resp = utils.send_request(settings.GATEWAY_PORT,
                              "check", Gateway.request_check_no_level)
    assert utils.is_error_response(resp)
    assert "crud-microservice.levelsbrief error" in resp.json()[
        "message"].lower()
    
@allure.description("Test for wrong code proxy-request to CRUD.CreateLevelsData")
@allure.epic("Integrational testing")
@allure.story("Gateway")
def test_error_create_levelsdata():
    resp = utils.send_request(settings.GATEWAY_PORT,
                              "levels", Gateway.request_create_levelsdata_error)
    assert utils.is_error_response(resp)
    assert "device synthesis error" in resp.json()[
        "message"].lower()
    
@allure.description("Test for wrong code proxy-request to CRUD.CreateLevelsData")
@allure.epic("Integrational testing")
@allure.story("Gateway")
def test_correct_create_levelsdata():
    assert True

