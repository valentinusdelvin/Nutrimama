-- ============================================
-- FIX DATA TYPE MISMATCHES FOR FOREIGN KEYS
-- ============================================

-- 1. Fix children table - pregnancy_id and mother_id
ALTER TABLE `children`
    MODIFY `pregnancy_id` bigint UNSIGNED NULL,
    MODIFY `mother_id` bigint UNSIGNED NULL;

-- 2. Fix mothers table - user_id
ALTER TABLE `mothers`
    MODIFY `user_id` bigint UNSIGNED NOT NULL;

-- 3. Fix pregnancies table - mother_id
ALTER TABLE `pregnancies`
    MODIFY `mother_id` bigint UNSIGNED NOT NULL;

-- 4. Fix consultation_sessions table
ALTER TABLE `consultation_sessions`
    MODIFY `mother_id` bigint UNSIGNED NOT NULL,
    MODIFY `consultant_id` bigint UNSIGNED NOT NULL;

-- 5. Fix messages table
ALTER TABLE `messages`
    MODIFY `consultant_id` bigint UNSIGNED NULL,
    MODIFY `mother_id` bigint UNSIGNED NULL;

-- 6. Fix daily_tracking table
ALTER TABLE `daily_tracking`
    MODIFY `tracking_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    MODIFY `mother_id` bigint UNSIGNED NULL,
    MODIFY `pregnancy_id` bigint UNSIGNED NULL,
    MODIFY `child_id` bigint UNSIGNED NULL;

-- 7. Fix weekly_tracking table
ALTER TABLE `weekly_tracking`
    MODIFY `tracking_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    MODIFY `mother_id` bigint UNSIGNED NULL,
    MODIFY `pregnancy_id` bigint UNSIGNED NULL,
    MODIFY `child_id` bigint UNSIGNED NULL;

-- 8. Fix monthly_tracking table
ALTER TABLE `monthly_tracking`
    MODIFY `tracking_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    MODIFY `mother_id` bigint UNSIGNED NULL,
    MODIFY `pregnancy_id` bigint UNSIGNED NULL,
    MODIFY `child_id` bigint UNSIGNED NULL;

-- 9. Fix nutrition_tracking table
ALTER TABLE `nutrition_tracking`
    MODIFY `track_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    MODIFY `mother_id` bigint UNSIGNED NULL,
    MODIFY `child_id` bigint UNSIGNED NULL;

-- 10. Fix meal_plans table
ALTER TABLE `meal_plans`
    MODIFY `meal_plan_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    MODIFY `mother_id` bigint UNSIGNED NOT NULL;

-- 11. Fix food_logs table
ALTER TABLE `food_logs`
    MODIFY `food_log_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    MODIFY `meal_plan_id` bigint UNSIGNED NOT NULL,
    MODIFY `food_id` bigint UNSIGNED NOT NULL;

-- 12. Fix foods table
ALTER TABLE `foods`
    MODIFY `food_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    MODIFY `nutrition_id` bigint UNSIGNED NULL;

-- 13. Fix user_bookmarks table
ALTER TABLE `user_bookmarks`
    MODIFY `bookmark_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    MODIFY `user_id` bigint UNSIGNED NOT NULL,
    MODIFY `edu_tools_id` bigint UNSIGNED NOT NULL;

-- 14. Fix questions table
ALTER TABLE `questions`
    MODIFY `question_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    MODIFY `schedule_id` bigint UNSIGNED NULL;

-- 15. Fix question_options table
ALTER TABLE `question_options`
    MODIFY `option_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    MODIFY `question_id` bigint UNSIGNED NULL;


-- ============================================
-- NOW ADD ALL FOREIGN KEY CONSTRAINTS
-- ============================================

-- 1. CHILDREN TABLE
ALTER TABLE `children`
ADD CONSTRAINT `fk_children_pregnancy` 
    FOREIGN KEY (`pregnancy_id`) REFERENCES `pregnancies`(`pregnancy_id`) 
    ON DELETE SET NULL ON UPDATE CASCADE,
ADD CONSTRAINT `fk_children_mother` 
    FOREIGN KEY (`mother_id`) REFERENCES `mothers`(`mother_id`) 
    ON DELETE SET NULL ON UPDATE CASCADE;

