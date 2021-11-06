CREATE DATABASE gin_db;

CREATE TABLE IF NOT EXISTS gin_db.sample_tables(ID int, name varchar(256), created_at datetime default current_timestamp);

INSERT INTO gin_db.sample_tables VALUES(1, "user name", "2017-10-06 07:25:42");
