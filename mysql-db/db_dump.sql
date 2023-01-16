/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */
;

/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */
;

/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */
;

/*!50503 SET NAMES utf8mb4 */
;

/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */
;

/*!40103 SET TIME_ZONE='+00:00' */
;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */
;

/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */
;

/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */
;

/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */
;

CREATE DATABASE
/*!32312 IF NOT EXISTS*/
`levels`
/*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */
/*!80016 DEFAULT ENCRYPTION='N' */
;

USE `levels`;

DROP TABLE IF EXISTS `LevelsBrief`;

/*!40101 SET @saved_cs_client     = @@character_set_client */
;

/*!50503 SET character_set_client = utf8mb4 */
;

CREATE TABLE `LevelsBrief` (
  `id` int NOT NULL AUTO_INCREMENT,
  `level_type` int NOT NULL,
  `seqnum` int NOT NULL,
  `cost` int NOT NULL,
  `is_active` tinyint(1) NOT NULL,
  `name` varchar(30) NOT NULL,
  `brief` varchar(280) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 2 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

/*!40101 SET character_set_client = @saved_cs_client */
;

LOCK TABLES `LevelsBrief` WRITE;

/*!40000 ALTER TABLE `LevelsBrief` DISABLE KEYS */
;

INSERT INTO
  `LevelsBrief`
VALUES
  (
    1,
    4,
    1,
    10,
    1,
    'Device lvl 1',
    'Device test lvl'
  ),
  (2, 1, 2, 0, 1, 'Text lvl 1', 'Text block'),
  (
    3,
    2,
    3,
    5,
    1,
    'Single lvl 1',
    'Singlechoice test'
  ),
  (
    4,
    3,
    4,
    7,
    1,
    'Multi lvl 1',
    'Multichoice test'
  );

/*!40000 ALTER TABLE `LevelsBrief` ENABLE KEYS */
;

UNLOCK TABLES;

DROP TABLE IF EXISTS `LevelsData`;

/*!40101 SET @saved_cs_client     = @@character_set_client */
;

/*!50503 SET character_set_client = utf8mb4 */
;

CREATE TABLE `LevelsData` (
  `id` int NOT NULL,
  `wide_description` text NOT NULL,
  `code` text NOT NULL,
  `question` text NOT NULL,
  `answer` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

/*!40101 SET character_set_client = @saved_cs_client */
;

LOCK TABLES `LevelsData` WRITE;

/*!40000 ALTER TABLE `LevelsData` DISABLE KEYS */
;

INSERT INTO
  `LevelsData`
VALUES
  (
    1,
    'This is wide desr of code',
    'Code example',
    'module adder_tb;  \n// Inputs  \nreg [3:0] A;  \nreg [3:0] B;  \nreg Cin;  \n// Outputs  \nwire [3:0] Sum;  \nwire Cout;  \n// Instantiate the Unit Under Test (UUT)  \nripple_adder_4bit uut (  \n.Sum(Sum),  \n.Cout(Cout),  \n.A(A),  \n.B(B),  \n.Cin(Cin)  \n);  \ninitial begin  \n// Initialize Inputs  \nA = 0;  \nB = 0;  \nCin = 0;  \n// Wait 100 ns for global reset to finish  \n#100;  \n// Add stimulus here  \nA=4\'b0001;B=4\'b0000;Cin=1\'b0;  \n#10 A=4\'b1010;B=4\'b0011;Cin=1\'b0;  \n#10 A=4\'b1101;B=4\'b1010;Cin=1\'b1;  \nend  \ninitial begin  \n$dumpfile(\"adder.vcd\");  \n$dumpvars;  \nend  \nendmodule\n',
    '[{\"data\":[\"b0\",\"b1\",\"b1101\",\"b1000\"],\"name\":\"Sum[0:3]\",\"wave\":\"=...................=.=.=.....\"},{\"data\":[],\"name\":\"Cout\",\"wave\":\"0.......................1.....\"},{\"data\":[\"b0\",\"b1\",\"b1010\",\"b1101\"],\"name\":\"A[0:3]\",\"wave\":\"=...................=.=.=.....\"},{\"data\":[\"b0\",\"b11\",\"b1010\"],\"name\":\"B[0:3]\",\"wave\":\"=.....................=.=.....\"},{\"data\":[],\"name\":\"Cin\",\"wave\":\"0.......................1.....\"}]'
  ),
  (
    2,
    'This is wide desr of text',
    'no code',
    'no question',
    'no answer'
  ),
  (
    3,
    'This is wide desr of single',
    'no code',
    '{\"caption\":\"SingleChoice test\",\"correct_answer_id\":3,\"answers\":[{\"text\":\"Умножение\",\"hint\":\"Название говорит само за себя\"},{\"text\":\"Вычитание\",\"hint\":\"Перечитай главу\"},{\"text\":\"Сложение\",\"hint\":\"Все верно\"}]}',
    '{\"user_answer_id\": 1}'
  ),
  (
    4,
    'This is wide desr of multi',
    'no code',
    '{\"caption\":\"MultiChoice test\",\"answers\": [\"Variant 1\", \"Variant 2\"]},\"correct_answers\": [true, true]}',
    '{\"user_answers\": [true, false]}'
  );

/*!40000 ALTER TABLE `LevelsData` ENABLE KEYS */
;

UNLOCK TABLES;

DROP TABLE IF EXISTS `SolutionEfforts`;

/*!40101 SET @saved_cs_client     = @@character_set_client */
;

/*!50503 SET character_set_client = utf8mb4 */
;

CREATE TABLE `SolutionEfforts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `level_id` int NOT NULL,
  `is_successful` tinyint(1) NOT NULL,
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 3 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

/*!40101 SET character_set_client = @saved_cs_client */
;

LOCK TABLES `SolutionEfforts` WRITE;

/*!40000 ALTER TABLE `SolutionEfforts` DISABLE KEYS */
;

INSERT INTO
  `SolutionEfforts`
VALUES
  (1, 1, 1, 1, '2020-01-10 14:53:01'),
  /* Trickster: 2 lvl is text block */
  (3, 3, 3, 1, '2022-01-10 14:53:01'),
  (4, 4, 4, 0, '2023-01-10 14:53:01');

/*!40000 ALTER TABLE `SolutionEfforts` ENABLE KEYS */
;

UNLOCK TABLES;

DROP TABLE IF EXISTS `Types`;

/*!40101 SET @saved_cs_client     = @@character_set_client */
;

/*!50503 SET character_set_client = utf8mb4 */
;

CREATE TABLE `Types` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` char(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 5 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

/*!40101 SET character_set_client = @saved_cs_client */
;

LOCK TABLES `Types` WRITE;

/*!40000 ALTER TABLE `Types` DISABLE KEYS */
;

INSERT INTO
  `Types`
VALUES
  (1, 'text'),
  (2, 'singlechoice_test'),
  (3, 'multichoice_test'),
  (4, 'program');

/*!40000 ALTER TABLE `Types` ENABLE KEYS */
;

UNLOCK TABLES;

DROP TABLE IF EXISTS `Users`;

/*!40101 SET @saved_cs_client     = @@character_set_client */
;

/*!50503 SET character_set_client = utf8mb4 */
;

CREATE TABLE `Users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nickname` varchar(30) NOT NULL,
  `is_admin` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 4 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

/*!40101 SET character_set_client = @saved_cs_client */
;

LOCK TABLES `Users` WRITE;

/*!40000 ALTER TABLE `Users` DISABLE KEYS */
;

INSERT INTO
  `Users`
VALUES
  (1, 'Deni', 1),
  (2, 'David', 1),
  (3, 'Mark', 0),
  (4, 'Johny', 0);

/*!40000 ALTER TABLE `Users` ENABLE KEYS */
;

UNLOCK TABLES;

/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */
;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */
;

/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */
;

/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */
;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */
;

/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */
;

/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */
;

/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */
;