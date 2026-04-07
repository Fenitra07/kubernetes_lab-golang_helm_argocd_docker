-- Fix encoding issues for the reservations table
-- Run this script to ensure proper UTF-8 encoding

-- Set database charset
ALTER DATABASE `admin-dashboard` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- Modify the reservations table charset
ALTER TABLE `reservations` CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- Ensure specific columns that might contain special characters are properly encoded
ALTER TABLE `reservations`
  MODIFY COLUMN `nom` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  MODIFY COLUMN `prenom` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  MODIFY COLUMN `route` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  MODIFY COLUMN `depart_ville` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  MODIFY COLUMN `arrivee_ville` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  MODIFY COLUMN `motif` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  MODIFY COLUMN `special` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  MODIFY COLUMN `notes` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;