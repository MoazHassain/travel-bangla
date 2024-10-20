create DATABASE Travelbangla;
use Travelbangla;
CREATE TABLE users (
	u_id INT AUTO_INCREMENT,
    u_name VARCHAR(30) NOT NULL,
    u_email VARCHAR(150) NOT NULL,pk_description
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
    pk_location VARCHAR(255) NOT NULL,
    pk_days INT NOT NULL,
    pk_persons INT NOT NULL,
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

ALTER TABLE package
ADD COLUMN pk_location VARCHAR(255) NOT NULL,
ADD COLUMN pk_days INT NOT NULL,
ADD COLUMN pk_persons INT NOT NULL;

INSERT INTO package (pk_package_id, pk_title, pk_description, pk_location, pk_days, pk_persons, pk_price, pk_thumbnail, pk_video_url, pk_date_added)
VALUES
(1, 'Cox\'s Bazar Tour', 'Enjoy a wonderful trip to the world\'s longest sea beach with luxury hotel stays.', 'Cox\'s Bazar', 3, 2, '5000 BDT', 'coxs_bazar_thumbnail.jpg', 'https://example.com/video1', CURDATE()),

(2, 'Rangamati Hill Adventure', 'Explore the serene beauty of Rangamati, boating on Kaptai Lake, and more.', 'Rangamati', 2, 4, '4000 BDT', 'rangamati_thumbnail.jpg', 'https://example.com/video2', CURDATE()),

(3, 'Sylhet Tea Gardens', 'Visit Sylhet\'s famous tea gardens, enjoy the lush green scenery and cool climate.', 'Sylhet', 2, 3, '4500 BDT', 'sylhet_thumbnail.jpg', 'https://example.com/video3', CURDATE()),

(4, 'Sundarbans Adventure', 'Discover the world\'s largest mangrove forest and experience wildlife safaris.', 'Sundarbans', 3, 6, '7000 BDT', 'sundarbans_thumbnail.jpg', 'https://example.com/video4', CURDATE()),

(5, 'Bandarban Hill Trek', 'Trek through the stunning hills of Bandarban, visit Nilgiri and Boga Lake.', 'Bandarban', 4, 4, '6500 BDT', 'bandarban_thumbnail.jpg', 'https://example.com/video5', CURDATE()),

(6, 'St. Martin\'s Island Getaway', 'Relax on the beautiful coral island of St. Martin\'s with white sandy beaches.', 'St. Martin\'s Island', 2, 2, '8000 BDT', 'st_martins_thumbnail.jpg', 'https://example.com/video6', CURDATE()),

(7, 'Kuakata Beach Sunrise & Sunset', 'Experience the unique beauty of both sunrise and sunset from the same beach.', 'Kuakata', 2, 3, '5000 BDT', 'kuakata_thumbnail.jpg', 'https://example.com/video7', CURDATE());
INSERT INTO category (ca_id, ca_name)
VALUES
(1, 'Adventure Tours'),
(2, 'Hill Trekking'),
(3, 'Beach Getaways'),
(4, 'Cultural Tours'),
(5, 'Tea Garden Tours'),
(6, 'Heritage Sites'),
(7, 'Archaeological Tours');