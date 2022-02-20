CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    id       uuid DEFAULT uuid_generate_v4(),
    name     varchar(64) UNIQUE,
    password varchar(64),
    PRIMARY KEY (id)
);

CREATE FUNCTION sign_in(IN p_name varchar(64), IN p_password varchar(64), OUT p_id varchar(128)) AS
$$
SELECT id
FROM users u
WHERE u.name = p_name AND u.password = p_password;
$$ LANGUAGE sql;

CREATE PROCEDURE sign_up(IN p_name varchar(64), IN p_password varchar(64)) AS
$$
INSERT INTO users(name, password)
VALUES (p_name, p_password);
$$ LANGUAGE sql;