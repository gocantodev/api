-- Create the users table
CREATE TABLE IF NOT EXISTS users (
    id INT NOT NULL,

    uuid varchar(32) unique NOT NULL,
    first_name varchar(250) NOT NULL,
    last_name varchar(250) NOT NULL,
    email varchar(250) unique NOT NULL,
    passowrd varchar(100) NOT NULL,

    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    PRIMARY KEY (id)
);
