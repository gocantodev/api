id SERIAL PRIMARY KEY,

    uuid varchar(32) unique NOT NULL,
    first_name varchar(250) NOT NULL,
    last_name varchar(250) NOT NULL,
    email varchar(250) unique NOT NULL,
    password varchar(100) NOT NULL,
    verified_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP

INSERT INTO users
    (id, uuid, first_name, last_name, email)
VALUES
    (nextval('books_sequence'), 'The Hobbit', 'Tolkien');
