### Bases de données

**Bases de données relationnelles :**
- PostgreSQL
- MySQL/MariaDB
- SQLite
- Microsoft SQL Server
- Oracle Database
- CockroachDB

**Bases de données NoSQL :**
- MongoDB
- Cassandra
- Redis
- CouchDB
- Neo4j (base de données graph)
- DynamoDB
- RethinkDB

**Bases de données en mémoire :**
- Redis
- Memcached

### Services de messagerie et files de messages

**Message brokers :**
- Kafka
- RabbitMQ
- ActiveMQ
- ZeroMQ
- NATS
- Pulsar
- NSQ

**Services de coordination :**
- cloudron.io
- Zookeeper : Nécessaire pour Kafka et autres services distribués.
- Consul
- etcd

### Cache

- Redis
- Memcached
- Hazelcast

### Recherche et analyse

**Moteurs de recherche :**
- Elasticsearch
- Solr
- Bleve

**Visualisation et monitoring :**
- Kibana (pour Elasticsearch)
- Grafana

### Monitoring et logging

**Monitoring :**
- https://github.com/siglens/siglens (new tech)
- Prometheus
- Grafana
- InfluxDB
- Telegraf

**Logging :**
- ELK Stack (Elasticsearch, Logstash, Kibana)
- Fluentd
- Graylog
- Loki
- Splunk

### Sécurité et authentification

- Keycloak (gestion des identités et des accès)
- Auth0
- OAuth2 Proxy
- Vault (gestion des secrets)
- Open Policy Agent (OPA)
- Dex (OIDC Identity)

### Outils DevOps et CI/CD

**Orchestration de conteneurs :**
- Kubernetes
- Docker Swarm
- Nomad

**CI/CD :**
- Jenkins
- GitLab CI
- Travis CI
- CircleCI
- Drone
- Spinnaker
- Argo CD

**Serveurs de build :**
- Docker
- Buildkite
- Bazel
- Packer

**Proxy inverse et serveurs web :**
- Nginx
- Apache
- Traefik
- Caddy
- Envoy

**Gestion de configuration :**
- Consul
- etcd
- Vault (pour la gestion des secrets)
- Spring Cloud Config

### Services de collaboration et de développement

**Contrôle de version :**
- Git
- GitHub
- GitLab
- Bitbucket
- Mercurial

**Outils de gestion de projet :**
- JIRA
- Trello
- Asana
- Clubhouse
- Monday.com

**Documentation :**
- Swagger/OpenAPI
- Redoc
- MkDocs
- Sphinx
- Slate

### Services d'API et Gateway

