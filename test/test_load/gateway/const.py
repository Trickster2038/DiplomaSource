class CRUD:
    request_read_levels_data = {
        "metainfo": {
            "obj_type": "levels_data",
            "action": "read"
        },
        "data": {
            "id": 3
        }
    }


class Stats:
    request_general_each_level_passed = {
        "scope": "general",
        "stat_type": "each_level_passed"
    }
    request_personal_each_level_passed = {
        "scope": "personal",
        "stat_type": "each_level_status",
        "user_id": 3
    }
    request_general_solutions_dist = {
        "scope": "general",
        "stat_type": "solutions_distribution"
    }


class Analyzer:
    request_check_program_ok = {
        "user_id": 1,
        "level_id": 1,
        "answer": "module half_adder(  \n    output S,C,  \n    input A,B  \n    );  \nxor(S,A,B);  \nand(C,A,B);  \nendmodule  \n \nmodule full_adder(  \n    output S,Cout,  \n    input A,B,Cin  \n    );  \nwire s1,c1,c2;  \nhalf_adder HA1(s1,c1,A,B);  \nhalf_adder HA2(S,c2,s1,Cin);  \nor OG1(Cout,c1,c2);  \n \nendmodule  \n \nmodule ripple_adder_4bit(  \n    output [3:0] Sum,  \n    output Cout,  \n    input [3:0] A,B,  \n    input Cin  \n    );  \nwire c1,c2,c3;  \nfull_adder FA1(Sum[0],c1,A[0],B[0],Cin),  \nFA2(Sum[1],c2,A[1],B[1],c1),  \nFA3(Sum[2],c3,A[2],B[2],c2),  \nFA4(Sum[3],Cout,A[3],B[3],c3);  \n \nendmodule\n"
    }
