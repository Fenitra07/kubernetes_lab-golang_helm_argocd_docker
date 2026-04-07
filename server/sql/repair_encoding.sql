-- Repair corrupted UTF-8 text caused by bad import encoding
SET NAMES utf8mb4;
SET character_set_client = utf8mb4;
SET character_set_results = utf8mb4;
SET character_set_connection = utf8mb4;

-- Rebuild route from depart_ville and arrivee_ville for corrupted rows
UPDATE reservations
SET route = CONCAT(depart_ville, ' → ', arrivee_ville)
WHERE depart_ville IS NOT NULL AND arrivee_ville IS NOT NULL
  AND (route LIKE '%?%' OR route LIKE '%ÔåÆ%' OR route LIKE '%-->%');

-- Fix common mojibake sequences in text columns
UPDATE reservations
SET
  document_type = REPLACE(REPLACE(REPLACE(REPLACE(document_type, '├®', 'é'), '├«', 'à'), '├Á', 'ù'), '├½', 'ê'),
  motif = REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(motif, '├ë', 'É'), '├®', 'é'), '├¿', 'è'), '├┤', 'ô'), '├á', 'à'), '├½', 'ê'), '├«', 'à'), '├╗', 'â'),
  notes = REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(notes, '├®', 'é'), '├¿', 'è'), '├┤', 'ô'), '├á', 'à'), '├½', 'ê'), '├«', 'à'), '├╗', 'â'), '├╕', 'â'), '├Á', 'ù'), '├¼', 'û')
WHERE document_type LIKE '%├%' OR motif LIKE '%├%' OR notes LIKE '%├%';

-- Last fallback correction for any remaining malformed motif values imported as literal question marks
UPDATE reservations
SET motif = 'Études'
WHERE motif = '??tudes';
