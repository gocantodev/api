DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS posts;

CREATE TABLE IF NOT EXISTS users (
    id INT NOT NULL,
    uuid varchar(32) unique NOT NULL,
    first_name varchar(250) NOT NULL,
    last_name varchar(250) NOT NULL,
    email varchar(250) unique NOT NULL,
    password varchar(100) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS posts (
    id INT NOT NULL,

    uuid varchar(32) unique NOT NULL,
    user_id INT NOT NULL,
    title varchar(250) NOT NULL,
    slug varchar(250) NOT NULL,
    body text NOT NULL,

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    PRIMARY KEY (id),
    CONSTRAINT user_fk
        FOREIGN KEY(user_id)
            REFERENCES users(id)
);
