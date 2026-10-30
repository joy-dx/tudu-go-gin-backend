--bun:split

CREATE TABLE users (
                       id INTEGER PRIMARY KEY,
                       username VARCHAR(255) NOT NULL,
                       password VARCHAR(255),
                       roles TEXT,
                       status TEXT
);

--bun:split

CREATE TABLE todos (
                       id INTEGER PRIMARY KEY,
                       title VARCHAR(255) NOT NULL,
                       content TEXT,
                       complete BOOLEAN NOT NULL DEFAULT FALSE
);