

-- Ibu Hamil - Harian
INSERT INTO questions (question_text, question_key, category, input_type, unit, is_required, display_order) VALUES
('🥩 Protein Hewani Utama (Ikan/Telur/Daging)', 'protein_hewani_harian', 'ibu_hamil', 'multiple_choice', NULL, 1, 1),
('💧 Air Putih (Min. 10 Gelas)', 'air_putih_harian', 'ibu_hamil', 'number', 'gelas', 1, 2),
('💊 Tablet Tambah Darah (TTD)', 'ttd_harian', 'ibu_hamil', 'boolean', NULL, 1, 3),
('💊 Asam Folat & Vit D3', 'asam_folat_vit_d3_harian', 'ibu_hamil', 'boolean', NULL, 1, 4),
('⚠️ Monitoring Gejala Klinis (Mual/Lemas)', 'gejala_klinis_harian', 'ibu_hamil', 'multiple_choice', NULL, 0, 5);

-- Balita < 2 Thn - Harian
INSERT INTO questions (question_text, question_key, category, input_type, unit, is_required, display_order) VALUES
('🤱 ASI Eksklusif (0-6 bln)', 'asi_eksklusif_harian', 'balita', 'boolean', NULL, 1, 1),
('🥩 MPASI Tinggi Prohe (6-24 bln)', 'mpasi_prohe_harian', 'balita', 'boolean', NULL, 1, 2),
('🧈 Lemak Tambahan (Minyak/Santan)', 'lemak_tambahan_harian', 'balita', 'boolean', NULL, 1, 3),
('🍚 Karbohidrat & Protein Nabati', 'karbohidrat_protein_nabati_harian', 'balita', 'boolean', NULL, 1, 4),
('🥦 Sayuran & Buah', 'sayur_buah_harian', 'balita', 'boolean', NULL, 1, 5),
('🧂 Garam Beryodium', 'garam_beryodium_harian', 'balita', 'boolean', NULL, 1, 6),
('💊 Vit D3 & Zat Besi (Iron Drops)', 'vit_d3_zat_besi_harian', 'balita', 'boolean', NULL, 1, 7),
('🚽 Monitoring Pola BAB & BAK', 'bab_bak_harian', 'balita', 'multiple_choice', NULL, 0, 8);


-- Ibu Hamil - Mingguan
INSERT INTO questions (question_text, question_key, category, input_type, unit, is_required, display_order) VALUES
('🐟 Konsumsi Omega-3 / Ikan Laut', 'omega3_ikan_mingguan_ibu', 'ibu_hamil', 'multiple_choice', NULL, 1, 1),
('⚖️ Penimbangan BB Mandiri', 'bb_mandiri_mingguan', 'ibu_hamil', 'number', 'kg', 1, 2),
('🏃 Stimulasi/Aktivitas Fisik Ringan', 'aktivitas_fisik_mingguan', 'ibu_hamil', 'boolean', NULL, 1, 3);

-- Balita < 2 Thn - Mingguan
INSERT INTO questions (question_text, question_key, category, input_type, unit, is_required, display_order) VALUES
('🐟 Konsumsi Omega-3 / Ikan Laut', 'omega3_ikan_mingguan_balita', 'balita', 'multiple_choice', NULL, 1, 1);



-- Ibu Hamil - Bulanan
INSERT INTO questions (question_text, question_key, category, input_type, unit, is_required, display_order) VALUES
('💊 Suplemen Kalsium', 'kalsium_bulanan', 'ibu_hamil', 'boolean', NULL, 1, 1),
('📏 Pengukuran LILA', 'lila_bulanan', 'ibu_hamil', 'number', 'cm', 1, 2),
('🩸 Cek Kadar Hemoglobin (Hb)', 'hb_bulanan_ibu', 'ibu_hamil', 'number', 'g/dL', 1, 3),
('🩺 Pemeriksaan Kehamilan (ANC)', 'anc_bulanan', 'ibu_hamil', 'single_choice', NULL, 1, 4),
('💓 Pemeriksaan Tensi (Tekanan Darah)', 'tensi_bulanan', 'ibu_hamil', 'text', 'mmHg', 1, 5);

