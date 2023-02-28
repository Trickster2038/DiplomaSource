import lib.settings as settings
import lib.utils as utils
from lib.consts import *

def test_correct_negative():
    resp = utils.send_request(settings.ANALYZER_PORT, \
        "check", Analyzer.single_valid_positive)
    assert utils.is_ok_response(resp)
    assert resp.json()["is_correct"] == True