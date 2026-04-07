SET NAMES utf8mb4;
SET character_set_client = utf8mb4;
SET character_set_results = utf8mb4;
SET character_set_connection = utf8mb4;

UPDATE `admin-dashboard`.reservations
SET route = CONCAT(depart_ville, CHAR(226,134,146), arrivee_ville)
WHERE route LIKE '%?%' OR route LIKE '%ÔåÆ%' OR route LIKE '%-->%';
