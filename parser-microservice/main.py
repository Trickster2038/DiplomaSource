import json
import os

from pyDigitalWaveTools.vcd.parser import VcdParser
from flask import Flask, request


# ===== File Parser =====


def parse_func(user_id, level_id, value_change_dump):
    fname = f"{user_id}_{level_id}.vcd"
    with open(fname, "w+") as vcd_file:
        vcd_file.write(value_change_dump)
    with open(fname) as vcd_file:
        vcd = VcdParser()
        vcd.parse(vcd_file)
        data = vcd.scope.toJson()

    os.remove(fname)

    return data


# ===== Flask Server =====
app = Flask(__name__)


@app.route('/parse', methods=['POST'])
def parse_handler():
    content_type = request.headers.get('Content-Type')
    if (content_type == 'application/json'):
        req = request.json
        if not (("user_id" in req) and ("level_id" in req)
                and ("data" in req)):
            return {
                "status_str": "error",
                "status_code": 400,
                "message": "JSON missing fields"
            }, 400
        else:
            try:
                vcd_parsed = parse_func(str(req["user_id"]), str(req["level_id"]),
                                        str(req["data"]))
                return {
                    "status_str": "ok",
                    "status_code": 200,
                    # FIXME (actual?): get only first child
                    "data": list(filter(lambda signal: signal["type"]["name"] != "struct", vcd_parsed["children"][0]["children"]))
                }, 200
            except Exception as e:
                return {
                    "status_str": "error",
                    "status_code": 400,
                    "message": f"VCD parsing error: {e}"
                }, 400
    else:
        return {
            "status_str": "error",
            "status_code": 400,
            "message": "invalid Content-Type"
        }, 400


if __name__ == "__main__":
    app.run(host="0.0.0.0")
