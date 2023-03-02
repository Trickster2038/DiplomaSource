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


class Parser:

    valid_request = {
        "user_id": "1",
        "level_id": "1",
        "data": "$date\n\tSun Jan 15 14:41:31 2023\n$end\n$version\n\tIcarus Verilog\n$end\n$timescale\n\t1s\n$end\n$scope module adder_tb $end\n$var wire 4 ! Sum [3:0] $end\n$var wire 1 \" Cout $end\n$var reg 4 # A [3:0] $end\n$var reg 4 $ B [3:0] $end\n$var reg 1 % Cin $end\n$scope module uut $end\n$var wire 4 & A [3:0] $end\n$var wire 4 ' B [3:0] $end\n$var wire 1 % Cin $end\n$var wire 1 ( c3 $end\n$var wire 1 ) c2 $end\n$var wire 1 * c1 $end\n$var wire 4 + Sum [3:0] $end\n$var wire 1 \" Cout $end\n$scope module FA1 $end\n$var wire 1 , A $end\n$var wire 1 - B $end\n$var wire 1 % Cin $end\n$var wire 1 * Cout $end\n$var wire 1 . s1 $end\n$var wire 1 / c2 $end\n$var wire 1 0 c1 $end\n$var wire 1 1 S $end\n$scope module HA1 $end\n$var wire 1 , A $end\n$var wire 1 - B $end\n$var wire 1 0 C $end\n$var wire 1 . S $end\n$upscope $end\n$scope module HA2 $end\n$var wire 1 . A $end\n$var wire 1 % B $end\n$var wire 1 / C $end\n$var wire 1 1 S $end\n$upscope $end\n$upscope $end\n$scope module FA2 $end\n$var wire 1 2 A $end\n$var wire 1 3 B $end\n$var wire 1 * Cin $end\n$var wire 1 ) Cout $end\n$var wire 1 4 s1 $end\n$var wire 1 5 c2 $end\n$var wire 1 6 c1 $end\n$var wire 1 7 S $end\n$scope module HA1 $end\n$var wire 1 2 A $end\n$var wire 1 3 B $end\n$var wire 1 6 C $end\n$var wire 1 4 S $end\n$upscope $end\n$scope module HA2 $end\n$var wire 1 4 A $end\n$var wire 1 * B $end\n$var wire 1 5 C $end\n$var wire 1 7 S $end\n$upscope $end\n$upscope $end\n$scope module FA3 $end\n$var wire 1 8 A $end\n$var wire 1 9 B $end\n$var wire 1 ) Cin $end\n$var wire 1 ( Cout $end\n$var wire 1 : s1 $end\n$var wire 1 ; c2 $end\n$var wire 1 < c1 $end\n$var wire 1 = S $end\n$scope module HA1 $end\n$var wire 1 8 A $end\n$var wire 1 9 B $end\n$var wire 1 < C $end\n$var wire 1 : S $end\n$upscope $end\n$scope module HA2 $end\n$var wire 1 : A $end\n$var wire 1 ) B $end\n$var wire 1 ; C $end\n$var wire 1 = S $end\n$upscope $end\n$upscope $end\n$scope module FA4 $end\n$var wire 1 > A $end\n$var wire 1 ? B $end\n$var wire 1 ( Cin $end\n$var wire 1 \" Cout $end\n$var wire 1 @ s1 $end\n$var wire 1 A c2 $end\n$var wire 1 B c1 $end\n$var wire 1 C S $end\n$scope module HA1 $end\n$var wire 1 > A $end\n$var wire 1 ? B $end\n$var wire 1 B C $end\n$var wire 1 @ S $end\n$upscope $end\n$scope module HA2 $end\n$var wire 1 @ A $end\n$var wire 1 ( B $end\n$var wire 1 A C $end\n$var wire 1 C S $end\n$upscope $end\n$upscope $end\n$upscope $end\n$upscope $end\n$enddefinitions $end\n#0\n$dumpvars\n0C\n0B\n0A\n0@\n0?\n0>\n0=\n0<\n0;\n0:\n09\n08\n07\n06\n05\n04\n03\n02\n01\n00\n0/\n0.\n0-\n0,\nb0 +\n0*\n0)\n0(\nb0 '\nb0 &\n0%\nb0 $\nb0 #\n0\"\nb0 !\n$end\n#100\nb1 !\nb1 +\n11\n1.\n1,\nb1 #\nb1 &\n#110\n1=\n1)\nb1101 !\nb1101 +\n1C\n16\n1@\n1-\n13\n0,\n12\n1>\nb11 $\nb11 '\nb1010 #\nb1010 &\n#120\n1(\n1C\n1\"\n15\n1)\n0=\n1;\n1*\n0@\n1B\n14\n06\n1:\nb1000 !\nb1000 +\n01\n1/\n0-\n1?\n1,\n02\n18\n1%\nb1010 $\nb1010 '\nb1101 #\nb1101 &\n"
    }
    valid_response = {
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
        ],
        "status_code": 200,
        "status_str": "ok"
    }
