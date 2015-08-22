CREATE TABLE IF NOT EXISTS account (
  email VARCHAR(32) NOT NULL PRIMARY KEY,
  username VARCHAR(32) NOT NULL UNIQUE,
  password CHAR(32) NOT NULL
);

INSERT INTO account VALUES('Christiano.Ronaldo@gmail.com', 'Christiano' , 'Ronaldo');