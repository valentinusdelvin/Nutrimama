-- Deleting Seed Data by specific identifiers
DELETE FROM `question_options` WHERE `option_label` IN ('Tidak makan sayur', '1 Porsi', '2 Porsi atau lebih');
DELETE FROM `questions` WHERE `question_key` IN ('water_intake', 'iron_tablet', 'vegetable_portion');
DELETE FROM `educational_tools` WHERE `slug` IN ('tips-makan-sehat-hamil', 'pentingnya-asam-folat', 'olahraga-ringan-bumil');
DELETE FROM `foods` WHERE `name` IN ('Nasi Putih', 'Ayam Goreng', 'Sayur Bayam', 'Telur Rebus', 'Pisang', 'Susu Ibu Hamil', 'Ikan Bakar', 'Tempe Bacem', 'Bubur Kacang Hijau');
DELETE FROM `consultants` WHERE `full_name` = 'Dr. Sarah Jenkins';
DELETE FROM `mothers` WHERE `full_name` = 'Jane Doe';
DELETE FROM `users` WHERE `email` IN ('admin@nutrimama.com', 'jane@gmail.com', 'sarah.consultant@nutrimama.com');
