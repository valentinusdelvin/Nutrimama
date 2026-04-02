DROP INDEX `idx_notifications_user_id` ON `notifications`;
DROP TABLE IF EXISTS `notifications`;
DROP TABLE IF EXISTS `push_subscriptions`;
ALTER TABLE `users` DROP COLUMN `timezone`;
