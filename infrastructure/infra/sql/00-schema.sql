-- Posts table for "Sandhedens stemme"
CREATE TABLE posts (

    id INT AUTO_INCREMENT PRIMARY KEY,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

);
