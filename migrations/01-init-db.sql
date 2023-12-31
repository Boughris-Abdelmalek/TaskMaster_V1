-- CREATE TABLE
DROP TABLE IF EXISTS todos;
CREATE TABLE todos(
    id SERIAL NOT NULL PRIMARY KEY,
    title VARCHAR(50),
    description VARCHAR(100),
    completed BOOLEAN
);

-- CREATE TABLE FOR USERS
DROP TABLE IF EXISTS users;
CREATE TABLE users(
    id SERIAL NOT NULL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    user_name VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    user_type VARCHAR(20) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL
);