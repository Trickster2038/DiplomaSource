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


class Synth:
    valid_request = {
        "user_id": 5,
        "level_id": 23,
        "device_src": "module half_adder(  \n    output S,C,  \n    input A,B  \n    );  \nxor(S,A,B);  \nand(C,A,B);  \nendmodule  \n \nmodule full_adder(  \n    output S,Cout,  \n    input A,B,Cin  \n    );  \nwire s1,c1,c2;  \nhalf_adder HA1(s1,c1,A,B);  \nhalf_adder HA2(S,c2,s1,Cin);  \nor OG1(Cout,c1,c2);  \n \nendmodule  \n \nmodule ripple_adder_4bit(  \n    output [3:0] Sum,  \n    output Cout,  \n    input [3:0] A,B,  \n    input Cin  \n    );  \nwire c1,c2,c3;  \nfull_adder FA1(Sum[0],c1,A[0],B[0],Cin),  \nFA2(Sum[1],c2,A[1],B[1],c1),  \nFA3(Sum[2],c3,A[2],B[2],c2),  \nFA4(Sum[3],Cout,A[3],B[3],c3);  \n \nendmodule\n",
        "tb_src": "module adder_tb;  \n// Inputs  \nreg [3:0] A;  \nreg [3:0] B;  \nreg Cin;  \n// Outputs  \nwire [3:0] Sum;  \nwire Cout;  \n// Instantiate the Unit Under Test (UUT)  \nripple_adder_4bit uut (  \n.Sum(Sum),  \n.Cout(Cout),  \n.A(A),  \n.B(B),  \n.Cin(Cin)  \n);  \ninitial begin  \n// Initialize Inputs  \nA = 0;  \nB = 0;  \nCin = 0;  \n// Wait 100 ns for global reset to finish  \n#100;  \n// Add stimulus here  \nA=4'b0001;B=4'b0000;Cin=1'b0;  \n#10 A=4'b1010;B=4'b0011;Cin=1'b0;  \n#10 A=4'b1101;B=4'b1010;Cin=1'b1;  \nend  \ninitial begin  \n$dumpfile(\"adder.vcd\");  \n$dumpvars;  \nend  \nendmodule\n"
    }
    valid_response = {
        "status_str": "ok",
        "status_code": 200,
        "message": "compiled successfully",
        "data": "$date\n\tThu Mar  2 15:03:05 2023\n$end\n$version\n\tIcarus Verilog\n$end\n$timescale\n\t1s\n$end\n$scope module adder_tb $end\n$var wire 4 ! Sum [3:0] $end\n$var wire 1 \" Cout $end\n$var reg 4 # A [3:0] $end\n$var reg 4 $ B [3:0] $end\n$var reg 1 % Cin $end\n$scope module uut $end\n$var wire 4 & A [3:0] $end\n$var wire 4 ' B [3:0] $end\n$var wire 1 % Cin $end\n$var wire 1 ( c3 $end\n$var wire 1 ) c2 $end\n$var wire 1 * c1 $end\n$var wire 4 + Sum [3:0] $end\n$var wire 1 \" Cout $end\n$scope module FA1 $end\n$var wire 1 , A $end\n$var wire 1 - B $end\n$var wire 1 % Cin $end\n$var wire 1 * Cout $end\n$var wire 1 . s1 $end\n$var wire 1 / c2 $end\n$var wire 1 0 c1 $end\n$var wire 1 1 S $end\n$scope module HA1 $end\n$var wire 1 , A $end\n$var wire 1 - B $end\n$var wire 1 0 C $end\n$var wire 1 . S $end\n$upscope $end\n$scope module HA2 $end\n$var wire 1 . A $end\n$var wire 1 % B $end\n$var wire 1 / C $end\n$var wire 1 1 S $end\n$upscope $end\n$upscope $end\n$scope module FA2 $end\n$var wire 1 2 A $end\n$var wire 1 3 B $end\n$var wire 1 * Cin $end\n$var wire 1 ) Cout $end\n$var wire 1 4 s1 $end\n$var wire 1 5 c2 $end\n$var wire 1 6 c1 $end\n$var wire 1 7 S $end\n$scope module HA1 $end\n$var wire 1 2 A $end\n$var wire 1 3 B $end\n$var wire 1 6 C $end\n$var wire 1 4 S $end\n$upscope $end\n$scope module HA2 $end\n$var wire 1 4 A $end\n$var wire 1 * B $end\n$var wire 1 5 C $end\n$var wire 1 7 S $end\n$upscope $end\n$upscope $end\n$scope module FA3 $end\n$var wire 1 8 A $end\n$var wire 1 9 B $end\n$var wire 1 ) Cin $end\n$var wire 1 ( Cout $end\n$var wire 1 : s1 $end\n$var wire 1 ; c2 $end\n$var wire 1 < c1 $end\n$var wire 1 = S $end\n$scope module HA1 $end\n$var wire 1 8 A $end\n$var wire 1 9 B $end\n$var wire 1 < C $end\n$var wire 1 : S $end\n$upscope $end\n$scope module HA2 $end\n$var wire 1 : A $end\n$var wire 1 ) B $end\n$var wire 1 ; C $end\n$var wire 1 = S $end\n$upscope $end\n$upscope $end\n$scope module FA4 $end\n$var wire 1 > A $end\n$var wire 1 ? B $end\n$var wire 1 ( Cin $end\n$var wire 1 \" Cout $end\n$var wire 1 @ s1 $end\n$var wire 1 A c2 $end\n$var wire 1 B c1 $end\n$var wire 1 C S $end\n$scope module HA1 $end\n$var wire 1 > A $end\n$var wire 1 ? B $end\n$var wire 1 B C $end\n$var wire 1 @ S $end\n$upscope $end\n$scope module HA2 $end\n$var wire 1 @ A $end\n$var wire 1 ( B $end\n$var wire 1 A C $end\n$var wire 1 C S $end\n$upscope $end\n$upscope $end\n$upscope $end\n$upscope $end\n$enddefinitions $end\n#0\n$dumpvars\n0C\n0B\n0A\n0@\n0?\n0>\n0=\n0<\n0;\n0:\n09\n08\n07\n06\n05\n04\n03\n02\n01\n00\n0/\n0.\n0-\n0,\nb0 +\n0*\n0)\n0(\nb0 '\nb0 &\n0%\nb0 $\nb0 #\n0\"\nb0 !\n$end\n#100\nb1 !\nb1 +\n11\n1.\n1,\nb1 #\nb1 &\n#110\n1=\n1)\nb1101 !\nb1101 +\n1C\n16\n1@\n1-\n13\n0,\n12\n1>\nb11 $\nb11 '\nb1010 #\nb1010 &\n#120\n1(\n1C\n1\"\n15\n1)\n0=\n1;\n1*\n0@\n1B\n14\n06\n1:\nb1000 !\nb1000 +\n01\n1/\n0-\n1?\n1,\n02\n18\n1%\nb1010 $\nb1010 '\nb1101 #\nb1101 &\n"
    }
    bad_device_request = {
        "user_id": 1,
        "level_id": 1,
        "device_src": "\\module half_adder(  \n    output S,C,  \n    input A,B  \n    );  \nxor(S,A,B);  \nand(C,A,B);  \nendmodule  \n \nmodule full_adder(  \n    output S,Cout,  \n    input A,B,Cin  \n    );  \nwire s1,c1,c2;  \nhalf_adder HA1(s1,c1,A,B);  \nhalf_adder HA2(S,c2,s1,Cin);  \nor OG1(Cout,c1,c2);  \n \nendmodule  \n \nmodule ripple_adder_4bit(  \n    output [3:0] Sum,  \n    output Cout,  \n    input [3:0] A,B,  \n    input Cin  \n    );  \nwire c1,c2,c3;  \nfull_adder FA1(Sum[0],c1,A[0],B[0],Cin),  \nFA2(Sum[1],c2,A[1],B[1],c1),  \nFA3(Sum[2],c3,A[2],B[2],c2),  \nFA4(Sum[3],Cout,A[3],B[3],c3);  \n \nendmodule\n",
        "tb_src": "module adder_tb;  \n// Inputs  \nreg [3:0] A;  \nreg [3:0] B;  \nreg Cin;  \n// Outputs  \nwire [3:0] Sum;  \nwire Cout;  \n// Instantiate the Unit Under Test (UUT)  \nripple_adder_4bit uut (  \n.Sum(Sum),  \n.Cout(Cout),  \n.A(A),  \n.B(B),  \n.Cin(Cin)  \n);  \ninitial begin  \n// Initialize Inputs  \nA = 0;  \nB = 0;  \nCin = 0;  \n// Wait 100 ns for global reset to finish  \n#100;  \n// Add stimulus here  \nA=4'b0001;B=4'b0000;Cin=1'b0;  \n#10 A=4'b1010;B=4'b0011;Cin=1'b0;  \n#10 A=4'b1101;B=4'b1010;Cin=1'b1;  \nend  \ninitial begin  \n$dumpfile(\"adder.vcd\");  \n$dumpvars;  \nend  \nendmodule\n"
    }
    bad_tb_request = {
        "user_id": 1,
        "level_id": 12,
        "device_src": "module half_adder(  \n    output S,C,  \n    input A,B  \n    );  \nxor(S,A,B);  \nand(C,A,B);  \nendmodule  \n \nmodule full_adder(  \n    output S,Cout,  \n    input A,B,Cin  \n    );  \nwire s1,c1,c2;  \nhalf_adder HA1(s1,c1,A,B);  \nhalf_adder HA2(S,c2,s1,Cin);  \nor OG1(Cout,c1,c2);  \n \nendmodule  \n \nmodule ripple_adder_4bit(  \n    output [3:0] Sum,  \n    output Cout,  \n    input [3:0] A,B,  \n    input Cin  \n    );  \nwire c1,c2,c3;  \nfull_adder FA1(Sum[0],c1,A[0],B[0],Cin),  \nFA2(Sum[1],c2,A[1],B[1],c1),  \nFA3(Sum[2],c3,A[2],B[2],c2),  \nFA4(Sum[3],Cout,A[3],B[3],c3);  \n \nendmodule\n",
        "tb_src": "module adder_tb;  \n// Inputs  \nreg [3:0] A;  \nreg [3:0] B;  \nreg Cin;  \n// Outputs  \nwire [3:0] Sum;  \nwire Cout;  \n// Instantiate the Unit Under Test (UUT)  \nripple_adder_4bit uut (  \n.Sum(Sum),  \n.Cout(Cout),  \n.A(A),  \n.B(B),  \n.Cin(Cin)  \n);  \ninitial begin  \n// Initialize Inputs  \nA = 0;  \nB = 0;  \nCin = 0;  \n// Wait 100 ns for global reset to finish  \n#100;  \n// Add stimulus here  \nA=4'b0001;B=4'b0000;Cin=1'b0;  \n#10 A=4'b1010;B=4'b0011;Cin=1'b0;  \n#10 A=4'b1101;B=4'b1010;Cin=1'b1;  \nend  \ninitial begin  \n$dumpile(\"adder.vcd\");  \n$dumpvars;  \nend  \nendmodule\n"
    }
    bad_tb_dumpvars_request = {
        "user_id": 1,
        "level_id": 1,
        "device_src": "module half_adder(  \n    output S,C,  \n    input A,B  \n    );  \nxor(S,A,B);  \nand(C,A,B);  \nendmodule  \n \nmodule full_adder(  \n    output S,Cout,  \n    input A,B,Cin  \n    );  \nwire s1,c1,c2;  \nhalf_adder HA1(s1,c1,A,B);  \nhalf_adder HA2(S,c2,s1,Cin);  \nor OG1(Cout,c1,c2);  \n \nendmodule  \n \nmodule ripple_adder_4bit(  \n    output [3:0] Sum,  \n    output Cout,  \n    input [3:0] A,B,  \n    input Cin  \n    );  \nwire c1,c2,c3;  \nfull_adder FA1(Sum[0],c1,A[0],B[0],Cin),  \nFA2(Sum[1],c2,A[1],B[1],c1),  \nFA3(Sum[2],c3,A[2],B[2],c2),  \nFA4(Sum[3],Cout,A[3],B[3],c3);  \n \nendmodule\n",
        "tb_src": "module adder_tb;  \n// Inputs  \nreg [3:0] A;  \nreg [3:0] B;  \nreg Cin;  \n// Outputs  \nwire [3:0] Sum;  \nwire Cout;  \n// Instantiate the Unit Under Test (UUT)  \nripple_adder_4bit uut (  \n.Sum(Sum),  \n.Cout(Cout),  \n.A(A),  \n.B(B),  \n.Cin(Cin)  \n);  \ninitial begin  \n// Initialize Inputs  \nA = 0;  \nB = 0;  \nCin = 0;  \n// Wait 100 ns for global reset to finish  \n#100;  \n// Add stimulus here  \nA=4'b0001;B=4'b0000;Cin=1'b0;  \n#10 A=4'b1010;B=4'b0011;Cin=1'b0;  \n#10 A=4'b1101;B=4'b1010;Cin=1'b1;  \nend  \ninitial begin  \n$dumpfile(\"adder.vcd\");   \nend  \nendmodule\n"
    }


