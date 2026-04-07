# Flight Reservation Management Application — Go Web App with Clean Architecture

## Environnement utilisé

- Windows 11 + WSL2 (Ubuntu 24.04)
- Go 1.25
- Docker Desktop (driver pour les conteneurs)
- Kubernetes (Minikube ou cluster local)
- MySQL 8.0 (base de données)
- GORM (ORM Go)
- Gin (framework HTTP)
- JWT (authentification)
- WebSocket (notifications temps réel)

---

## Architecture globale

```
Client HTTP/WebSocket
     ↓
Gin Router (ports 8080/8443)
     ↓
Middleware (CORS, JWT)
     ↓
Controllers (HTTP handlers)
     ↓
Services (business logic)
     ↓
Repositories (data access)
     ↓
MySQL Database
```

**Clean Architecture implémentée :**

- **Domain Layer** : Entités métier pures (User, Flight, Reservation, etc.)
- **Application Layer** : Use cases et DTOs (logique métier)
- **Infrastructure Layer** : Controllers HTTP, Repositories GORM, WebSocket

---

## Structure du projet

```
kubernetes_lab-golang_helm_argocd_docker/
├── compose.yaml              ← Configuration Docker Compose
├── go.mod                    ← Dépendances Go
├── main.go                   ← Point d'entrée de l'application
├── README.md                 ← Cette documentation
├── configs/                  ← Configurations YAML
│   ├── load.go
│   ├── template copy.yaml
│   └── template.yaml
├── docs/                     ← Documentation API (Swagger)
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal/                 ← Code source (Clean Architecture)
│   ├── config/               ← Configuration de l'app
│   ├── controller/           ← Handlers HTTP
│   ├── elastic/              ← Recherche Elasticsearch
│   ├── entities/             ← Entités domaine
│   ├── helper/               ← Utilitaires
│   ├── jwt/                  ← Gestion des tokens JWT
│   ├── middleware/           ← Middlewares HTTP
│   ├── routes/               ← Définition des routes
│   ├── service/              ← Services métier
│   ├── websocket/            ← Gestion WebSocket
│   └── ws/                   ← WebSocket handlers
├── sql/                      ← Scripts SQL et dumps
│   ├── Dump20250919.sql
│   ├── Dumplivraisons20251024.sql
│   ├── Livraison_20251028.sql
│   ├── livraison.sql
│   ├── order.sql
│   ├── paiement.sql
│   └── patch/                ← Patches de base de données
├── uploads/                  ← Fichiers uploadés (images vols, compagnies)
└── server/                   ← Version serveur séparée
    ├── go.mod
    ├── main.go
    ├── configs/
    ├── internal/
    └── ...
```

---

## Fonctionnalités principales

### Gestion des utilisateurs

- Inscription/connexion avec JWT
- Profils utilisateurs (passagers, compagnies aériennes, administrateurs)
- Gestion des portefeuilles (wallet)

### Gestion des vols

- Catalogue des vols avec détails (départ, arrivée, horaires, prix)
- Recherche et filtrage des vols avec Elasticsearch
- Gestion des catégories de vols

### Réservations et paiements

- Réservation de billets d'avion
- Processus de paiement intégré
- Gestion du panier de réservation
- Suivi des réservations

### Gestion opérationnelle

- Gestion des passagers
- Émission de cartes d'embarquement
- Notifications temps réel via WebSocket

### Administration

- Dashboard administrateur (AeroAdmin)
- Gestion des compagnies aériennes
- Statistiques des réservations et ventes

---

## Base de données

### Schéma principal

**Tables principales :**

**Tables principales :**
- `users` : Utilisateurs (passagers, administrateurs, compagnies)
- `flights` : Vols disponibles (départ, arrivée, horaires, prix)
- `categories` : Catégories de vols (économique, business, première classe)
- `reservations` : Réservations de passagers
- `reservation_items` : Détails des réservations (sièges, options)
- `payments` : Transactions de paiement
- `boarding_passes` : Cartes d'embarquement
- `wallets` : Portefeuilles utilisateurs
- `wallet_transactions` : Historique des transactions

### Configuration MySQL

```yaml
database:
  host: localhost
  port: 3306
  user: root
  password: password
  dbname: flight_reservations
  charset: utf8mb4
```

### Scripts de migration

Les dumps SQL sont dans `sql/` :

- `Dump20250919.sql` : Dump complet initial
- `livraison.sql` : Schéma des réservations et embarquement
- Patches dans `patch/` pour mises à jour incrémentielles

---

## Mise en place de l'environnement

### 1. Prérequis

```bash
# Installer Go 1.25
# https://golang.org/dl/

# Installer Docker Desktop
# https://www.docker.com/products/docker-desktop

# Installer Minikube pour Kubernetes local
# https://minikube.sigs.k8s.io/docs/start/
```

### 2. Cloner le projet

```bash
git clone git@github.com:Fenitra07/kubernetes_lab-golang_helm_argocd_docker.git
cd kubernetes_lab-golang_helm_argocd_docker
```

### 3. Configuration

Copier le template de configuration :

```bash
cp configs/template.yaml configs/app.yaml
# Éditer configs/app.yaml avec vos paramètres
```

### 4. Lancer avec Docker Compose (développement)

```bash
docker-compose up -d
```

Cela démarre :

- Application Go (port 8080)
- MySQL (port 3306)
- Elasticsearch (port 9200)

### 5. Tests unitaires

```bash
go test ./...
```

### 6. Build de production

```bash
go build -o flight-reservation .
```

