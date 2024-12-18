CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY
);

CREATE TABLE segments (
    id SERIAL PRIMARY KEY,
    slug VARCHAR(255) UNIQUE NOT NULL,
    auto_add_percent SMALLINT DEFAULT 0
);

CREATE TABLE user_segments (
    user_id BIGINT NOT NULL,
    segment_id INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    ttl TIMESTAMP WITH TIME ZONE,
    PRIMARY KEY (user_id, segment_id, created_at),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (segment_id) REFERENCES segments(id) 
);


