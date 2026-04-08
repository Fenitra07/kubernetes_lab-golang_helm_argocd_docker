{{/*
Nom complet du chart
*/}}
{{- define "admin-dashboard.fullname" -}}
{{ .Release.Name }}-{{ .Chart.Name }}
{{- end }}

{{/*
Labels communs appliqués à toutes les ressources
*/}}
{{- define "admin-dashboard.labels" -}}
helm.sh/chart: {{ .Chart.Name }}-{{ .Chart.Version }}
app.kubernetes.io/name: {{ .Chart.Name }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/version: {{ .Chart.AppVersion }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels pour les Deployments/Services
*/}}
{{- define "admin-dashboard.selectorLabels" -}}
app.kubernetes.io/name: {{ .Chart.Name }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Selector labels pour MySQL
*/}}
{{- define "admin-dashboard.mysqlSelectorLabels" -}}
app.kubernetes.io/name: {{ .Chart.Name }}-mysql
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}
