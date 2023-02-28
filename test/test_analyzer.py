import utils.settings as settings
import utils.utils as utils
from utils.consts import *

def test_correct_positive():
    resp = utils.send_request(settings.ANALYZER_PORT, \
        "check", Analyzer.single_valid_positive)
    assert utils.is_ok_response(resp)
    assert resp.json()["is_correct"] == True