class StatsGeneral:
    request_each_level_passed = {
        "stat_type": "each_level_passed"
    }
    response_each_level_passed = {
        "status_str": "ok",
        "status_code": 200,
        "message": "",
        "data": [
            {
                "id": 1,
                "name": "Device lvl 1",
                "seqnum": 1,
                "solutions": 1
            },
            {
                "id": 2,
                "name": "Text lvl 1",
                "seqnum": 2,
                "solutions": 0
            },
            {
                "id": 3,
                "name": "Single lvl 1",
                "seqnum": 3,
                "solutions": 1
            },
            {
                "id": 4,
                "name": "Multi lvl 1",
                "seqnum": 4,
                "solutions": 0
            }
        ]
    }
    request_each_avg_efforts = {
        "stat_type": "each_level_avg_efforts"
    }
    response_each_avg_efforts = {
        "status_str": "ok",
        "status_code": 200,
        "message": "",
        "data": [
            {
                "id": 1,
                "name": "Device lvl 1",
                "seqnum": 1,
                "avg_efforts": 1
            },
            {
                "id": 2,
                "name": "Text lvl 1",
                "seqnum": 2,
                "avg_efforts": 0
            },
            {
                "id": 3,
                "name": "Single lvl 1",
                "seqnum": 3,
                "avg_efforts": 1
            },
            {
                "id": 4,
                "name": "Multi lvl 1",
                "seqnum": 4,
                "avg_efforts": 0
            }
        ]
    }
    request_activity_by_months = {
        "stat_type": "activity_by_month"
    }
    response_activity_by_months = {
        "status_str": "ok",
        "status_code": 200,
        "message": "",
        "data": [
            {
                "year": 2020,
                "month": 1,
                "solutions_count": 1,
                "efforts_count": 1
            },
            {
                "year": 2022,
                "month": 1,
                "solutions_count": 1,
                "efforts_count": 1
            },
            {
                "year": 2022,
                "month": 2,
                "solutions_count": 0,
                "efforts_count": 1
            },
            {
                "year": 2023,
                "month": 1,
                "solutions_count": 0,
                "efforts_count": 1
            }
        ]
    }
    request_top_active_users = {
        "stat_type": "top_month_users"
    }
    response_top_active_users = {
        "status_str": "ok",
        "status_code": 200,
        "message": "",
        "data": [
            {
                "id": 4,
                "nickname": "Johny",
                "solutions_count": 0,
                "efforts_count": 1
            },
            {
                "id": 1,
                "nickname": "Deni",
                "solutions_count": 0,
                "efforts_count": 0
            },
            {
                "id": 2,
                "nickname": "David",
                "solutions_count": 0,
                "efforts_count": 0
            },
            {
                "id": 3,
                "nickname": "Mark",
                "solutions_count": 0,
                "efforts_count": 0
            }
        ]
    }


