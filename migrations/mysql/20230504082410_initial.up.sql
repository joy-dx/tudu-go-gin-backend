--bun:split

CREATE TABLE users (
                       id BIGINT NOT NULL AUTO_INCREMENT,
                       username VARCHAR(255) NOT NULL,
                       password VARCHAR(255),
                       roles JSON,
                       status VARCHAR(255),
                       PRIMARY KEY (id)
) ENGINE=InnoDB;

--bun:split

CREATE TABLE todos (
                       id BIGINT NOT NULL AUTO_INCREMENT,
                       title VARCHAR(255) NOT NULL,
                       content TEXT,
                       complete BOOLEAN NOT NULL DEFAULT FALSE,
                       PRIMARY KEY (id)
) ENGINE=InnoDB;