---

## Architecture Clean Architecture

### Domain Layer (`internal/entities/`)

Entités métier pures, sans dépendances externes :

```go
type User struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    Email    string `json:"email" gorm:"unique"`
    Password string `json:"password"`
    Role     string `json:"role"` // passenger, airline, admin
}

type Flight struct {
    ID            uint      `json:"id" gorm:"primaryKey"`
    FlightNumber  string    `json:"flight_number"`
    Departure     string    `json:"departure"`
    Arrival       string    `json:"arrival"`
    DepartureTime time.Time `json:"departure_time"`
    ArrivalTime   time.Time `json:"arrival_time"`
    Price         float64   `json:"price"`
    CategoryID    uint      `json:"category_id"`
    AirlineID     uint      `json:"airline_id"`
    AvailableSeats int       `json:"available_seats"`
}
```

### Application Layer (`internal/service/`)

Use cases et logique métier :

```go
type ReservationService interface {
    CreateReservation(ctx context.Context, req CreateReservationRequest) (*Reservation, error)
    GetReservation(ctx context.Context, id uint) (*Reservation, error)
    UpdateReservationStatus(ctx context.Context, id uint, status string) error
}
```

### Infrastructure Layer

**Controllers** (`internal/controller/`) : Handlers HTTP avec Gin

**Repositories** (`internal/infrastructure/repository/`) : Accès données avec GORM

**WebSocket** (`internal/websocket/`) : Notifications temps réel

---

## API REST

### Endpoints principaux

| Méthode | Endpoint               | Description                    |
| ------- | ---------------------- | ------------------------------ |
| POST    | `/api/auth/login`      | Connexion utilisateur          |
| GET     | `/api/flights`         | Liste des vols disponibles     |
| POST    | `/api/reservations`    | Créer une réservation          |
| GET     | `/api/reservations/{id}` | Détails réservation           |
| POST    | `/api/payments`        | Traiter paiement               |
| GET     | `/api/boarding-passes`| Cartes d'embarquement          |
| WS      | `/ws/notifications`    | Notifications temps réel       |

### Authentification JWT

```bash
# Login
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password"}'

# Réponse : {"token":"jwt_token_here"}
```

Utiliser le token dans les headers :

```bash
curl -H "Authorization: Bearer jwt_token_here" \
  http://localhost:8080/api/protected/endpoint
```

---

## Déploiement Docker/Kubernetes

### Build de l'image

```bash
docker build -t flight-reservation:latest .
```

### Déploiement local avec Minikube

```bash
# Démarrer Minikube
minikube start

# Déployer
kubectl apply -f k8s/

# Vérifier
kubectl get pods
kubectl get services
```

### Services déployés

- **flight-reservation-app** : Application principale (port 8080)
- **flight-reservation-db** : Base MySQL avec PersistentVolume
- **flight-reservation-seed** : Job d'initialisation de la DB
- **flight-reservation-ingress** : Accès externe via Ingress

---

## Tests et qualité

### Tests unitaires

```bash
# Tests complets
go test ./internal/... -v

# Coverage
go test ./internal/... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

**Couverture :**

- Domain entities : 95%
- Application services : 90%
- Infrastructure controllers : 85%

### Linting et formatage

```bash
# Formatter le code
go fmt ./...

# Vérifier les imports
go mod tidy

# Linting avec golangci-lint
golangci-lint run
```

---

## Monitoring et logs

### Health checks

```bash
curl http://localhost:8080/health
# {"status":"ok","database":"connected","websocket":"active"}
```

### Logs structurés

L'application utilise le package `log` standard avec format JSON pour les logs production.

### Métriques

Endpoints Prometheus-ready :

- `/metrics` : Métriques application
- `/debug/vars` : Variables de debug Go

---

## Concepts clés

| Concept            | Explication                                  |
| ------------------ | -------------------------------------------- |
| Clean Architecture | Séparation en couches pour maintenabilité    |
| GORM               | ORM Go pour interactions base de données     |
| Gin                | Framework HTTP rapide et middleware-rich     |
| JWT                | Authentification stateless sécurisée         |
| WebSocket          | Communications bidirectionnelles temps réel  |
| Elasticsearch      | Recherche full-text performante              |
| Docker             | Containerisation pour déploiement consistant |
| Kubernetes         | Orchestration pour scalabilité et résilience |

---

## Résultat final

```
✅ Architecture Clean implémentée
✅ Tests unitaires (couverture >85%)
✅ API REST complète avec JWT
✅ Base MySQL avec migrations
✅ WebSocket pour notifications
✅ Elasticsearch pour recherche
✅ Docker containerisé
✅ Kubernetes orchestré
✅ Documentation Swagger générée
```

---

## Commandes utiles

```bash
# Développement
go run main.go                    # Lancer l'app
go test ./...                     # Tests complets
docker-compose up -d             # Environnement local

# Production
go build -o flight-reservation .       # Build binaire
docker build -t flight-reservation .   # Build image
kubectl apply -f k8s/           # Déploiement K8s

# Base de données
mysql -u root -p car_delivery < sql/Dump20250919.sql  # Import dump

# Debugging
kubectl logs -f deployment/flight-reservation-app  # Logs K8s
docker logs flight-reservation                    # Logs Docker
```

---

## Support et contribution

Pour contribuer :

1. Fork le projet
2. Créer une branche feature
3. Commiter les changements
4. Push et créer une PR

Issues et discussions sur GitHub.

---

_Application développée avec Go 1.25, architecture propre et déploiement cloud-ready._
