
CREATE DATABASE golangtodos;

USE golangtodos;

CREATE TABLE users(
    userId VARCHAR(255) NOT NULL,
    lastname VARCHAR(100) NOT NULL,
    firstname VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    mailConfirmed BOOLEAN DEFAULT FALSE,
    PRIMARY KEY(userId)
);

CREATE TABLE projects(
    projectId VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    userId VARCHAR(255),
    PRIMARY KEY(projectId),
    FOREIGN KEY(userId) REFERENCES users(userId) ON DELETE CASCADE
);