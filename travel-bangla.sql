create DATABASE Travelbangla;
use Travelbangla;
CREATE TABLE users (
	u_id INT AUTO_INCREMENT,
    u_name VARCHAR(30) NOT NULL,
    u_email VARCHAR(150) NOT NULL,
    u_password VARCHAR(30) NOT NULL,
    u_phone VARCHAR(20) NOT NULL,
    u_social_links LONGTEXT,
    u_wishlist LONGTEXT,
    u_watch_history LONGTEXT,
    PRIMARY KEY(u_id)
);

CREATE TABLE category (
	ca_id INT AUTO_INCREMENT,
    ca_name VARCHAR(30) NOT NULL,
    ca_icon VARCHAR(255),
    PRIMARY KEY(ca_id)
);

CREATE TABLE package (
	pk_id INT AUTO_INCREMENT,
    pk_package_id INT,
    pk_title VARCHAR(255) NOT NULL,
    pk_description LONGTEXT,
    pk_price VARCHAR(100) DEFAULT "Free",
    pk_thumbnail VARCHAR(255),
    pk_video_url VARCHAR(512) CHARACTER SET 'ascii' COLLATE 'ascii_general_ci',
    pk_date_added DATE,
    PRIMARY KEY(pk_id),
    FOREIGN KEY(pk_package_id) REFERENCES category(ca_id)
);

CREATE TABLE enroll (
	en_id INT AUTO_INCREMENT,
    en_user_id INT NOT NULL,
    en_package_id INT NOT NULL,
    PRIMARY KEY(en_id),
    FOREIGN KEY(en_user_id) REFERENCES users(u_id),
    FOREIGN KEY(en_package_id) REFERENCES package(pk_id)
);