- Kong
- Apigee
- Tyk
- AWS API Gateway
- Traefik (pour l'API Gateway)
- Ambassador
- NGINX Plus

### Stockage et CDN

**Stockage d'objets :**
- AWS S3
- MinIO
- Google Cloud Storage
- Ceph
- DigitalOcean Spaces

**Content Delivery Network (CDN) :**
- Cloudflare
- Akamai
- Fastly
- AWS CloudFront
- Netlify
- Varnish

### Email et communication

**Serveurs de messagerie :**
- Postfix
- Sendmail
- Exim

**Email Testing :**
- Mailhog
- Mailcatcher
- Ethereal Email

**Services de notification :**
- Twilio (SMS, appels vocaux)
- SendGrid (emails)
- Mailgun (emails)
- Pusher (notifications en temps réel)
- Firebase Cloud Messaging (FCM)

### Outils de développement web

**Serveurs de développement :**
- Hugo (générateur de site statique)
- Jekyll (générateur de site statique)
- Gatsby
- Next.js
- Nuxt.js

### Outils spécifiques pour Go

**Web UI Testing**
- Inkmi : https://www.inkmi.com/blog/web-ui-tests-in-golang

**Hot Reloading :**
- Air : Hot reloading pour les applications Go.
- Reflex : Outil pour exécuter des commandes sur les changements de fichiers.
- Fresh : Hot reloading pour les applications Go.

**Frameworks Web :**
- Gin : Framework web rapide et minimaliste.
- Echo : Framework web hautes performances.
- Fiber : Framework web inspiré de Express.js.
- Revel : Framework web complet et performant.
- Beego : Framework web complet pour Go.
- Buffalo : Outils pour le développement web rapide.

**Templates :**
- Go Template : Package standard pour le rendu des templates.
- Pongo2 : Système de templates basé sur Django pour Go.
- Jet : Moteur de templates rapide et sécurisé pour Go.
- Ace : Moteur de templates basé sur Go.

**ORM (Object-Relational Mapping) :**
- GORM : ORM populaire pour Go.
- SQLBoiler : Générateur de code ORM basé sur des modèles SQL.
- Ent : Framework d'entités et d'ORM pour Go.
- XORM : Simple et puissant ORM pour Go.
- Pop/soda : ORM et migrations de base de données pour Go.

**Validation :**
- Validator : Bibliothèque de validation pour Go.
- Ozzo Validation : Un autre package de validation pour Go.
- Govalidator : Validateur de données.
- V8n : Validateur extensible pour Go.

**Authentification et Autorisation :**
- Casbin : Gestionnaire de contrôle d'accès.
- Authboss : Kit de gestion d'authentification pour Go.
- GoGuardian : Middleware d'authentification et d'autorisation.
- Paseto : Tokens sécurisés basés sur JSON.

**Middleware :**
- Negroni : Middleware minimaliste.
- Alice : Chaînage de middleware.
- CORS : Middleware pour gérer les requêtes CORS.
- Tollbooth : Middleware de limitation de débit.

**Testing :**
- Testify : Bibliothèque de test pour Go.
- Ginkgo : Framework de test BDD pour Go.
- GoMock : Bibliothèque de mocking pour Go.
- Gomega : Bibliothèque de matchers pour Ginkgo.
- Convey : Spécification de tests en BDD.
- httpexpect : Tests d'API HTTP en Go.

**Benchmarking :**
- GoBench : Outils de benchmarking.
- benchstat : Analyse comparative des benchmarks.
- Vegeta : Testeur de charge HTTP.
- hey : Testeur de charge HTTP simple.

**Profiling :**
- pprof : Outil intégré de profilage pour Go.
- Go-Torch : Générateur de graphiques Flame pour pprof.
- tracelog : Simple logger pour le traçage des applications.
- gops : Diagnostics pour les processus Go.

**Linting et Analyse Statique :**
- golangci-lint : Linter multi-outils pour Go.
- staticcheck : Analyse statique pour détecter les bugs et améliorer la qualité du code.
- revive : Outil de linting rapide et configurable pour Go.
- ineffassign : Détecte les affectations inefficaces.
- gosimple : Détecte les constructions de code simplifiables.
- errcheck : Vérifie que les erreurs sont bien gérées.

**Gestion de dépendances :**
- Go Modules : Système de gestion des dépendances officiel pour Go.
- dep : Ancien gestionnaire de dépendances pour Go.
- Glide : Gestionnaire de dépendances pour Go.
- Godep : Gestionnaire de dépendances pour Go.

**Makefile :**
- Mage : Alternative à Make pour les scripts de build en Go.

**Gestion des versions :**
- goreleaser : Outil pour créer et publier des releases Go.

**Génération de documentation :**
- Swagger : Outils pour générer de la documentation API à partir de code Go.
- godoc : Serveur de documentation et documentation en ligne pour Go.
- Go Swagger : Générateur de code OpenAPI et outil de documentation.

**CI/CD :**
- GitHub Actions : Intégration continue avec GitHub.
- GitLab CI : CI/CD intégré à GitLab.
- Travis CI : Service d'intégration continue.
- CircleCI : Outils de CI/CD basés sur le cloud.
- Drone : Plateforme CI/CD open-source.

**Déploiement :**
- Docker : Conteneurisation des applications.
- Kubernetes : Orchestration de conteneurs.
- Helm : Gestionnaire de paquets pour Kubernetes.
- Terraform : Infrastructure as Code.
- Ansible : Automation et gestion de la configuration.
- Pulumi : Infrastructure as Code avec des langages de programmation.

**Génération de code :**
- Stringer : Génération de code pour les types string.
- Go Generate : Génération de code via directives dans le code source.
- Mockery : Générateur de mocks pour Go utilisant les interfaces.
- Enumer : Génération de code pour les énumérations en Go.
- Vecty : Génération de composants pour les interfaces web en Go.

**Débogage :**
- Delve : Débogueur pour Go.

**Autres outils :**
- gops : Diagnostics pour les processus Go.
- realize : Outil de génération et de gestion de projet avec rechargement à chaud.
- Fyne : Toolkit pour créer des interfaces graphiques pour Go.
- Lorca : Interface utilisateur graphique pour Go utilisant le navigateur.
- Gio : Toolkit pour créer des interfaces graphiques pour Go.
- termui : Interface utilisateur basée sur le terminal.
- tview : Widgets pour les interfaces utilisateur basées sur le terminal.

### Microservice clould (new tech)
- https://www.jaegertracing.io/
- https://www.cncf.io/
