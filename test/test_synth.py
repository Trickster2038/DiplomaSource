import utils.settings as settings
import utils.utils as utils
from utils.consts import *

import allure
import copy

@allure.description("Test for valid source files")
@allure.epic("Unit-testing")
@allure.story("Synthesizer")
def test_correct_positive():
    resp = utils.send_request(settings.SYNTH_PORT,
                              "build", Synth.valid_request)
    assert utils.is_ok_response(resp)
    # cut date at the start of VCD file
    assert Synth.valid_response["data"][32:] in resp.json()["data"] 

@allure.description("Test for invalid device.v")
@allure.epic("Unit-testing")
@allure.story("Synthesizer")
def test_error_in_device():
    resp = utils.send_request(settings.SYNTH_PORT,
                              "build", Synth.bad_device_request)
    assert utils.is_error_response(resp)
    assert "synthethis error" in resp.json()["message"].lower()

@allure.description("Test for invalid testbecnch.v")
@allure.epic("Unit-testing")
@allure.story("Synthesizer")
def test_error_in_testbecnch():
    resp = utils.send_request(settings.SYNTH_PORT,
                              "build", Synth.bad_tb_request)
    assert utils.is_error_response(resp)
    assert "simulation error" in resp.json()["message"].lower()

@allure.description("Test for invalid testbecnch.v without \"$dumpvars\"")
@allure.epic("Unit-testing")
@allure.story("Synthesizer")
def test_error_in_testbecnch_no_dumpvars():
    resp = utils.send_request(settings.SYNTH_PORT,
                              "build", Synth.bad_tb_dumpvars_request)
    assert utils.is_error_response(resp)
    assert "testbench without \"$dumpvars\"" in resp.json()["message"].lower()