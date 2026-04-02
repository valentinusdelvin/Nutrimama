CREATE TABLE user_bookmarks (
    bookmark_id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    edu_tools_id BIGINT UNSIGNED NOT NULL REFERENCES educational_tools(edu_tools_id) ON DELETE CASCADE,
    UNIQUE(user_id, edu_tools_id)
);