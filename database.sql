/**
  This is the SQL script that will be used to initialize the database schema.
  We will evaluate you based on how well you design your database.
  1. How you design the tables.
  2. How you choose the data types and keys.
  3. How you name the fields.
  In this assignment we will use PostgreSQL as the database.
  */

/** This is test table. Remove this table and replace with your own tables. */
CREATE TABLE test (
	id serial PRIMARY KEY,
	name VARCHAR ( 50 ) UNIQUE NOT NULL
);

INSERT INTO test (name) VALUES ('test1');
INSERT INTO test (name) VALUES ('test2');


CREATE TABLE user_table (
  id SERIAL,
  phone_number varchar(16) NOT NULL,
  fullname varchar(60) NOT NULL,
  token varchar(255) NULL,
  password varchar(255) NOT NULL) ;

ALTER TABLE user_table
  ADD PRIMARY KEY (id),
  ADD CONSTRAINT phone_number UNIQUE (phone_number);

