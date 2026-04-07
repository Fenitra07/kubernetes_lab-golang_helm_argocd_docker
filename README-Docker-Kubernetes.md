# Docker & Kubernetes lab pour le projet admin-dashboard

Ce guide décrit les fichiers créés pour dockeriser et déployer l'application sur Docker et Kubernetes.

## Docker

### Construire l'image

Depuis le dossier racine :

```bash
docker compose build
```

### Lancer les services

```bash
docker compose up -d
```

### Accéder à l'application

- API backend : `http://localhost:8081`
- MySQL : `localhost:3306`

### Seed de données

Le projet contient `server/sql/seed_reservations.sql`.
Après le démarrage de MySQL et du backend, exécutez :

```bash
docker exec -i admin-dashboard-db mysql -uroot -proot admin-dashboard < server/sql/seed_reservations.sql
```

---

## Kubernetes

Les manifests Kubernetes sont dans `k8s/`.

### Étapes rapides

1. Construire l'image Docker locale :

```bash
docker build -t admin-dashboard:latest -f server/Dockerfile .
```

2. Appliquer les manifests :

```bash
kubectl apply -f k8s/
```

3. Vérifier les pods :

```bash
kubectl get pods
```

4. Vérifier le service :

```bash
kubectl get svc
```

5. Accéder à l'application :

- avec `NodePort` : `http://<NODE_IP>:30081`
- avec Ingress : `http://admin-dashboard.local` (si ingress controller installé)

### Seed automatique dans Kubernetes

Un job est disponible dans `k8s/seed-job.yaml`.
Après le déploiement, exécutez :

```bash
kubectl apply -f k8s/seed-job.yaml
```

### Ingress optionnel

Si vous avez un ingress controller, appliquez :

```bash
kubectl apply -f k8s/ingress.yaml
```

Puis ajoutez dans votre `/etc/hosts` :

```text
127.0.0.1 admin-dashboard.local
```

### Pousser l'image vers un registre

Pour un VPS ou un vrai cluster :

```bash
docker tag admin-dashboard:latest <REGISTRY>/admin-dashboard:latest

docker push <REGISTRY>/admin-dashboard:latest
```

Puis modifiez `k8s/app-deployment.yaml` pour utiliser cette image.

### Remarque

Le manifeste Kubernetes utilise une base MySQL interne et une image nommée `admin-dashboard:latest`.
Sur un vrai cluster, utilisez un registre Docker et des secrets pour les mots de passe.