-- Balita < 2 Thn - Bulanan
INSERT INTO questions (question_text, question_key, category, input_type, unit, is_required, display_order) VALUES
('💊 Vitamin A (Februari & Agustus)', 'vitamin_a_bulanan', 'balita', 'single_choice', NULL, 1, 1),
('⚖️ Penimbangan Berat Badan (BB)', 'bb_bulanan_balita', 'balita', 'number', 'kg', 1, 2),
('📏 Pengukuran Tinggi/Panjang Badan (TB)', 'tb_bulanan', 'balita', 'number', 'cm', 1, 3),
('🧠 Pengukuran Lingkar Kepala (LK)', 'lk_bulanan', 'balita', 'number', 'cm', 1, 4),
('💉 Imunisasi Dasar Lengkap', 'imunisasi_bulanan', 'balita', 'single_choice', NULL, 1, 5),
('🩸 Cek Kadar Hemoglobin (Hb)', 'hb_bulanan_balita', 'balita', 'number', 'g/dL', 1, 6);

-- ===================================
-- OPTION
-- ===================================

-- ==========================================
-- OPTIONS FOR HARIAN
-- ==========================================

-- ID 1: Protein Hewani Options (assuming question_id = 1)
INSERT INTO question_option (question_id, option_value, option_label, icon_emoji, display_order, is_default) VALUES
(1, 'ikan', 'Ikan', '🐟', 1, 0),
(1, 'telur', 'Telur', '🥚', 2, 0),
(1, 'daging_ayam', 'Daging Ayam', '🍗', 3, 0),
(1, 'daging_merah', 'Daging Merah', '🥩', 4, 0),
(1, 'seafood', 'Seafood Lain', '🦐', 5, 0);

-- ID 5: Monitoring Gejala Klinis (assuming question_id = 5)
-- Note: This allows multiple selections
INSERT INTO question_option (question_id, option_value, option_label, icon_emoji, display_order, is_default) VALUES
(5, 'tidak_ada', 'Tidak Ada Gejala', '✅', 1, 1),
(5, 'mual', 'Mual', '🤢', 2, 0),
(5, 'muntah', 'Muntah', '🤮', 3, 0),
(5, 'lemas', 'Lemas/Lelah', '😴', 4, 0),
(5, 'pusing', 'Pusing', '😵‍💫', 5, 0),
(5, 'sesak', 'Sesak Napas', '😮‍💨', 6, 0),
(5, 'bengkak', 'Bengkak Kaki/Tangan', '🦶', 7, 0),
(5, 'pendarahan', 'Pendarahan/Spotting', '🩸', 8, 0);

-- ID 13: Monitoring Pola BAB & BAK (assuming question_id = 13)
INSERT INTO question_option (question_id, option_value, option_label, icon_emoji, display_order, is_default) VALUES
(13, 'normal', 'Normal', '✅', 1, 0),
(13, 'sembelit', 'Sembelit/Kesulitan BAB', '💩', 2, 0),
(13, 'diare', 'Diare', '🚽', 3, 0),
(13, 'bak_normal', 'BAK Normal', '💧', 4, 0),
(13, 'bak_berlebihan', 'BAK Berlebihan', '🚰', 5, 0),
(13, 'bak_sedikit', 'BAK Sedikit', '⚠️', 6, 0);

-- ==========================================
-- OPTIONS FOR MINGGUAN
-- ==========================================

-- ID 14: Omega-3 (assuming question_id = 14)
INSERT INTO question_option (question_id, option_value, option_label, icon_emoji, display_order, is_default) VALUES
(14, 'tuna', 'Tuna/Salmon', '🐟', 1, 0),
(14, 'sarden', 'Sarden', '🥫', 2, 0),
(14, 'mackerel', 'Mackerel/Ikan Kembung', '🎣', 3, 0),
(14, 'suplemen', 'Suplemen Minyak Ikan', '💊', 4, 0),
(14, 'belum', 'Belum Konsumsi', '❌', 5, 0);

