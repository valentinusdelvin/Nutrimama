ALTER TABLE `users` ADD COLUMN `timezone` VARCHAR(50) DEFAULT 'UTC';

CREATE TABLE `push_subscriptions` (
    `subscription_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id` bigint UNSIGNED NOT NULL,
    `device_token` VARCHAR(500) NOT NULL UNIQUE,
    `platform` VARCHAR(50),
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT `fk_push_user` FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`) ON DELETE CASCADE
);

CREATE TABLE `notifications` (
    `notification_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id` bigint UNSIGNED NOT NULL,
    `title` VARCHAR(255) NOT NULL,
    `message` TEXT NOT NULL,
    `is_read` BOOLEAN DEFAULT FALSE,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT `fk_notification_user` FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`) ON DELETE CASCADE
);

CREATE INDEX `idx_notifications_user_id` ON `notifications`(`user_id`, `is_read`);
