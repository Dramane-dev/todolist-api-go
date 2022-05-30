
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

CREATE TABLE tasks(
    taskId VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    status VARCHAR(20) DEFAULT "todo",
    created_at VARCHAR(255),
    projectId VARCHAR(255) NOT NULL,
    PRIMARY KEY(taskId),
    FOREIGN KEY(projectId) REFERENCES projects(projectId) ON DELETE CASCADE
);

CREATE TABLE attachments(
    attachmentId VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(255),
    path VARCHAR(255) NOT NULL,
    projectId VARCHAR(255) NOT NULL,
    PRIMARY KEY(attachmentId),
    FOREIGN KEY(projectId) REFERENCES projects(projectId) ON DELETE CASCADE
);

CREATE TABLE subscriptions(
    subscriptionId VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    amount VARCHAR(255) NOT NULL,
    userId VARCHAR(255) NOT NULL,
    PRIMARY KEY(subscriptionId),
    FOREIGN KEY(userId) REFERENCES users(userId) ON DELETE CASCADE
);