-- ==========================================
-- OPTIONS FOR BULANAN
-- ==========================================

-- ID 21: ANC Check (assuming question_id = 21)
INSERT INTO question_option (question_id, option_value, option_label, icon_emoji, display_order, is_default) VALUES
(21, 'sudah_trimester1', 'Sudah - Trimester 1', '🏥', 1, 0),
(21, 'sudah_trimester2', 'Sudah - Trimester 2', '🏥', 2, 0),
(21, 'sudah_trimester3', 'Sudah - Trimester 3', '🏥', 3, 0),
(21, 'belum', 'Belum Periksa', '⚠️', 4, 0),
(21, 'tidak_perlu', 'Tidak Perlu/Bukan Jadwal', '📅', 5, 0);

-- ID 23: Vitamin A (assuming question_id = 23)
INSERT INTO question_option (question_id, option_value, option_label, icon_emoji, display_order, is_default) VALUES
(23, 'sudah_feb', 'Sudah - Dosis Februari', '💊', 1, 0),
(23, 'sudah_aug', 'Sudah - Dosis Agustus', '💊', 2, 0),
(23, 'belum_wajib', 'Belum (Belum Jadwal)', '⏳', 3, 1),
(23, 'belum_lupa', 'Belum (Terlewat)', '❌', 4, 0),
(23, 'tidak_berlaku', 'Tidak Berlaku (Usia > 2th)', '➖', 5, 0);

-- ID 27: Imunisasi Dasar (assuming question_id = 27)
INSERT INTO question_option (question_id, option_value, option_label, icon_emoji, display_order, is_default) VALUES
(27, 'lengkap', 'Lengkap Sesuai Jadwal', '✅', 1, 0),
(27, 'sebagian', 'Sebagian (Ada yang Terlewat)', '⚠️', 2, 0),
(27, 'belum', 'Belum Imunisasi', '❌', 3, 0),
(27, 'tidak_sesuai', 'Tidak Sesuai Jadwal', '📋', 4, 0);

-- ID 28: Hemoglobin for Balita (interpretation options) (assuming question_id = 28)
INSERT INTO question_option (question_id, option_value, option_label, icon_emoji, display_order, is_default) VALUES
(28, 'normal', 'Normal (≥11 g/dL)', '🟢', 1, 0),
(28, 'rendah', 'Rendah (10-10.9 g/dL)', '🟡', 2, 0),
(28, 'anemia', 'Anemia (<10 g/dL)', '🔴', 3, 0);

