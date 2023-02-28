import requests

def send_request(port, url, payload):
    resp = requests.post(f"http://127.0.0.1:{port}/{url}", json = payload)
    return resp

def is_ok_response(response):
    return response.status_code == 200 \
        and response.json()["status_code"] == 200 \
            and response.json()["status_str"] == "ok"