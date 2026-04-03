ALTER TABLE `push_subscriptions`
    ADD COLUMN `endpoint` TEXT NOT NULL AFTER `device_token`,
    ADD COLUMN `p256dh` TEXT NOT NULL AFTER `endpoint`,
    ADD COLUMN `auth` VARCHAR(255) NOT NULL DEFAULT '' AFTER `p256dh`;
