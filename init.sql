CREATE DATABASE IF NOT EXISTS notes_application;
USE notes_application;
CREATE TABLE IF NOT EXISTS users (
	user_id INT AUTO_INCREMENT,
	name varchar(255) NOT NULL,
	email varchar(255),
	password varchar(255) NOT NULL,
	PRIMARY KEY (user_id)
);
CREATE INDEX idx_users_email_password on users(email,password);
CREATE TABLE IF NOT EXISTS notes (
	note_id INT AUTO_INCREMENT,
	created_by varchar(255) NOT NULL,
	note varchar(255) NOT NULL,
	PRIMARY KEY (note_id),
	FOREIGN KEY (created_by) REFERENCES users(user_id)
);
CREATE TABLE IF NOT EXISTS user_session (
	user_id INT, 
	session_id varchar(255),
	PRIMARY KEY (session_id),
	FOREIGN KEY (user_id) REFERENCES users(user_id)
);