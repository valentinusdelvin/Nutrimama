

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

INSERT INTO question_options (question_id, option_value, option_label, icon_emoji, display_order, is_default) VALUES
(1, 'ikan', 'Ikan', '🐟', 1, 0),
(1, 'telur', 'Telur', '🥚', 2, 0),
(1, 'daging_ayam', 'Daging Ayam', '🍗', 3, 0),
(1, 'daging_merah', 'Daging Merah', '🥩', 4, 0),
(1, 'seafood', 'Seafood Lain', '🦐', 5, 0);

INSERT INTO question_options (question_id, option_value, option_label, icon_emoji, display_order, is_default) VALUES
(5, 'tidak_ada', 'Tidak Ada Gejala', '✅', 1, 1),
(5, 'mual', 'Mual', '🤢', 2, 0),
(5, 'muntah', 'Muntah', '🤮', 3, 0),
(5, 'lemas', 'Lemas/Lelah', '😴', 4, 0),
(5, 'pusing', 'Pusing', '😵‍💫', 5, 0),
(5, 'sesak', 'Sesak Napas', '😮‍💨', 6, 0),
(5, 'bengkak', 'Bengkak Kaki/Tangan', '🦶', 7, 0),
(5, 'pendarahan', 'Pendarahan/Spotting', '🩸', 8, 0);

INSERT INTO question_options (question_id, option_value, option_label, icon_emoji, display_order, is_default) VALUES
(13, 'normal', 'Normal', '✅', 1, 0),
(13, 'sembelit', 'Sembelit/Kesulitan BAB', '💩', 2, 0),
(13, 'diare', 'Diare', '🚽', 3, 0),
(13, 'bak_normal', 'BAK Normal', '💧', 4, 0),
(13, 'bak_berlebihan', 'BAK Berlebihan', '🚰', 5, 0),
(13, 'bak_sedikit', 'BAK Sedikit', '⚠️', 6, 0);

-- ==========================================
-- OPTIONS FOR MINGGUAN
-- ==========================================

INSERT INTO question_options (question_id, option_value, option_label, icon_emoji, display_order, is_default) VALUES
(14, 'tuna', 'Tuna/Salmon', '🐟', 1, 0),
(14, 'sarden', 'Sarden', '🥫', 2, 0),
(14, 'mackerel', 'Mackerel/Ikan Kembung', '🎣', 3, 0),
(14, 'suplemen', 'Suplemen Minyak Ikan', '💊', 4, 0),
(14, 'belum', 'Belum Konsumsi', '❌', 5, 0);

-- ==========================================
-- OPTIONS FOR BULANAN
-- ==========================================

INSERT INTO question_options (question_id, option_value, option_label, icon_emoji, display_order, is_default) VALUES
(21, 'sudah_trimester1', 'Sudah - Trimester 1', '🏥', 1, 0),
(21, 'sudah_trimester2', 'Sudah - Trimester 2', '🏥', 2, 0),
(21, 'sudah_trimester3', 'Sudah - Trimester 3', '🏥', 3, 0),
(21, 'belum', 'Belum Periksa', '⚠️', 4, 0),
(21, 'tidak_perlu', 'Tidak Perlu/Bukan Jadwal', '📅', 5, 0);

INSERT INTO question_options (question_id, option_value, option_label, icon_emoji, display_order, is_default) VALUES
(23, 'sudah_feb', 'Sudah - Dosis Februari', '💊', 1, 0),
(23, 'sudah_aug', 'Sudah - Dosis Agustus', '💊', 2, 0),
(23, 'belum_wajib', 'Belum (Belum Jadwal)', '⏳', 3, 1),
(23, 'belum_lupa', 'Belum (Terlewat)', '❌', 4, 0),
(23, 'tidak_berlaku', 'Tidak Berlaku (Usia > 2th)', '➖', 5, 0);

