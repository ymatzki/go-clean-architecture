CREATE DATABASE test;

\c test

CREATE TABLE IF NOT EXISTS Greeting
(
    word TEXT
);

INSERT INTO Greeting VALUES ('Hello');
