ALTER TABLE `push_subscriptions`
    ADD COLUMN `endpoint` TEXT NOT NULL DEFAULT '' AFTER `device_token`,
    ADD COLUMN `p256dh` TEXT NOT NULL DEFAULT '' AFTER `endpoint`,
    ADD COLUMN `auth` VARCHAR(255) NOT NULL DEFAULT '' AFTER `p256dh`;
