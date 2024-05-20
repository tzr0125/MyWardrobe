CREATE TABLE IF NOT EXISTS engine_image (
    image_id INT AUTO_INCREMENT PRIMARY KEY,
    file_path VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS engine_clothes (
    clothes_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    image_id INT,
    FOREIGN KEY (image_id) REFERENCES engine_image(image_id)
);

CREATE TABLE IF NOT EXISTS engine_tag (
    tag_id INT AUTO_INCREMENT PRIMARY KEY,
    tag_name VARCHAR(255) NOT NULL,
    rank INT
);

CREATE TABLE IF NOT EXISTS engine_class (
    class_id INT AUTO_INCREMENT PRIMARY KEY,
    class_name VARCHAR(255) NOT NULL,
    rank INT
);

CREATE TABLE IF NOT EXISTS engine_description (
    description_id INT AUTO_INCREMENT PRIMARY KEY,
    clothes_id INT,
    description TEXT,
    FOREIGN KEY (clothes_id) REFERENCES engine_clothes(clothes_id)
);