class StatsPersonal:
    request_general_progress = {
        "user_id": 3,
        "stat_type": "general_progress"
    }
    response_general_progress = {
        "status_str": "ok",
        "status_code": 200,
        "message": "",
        "data": {
            "actual_points": 5,
            "total_points": 22,
            "actual_levels": 1,
            "total_levels": 4,
            "pass_status": "not_passed"
        }
    }
    request_each_level_status = {
        "user_id": 1,
        "stat_type": "each_level_status"
    }
    response_each_level_status = {
        "status_str": "ok",
        "status_code": 200,
        "message": "",
        "data": [
            {
                "id": 1,
                "seqnum": 1,
                "cost": 10,
                "level_name": "Device lvl 1",
                "brief": "Device test lvl",
                "is_succesful": True,
                "level_type": "program"
            },
            {
                "id": 2,
                "seqnum": 2,
                "cost": 0,
                "level_name": "Text lvl 1",
                "brief": "Text block",
                "is_succesful": False,
                "level_type": "text"
            },
            {
                "id": 3,
                "seqnum": 3,
                "cost": 5,
                "level_name": "Single lvl 1",
                "brief": "Singlechoice test",
                "is_succesful": False,
                "level_type": "singlechoice_test"
            },
            {
                "id": 4,
                "seqnum": 4,
                "cost": 7,
                "level_name": "Multi lvl 1",
                "brief": "Multichoice test",
                "is_succesful": False,
                "level_type": "multichoice_test"
            }
        ]
    }
    request_avg_efforts = {
        "user_id": 3,
        "stat_type": "avg_efforts"
    }
    response_avg_efforts = {
        "status_str": "ok",
        "status_code": 200,
        "message": "",
        "data": {
            "avg_efforts": 1
        }
    }
    request_monthly_activity = {
        "user_id": 3,
        "stat_type": "monthly_activity"
    }
    response_monthly_activity = {
        "status_str": "ok",
        "status_code": 200,
        "message": "",
        "data": {
            "accepted_tasks": 0,
            "earned_points": 0,
            "efforts": 0
        }
    }
    request_activity_borders = {
        "user_id": 3,
        "stat_type": "activity_borders"
    }
    response_activity_borders = {
        "status_str": "ok",
        "status_code": 200,
        "message": "",
        "data": {
            "first_efforts": "2022-01-10 14:53:01",
            "first_solved": "2022-01-10 14:53:01",
            "last_efforts": "2022-01-10 14:53:01",
            "last_solved": "2022-01-10 14:53:01"
        }
    }


