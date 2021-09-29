-- DROP DATABASE IF EXISTS todo;
CREATE DATABASE IF NOT EXISTS todo;
use todo;

CREATE TABLE IF NOT EXISTS users
(
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    uuid VARCHAR(36) NOT NULL UNIQUE, 
    name VARCHAR(225) NOT NULL,
    email VARCHAR(225) NOT NULL UNIQUE, 
    password VARCHAR(225) NOT NULL, 
    create_at DATE NOT NULL,
    update_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS sessions
(
    session_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    uuid VARCHAR(36) NOT NULL UNIQUE, 
    email VARCHAR(255),
    create_at DATE,
    FOREIGN KEY(user_id) REFERENCES users(user_id)
);


CREATE TABLE IF NOT EXISTS todos
(
    todo_id INT AUTO_INCREMENT PRIMARY KEY, 
    user_id INT NOT NULL, 
    content VARCHAR(255) NOT NULL, 
    create_at DATE NOT NULL, 
    update_at TIMESTAMP NOT NULL,
    deadline DATETIME NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(user_id)
);
