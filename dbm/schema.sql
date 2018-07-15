DROP DATABASE IF EXISTS jsontest;
CREATE DATABASE IF NOT EXISTS jsontest;
USE jsontest;

CREATE TABLE Applicant (
    AID BIGINT(20) NOT NULL AUTO_INCREMENT,
    Name VARCHAR(100) NOT NULL DEFAULT '',
    Email VARCHAR(100) NOT NULL DEFAULT '',
    CellPhone VARCHAR(100) NOT NULL DEFAULT '',
    Address VARCHAR(100) NOT NULL DEFAULT '',
    PRIMARY KEY (AID)
);

CREATE TABLE JSONDoc (
    DocID BIGINT(20) NOT NULL AUTO_INCREMENT,
    Data JSON DEFAULT NULL,
    PRIMARY KEY (DocID)
);