SET NAMES utf8mb4;
SET character_set_client = utf8mb4;
SET character_set_results = utf8mb4;
SET character_set_connection = utf8mb4;

UPDATE `admin-dashboard`.reservations
SET route = CONCAT(depart_ville, CHAR(226,134,146), arrivee_ville)
WHERE route LIKE '%?%' OR route LIKE '%ÔåÆ%' OR route LIKE '%-->%';

UPDATE `admin-dashboard`.reservations
SET classe = CONCAT(CHAR(195,137), 'conomique')
WHERE classe = '??conomique' OR classe = '?conomique';

UPDATE `admin-dashboard`.reservations
SET motif = CONCAT(CHAR(195,137), 'tudes')
WHERE motif IN ('╔tudes', '??tudes');

UPDATE `admin-dashboard`.reservations
SET notes = CONCAT('Souhaite si', CHAR(195,168), 'ge c', CHAR(195,180), 't', CHAR(195,169), ' hublot')
WHERE notes = 'Souhaite si?ge c?t? hublot';

UPDATE `admin-dashboard`.reservations
SET notes = CONCAT('Annulation de derni', CHAR(195,168), 're minute')
WHERE notes = 'Annulation de derni?re minute';

UPDATE `admin-dashboard`.reservations
SET notes = CONCAT('Court s', CHAR(195,169), 'jour')
WHERE notes = 'Court s?jour';

UPDATE `admin-dashboard`.reservations
SET notes = CONCAT(CHAR(195,137), 'placement ', CHAR(195,169), 'tudiant')
WHERE notes = 'D?placement ?tudiant';

UPDATE `admin-dashboard`.reservations
SET notes = CONCAT('Demande assistance ', CHAR(195,160), ' l', CHAR(39), 'a', CHAR(195,180), 'roport')
WHERE notes = 'Demande assistance ?? l\'a??roport';

UPDATE `admin-dashboard`.reservations
SET notes = CONCAT('Doit transporter mat', CHAR(195,169), 'riel m', CHAR(195,169), 'dical')
WHERE notes = 'Doit transporter mat??riel m??dical';

UPDATE `admin-dashboard`.reservations
SET notes = CONCAT('Enfants ', CHAR(195,160), ' bord')
WHERE notes = 'Enfants ?? bord';

UPDATE `admin-dashboard`.reservations
SET notes = CONCAT(CHAR(195,137), 'servation pour famille de 4')
WHERE notes = 'R??servation pour famille de 4';

UPDATE `admin-dashboard`.reservations
SET notes = CONCAT('Voyage m', CHAR(195,169), 'dical')
WHERE notes = 'Voyage m??dical';
