# AeroAdmin — Flight Reservation Management Application

Application web de gestion de réservations aériennes, construite avec Go (Clean Architecture), déployée via Docker Compose en local et Kubernetes (Helm) en production, avec CI/CD GitHub Actions.

---

## Stack technique

| Composant        | Technologie                       |
| ---------------- | --------------------------------- |
| Backend          | Go 1.25 + Gin (HTTP) + GORM (ORM) |
| Frontend         | HTML/CSS/JS servi par nginx       |
| Base de données  | MySQL 8.1                         |
| Containerisation | Docker + Docker Compose           |
| Orchestration    | Kubernetes (Docker Desktop / EKS) |
| Packaging K8s    | Helm                              |
| CI/CD            | GitHub Actions → Docker Hub       |
| GitOps           | ArgoCD (prévu)                    |

---

## Architecture globale

```
Client HTTP (navigateur)
     ↓
nginx (frontend + reverse proxy /api/*)
     ↓
Gin Router (port 8081)
     ↓
Middleware (CORS)
     ↓
Controllers (HTTP handlers)
     ↓
Use Cases (logique métier)
     ↓
Repositories (accès données GORM)
     ↓
MySQL Database
```

### Clean Architecture

```
server/internal/
├── domain/entities/        ← Entités métier pures
├── application/
│   ├── usecases/           ← Logique métier (login, réservations)
│   └── dtos/               ← Data Transfer Objects
├── interfaces/
│   ├── presenters/         ← Handlers HTTP (Gin)
│   └── repository/         ← Interfaces repository
└── infrastructure/
    ├── repository/         ← Implémentation GORM
    ├── controllers/        ← Auth + Reservation controllers
    └── routes/             ← Router Gin
```

---

## Structure du projet

```
kubernetes_lab-golang_helm_argocd_docker/
├── compose.yaml                  ← Docker Compose (dev local)
├── server/                       ← Backend Go
│   ├── Dockerfile
│   ├── main.go
│   ├── go.mod
│   ├── configs/                  ← config.yaml (DSN, ports, CORS)
│   ├── internal/                 ← Clean Architecture (voir ci-dessus)
│   └── sql/
│       └── seed_reservations.sql ← Données initiales
├── webApp/                       ← Frontend nginx
│   ├── Dockerfile
│   ├── nginx.conf                ← Reverse proxy /api/ → backend
│   └── html/                    ← Fichiers statiques
├── helm/
│   └── admin-dashboard/          ← Chart Helm
│       ├── Chart.yaml
│       ├── values.yaml
│       └── templates/
│           ├── app-deployment.yaml
│           ├── app-service.yaml
│           ├── webapp-deployment.yaml
│           ├── webapp-service.yaml
│           ├── mysql-deployment.yaml
│           ├── mysql-service.yaml
│           ├── mysql-pvc.yaml
│           ├── mysql-secret.yaml
│           ├── ingress.yaml
│           └── seed-job.yaml
└── .github/workflows/            ← CI GitHub Actions
    └── ci.yaml                   ← Build & Push Docker Hub
```

---

## API REST

| Méthode | Endpoint                | Description               |
| ------- | ----------------------- | ------------------------- |
| POST    | `/api/auth/login`       | Connexion                 |
| GET     | `/api/reservations`     | Liste des réservations    |
| POST    | `/api/reservations`     | Créer une réservation     |
| PUT     | `/api/reservations/:id` | Modifier une réservation  |
| DELETE  | `/api/reservations/:id` | Supprimer une réservation |
| GET     | `/health`               | Health check              |

### Exemple login

```bash
curl -X POST http://localhost/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"login":"test@","password":"test"}'

# Réponse :
# {"id":1,"mail":"test@","motdepasse":"test","status":200,"message":"Connexion réussie"}
```

---

## Base de données

### Schéma

**Table `login`**

| Colonne    | Type              | Description    |
| ---------- | ----------------- | -------------- |
| id         | bigint (PK, auto) | Identifiant    |
| mail       | longtext          | Adresse e-mail |
| motdepasse | longtext          | Mot de passe   |

**Table `reservations`**

Voir `server/sql/seed_reservations.sql` pour le schéma complet.

### Ajouter un utilisateur manuellement

```bash
docker exec -i admin-dashboard-db mysql -uroot -proot admin-dashboard -e "
INSERT INTO login (mail, motdepasse) VALUES ('admin@example.com', 'monmotdepasse');
"
```

---

## CI/CD — GitHub Actions

À chaque `git push` sur `main` :

1. **Tests Go** — `go test ./...`
2. **Build & Push** des images Docker sur Docker Hub :
   - `fenitra0011/admin-dashboard:latest` (backend Go)
   - `fenitra0011/admin-dashboard-webapp:latest` (frontend nginx)

Les images sont ensuite utilisées par le chart Helm via `values.yaml`.

---

## Concepts clés

| Concept            | Rôle dans ce projet                              |
| ------------------ | ------------------------------------------------ |
| Clean Architecture | Séparation domain / application / infrastructure |
| GORM               | ORM Go pour MySQL                                |
| Gin                | Framework HTTP, routing, middlewares CORS        |
| Docker Compose     | Environnement de développement local             |
| Helm               | Packaging et déploiement Kubernetes              |
| GitHub Actions     | Build et push automatiques des images            |
| ArgoCD             | GitOps — déploiement automatique K8s (à venir)   |

---

## Status du 09/04/2026 à 23h11

✅ Docker Compose — validation locale
✅ GitHub Actions CI — build & push Docker Hub automatique
✅ Helm Chart — déployé sur Kubernetes docker-desktop
✅ nginx.conf — fix upstream Docker vs K8s
✅ app-deployment.yaml — fix ordre variables d'env
✅ MySQL + seed — données persistantes sur PVC
✅ Stack complète — 3 pods Running, 0 restart
✅ README mis à jour — documentation complète

_Application développée avec Go 1.25 — architecture propre, déploiement cloud-ready._