-- 2. MOTHERS TABLE
ALTER TABLE `mothers`
ADD CONSTRAINT `fk_mothers_user` 
    FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`) 
    ON DELETE CASCADE ON UPDATE CASCADE;

-- 3. PREGNANCIES TABLE
ALTER TABLE `pregnancies`
ADD CONSTRAINT `fk_pregnancies_mother` 
    FOREIGN KEY (`mother_id`) REFERENCES `mothers`(`mother_id`) 
    ON DELETE CASCADE ON UPDATE CASCADE;

-- 4. CONSULTATION_SESSIONS TABLE
ALTER TABLE `consultation_sessions`
ADD CONSTRAINT `fk_consultation_mother` 
    FOREIGN KEY (`mother_id`) REFERENCES `mothers`(`mother_id`) 
    ON DELETE CASCADE ON UPDATE CASCADE,
ADD CONSTRAINT `fk_consultation_consultant` 
    FOREIGN KEY (`consultant_id`) REFERENCES `consultants`(`consultant_id`) 
    ON DELETE CASCADE ON UPDATE CASCADE;

-- 5. MESSAGES TABLE
ALTER TABLE `messages`
ADD CONSTRAINT `fk_messages_consultant` 
    FOREIGN KEY (`consultant_id`) REFERENCES `consultants`(`consultant_id`) 
    ON DELETE SET NULL ON UPDATE CASCADE,
ADD CONSTRAINT `fk_messages_mother` 
    FOREIGN KEY (`mother_id`) REFERENCES `mothers`(`mother_id`) 
    ON DELETE SET NULL ON UPDATE CASCADE;

-- 6. DAILY_TRACKING TABLE
ALTER TABLE `daily_tracking`
ADD CONSTRAINT `fk_daily_tracking_mother` 
    FOREIGN KEY (`mother_id`) REFERENCES `mothers`(`mother_id`) 
    ON DELETE SET NULL ON UPDATE CASCADE,
ADD CONSTRAINT `fk_daily_tracking_pregnancy` 
    FOREIGN KEY (`pregnancy_id`) REFERENCES `pregnancies`(`pregnancy_id`) 
    ON DELETE SET NULL ON UPDATE CASCADE,
ADD CONSTRAINT `fk_daily_tracking_child` 
    FOREIGN KEY (`child_id`) REFERENCES `children`(`child_id`) 
    ON DELETE SET NULL ON UPDATE CASCADE;

-- 7. WEEKLY_TRACKING TABLE
ALTER TABLE `weekly_tracking`
ADD CONSTRAINT `fk_weekly_tracking_mother` 
    FOREIGN KEY (`mother_id`) REFERENCES `mothers`(`mother_id`) 
    ON DELETE SET NULL ON UPDATE CASCADE,
ADD CONSTRAINT `fk_weekly_tracking_pregnancy` 
    FOREIGN KEY (`pregnancy_id`) REFERENCES `pregnancies`(`pregnancy_id`) 
    ON DELETE SET NULL ON UPDATE CASCADE,
ADD CONSTRAINT `fk_weekly_tracking_child` 
    FOREIGN KEY (`child_id`) REFERENCES `children`(`child_id`) 
    ON DELETE SET NULL ON UPDATE CASCADE;

-- 8. MONTHLY_TRACKING TABLE
ALTER TABLE `monthly_tracking`
ADD CONSTRAINT `fk_monthly_tracking_mother` 
    FOREIGN KEY (`mother_id`) REFERENCES `mothers`(`mother_id`) 
    ON DELETE SET NULL ON UPDATE CASCADE,
ADD CONSTRAINT `fk_monthly_tracking_pregnancy` 
    FOREIGN KEY (`pregnancy_id`) REFERENCES `pregnancies`(`pregnancy_id`) 
    ON DELETE SET NULL ON UPDATE CASCADE,
ADD CONSTRAINT `fk_monthly_tracking_child` 
    FOREIGN KEY (`child_id`) REFERENCES `children`(`child_id`) 
    ON DELETE SET NULL ON UPDATE CASCADE;

-- 9. NUTRITION_TRACKING TABLE
ALTER TABLE `nutrition_tracking`
ADD CONSTRAINT `fk_nutrition_tracking_mother` 
    FOREIGN KEY (`mother_id`) REFERENCES `mothers`(`mother_id`) 
    ON DELETE SET NULL ON UPDATE CASCADE,
ADD CONSTRAINT `fk_nutrition_tracking_child` 
    FOREIGN KEY (`child_id`) REFERENCES `children`(`child_id`) 
    ON DELETE SET NULL ON UPDATE CASCADE;

-- 10. MEAL_PLANS TABLE
ALTER TABLE `meal_plans`
ADD CONSTRAINT `fk_meal_plans_mother` 
    FOREIGN KEY (`mother_id`) REFERENCES `mothers`(`mother_id`) 
    ON DELETE CASCADE ON UPDATE CASCADE;

-- 11. FOOD_LOGS TABLE
ALTER TABLE `food_logs`
ADD CONSTRAINT `fk_food_logs_meal_plan` 
    FOREIGN KEY (`meal_plan_id`) REFERENCES `meal_plans`(`meal_plan_id`) 
    ON DELETE CASCADE ON UPDATE CASCADE,
ADD CONSTRAINT `fk_food_logs_food` 
    FOREIGN KEY (`food_id`) REFERENCES `foods`(`food_id`) 
    ON DELETE CASCADE ON UPDATE CASCADE;

-- 12. USER_BOOKMARKS TABLE
ALTER TABLE `user_bookmarks`
ADD CONSTRAINT `fk_bookmarks_user` 
    FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`) 
    ON DELETE CASCADE ON UPDATE CASCADE,
ADD CONSTRAINT `fk_bookmarks_edu_tool` 
    FOREIGN KEY (`edu_tools_id`) REFERENCES `educational_tools`(`edu_tools_id`) 
    ON DELETE CASCADE ON UPDATE CASCADE;

-- 13. QUESTION_OPTIONS TABLE
ALTER TABLE `question_options`
ADD CONSTRAINT `fk_options_question` 
    FOREIGN KEY (`question_id`) REFERENCES `questions`(`question_id`) 
    ON DELETE CASCADE ON UPDATE CASCADE;