class CRUD:
    request_read_user = {
        "metainfo": {
            "action": "read",
            "obj_type": "user"
        },
        "data": {
            "id": 1
        }
    }
    response_read_user = {
        "status_str": "ok",
        "status_code": 200,
        "message": "",
        "data": {
            "id": 1,
            "nickname": "Deni",
            "is_admin": True
        }
    }
    request_read_level_brief = {
        "metainfo": {
            "action": "read",
            "obj_type": "levels_brief"
        },
        "data": {
            "id": 1
        }
    }
    response_read_level_brief = {
        "status_str": "ok",
        "status_code": 200,
        "message": "",
        "data": {
            "id": 1,
            "level_type": 4,
            "seqnum": 1,
            "cost": 10,
            "is_active": True,
            "name": "Device lvl 1",
            "brief": "Device test lvl",
            "level_type_name": "program"
        }
    }
    request_read_level_data = {
        "metainfo": {
            "action": "read",
            "obj_type": "levels_data"
        },
        "data": {
            "id": 1
        }
    }
    response_read_level_data = {
        "status_str": "ok",
        "status_code": 200,
        "message": "",
        "data": {
            "id": 1,
            "wide_description": "This is wide descr of code",
            "code": "Code example",
            "question": "module adder_tb;  \n// Inputs  \nreg [3:0] A;  \nreg [3:0] B;  \nreg Cin;  \n// Outputs  \nwire [3:0] Sum;  \nwire Cout;  \n// Instantiate the Unit Under Test (UUT)  \nripple_adder_4bit uut (  \n.Sum(Sum),  \n.Cout(Cout),  \n.A(A),  \n.B(B),  \n.Cin(Cin)  \n);  \ninitial begin  \n// Initialize Inputs  \nA = 0;  \nB = 0;  \nCin = 0;  \n// Wait 100 ns for global reset to finish  \n#100;  \n// Add stimulus here  \nA=4'b0001;B=4'b0000;Cin=1'b0;  \n#10 A=4'b1010;B=4'b0011;Cin=1'b0;  \n#10 A=4'b1101;B=4'b1010;Cin=1'b1;  \nend  \ninitial begin  \n$dumpfile(\"adder.vcd\");  \n$dumpvars;  \nend  \nendmodule\n",
            "answer": "[{\"data\":[\"b0\",\"b1\",\"b1101\",\"b1000\"],\"name\":\"Sum[0:3]\",\"wave\":\"=...................=.=.=.....\"},{\"data\":[],\"name\":\"Cout\",\"wave\":\"0.......................1.....\"},{\"data\":[\"b0\",\"b1\",\"b1010\",\"b1101\"],\"name\":\"A[0:3]\",\"wave\":\"=...................=.=.=.....\"},{\"data\":[\"b0\",\"b11\",\"b1010\"],\"name\":\"B[0:3]\",\"wave\":\"=.....................=.=.....\"},{\"data\":[],\"name\":\"Cin\",\"wave\":\"0.......................1.....\"}]"
        }
    }
    request_read_all_level_brief = {
        "metainfo": {
            "action": "read_all",
            "obj_type": "levels_brief"
        }
    }
    response_read_all_level_brief = {
        "status_str": "ok",
        "status_code": 200,
        "message": "",
        "data": {
            "levels": [
                {
                    "id": 1,
                    "level_type": 4,
                    "seqnum": 1,
                    "cost": 10,
                    "is_active": True,
                    "name": "Device lvl 1",
                    "brief": "Device test lvl",
                    "level_type_name": "program"
                },
                {
                    "id": 2,
                    "level_type": 1,
                    "seqnum": 2,
                    "cost": 0,
                    "is_active": True,
                    "name": "Text lvl 1",
                    "brief": "Text block",
                    "level_type_name": "text"
                },
                {
                    "id": 3,
                    "level_type": 2,
                    "seqnum": 3,
                    "cost": 5,
                    "is_active": True,
                    "name": "Single lvl 1",
                    "brief": "Singlechoice test",
                    "level_type_name": "singlechoice_test"
                },
                {
                    "id": 4,
                    "level_type": 3,
                    "seqnum": 4,
                    "cost": 7,
                    "is_active": True,
                    "name": "Multi lvl 1",
                    "brief": "Multichoice test",
                    "level_type_name": "multichoice_test"
                }
            ]
        }
    }
    request_check_succesful = {
        "metainfo": {
            "action": "check_successful",
            "obj_type": "solution_effort"
        },
        "data": {
            "user_id": 1,
            "level_id": 1
        }
    }
    response_check_succesful = {
        "status_str": "ok",
        "status_code": 200,
        "message": "",
        "data": {
            "is_successful": True
        }
    }


class Gateway:
    request_stats_each_level_passed = {
        "scope": "general",
        "stat_type": "each_level_passed"
    }
