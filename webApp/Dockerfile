# ── Frontend — Nginx ──────────────────────────────────────
# Image légère Nginx Alpine
FROM nginx:alpine

# Supprimer la config Nginx par défaut
RUN rm /etc/nginx/conf.d/default.conf

# Copier notre config reverse proxy
COPY nginx.conf /etc/nginx/conf.d/app.conf

# Copier les fichiers statiques du frontend
COPY admin-dashboard.html /usr/share/nginx/html/admin-dashboard.html

# Plus tard pour React, remplacer la ligne ci-dessus par :
# COPY dist/ /usr/share/nginx/html/

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
