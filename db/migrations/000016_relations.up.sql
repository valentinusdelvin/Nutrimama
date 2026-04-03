
ALTER TABLE `children`
    MODIFY `pregnancy_id` bigint UNSIGNED NULL,
    MODIFY `mother_id` bigint UNSIGNED NULL;

ALTER TABLE `mothers`
    MODIFY `user_id` bigint UNSIGNED NOT NULL;

ALTER TABLE `pregnancies`
    MODIFY `mother_id` bigint UNSIGNED NOT NULL;

ALTER TABLE `consultation_sessions`
    MODIFY `mother_id` bigint UNSIGNED NOT NULL,
    MODIFY `consultant_id` bigint UNSIGNED NOT NULL;

ALTER TABLE `messages`
    MODIFY `consultant_id` bigint UNSIGNED NULL,
    MODIFY `mother_id` bigint UNSIGNED NULL;

ALTER TABLE `user_tracking_logs`
    MODIFY `tracking_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    MODIFY `mother_id` bigint UNSIGNED NULL,
    MODIFY `pregnancy_id` bigint UNSIGNED NULL,
    MODIFY `child_id` bigint UNSIGNED NULL;

ALTER TABLE `nutrition_tracking`
    MODIFY `track_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    MODIFY `mother_id` bigint UNSIGNED NULL,
    MODIFY `child_id` bigint UNSIGNED NULL;

ALTER TABLE `meal_plans`
    MODIFY `meal_plan_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    MODIFY `mother_id` bigint UNSIGNED NOT NULL;

ALTER TABLE `food_logs`
    MODIFY `food_log_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    MODIFY `meal_plan_id` bigint UNSIGNED NOT NULL,
    MODIFY `food_id` bigint UNSIGNED NOT NULL;

ALTER TABLE `foods`
    MODIFY `food_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT;

ALTER TABLE `user_bookmarks`
    MODIFY `bookmark_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    MODIFY `user_id` bigint UNSIGNED NOT NULL,
    MODIFY `edu_tools_id` bigint UNSIGNED NOT NULL;

ALTER TABLE `questions`
    MODIFY `question_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT;

ALTER TABLE `question_options`
    MODIFY `option_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    MODIFY `question_id` bigint UNSIGNED NULL;



ALTER TABLE `children`
ADD CONSTRAINT `fk_children_pregnancy` 
    FOREIGN KEY (`pregnancy_id`) REFERENCES `pregnancies`(`pregnancy_id`) 
    ON DELETE SET NULL ON UPDATE CASCADE,
ADD CONSTRAINT `fk_children_mother` 
    FOREIGN KEY (`mother_id`) REFERENCES `mothers`(`mother_id`) 
    ON DELETE SET NULL ON UPDATE CASCADE;

ALTER TABLE `mothers`
ADD CONSTRAINT `fk_mothers_user` 
    FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`) 
    ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE `pregnancies`
ADD CONSTRAINT `fk_pregnancies_mother` 
    FOREIGN KEY (`mother_id`) REFERENCES `mothers`(`mother_id`) 
    ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE `consultation_sessions`
ADD CONSTRAINT `fk_consultation_mother` 
    FOREIGN KEY (`mother_id`) REFERENCES `mothers`(`mother_id`) 
    ON DELETE CASCADE ON UPDATE CASCADE,
ADD CONSTRAINT `fk_consultation_consultant` 
    FOREIGN KEY (`consultant_id`) REFERENCES `consultants`(`consultant_id`) 
    ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE `messages`
ADD CONSTRAINT `fk_messages_consultant` 
    FOREIGN KEY (`consultant_id`) REFERENCES `consultants`(`consultant_id`) 
    ON DELETE SET NULL ON UPDATE CASCADE,
ADD CONSTRAINT `fk_messages_mother` 
    FOREIGN KEY (`mother_id`) REFERENCES `mothers`(`mother_id`) 
    ON DELETE SET NULL ON UPDATE CASCADE;

ALTER TABLE `user_tracking_logs`
ADD CONSTRAINT `fk_tracking_logs_mother` 
    FOREIGN KEY (`mother_id`) REFERENCES `mothers`(`mother_id`) 
    ON DELETE SET NULL ON UPDATE CASCADE,
ADD CONSTRAINT `fk_tracking_logs_pregnancy` 
    FOREIGN KEY (`pregnancy_id`) REFERENCES `pregnancies`(`pregnancy_id`) 
    ON DELETE SET NULL ON UPDATE CASCADE,
ADD CONSTRAINT `fk_tracking_logs_child` 
    FOREIGN KEY (`child_id`) REFERENCES `children`(`child_id`) 
    ON DELETE SET NULL ON UPDATE CASCADE;


ALTER TABLE `meal_plans`
ADD CONSTRAINT `fk_meal_plans_mother` 
    FOREIGN KEY (`mother_id`) REFERENCES `mothers`(`mother_id`) 
    ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE `food_logs`
ADD CONSTRAINT `fk_food_logs_meal_plan` 
    FOREIGN KEY (`meal_plan_id`) REFERENCES `meal_plans`(`meal_plan_id`) 
    ON DELETE CASCADE ON UPDATE CASCADE,
ADD CONSTRAINT `fk_food_logs_food` 
    FOREIGN KEY (`food_id`) REFERENCES `foods`(`food_id`) 
    ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE `user_bookmarks`
ADD CONSTRAINT `fk_bookmarks_user` 
    FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`) 
    ON DELETE CASCADE ON UPDATE CASCADE,
ADD CONSTRAINT `fk_bookmarks_edu_tool` 
    FOREIGN KEY (`edu_tools_id`) REFERENCES `educational_tools`(`edu_tools_id`) 
    ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE `question_options`
ADD CONSTRAINT `fk_options_question` 
    FOREIGN KEY (`question_id`) REFERENCES `questions`(`question_id`) 
    ON DELETE CASCADE ON UPDATE CASCADE;
