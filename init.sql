CREATE DATABASE IF NOT EXISTS notes_application;
USE notes_application;
CREATE TABLE IF NOT EXISTS users (
	name varchar(255) NOT NULL,
	email varchar(255),
	password varchar(255) NOT NULL,
	session_id varchar(255) NOT NULL,
	PRIMARY KEY (email,password)
);
CREATE TABLE IF NOT EXISTS notes (
	note_id INT AUTO_INCREMENT,
	created_by varchar(255) NOT NULL,
	note varchar(255) NOT NULL,
	PRIMARY KEY (note_id),
	FOREIGN KEY (created_by) REFERENCES users(session_id)
);