INSERT INTO `foods` (`food_id`, `name`, `carbohydrates`, `fat`, `category`, `nutrition`, `image`) VALUES
(1, 'Nasi Goreng Sayur', 45.50, 12.30, 'Indonesian Main', '{\"iron\": 2.1, \"fiber\": 4.2, \"folate\": 85, \"calcium\": 35, \"protein\": 8.5, \"calories\": 320, \"vitamin_c\": 12}', NULL),
(2, 'Gado-Gado', 28.40, 18.50, 'Indonesian Main', '{\"iron\": 3.2, \"fiber\": 7.5, \"folate\": 145, \"calcium\": 120, \"protein\": 12.0, \"calories\": 295, \"vitamin_c\": 28}', NULL),
(3, 'Pepes Ikan', 5.20, 8.40, 'Indonesian Main', '{\"iron\": 1.8, \"fiber\": 1.2, \"folate\": 25, \"calcium\": 85, \"protein\": 18.5, \"calories\": 145, \"vitamin_c\": 8}', NULL),
(4, 'Ayam Goreng Kalasan', 12.80, 15.60, 'Indonesian Main', '{\"iron\": 2.5, \"fiber\": 0.8, \"folate\": 18, \"calcium\": 28, \"protein\": 22.0, \"calories\": 245, \"vitamin_c\": 5}', NULL),
(5, 'Ikan Bakar', 3.50, 10.20, 'Indonesian Main', '{\"iron\": 1.5, \"fiber\": 0.5, \"folate\": 22, \"calcium\": 45, \"protein\": 20.5, \"calories\": 165, \"vitamin_c\": 6}', NULL),
(6, 'Rendang Daging', 8.60, 22.40, 'Indonesian Main', '{\"iron\": 4.2, \"fiber\": 2.1, \"folate\": 35, \"calcium\": 32, \"protein\": 19.5, \"calories\": 285, \"vitamin_c\": 8}', NULL),
(7, 'Opor Ayam', 9.20, 16.80, 'Indonesian Main', '{\"iron\": 2.1, \"fiber\": 1.5, \"folate\": 28, \"calcium\": 42, \"protein\": 16.5, \"calories\": 215, \"vitamin_c\": 6}', NULL),
(8, 'Sate Ayam', 6.40, 11.20, 'Indonesian Main', '{\"iron\": 2.8, \"fiber\": 0.8, \"folate\": 22, \"calcium\": 25, \"protein\": 17.8, \"calories\": 175, \"vitamin_c\": 4}', NULL),
(9, 'Sayur Asem', 12.50, 3.20, 'Indonesian Main', '{\"iron\": 1.8, \"fiber\": 4.8, \"folate\": 65, \"calcium\": 55, \"protein\": 3.5, \"calories\": 85, \"vitamin_c\": 22}', NULL),
(10, 'Sayur Lodeh', 14.80, 8.50, 'Indonesian Main', '{\"iron\": 2.1, \"fiber\": 3.5, \"folate\": 55, \"calcium\": 65, \"protein\": 4.2, \"calories\": 145, \"vitamin_c\": 18}', NULL),
(11, 'Nasi Uduk', 52.40, 14.20, 'Indonesian Main', '{\"iron\": 2.5, \"fiber\": 2.8, \"folate\": 42, \"calcium\": 28, \"protein\": 6.5, \"calories\": 345, \"vitamin_c\": 4}', NULL),
(12, 'Nasi Liwet', 48.60, 10.80, 'Indonesian Main', '{\"iron\": 2.8, \"fiber\": 3.2, \"folate\": 48, \"calcium\": 35, \"protein\": 7.2, \"calories\": 305, \"vitamin_c\": 6}', NULL),
(13, 'Soto Ayam', 15.20, 9.40, 'Indonesian Soup', '{\"iron\": 2.2, \"fiber\": 2.1, \"folate\": 32, \"calcium\": 38, \"protein\": 14.5, \"calories\": 165, \"vitamin_c\": 8}', NULL),
(14, 'Sop Buntut', 8.50, 12.60, 'Indonesian Soup', '{\"iron\": 3.5, \"fiber\": 1.5, \"folate\": 28, \"calcium\": 45, \"protein\": 16.8, \"calories\": 185, \"vitamin_c\": 6}', NULL),
(15, 'Rawon', 12.80, 11.20, 'Indonesian Soup', '{\"iron\": 4.5, \"fiber\": 2.8, \"folate\": 35, \"calcium\": 42, \"protein\": 15.2, \"calories\": 175, \"vitamin_c\": 5}', NULL),
(16, 'Bakso', 18.40, 10.50, 'Indonesian Soup', '{\"iron\": 2.8, \"fiber\": 1.8, \"folate\": 25, \"calcium\": 35, \"protein\": 12.5, \"calories\": 195, \"vitamin_c\": 4}', NULL),
(17, 'Siomay', 22.50, 6.80, 'Indonesian Side', '{\"iron\": 2.2, \"fiber\": 2.5, \"folate\": 32, \"calcium\": 55, \"protein\": 10.5, \"calories\": 165, \"vitamin_c\": 8}', NULL),
(18, 'Batagor', 28.60, 14.20, 'Indonesian Side', '{\"iron\": 2.5, \"fiber\": 2.8, \"folate\": 38, \"calcium\": 48, \"protein\": 12.8, \"calories\": 265, \"vitamin_c\": 6}', NULL),
(19, 'Perkedel Kentang', 32.40, 12.50, 'Indonesian Side', '{\"iron\": 1.8, \"fiber\": 3.5, \"folate\": 28, \"calcium\": 25, \"protein\": 5.5, \"calories\": 255, \"vitamin_c\": 18}', NULL),
(20, 'Tempe Goreng', 12.80, 11.60, 'Indonesian Side', '{\"iron\": 3.2, \"fiber\": 4.5, \"folate\": 55, \"calcium\": 85, \"protein\": 14.2, \"calories\": 165, \"vitamin_c\": 4}', NULL),
(21, 'Grilled Salmon', 0.50, 13.40, 'International', '{\"iron\": 0.8, \"fiber\": 0.0, \"folate\": 35, \"calcium\": 12, \"protein\": 20.0, \"calories\": 208, \"vitamin_c\": 4}', NULL),
(22, 'Chicken Breast', 0.00, 3.60, 'International', '{\"iron\": 1.5, \"fiber\": 0.0, \"folate\": 8, \"calcium\": 15, \"protein\": 31.0, \"calories\": 165, \"vitamin_c\": 0}', NULL),
(23, 'Oatmeal', 54.00, 6.50, 'International', '{\"iron\": 4.7, \"fiber\": 10.6, \"folate\": 56, \"calcium\": 54, \"protein\": 16.9, \"calories\": 389, \"vitamin_c\": 0}', NULL),
(24, 'Greek Yogurt', 9.00, 10.00, 'International', '{\"iron\": 0.1, \"fiber\": 0.0, \"folate\": 12, \"calcium\": 110, \"protein\": 9.9, \"calories\": 146, \"vitamin_c\": 1}', NULL),
(25, 'Quinoa Salad', 39.40, 8.40, 'International', '{\"iron\": 4.5, \"fiber\": 5.2, \"folate\": 85, \"calcium\": 65, \"protein\": 9.5, \"calories\": 280, \"vitamin_c\": 22}', NULL),
(26, 'Lentil Soup', 30.00, 3.00, 'International', '{\"iron\": 6.5, \"fiber\": 7.8, \"folate\": 180, \"calcium\": 35, \"protein\": 12.0, \"calories\": 165, \"vitamin_c\": 8}', NULL),
(27, 'Avocado Toast', 32.50, 16.80, 'International', '{\"iron\": 2.8, \"fiber\": 9.5, \"folate\": 125, \"calcium\": 85, \"protein\": 8.5, \"calories\": 315, \"vitamin_c\": 18}', NULL),
(28, 'Spinach Omelette', 3.20, 14.50, 'International', '{\"iron\": 3.2, \"fiber\": 1.8, \"folate\": 145, \"calcium\": 95, \"protein\": 12.5, \"calories\": 175, \"vitamin_c\": 12}', NULL),
(29, 'Grilled Chicken Salad', 12.80, 8.50, 'International', '{\"iron\": 2.5, \"fiber\": 4.5, \"folate\": 85, \"calcium\": 55, \"protein\": 22.0, \"calories\": 185, \"vitamin_c\": 35}', NULL),
(30, 'Baked Sweet Potato', 41.40, 0.30, 'International', '{\"iron\": 1.2, \"fiber\": 6.6, \"folate\": 32, \"calcium\": 68, \"protein\": 4.0, \"calories\": 180, \"vitamin_c\": 35}', NULL),
(31, 'Tuna Sandwich', 38.50, 8.20, 'International', '{\"iron\": 2.5, \"fiber\": 4.2, \"folate\": 45, \"calcium\": 85, \"protein\": 18.5, \"calories\": 285, \"vitamin_c\": 8}', NULL),
(32, 'Vegetable Stir Fry', 18.50, 9.80, 'International', '{\"iron\": 2.8, \"fiber\": 5.8, \"folate\": 95, \"calcium\": 75, \"protein\": 6.5, \"calories\": 175, \"vitamin_c\": 65}', NULL),
(33, 'Bayam Rebus', 6.80, 0.80, 'Vegetable', '{\"iron\": 3.6, \"fiber\": 3.8, \"folate\": 195, \"calcium\": 145, \"protein\": 4.0, \"calories\": 35, \"vitamin_c\": 45}', NULL),
(34, 'Brokoli Kukus', 11.20, 0.60, 'Vegetable', '{\"iron\": 1.2, \"fiber\": 5.2, \"folate\": 108, \"calcium\": 62, \"protein\": 3.7, \"calories\": 55, \"vitamin_c\": 89}', NULL),
(35, 'Wortel Rebus', 14.80, 0.40, 'Vegetable', '{\"iron\": 0.6, \"fiber\": 4.2, \"folate\": 28, \"calcium\": 55, \"protein\": 1.5, \"calories\": 65, \"vitamin_c\": 8}', NULL),
(36, 'Kangkung Tumis', 8.50, 5.20, 'Vegetable', '{\"iron\": 3.8, \"fiber\": 3.5, \"folate\": 105, \"calcium\": 125, \"protein\": 3.2, \"calories\": 85, \"vitamin_c\": 55}', NULL),
(37, 'Labu Siam', 9.20, 0.50, 'Vegetable', '{\"iron\": 1.2, \"fiber\": 3.8, \"folate\": 45, \"calcium\": 42, \"protein\": 2.0, \"calories\": 45, \"vitamin_c\": 18}', NULL),
(38, 'Terong Balado', 12.50, 8.80, 'Vegetable', '{\"iron\": 1.5, \"fiber\": 6.5, \"folate\": 65, \"calcium\": 35, \"protein\": 2.8, \"calories\": 135, \"vitamin_c\": 22}', NULL),
(39, 'Capcay', 15.80, 7.50, 'Vegetable', '{\"iron\": 2.5, \"fiber\": 4.8, \"folate\": 85, \"calcium\": 85, \"protein\": 5.5, \"calories\": 145, \"vitamin_c\": 85}', NULL),
(40, 'Asparagus', 7.40, 0.40, 'Vegetable', '{\"iron\": 2.1, \"fiber\": 3.6, \"folate\": 135, \"calcium\": 32, \"protein\": 3.9, \"calories\": 35, \"vitamin_c\": 12}', NULL),
(41, 'Pisang', 27.00, 0.30, 'Fruit', '{\"iron\": 0.6, \"fiber\": 3.1, \"folate\": 24, \"calcium\": 5, \"protein\": 1.3, \"calories\": 105, \"vitamin_c\": 10}', NULL),
(42, 'Alpukat', 17.00, 29.50, 'Fruit', '{\"iron\": 1.1, \"fiber\": 13.5, \"folate\": 163, \"calcium\": 24, \"protein\": 4.0, \"calories\": 322, \"vitamin_c\": 20}', NULL),
(43, 'Jeruk', 17.50, 0.20, 'Fruit', '{\"iron\": 0.2, \"fiber\": 4.4, \"folate\": 54, \"calcium\": 61, \"protein\": 1.3, \"calories\": 73, \"vitamin_c\": 82}', NULL),
(44, 'Pepaya', 15.00, 0.50, 'Fruit', '{\"iron\": 0.4, \"fiber\": 2.5, \"folate\": 38, \"calcium\": 24, \"protein\": 0.7, \"calories\": 62, \"vitamin_c\": 88}', NULL),
(45, 'Mangga', 25.00, 0.60, 'Fruit', '{\"iron\": 0.3, \"fiber\": 3.0, \"folate\": 68, \"calcium\": 18, \"protein\": 0.8, \"calories\": 107, \"vitamin_c\": 60}', NULL),
(46, 'Apel', 25.00, 0.30, 'Fruit', '{\"iron\": 0.2, \"fiber\": 4.4, \"folate\": 6, \"calcium\": 11, \"protein\": 0.5, \"calories\": 95, \"vitamin_c\": 8}', NULL),
(47, 'Semangka', 11.50, 0.20, 'Fruit', '{\"iron\": 0.4, \"fiber\": 0.6, \"folate\": 5, \"calcium\": 11, \"protein\": 0.9, \"calories\": 46, \"vitamin_c\": 12}', NULL),
(48, 'Buah Naga', 22.00, 0.60, 'Fruit', '{\"iron\": 1.9, \"fiber\": 5.0, \"folate\": 8, \"calcium\": 18, \"protein\": 2.0, \"calories\": 102, \"vitamin_c\": 20}', NULL),
(49, 'Telur Rebus', 1.10, 11.00, 'Dairy & Egg', '{\"iron\": 1.2, \"fiber\": 0.0, \"folate\": 47, \"calcium\": 50, \"protein\": 13.0, \"calories\": 155, \"vitamin_c\": 0}', NULL),
(50, 'Susu UHT', 12.00, 8.00, 'Dairy & Egg', '{\"iron\": 0.1, \"fiber\": 0.0, \"folate\": 12, \"calcium\": 276, \"protein\": 8.0, \"calories\": 146, \"vitamin_c\": 2}', NULL),
(51, 'Keju Cheddar', 3.40, 33.30, 'Dairy & Egg', '{\"iron\": 0.7, \"fiber\": 0.0, \"folate\": 27, \"calcium\": 710, \"protein\": 25.0, \"calories\": 402, \"vitamin_c\": 0}', NULL),
(52, 'Yogurt Plain', 7.00, 3.30, 'Dairy & Egg', '{\"iron\": 0.1, \"fiber\": 0.0, \"folate\": 7, \"calcium\": 110, \"protein\": 10.0, \"calories\": 59, \"vitamin_c\": 1}', NULL),
(53, 'Nasi Putih', 53.20, 0.90, 'Grains & Legumes', '{\"iron\": 0.8, \"fiber\": 0.6, \"folate\": 5, \"calcium\": 10, \"protein\": 4.4, \"calories\": 242, \"vitamin_c\": 0}', NULL),
(54, 'Nasi Merah', 45.80, 1.80, 'Grains & Legumes', '{\"iron\": 0.8, \"fiber\": 3.5, \"folate\": 8, \"calcium\": 20, \"protein\": 5.0, \"calories\": 216, \"vitamin_c\": 0}', NULL),
(55, 'Kacang Hijau', 38.70, 0.50, 'Grains & Legumes', '{\"iron\": 5.8, \"fiber\": 15.4, \"folate\": 159, \"calcium\": 55, \"protein\": 14.0, \"calories\": 170, \"vitamin_c\": 3}', NULL),
(56, 'Kacang Merah', 40.40, 0.90, 'Grains & Legumes', '{\"iron\": 5.2, \"fiber\": 15.2, \"folate\": 180, \"calcium\": 65, \"protein\": 12.0, \"calories\": 180, \"vitamin_c\": 2}', NULL),
(57, 'Kacang Almond', 21.60, 49.90, 'Snack', '{\"iron\": 3.7, \"fiber\": 12.5, \"folate\": 44, \"calcium\": 269, \"protein\": 21.2, \"calories\": 579, \"vitamin_c\": 0}', NULL),
(58, 'Kismis', 79.20, 0.50, 'Snack', '{\"iron\": 1.9, \"fiber\": 3.7, \"folate\": 5, \"calcium\": 50, \"protein\": 3.1, \"calories\": 299, \"vitamin_c\": 2}', NULL),
(59, 'Roti Gandum', 41.30, 3.50, 'Snack', '{\"iron\": 3.6, \"fiber\": 7.0, \"folate\": 85, \"calcium\": 160, \"protein\": 13.0, \"calories\": 247, \"vitamin_c\": 0}', NULL),
(60, 'Puding Chia', 18.50, 9.80, 'Snack', '{\"iron\": 2.2, \"fiber\": 12.5, \"folate\": 25, \"calcium\": 180, \"protein\": 5.5, \"calories\": 180, \"vitamin_c\": 2}', NULL);

