SET NAMES utf8mb4;
SET character_set_client = utf8mb4;
SET character_set_results = utf8mb4;
SET character_set_connection = utf8mb4;

UPDATE `admin-dashboard`.reservations
SET route = CONCAT(depart_ville, ' → ', arrivee_ville)
WHERE route LIKE '%?%' OR route LIKE '%ÔåÆ%' OR route LIKE '%-->%';

UPDATE `admin-dashboard`.reservations
SET classe = 'Économique'
WHERE classe = '??conomique' OR classe = '?conomique';

UPDATE `admin-dashboard`.reservations
SET motif = 'Études'
WHERE motif IN ('╔tudes', '??tudes');

UPDATE `admin-dashboard`.reservations
SET notes = 'Souhaite siège côté hublot'
WHERE notes = 'Souhaite si??ge c??t?? hublot';

UPDATE `admin-dashboard`.reservations
SET notes = 'Annulation de dernière minute'
WHERE notes = 'Annulation de derni??re minute';

UPDATE `admin-dashboard`.reservations
SET notes = 'Court séjour'
WHERE notes = 'Court s??jour';

UPDATE `admin-dashboard`.reservations
SET notes = 'Déplacement étudiant'
WHERE notes = 'D??placement ??tudiant';
