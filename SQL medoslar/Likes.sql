CREATE TABLE likes_post (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) NOT NULL,
    post_id INT REFERENCES posts(id) NOT NULL,
    liked_user_id INT REFERENCES users(id) NOT NULL,
    liked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE likes_comment (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) NOT NULL,
    post_id INT REFERENCES posts(id) NOT NULL,
    comment_id INT REFERENCES posts(id) NOT NULL,
    liked_user_id INT REFERENCES users(id) NOT NULL,
    liked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);