INSERT INTO question_options (question_id, option_value, option_label, icon_emoji, display_order, is_default) VALUES
(27, 'lengkap', 'Lengkap Sesuai Jadwal', '✅', 1, 0),
(27, 'sebagian', 'Sebagian (Ada yang Terlewat)', '⚠️', 2, 0),
(27, 'belum', 'Belum Imunisasi', '❌', 3, 0),
(27, 'tidak_sesuai', 'Tidak Sesuai Jadwal', '📋', 4, 0);

INSERT INTO question_options (question_id, option_value, option_label, icon_emoji, display_order, is_default) VALUES
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


CREATE TABLE `educational_tools` (
  `edu_tools_id` bigint UNSIGNED NOT NULL,
  `publisher` varchar(255) DEFAULT NULL,
  `title` varchar(255) NOT NULL,
  `slug` varchar(255) NOT NULL,
  `content` text,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `category` varchar(100) DEFAULT NULL,
  `thumbnail` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


INSERT INTO `educational_tools` (`edu_tools_id`, `publisher`, `title`, `slug`, `content`, `created_at`, `category`, `thumbnail`) VALUES
(19, 'Dr. Sarah Wijaya, Sp.OG, M.Kes', 'Panduan Lengkap Masa Kehamilan: Dari Trimester Pertama hingga Ketiga', 'panduan-lengkap-masa-kehamilan-trimester-1-2-3', 'Kehamilan adalah perjalanan transformasi yang luar biasa selama kurang lebih 40 minggu, dibagi menjadi tiga trimester utama. Setiap fase membawa perkembangan unik baik bagi ibu maupun janin yang sedang tumbuh. Memahami perubahan yang terjadi di setiap trimester sangat penting untuk memastikan kehamilan yang sehat dan persalinan yang lancar.\r\n\r\n**TRIMESTER PERTAMA (Minggu 1-12): Fondasi Kehidupan**\r\n\r\nTrimester pertama adalah periode paling krusial dalam perkembangan janin. Pada fase ini, semua organ vital mulai terbentuk, sehingga ibu hamil perlu berhati-hati dengan asupan makanan dan gaya hidup. Tanda-tanda awal kehamilan yang umum meliputi telat haid, mual dan muntah terutama di pagi hari (morning sickness), kelelahan yang berlebihan, payudara yang sensitif dan membesar, serta perubahan emosional yang signifikan akibat lonjakan hormon progesteron dan estrogen.\r\n\r\nPerkembangan janin di trimester pertama sangat pesat. Minggu ke-4, embrio mulai menempel pada dinding rahim. Minggu ke-8, jantung janin mulai berdetak dan struktur dasar otak terbentuk. Pada akhir minggu ke-12, janin sudah memiliki bentuk manusia lengkap dengan organ-organ utama yang mulai berfungsi, meskipun ukurannya masih sekecil jeruk nipis.\r\n\r\nNutrisi yang sangat penting di trimester pertama adalah asam folat (400-600 mcg per hari) untuk mencegah cacat tabung saraf, zat besi untuk pembentukan darah, dan kalsium untuk perkembangan tulang. Ibu hamil disarankan untuk mengonsumsi sayuran berdaun hijau seperti bayam dan kangkung, buah-buahan segar, protein berkualitas dari ikan, ayam, tahu, dan tempe, serta produk susu rendah lemak. Hindari makanan mentah seperti sushi, seafood bermerkuri tinggi seperti tuna besar, dan makanan olahan yang tinggi pengawet.\r\n\r\n**TRIMESTER KEDUA (Minggu 13-27): Periode Keemasan**\r\n\r\nTrimester kedua sering disebut sebagai \"honeymoon period\" kehamilan karena gejala morning sickness biasanya mereda, energi kembali meningkat, dan risiko keguguran menurun signifikan. Perut mulai terlihat membesar secara nyata, dan sekitar minggu ke-18-20, ibu bisa merasakan gerakan janin pertama yang disebut quickening.\r\n\r\nPerkembangan janin di trimester kedua sangat menakjubkan. Minggu ke-16, janin bisa menggerakkan jari-jari tangan dan membentuk ekspresi wajah. Minggu ke-20, detak jantung janin dapat didengar dengan stetoskop, dan USG morfologi bisa menunjukkan jenis kelamin. Minggu ke-24, paru-paru mulai berkembang dengan pembentukan saccules yang akan digunakan untuk pertukaran oksigen setelah lahir. Janin juga mulai bisa merasakan suara dari luar, sehingga membaca buku atau berbicara dengan janin mulai bisa dilakukan. Pada minggu ke-27, mata janin bisa terbuka dan janin mulai mengenali suara ibu.\r\n\r\nTrimester kedua adalah waktu ideal untuk memulai atau melanjutkan rutinitas olahraga. Aktivitas fisik yang dianjurkan meliputi jalan kaki 30 menit setiap hari, renang yang sangat baik karena mengurangi beban sendi, yoga prenatal untuk fleksibilitas dan relaksasi, pilates modifikasi untuk kekuatan core, dan senam hamil ringan. Olahraga membantu mengontrol berat badan, mengurangi nyeri punggung, meningkatkan sirkulasi darah, dan mempersiapkan tubuh untuk persalinan.\r\n\r\nPemeriksaan penting di trimester kedua meliputi USG morfologi detail antara minggu 18-22 untuk memeriksa struktur anatomi janin, tes darah lengkap untuk memantau kadar hemoglobin, dan skrining diabetes gestasional antara minggu 24-28 menggunakan tes toleransi glukosa oral.\r\n\r\n**TRIMESTER KETIGA (Minggu 28-40): Persiapan Bertemu Buah Hati**\r\n\r\nTrimester ketiga adalah fase akhir kehamilan yang penuh antisipasi dan persiapan. Tubuh mengalami perubahan signifikan dengan perut yang semakin besar, sering buang air kecil akibat tekanan pada kandung kemih, nyeri punggung bagian bawah, kram kaki terutama di malam hari, dan kontraksi Braxton Hicks yang merupakan latihan otot rahim untuk persalinan yang sesungguhnya.\r\n\r\nPertumbuhan janin di trimester ketiga sangat pesat. Berat badan janin bisa bertambah 200-250 gram per minggu. Organ-organ terus matang, lapisan lemak di bawah kulit bertambah untuk mengatur suhu tubuh setelah lahir, dan sistem saraf berkembang pesat. Janin mulai menempati posisi kepala di bawah (lightening) sekitar 2-4 minggu sebelum persalinan pada kehamilan pertama.\r\n\r\nNutrisi di trimester ketiga membutuhkan peningkatan asupan kalori sekitar 450 kcal per hari. Fokus pada DHA dan omega-3 untuk perkembangan otak dan retina dari ikan salmon, sarden, dan makarel. Protein dibutuhkan 75-100 gram per hari untuk pertumbuhan jaringan. Kalsium 1000 mg untuk tulang dan gigi, zat besi 27 mg untuk mencegah anemia, dan serat tinggi untuk mencegah konstipasi yang umum terjadi. Minum air putih minimal 10-12 gelas sehari sangat penting untuk menjaga hidrasi dan volume cairan ketuban.\r\n\r\nTanda-tanda persalinan yang perlu dikenali meliputi lightening atau turunnya kepala janin ke panggul, show atau keluarnya lendir berdarah dari serviks yang menandakan pembukaan, kontraksi yang teratur semakin kuat dan tidak hilang dengan berubah posisi, pecah ketuban yang bisa berupa pocesan deras atau hanya rembesan, serta penurunan gerakan janin yang signifikan.\r\n\r\nPersiapan melahirkan yang perlu dilakukan sejak awal trimester ketiga meliputi menyiapkan tas persalinan dengan pakaian ganti, popok bayi, handuk, dan dokumen penting; menentukan metode persalinan dan membuat birth plan; mengenali rute tercepat ke rumah sakit atau bidan; memastikan dokumen lengkap seperti KTP, KK, BPJS, dan kartu hamil; serta mengikuti kelas persiapan persalinan dan perawatan bayi baru lahir.\r\n\r\nKecemasan menjelang persalinan (tokophobia) adalah hal normal, tetapi bisa dikelola dengan hipnobirthing, meditasi mindfulness, latihan pernapasan 4-7-8, pijatan prenatal, dan dukungan emosional dari pasangan serta keluarga. Ingatlah bahwa tubuh wanita secara biologis dirancang untuk melahirkan, dan jutaan wanita telah berhasil melewati proses ini sebelumnya.', '2024-01-18 17:00:00', 'Masa Kehamilan', '/uploads/image1.jpg'),
(20, 'Dr. Rina Susanto, Sp.OG & Tim Konselor Laktasi IBI', 'Komplikasi Kehamilan: Pencegahan, Deteksi Dini, dan Penanganan yang Tepat', 'komplikasi-kehamilan-pencegahan-deteksi-dini-penanganan', 'Kehamilan merupakan proses fisiologis yang kompleks dan unik bagi setiap wanita. Meskipun sebagian besar kehamilan berlangsung tanpa masalah serius, sekitar 15-20% kehamilan mengalami komplikasi yang memerlukan perhatian medis khusus. Memahami tanda-tanda peringatan, faktor risiko, dan strategi pencegahan sangat penting untuk memastikan keselamatan ibu dan janin. Artikel ini membahas komplikasi kehamilan yang paling umum beserta pendekatan evidence-based untuk penanganannya.\r\n\r\n**ANEMIA DALAM KEHAMILAN**\r\n\r\nAnemia adalah kondisi paling umum selama kehamilan, terjadi pada 40-60% ibu hamil di Indonesia, terutama akibat defisiensi zat besi. Volume darah meningkat hingga 50% selama kehamilan, sementara massa sel darah merah hanya naik 20-30%, menyebabkan hemodilusi fisiologis. Namun, ketika kadar hemoglobin turun di bawah 11 g/dL pada trimester pertama dan ketiga, atau di bawah 10.5 g/dL pada trimester kedua, diagnosis anemia ditegakkan.\r\n\r\nGejala anemia meliputi kelelahan ekstrem yang tidak proporsional dengan aktivitas, pucat pada kulit dan selaput lendir, pusing dan sakit kepala, jantung berdebar (palpitasi), sesak napas saat aktivitas ringan, dan rasa dingin pada tangan dan kaki. Anemia berat (Hb < 7 g/dL) berisiko bayi berat lahir rendah, kelahiran prematur, dan mortalitas perinatal meningkat.\r\n\r\nPencegahan dimulai sejak pra-konsepsi dengan memastikan status zat besi optimal. Selama kehamilan, konsumsi zat besi 27-30 mg per hari melalui suplemen prenatal dan makanan kaya zat besi seperti daging merah, hati ayam, bayam, dan kacang-kacangan. Vitamin C meningkatkan penyerapan zat besi non-heme hingga 6 kali, jadi kombinasikan sayuran hijau dengan jeruk atau strawberry. Hindari teh atau kopi bersamaan dengan makanan kaya zat besi karena tannin menghambat penyerapan.\r\n\r\nPenanganan anemia ringan-sedang melibatkan suplementasi zat besi oral 60-120 mg per hari elemental iron, yang bisa ditingkatkan menjadi 2 kali sehari jika respons lambat. Efek samping seperti mual, konstipasi, atau tinja hitam bisa diminimalkan dengan konsumsi bersama makanan atau sebelum tidur. Anemia berat atau yang tidak responsif dengan terapi oral memerlukan transfusi darah atau terapi intravena di rumah sakit.\r\n\r\n**DIABETES GESTASIONAL (GDM)**\r\n\r\nDiabetes gestasional adalah intoleransi glukosa yang pertama kali dikenali selama kehamilan, terjadi pada 2-10% kehamilan. Hormon plasenta seperti human placental lactogen (hPL) dan hormon kontra-insulin lainnya meningkat seiring kehamilan, menyebabkan insulin resistance fisiologis. Pada wanita dengan predisposisi genetik atau kapasitas pankreas yang terbatas, resistensi ini melebihi kemampuan kompensasi, menyebabkan hiperglikemia.\r\n\r\nFaktor risiko GDM meliputi obesitas pra-kehamilan (BMI > 30), riwayat GDM pada kehamilan sebelumnya, usia > 35 tahun, riwayat keluarga diabetes tipe 2, riwayat melahirkan bayi makrosomia (> 4 kg), riwayat keguguran berulang atau kelahiran mati intrauterine, dan sindrom ovarium polikistik (PCOS).\r\n\r\nSkrining universal dilakukan pada semua ibu hamil antara minggu 24-28 menggunakan tes toleransi glukosa oral 75 gram (OGTT). Kriteria diagnosis GDM menurut IADPSG adalah glukosa puasa ≥ 92 mg/dL, 1 jam ≥ 180 mg/dL, atau 2 jam ≥ 153 mg/dL. Satu nilai abnormal sudah cukup untuk diagnosis.\r\n\r\nKomplikasi GDM yang tidak terkontrol meliputi polihidramnion (cairan ketuban berlebihan), bayi makrosomia yang meningkatkan risiko distosia bahu dan trauma lahir, hipoglikemia neonatal akibat hiperinsulinemia janin, dan risiko preeklampsia yang lebih tinggi. Jangka panjang, ibu dengan GDM memiliki risiko 50% mengembangkan diabetes tipe 2 dalam 10 tahun, dan anak mereka berisiko obesitas dan diabetes lebih tinggi.\r\n\r\nManajemen GDM berfokus pada gaya hidup sebagai terapi pertama. Diet mediterania modifikasi dengan karbohidrat kompleks, serat tinggi, lemak sehat, dan protein berkualitas mengontrol glukosa pascaprandial. Monitoring glukosa mandiri 4 kali sehari (puasa dan 2 jam pasca makan) memberikan data untuk penyesuaian diet. Olahraga 30 menit moderat sehari meningkatkan sensitivitas insulin. Jika target glukosa tidak tercapai (puasa < 95 mg/dL, 2 jam pasca makan < 120 mg/dL), terapi insulin adalah gold standard karena tidak melewati plasenta. Metformin bisa dipertimbangkan tapi data keamanan jangka panjang masih terbatas.\r\n\r\n**PREEKLAMPSIA DAN EKLAMPSIA**\r\n\r\nPreeklampsia adalah gangguan multisistem yang didefinisikan sebagai hipertensi baru (tekanan darah ≥ 140/90 mmHg) setelah minggu ke-20 disertai proteinuria atau disfungsi organ akut. Kondisi ini terjadi pada 2-8% kehamilan secara global dan merupakan penyebab utama morbiditas maternal dan mortalitas perinatal. Eklampsia adalah komplikasi paling serius berupa kejang pada ibu dengan preeklampsia yang tidak memiliki penyebab neurologis lain.\r\n\r\nPatofisiologi preeklampsia berpusat pada disfungsi plasenta akibat implantasi abnormal trofoblas pada trimester pertama, menyebabkan pelepasan faktor antiangiogenik seperti soluble fms-like tyrosine kinase 1 (sFlt-1) dan soluble endoglin ke sirkulasi maternal. Faktor-faktor ini mengganggu pembentukan pembuluh darah, menyebabkan hipertensi, proteinuria, dan disfungsi organ.\r\n\r\nFaktor risiko meliputi nulliparitas (kehamilan pertama), riwayat preeklampsia sebelumnya (risiko 7 kali lebih tinggi), usia ekstrem (< 18 atau > 35 tahun), obesitas, hipertensi kronis, diabetes pre-existing atau GDM, penyakit ginjal, riwayat keluarga preeklampsia, kehamilan multiple, dan sindrom antifosfolipid.\r\n\r\nTanda-tanda peringatan yang harus segera dilaporkan meliputi sakit kepala parah yang tidak membaik dengan analgesik, gangguan penglihatan seperti silau atau pandangan kabur, nyeri epigastrium atau di bagian kanan atas perut (tanda gangguan hati), mual dan muntah yang parah, pembengkakan tiba-tiba pada wajah dan tangan, penurunan produksi urine, dan penurunan gerakan janin.\r\n\r\nPencegangan preeklampsia pada wanita berisiko tinggi melibatkan aspirin dosis rendes (75-150 mg) mulai minggu ke-12 hingga melahirkan, yang terbukti mengurangi insiden preeklampsia hingga 10%. Suplementasi kalsium 1-1.5 gram per hari pada populasi dengan asupan rendes juga efektif.\r\n\r\nPenanganan preeklampsia tergantung paritas dan usia kehamilan. Preeklampsia tanpa gejala parah (mild) bisa dipantau ambulatory dengan kontrol tekanan darah rutin, tes urine, dan pemeriksaan darah mingguan hingga usia kehamilan memungkinkan persalinan aman. Preeklampsia dengan gejala parah (severe) atau eklampsia memerlukan hospitalisasi segera, stabilisasi dengan magnesium sulfat untuk pencegahan kejang, pengontrolan tekanan darah dengan hidralazin atau labetalol, dan terminasi kehamilan melalui induksi atau seksio sesarea setelah stabil.\r\n\r\n**Keguguran DAN KEHAMILAN EKTOPik**\r\n\r\nKeguguran (abortus spontan) adalah terminasi kehamilan spontan sebelum minggu ke-20, terjadi pada 10-15% kehamilan yang diketahui dan 30-50% semua konsepsi. Sebagian besar (80%) terjadi pada trimester pertama, seringkali akibat aneuploidy kromosom yang fatal. Faktor risiko maternal meliputi usia > 35 tahun (risiko 20-35%), riwayat keguguran sebelumnya, kelainan uterus seperti septum atau fibroid, kelainan hormonal seperti hipotiroidisme atau hiperprolaktinemia, trombofilia, dan infeksi TORCH.\r\n\r\nTanda-tanda meliputi perdarahan pervaginam berwarna merah cerah atau coklat, nyeri perut bagian bawah seperti kram menstruasi, dan penurunan gejala kehamilan. Diagnosis USG transvaginal menunjukkan gestational sac tanpa fetal pole (blighted ovum), fetal bradycardia, atau tidak ada pertumbuhan sesuai usia kehamilan. Manajemen tergantung klinis: expectant management (tunggu alami) untuk keguguran iminen, medikal dengan misoprostol untuk keguguran incomplete, atau kuretase untuk retained products of conception atau perdarahan berat.\r\n\r\nKehamilan ektopik adalah implantasi embrio di luar rongga uterus, 95% di tuba falopii. Kondisi ini mengancam jiwa karena risiko ruptur dan hemoperitoneum. Faktor risiko meliputi riwayat infeksi pelvic inflammatory disease (PID), kehamilan ektopik sebelumnya, riwayat tuboplasty atau steril reversi, endometriosis, dan konsepsi dengan assisted reproductive technology (ART). Trias klasik adalah amenore (telat haid), perdarahan pervaginam, dan nyeri perut unilateral, meskipun presentasi bisa atipikal. Diagnosis ditegakkan dengan USG transvaginal menunjukkan gestational sac di luar uterus dan kadar beta-hCG yang tidak naik normal. Penanganan medikal dengan methotrexate untuk kehamilan ektopik unruptured dan stabil, atau bedah laparoskopi/laparotomi untuk rupture atau hemodynamic instability.\r\n\r\n**PREMATURITAS DAN KELAHIRAN MATI INTRAUTERINE (IUFD)**\r\n\r\nKelahiran prematur, didefinisikan sebagai kelahiran sebelum minggu ke-37, adalah penyebab utama morbiditas neonatal dan mortalitas bayi. Faktor risiko meliputi riwayat kelahiran prematur sebelumnya, kehamilan multiple, infeksi genital seperti bakterial vaginosis, cervix incompetence, stres psikososial berat, merokok, dan kondisi medis seperti preeklampsia atau diabetes.', '2026-04-03 11:21:28', 'Pasca Kelahiran', '/uploads/image2.jpg'),
(21, 'Dr. Siti Rahayu, Sp.OG & Psikolog Anak', 'Program Hamil Berikutnya: Interval, Persiapan, dan Pertimbangan Kesehatan', 'program-hamil-berikutnya-interval-persiapan-pertimbangan', 'Merencanakan kehamilan berikutnya memerlukan pertimbangan matang untuk kesehatan ibu dan bayi. Interval ideal antar kehamilan adalah 18-24 bulan untuk pemulihan tubuh penuh, mengurangi risiko komplikasi seperti plasenta previa, preeklampsia, dan bayi berat lahir rendah. Evaluasi kesehatan pasca melahirkan mencakup pemulihan luka persalinan, status anemia, kesehatan tiroid, dan keseimbangan mental. Perencanaan kontrasepsi selama interval melibatkan pilihan sesuai kebutuhan: IUD, implant, suntik, atau pil progestin. Persiapan fisik meliputi optimasi BMI, asupan asam folat 3 bulan sebelum konsepsi, evaluasi kondisi kronis seperti diabetes atau hipertensi, dan vaksinasi booster. Aspek psikologis juga penting: pastikan attachment dengan bayi pertama terbentuk baik, dukungan sistem tersedia, dan kesiapan menangani sibling rivalry. Keluarga berencana bukan hanya tentang jarak, tapi kualitas kesehatan dan kesiapan holistik.', '2024-06-14 17:00:00', 'Pasca Kelahiran', '/uploads/image3.jpg');

INSERT INTO `users` (`user_id`, `email`, `phone_number`, `password`, `role`, `created_at`, `timezone`) VALUES
(1, 'cons@cons.com', '012345654321', '$2a$10$dWFLYjXLkkbIp6AnEAjir.tACgtZFD4eRRU3l.ioj.V9gNQgST3PW', 'consultant', '2026-04-02 01:54:08', 'UTC'),
(2, 'user@user.com', '012345654321', '$2a$10$1k5IvsiGWC5OSFNROTy6L.Mc4R4iAtAZwLzNr0L/FyT16Sj/okoiS', 'mother', '2026-04-02 01:54:23', 'UTC'),
(4, 'admin@gmail.com', '085362310682', '$2a$10$qUwtnNnwbetqeaBSi8HOg.8Dg4DWFvM6m4SL0qBe9p82MTDoAALSm', 'admin', '2026-04-02 06:55:57', 'UTC'),
(5, 'admin@user.com', '012345654321', '$2a$10$7BNU3xC9GLU6QgeSlB3rpu/V7vviY6/mVFXcAvQVv76JfWQyg2RUW', 'admin', '2026-04-02 07:30:56', 'UTC'),
(10, 'richard@gmail.com', '', '$2a$10$hU5ZKPdfnFygMpJIFvu.mOsEFz7AQ1nlOSSQtIZN0GRgSZdoJUkgm', 'mother', '2026-04-03 02:00:28', 'UTC');
