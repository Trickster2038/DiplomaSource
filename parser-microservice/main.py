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
    # return json.dumps(data, indent=4, sort_keys=True)


# ===== Flask Server =====
app = Flask(__name__)


@app.route('/parse', methods=['POST'])
def parse_handler():
    content_type = request.headers.get('Content-Type')
    if (content_type == 'application/json'):
        req = request.json
        if not(("user_id" in req) and ("level_id" in req)
               and ("value_change_dump" in req)):
            return {
                "status": "error",
                "code": 400,
                "message": "JSON missing fields"
            }, 400
        else:
            try:
                vcd_parsed = parse_func(str(req["user_id"]), str(req["level_id"]),
                                   str(req["value_change_dump"]))
                return {
                    "status": "ok",
                    "code": 200,
                    "vcd_parsed": vcd_parsed #FIXME: get only first child
                }, 200
            except Exception as e:
                return {
                    "status": "error",
                    "code": 400,
                    "message": f"VCD parse error: {e}"
                }, 400
    else:
        return {
            "status": "error",
            "code": 400,
            "message": "invalid Content-Type"
        }, 400


if __name__ == "__main__":
    app.run(host="0.0.0.0")