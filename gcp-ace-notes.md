# Google Cloud Platform

## Google Cloud Design and Structure

### Google Cloud Fiber Network (At the time of writing notes)

- 28 regions
- 85 zones
- 146 edge location around the world

### 3 Main Pillar of GCP Design

- Provide Trust and Security
- Open Cloud - Open Source
- Analytics and Artificial Intelligence

## Understanding Networking and Security in GCP

### 3-Layer Network that build Google Cloud

- VPC (Virtual Private Cloud)
- Subnet (Division of resources that gets private IP addresses)
- Router (Which traffic should go to where)

### IAM and Firewall rules

IAM and firewall rules allow you to set Granular policies to practice the "principle of least privilege"

- IAM
  - Who enable and protect his neighborhood and make sure everything is safe
- Firewall rule
  - Which traffic is allowed and which it not
  - We term this as Ingress and Egress traffic of our network

## Security in Google Cloud

- Trust and Security
- 6 Layers of Security in Trusted Cloud Infrastructure
  - Device
  - Internet
  - Identity
  - Storage
  - Deploy
  - Hardware
- Encryption at rest

## Google Cloud APIs

- Google Cloud APIs allow to interact with Google cloud service programmatically through the console
- API management console is used to monitor requests, traffic, error, and latency on any enabled API.
- API Keys (OR API credentials) that we can get benefit from without using any service account or without any Google-managed user account
- We have API library that provides tons of services to interact with GCP service programmatically
- We can enable and create API credentials for any API

## Database Overview of Google Cloud

- Relational Database
  - Cloud SQL
  - Cloud Spanner
    - Enterprise-grade, globally-distributed, and strongly-consistent database
- Non-Relational Database
  - Cloud Bigtable -> NoSQL based database
  - Cloud Firestore
    - develop mobile, web and IoT application
  - Firebase
  - Memorystore
    - Fully managed Redis and Memcached database for millisecond data access

## Google Cloud Shell

- Cloud Shell runs on Linux server
  - Its a linu-based container
- It holds up to 5 GB of storage for us
- Comes with all of our favorite tools preinstalled like gcloud, MySql, Kubernetes, Docker, and much more...

### Google Cloud Shell Editor

- We can open Cloud Shell Editor in browser
- We can interact with any type of files right from the browser

### Google Cloud SDK

Set of command-line tools that helps us to manage GCP resources through the terminal

- gcloud
- gsutil
- bq
- kubectl

These tools can be run in interactive mode as well as automated scripts.

#### Authorization to GCP

- User account
  - Single machine
  - An account for a user
- Service account
  - google account
  - Multiple machines
  - Can be used programmatically

#### Set of commands to initialize, authorize, and setup

- gcloud init
- gcloud auth login
  - Authorize our access for gcloud with user's google credentials and sets the current account as default
  - This is similar to Connect-AzAccount in Microsoft Azure
- gcloud auth list
  - Ensures that you are authenticated with the right account by running gcloud auth list. The active account will be marked with an asterisk (*)
- gcloud config set account ACCOUNT
  - gcloud config set account abc@go.dev
- gcloud config
  - used to configure accounts and projects
- gcloud config list
  - it'll show the properties and active configuration
- gcloud components
  - install, update, delete sdk components that gives us more flexibility with different resources
  - We can refer "Components" in GCP as "Resource Provider" in Microsoft Azure

#### gcloud command format

gcloud [COMPONENT] [ENTITY] [OPERATION] [POSITIONAL_ARGUMENTS] [FLAGS]
gcloud   compute   instance   create     demo-instance-1       --zone=us-central1-a

## gcloud commands

- gcloud services enable storage.googleapis.com
- gcloud services enable pubsub.googleapis.com