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
