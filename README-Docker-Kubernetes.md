# Docker & Kubernetes — Guide complet AeroAdmin

Ce guide couvre l'installation complète, le déploiement Docker Compose (local) et Kubernetes avec Helm.

---

## Prérequis

| Outil | Version | Usage |
|---|---|---|
| Docker Desktop | 4.x+ | Conteneurs + cluster K8s local |
| kubectl | 1.28+ | CLI Kubernetes |
| Helm | 3.x+ | Déploiement chart K8s |
| Git | 2.x+ | Cloner et pousser le code |

Vérifier les installations :

```bash
docker --version
kubectl version --client
helm version
git --version
```

---

## 1. Cloner le projet

```bash
git clone https://github.com/Fenitra07/kubernetes_lab-golang_helm_argocd_docker.git
cd kubernetes_lab-golang_helm_argocd_docker
```

---

## 2. Docker Compose — Développement local

### Lancer l'application

```bash
docker compose up -d --build
```

### Vérifier que tout tourne

```bash
docker compose ps
# Attendu :
# admin-dashboard-db      Healthy   port 3306
# admin-dashboard-app     Running   port 8081
# admin-dashboard-webapp  Running   (pas de port exposé — via nginx)
```

### Accéder à l'application

| Service | URL |
|---|---|
| Frontend | http://localhost |
| Backend API | http://localhost:8081 |
| MySQL | localhost:3306 |

### Insérer un utilisateur de test

```bash
docker exec -i admin-dashboard-db mysql -uroot -proot admin-dashboard -e "
INSERT INTO login (mail, motdepasse) VALUES ('test@', 'test');
"
```

### Tester le login (API)

```bash
curl -X POST http://localhost/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"login":"test@","password":"test"}'
```

### Voir les logs

```bash
# Tous les services
docker compose logs -f

# Par service
docker compose logs -f admin-dashboard-app
docker compose logs -f admin-dashboard-webapp
docker compose logs -f admin-dashboard-db
```

### Arrêter

```bash
docker compose down

# Arrêter ET supprimer les volumes (reset DB)
docker compose down -v
```

### ⚠️ Note nginx — Docker Compose vs Kubernetes

Le fichier `webApp/nginx.conf` a deux configurations possibles pour le upstream backend :

```nginx
# Docker Compose (service nommé "app" dans compose.yaml)
proxy_pass http://app:8081/api/;

# Kubernetes (service nommé "admin-dashboard-service" dans app-service.yaml)
proxy_pass http://admin-dashboard-service:8081/api/;
```

L'image Docker Hub (`fenitra0011/admin-dashboard-webapp`) est buildée avec `admin-dashboard-service` — correcte pour K8s. Pour Docker Compose local, changer en `app` si besoin de rebuild.

---

## 3. Kubernetes avec Helm

### Prérequis K8s

Activer Kubernetes dans Docker Desktop : **Settings → Kubernetes → Enable Kubernetes**

Vérifier le contexte actif :

```bash
kubectl config current-context
# Attendu : docker-desktop
```

Lister les contextes disponibles :

```bash
kubectl config get-contexts
```

Changer de contexte :

```bash
kubectl config use-context docker-desktop
```

### Valider le chart Helm avant déploiement

```bash
helm lint helm/admin-dashboard
# Attendu : 1 chart(s) linted, 0 chart(s) failed
```

Prévisualiser les manifests générés sans déployer :

```bash
helm template admin-dashboard helm/admin-dashboard
```

### Premier déploiement

```bash
helm install admin-dashboard helm/admin-dashboard
```

### Mise à jour (après modification du chart)

```bash
helm upgrade admin-dashboard helm/admin-dashboard
```

### Vérifier le déploiement

```bash
# Tous les pods
kubectl get pods
# Attendu :
# admin-dashboard-app-xxx      1/1   Running
# admin-dashboard-mysql-xxx    1/1   Running
# admin-dashboard-webapp-xxx   1/1   Running

# Services
kubectl get svc

# Deployments
kubectl get deployments

# Tout en une commande
kubectl get all
```

### Accéder à l'application sur K8s

Docker Desktop expose automatiquement les services sur `localhost` :

```
http://localhost
```

