import requests
import allure
import json

@allure.step("Send request")
def send_request(port, url, payload):
    path = f"http://127.0.0.1:{port}/{url}"
    allure.attach(path, 'Request URL', allure.attachment_type.TEXT)
    allure.attach(json.dumps(payload, indent=4, ensure_ascii=False).encode(), 'Request payload', allure.attachment_type.TEXT)
    resp = requests.post(path, json = payload)
    allure.attach(json.dumps(resp.json(), indent=4, ensure_ascii=False).encode(), 'Response payload', allure.attachment_type.TEXT)
    return resp

@allure.step("Check response [ok]")
def is_ok_response(response):
    allure.attach(json.dumps(response.json(), indent=4, ensure_ascii=False).encode(), 'Response payload', allure.attachment_type.TEXT)
    return response.status_code == 200 \
        and response.json()["status_code"] == 200 \
            and response.json()["status_str"] == "ok"

@allure.step("Check response [error]")
def is_error_response(response):
    allure.attach(json.dumps(response.json(), indent=4, ensure_ascii=False).encode(), 'Response payload', allure.attachment_type.TEXT)
    return response.status_code == 400 \
        and response.json()["status_code"] == 400 \
            and response.json()["status_str"] == "error"