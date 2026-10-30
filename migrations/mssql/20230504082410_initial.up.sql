--bun:split

CREATE TABLE users (
                       id BIGINT IDENTITY(1, 1) NOT NULL,
                       username NVARCHAR(255) NOT NULL,
                       password NVARCHAR(255),
                       roles NVARCHAR(MAX),
                       status NVARCHAR(255),
                       CONSTRAINT pk_users PRIMARY KEY (id)
);

--bun:split

CREATE TABLE todos (
                       id BIGINT IDENTITY(1, 1) NOT NULL,
                       title NVARCHAR(255) NOT NULL,
                       content NVARCHAR(MAX),
                       complete BIT NOT NULL
                           CONSTRAINT df_todos_complete DEFAULT 0,
                       CONSTRAINT pk_todos PRIMARY KEY (id)
);