Pour accéder via le nom de domaine `admin-dashboard.local` (nécessite l'Ingress) :

```bash
# Ajouter dans /etc/hosts (WSL)
echo "127.0.0.1  admin-dashboard.local" | sudo tee -a /etc/hosts
```

Puis : http://admin-dashboard.local

### Logs

```bash
# Backend Go
kubectl logs -f deployment/admin-dashboard-app

# Frontend nginx
kubectl logs -f deployment/admin-dashboard-webapp

# MySQL
kubectl logs -f deployment/admin-dashboard-mysql

# Logs d'un pod spécifique
kubectl logs -f <nom-du-pod>
```

### Redémarrer un déploiement (forcer re-pull de l'image)

```bash
kubectl rollout restart deployment/admin-dashboard-app
kubectl rollout restart deployment/admin-dashboard-webapp
kubectl rollout restart deployment/admin-dashboard-mysql
```

### Surveiller le redémarrage en temps réel

```bash
kubectl get pods -w
# Ctrl+C pour arrêter
```

### Vérifier l'historique des déploiements Helm

```bash
helm history admin-dashboard
```

### Rollback Helm vers une révision précédente

```bash
helm rollback admin-dashboard 1
```

### Désinstaller

```bash
helm uninstall admin-dashboard
```

---

## 4. Debugging

### Inspecter un pod en erreur

```bash
# Voir les événements K8s
kubectl describe pod <nom-du-pod>

# Logs d'un pod crashé (dernière exécution)
kubectl logs <nom-du-pod> --previous
```

### Vérifier les variables d'environnement d'un déploiement

```bash
kubectl get deployment admin-dashboard-app -o yaml | grep -A 30 env
```

### Vérifier les secrets

```bash
kubectl get secret mysql-secret -o jsonpath='{.data}' | \
  python3 -c "import sys,json,base64; d=json.load(sys.stdin); [print(k+':', base64.b64decode(v).decode()) for k,v in d.items()]"
```

### Tester la résolution DNS dans le cluster

```bash
kubectl run test-dns --image=busybox --rm -it --restart=Never -- nslookup mysql
```

### Accéder à MySQL directement dans K8s

```bash
kubectl exec -it deployment/admin-dashboard-mysql -- mysql -uroot -proot admin-dashboard
```

### Tester le health check du backend

```bash
curl http://localhost:8081/health
# ou via le service K8s
kubectl port-forward svc/admin-dashboard-service 8081:8081
curl http://localhost:8081/health
```

### Vérifier les jobs (seed)

```bash
kubectl get jobs
kubectl logs job/admin-dashboard-seed
```

---

## 5. Seed des données

Le job de seed s'exécute automatiquement au premier `helm install`. Pour le relancer manuellement :

```bash
# Supprimer l'ancien job
kubectl delete job admin-dashboard-seed 2>/dev/null

# Relancer via Helm
helm upgrade admin-dashboard helm/admin-dashboard
```

Pour désactiver le seed sur les upgrades suivants :

```bash
helm upgrade admin-dashboard helm/admin-dashboard --set seedJob.enabled=false
```

---

## 6. CI/CD — GitHub Actions

La pipeline se déclenche à chaque `git push` sur `main` :

```
git push origin main
      ↓
GitHub Actions
      ↓  Tests Go
      ↓  Build image backend  → fenitra0011/admin-dashboard:latest
      ↓  Build image webapp   → fenitra0011/admin-dashboard-webapp:latest
Docker Hub
      ↓
kubectl rollout restart (manuel ou ArgoCD)
```

Pour forcer K8s à reprendre les nouvelles images après un push :

```bash
kubectl rollout restart deployment/admin-dashboard-app
kubectl rollout restart deployment/admin-dashboard-webapp
```

---

## 7. Commandes de référence rapide

```bash
# ── Docker Compose ──────────────────────────────────────
docker compose up -d --build        # Lancer (avec rebuild)
docker compose ps                   # État des containers
docker compose logs -f              # Logs temps réel
docker compose down                 # Arrêter
docker compose down -v              # Arrêter + reset DB

# ── Kubernetes ──────────────────────────────────────────
kubectl get pods                    # État des pods
kubectl get pods -w                 # Pods en temps réel
kubectl get all                     # Tout voir
kubectl describe pod <nom>          # Détails + événements
kubectl logs -f deployment/<nom>    # Logs temps réel
kubectl rollout restart deployment/<nom>  # Redémarrer

# ── Helm ────────────────────────────────────────────────
helm lint helm/admin-dashboard      # Valider le chart
helm install admin-dashboard helm/admin-dashboard    # Premier déploiement
helm upgrade admin-dashboard helm/admin-dashboard    # Mise à jour
helm history admin-dashboard        # Historique
helm rollback admin-dashboard 1     # Rollback révision 1
helm uninstall admin-dashboard      # Désinstaller

# ── Contextes K8s ───────────────────────────────────────
kubectl config get-contexts         # Lister les contextes
kubectl config current-context      # Contexte actif
kubectl config use-context docker-desktop  # Changer de contexte
```