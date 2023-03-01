class Analyzer:

    single_valid_positive = {
        "type": "singlechoice_test",
        "data": {
            "user_answer_id": 2,
            "task": {
                "correct_answer_id": 2,
                "answers": [
                    {
                        "text": "Умножение",
                        "hint": "Название говорит само за себя"
                    },
                    {
                        "text": "Вычитание",
                        "hint": "Перечитай главу"
                    },
                    {
                        "text": "Сложение",
                        "hint": "Все верно"
                    }
                ]
            }
        }
    }
    multi_valid_positive = {
        "type": "multichoice_test",
        "data": {
            "user_answers": [
                True,
                True
            ],
            "task": {
                "correct_answers": [
                    True,
                    True
                ]
            }
        }
    }
    code_valid_positive = {
        "type": "program",
        "data": {
            "user_signals": [
                {
                    "name": "Sum[0:3]",
                    "wave": "=...................=.=.=.....",
                    "data": [
                        "b0",
                        "b1",
                        "b1101",
                        "b1000"
                    ]
                },
                {
                    "name": "Cout",
                    "wave": "0.......................1.....",
                    "data": []
                },
                {
                    "name": "A[0:3]",
                    "wave": "=...................=.=.=.....",
                    "data": [
                        "b0",
                        "b1",
                        "b1010",
                        "b1101"
                    ]
                },
                {
                    "name": "B[0:3]",
                    "wave": "=.....................=.=.....",
                    "data": [
                        "b0",
                        "b11",
                        "b1010"
                    ]
                },
                {
                    "name": "Cin",
                    "wave": "0.......................1.....",
                    "data": []
                }
            ],
            "correct_signals": [
                {
                    "name": "Sum[0:3]",
                    "wave": "=...................=.=.=.....",
                    "data": [
                        "b0",
                        "b1",
                        "b1101",
                        "b1000"
                    ]
                },
                {
                    "name": "Cout",
                    "wave": "0.......................1.....",
                    "data": []
                },
                {
                    "name": "A[0:3]",
                    "wave": "=...................=.=.=.....",
                    "data": [
                        "b0",
                        "b1",
                        "b1010",
                        "b1101"
                    ]
                },
                {
                    "name": "B[0:3]",
                    "wave": "=.....................=.=.....",
                    "data": [
                        "b0",
                        "b11",
                        "b1010"
                    ]
                },
                {
                    "name": "Cin",
                    "wave": "0.......................1.....",
                    "data": []
                }
            ]
        }
    }
    code_valid_negative = {
        "type": "program",
        "data": {
            "user_signals": [
                {
                    "name": "Sum[0:3]",
                    "wave": "=...................=.=.=.....",
                    "data": [
                        "b01",
                        "b1",
                        "b1101",
                        "b1000"
                    ]
                },
                {
                    "name": "Cout1",
                    "wave": "0.......................1.....",
                    "data": []
                },
                {
                    "name": "A[0:3]1",
                    "wave": "=...................=.=.=.....",
                    "data": [
                        "b0",
                        "b1",
                        "b1010",
                        "b1101"
                    ]
                },
                {
                    "name": "B[0:3]",
                    "wave": "=.....................=.=.....",
                    "data": [
                        "b0",
                        "b11",
                        "b1010"
                    ]
                },
                {
                    "name": "Cin",
                    "wave": "0.......................1.....",
                    "data": []
                }
            ],
            "correct_signals": [
                {
                    "name": "Sum[0:3]",
                    "wave": "=...................=.=.=.....",
                    "data": [
                        "b0",
                        "b1",
                        "b1101",
                        "b1000"
                    ]
                },
                {
                    "name": "Cout",
                    "wave": "0.......................1.....",
                    "data": []
                },
                {
                    "name": "A[0:3]",
                    "wave": "=...................=.=.=.....",
                    "data": [
                        "b0",
                        "b1",
                        "b1010",
                        "b1101"
                    ]
                },
                {
                    "name": "B[0:3]",
                    "wave": "=.....................=.=.....",
                    "data": [
                        "b0",
                        "b11",
                        "b1010"
                    ]
                },
                {
                    "name": "Cin",
                    "wave": "0.......................1.....",
                    "data": []
                }
            ]
        }
    }


class Wavedrom:

    valid_request = {
        "data": [
            {
                "data": [
                    [
                        0,
                        "b0"
                    ],
                    [
                        100,
                        "b1"
                    ],
                    [
                        110,
                        "b1101"
                    ],
                    [
                        120,
                        "b1000"
                    ]
                ],
                "name": "Sum",
                "type": {
                    "name": "wire",
                    "width": 4
                }
            },
            {
                "data": [
                    [
                        0,
                        "0"
                    ],
                    [
                        120,
                        "1"
                    ]
                ],
                "name": "Cout",
                "type": {
                    "name": "wire",
                    "width": 1
                }
            },
            {
                "data": [
                    [
                        0,
                        "b0"
                    ],
                    [
                        100,
                        "b1"
                    ],
                    [
                        110,
                        "b1010"
                    ],
                    [
                        120,
                        "b1101"
                    ]
                ],
                "name": "A",
                "type": {
                    "name": "reg",
                    "width": 4
                }
            },
            {
                "data": [
                    [
                        0,
                        "b0"
                    ],
                    [
                        110,
                        "b11"
                    ],
                    [
                        120,
                        "b1010"
                    ]
                ],
                "name": "B",
                "type": {
                    "name": "reg",
                    "width": 4
                }
            },
            {
                "data": [
                    [
                        0,
                        "0"
                    ],
                    [
                        120,
                        "1"
                    ]
                ],
                "name": "Cin",
                "type": {
                    "name": "reg",
                    "width": 1
                }
            }
        ]
    }
    valid_response = {
        "status_str": "ok",
        "status_code": 200,
        "data": [
            {
                "name": "Sum[0:3]",
                "wave": "=...................=.=.=.....",
                "data": [
                    "b0",
                    "b1",
                    "b1101",
                    "b1000"
                ]
            },
            {
                "name": "Cout",
                "wave": "0.......................1.....",
                "data": []
            },
            {
                "name": "A[0:3]",
                "wave": "=...................=.=.=.....",
                "data": [
                    "b0",
                    "b1",
                    "b1010",
                    "b1101"
                ]
            },
            {
                "name": "B[0:3]",
                "wave": "=.....................=.=.....",
                "data": [
                    "b0",
                    "b11",
                    "b1010"
                ]
            },
            {
                "name": "Cin",
                "wave": "0.......................1.....",
                "data": []
            }
        ]
    }