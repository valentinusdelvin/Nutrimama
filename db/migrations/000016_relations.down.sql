-- ============================================
-- 000016_relations.down.sql
-- DROP ALL FOREIGN KEY CONSTRAINTS
-- ============================================

-- 1. CHILDREN TABLE
ALTER TABLE `children`
DROP FOREIGN KEY `fk_children_pregnancy`,
DROP FOREIGN KEY `fk_children_mother`;

-- 2. MOTHERS TABLE
ALTER TABLE `mothers`
DROP FOREIGN KEY `fk_mothers_user`;

-- 3. PREGNANCIES TABLE
ALTER TABLE `pregnancies`
DROP FOREIGN KEY `fk_pregnancies_mother`;

-- 4. CONSULTATION_SESSIONS TABLE
ALTER TABLE `consultation_sessions`
DROP FOREIGN KEY `fk_consultation_mother`,
DROP FOREIGN KEY `fk_consultation_consultant`;

-- 5. MESSAGES TABLE
ALTER TABLE `messages`
DROP FOREIGN KEY `fk_messages_consultant`,
DROP FOREIGN KEY `fk_messages_mother`;

-- 6. USER_TRACKING_LOGS TABLE
ALTER TABLE `user_tracking_logs`
DROP FOREIGN KEY `fk_tracking_logs_mother`,
DROP FOREIGN KEY `fk_tracking_logs_pregnancy`,
DROP FOREIGN KEY `fk_tracking_logs_child`;

-- 7. NUTRITION_TRACKING TABLE
ALTER TABLE `nutrition_tracking`
DROP FOREIGN KEY `fk_nutrition_tracking_mother`,
DROP FOREIGN KEY `fk_nutrition_tracking_child`;

-- 8. MEAL_PLANS TABLE
ALTER TABLE `meal_plans`
DROP FOREIGN KEY `fk_meal_plans_mother`;

-- 9. FOOD_LOGS TABLE
ALTER TABLE `food_logs`
DROP FOREIGN KEY `fk_food_logs_meal_plan`,
DROP FOREIGN KEY `fk_food_logs_food`;

-- 10. USER_BOOKMARKS TABLE
ALTER TABLE `user_bookmarks`
DROP FOREIGN KEY `fk_bookmarks_user`,
DROP FOREIGN KEY `fk_bookmarks_edu_tool`;

-- 11. QUESTION_OPTIONS TABLE
ALTER TABLE `question_options`
DROP FOREIGN KEY `fk_options_question`;
