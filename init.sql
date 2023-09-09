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
	id INT AUTO_INCREMENT,
	created_by varchar(255) NOT NULL,
	created_at DATETIME,
	note varchar(255) NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (created_by) REFERENCES users(user_id)
);
CREATE INDEX idx_created_by on notes(created_by);
CREATE TABLE IF NOT EXISTS user_session (
	user_id INT, 
	session_id varchar(255),
	PRIMARY KEY (user_id),
	FOREIGN KEY (user_id) REFERENCES users(user_id)
);
CREATE INDEX idx_users_session_id on user_session(session_id);