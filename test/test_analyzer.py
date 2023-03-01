import utils.settings as settings
import utils.utils as utils
from utils.consts import *

import allure
import copy

@allure.description("Test for singlechoice right-answered task")
@allure.epic("Unit-testing")
@allure.story("Analyzer")
def test_single_correct_positive():
    resp = utils.send_request(settings.ANALYZER_PORT, \
        "check", Analyzer.single_valid_positive)
    assert utils.is_ok_response(resp)
    assert resp.json()["is_correct"] == True

@allure.description("Test for singlechoice wrong-answered task")
@allure.epic("Unit-testing")
@allure.story("Analyzer")
def test_single_correct_negative():
    payload = copy.deepcopy(Analyzer.single_valid_positive)
    payload["data"]["user_answer_id"] = 1
    resp = utils.send_request(settings.ANALYZER_PORT, \
        "check", payload)
    assert utils.is_ok_response(resp)
    assert resp.json()["is_correct"] == False
    assert resp.json()["data"]["hint"] ==  payload["data"]["task"]["answers"][1]["hint"]

@allure.description("Test for singlechoice task with overflow for answer IDs")
@allure.epic("Unit-testing")
@allure.story("Analyzer")
def test_single_error_overflow():
    payload = copy.deepcopy(Analyzer.single_valid_positive)
    payload["data"]["user_answer_id"] = 6
    payload["data"]["task"]["correct_answer_id"] = 6
    resp = utils.send_request(settings.ANALYZER_PORT, \
        "check", payload)
    assert resp.json()["is_correct"] == False

@allure.description("Test for multichoice right-answered task")
@allure.epic("Unit-testing")
@allure.story("Analyzer")
def test_multi_correct_positive():
    resp = utils.send_request(settings.ANALYZER_PORT, \
        "check", Analyzer.multi_valid_positive)
    assert utils.is_ok_response(resp)
    assert resp.json()["is_correct"] == True
    assert resp.json()["data"]["false_positive"] == False
    assert resp.json()["data"]["false_negative"] == False

@allure.description("Test for multichoice false positive answered task")
@allure.epic("Unit-testing")
@allure.story("Analyzer")
def test_multi_correct_false_positive():
    payload = copy.deepcopy(Analyzer.multi_valid_positive)
    payload["data"]["task"]["correct_answers"][0] = False
    resp = utils.send_request(settings.ANALYZER_PORT, \
        "check", payload)
    assert utils.is_ok_response(resp)
    assert resp.json()["is_correct"] == False
    assert resp.json()["data"]["false_positive"] == True
    assert resp.json()["data"]["false_negative"] == False

@allure.description("Test for multichoice false negative answered task")
@allure.epic("Unit-testing")
@allure.story("Analyzer")
def test_multi_correct_false_negative():
    payload = copy.deepcopy(Analyzer.multi_valid_positive)
    payload["data"]["user_answers"][0] = False
    resp = utils.send_request(settings.ANALYZER_PORT, \
        "check", payload)
    assert utils.is_ok_response(resp)
    assert resp.json()["is_correct"] == False
    assert resp.json()["data"]["false_positive"] == False
    assert resp.json()["data"]["false_negative"] == True

@allure.description("Test for multichoice task with mismatching answer and reference size")
@allure.epic("Unit-testing")
@allure.story("Analyzer")
def test_multi_error_size_mismatch():
    payload = copy.deepcopy(Analyzer.multi_valid_positive)
    payload["data"]["user_answers"].append(True)
    resp = utils.send_request(settings.ANALYZER_PORT, \
        "check", payload)
    assert utils.is_error_response(resp)
    assert resp.json()["is_correct"] == False

