# Google Cloud Platform

## Region & Zones - Section 1

### Region

- Region is a geographical location (like US, Asia, Europe, Japan, etc) where gcp is hosting multiple zones.
- Zone is discrete data center connected with low latency.
- Region can have 3 or more zones. For example, us-west-1, us-west-2

### Select Region that nearest to customer

- GCP has a nice tool called *[Region Picker](https://googlecloudplatform.github.io/region-picker/)*.
  - We can pick the right region based on the incoming traffic from our application.
  - We can use this tool by selecting three different parameters.
    - Carbon footprint
    - Price
    - Latency

## Hybrid and Multi-Cloud Environment

### Hybrid Cloud

- A setup (of cloud) with the combination of private cloud (or on-prem data center) while leveraging public cloud like GCP.

### Multi-Cloud

- A setup of Multi-cloud in which we do work with Azure, AWS, and GCP. For example hosting 1 service in Azure and other services on GCP.

## GCP account and Resources - Section 2

### Use Kodekloud GCP environment

- To pass the Google Cloud Digital Leader certification, it is enough to use KodeKloud environment for GCP.
  
### Setup Personal GCP account

- GCP provides $300 free credit for the 90 days
- 20 service and free to use out of 92
- console.cloud.google.com -> GCP cloud console -> Similar to Azure Portal
- Make sure to disable the console account once trial period has expired or delete the details from console that were provided during setup personal account.

### GCP Dashboard

- Project number and Project ID is unique
- Project is similar to Resource Group in Microsoft Azure.

### GCP Resource Hierarchy

- Following is the resource hierarchy in GCP
  - Organization -> This is the root node -> We can also call this as '*company*'
  - Folder -> Here comes some departments of Shared stuff, for example, BI Team, SRE, OPS, etc
  - Project -> We can create one or many projects inside each above folder, for example, dev project, prod project etc
  - Resources -> These are the resources that we create inside the project. For example, VM, DB, and Firewall etc
- Sample Hierarchy in 'tree' format:
- acloudtechie.com -> Organization
  - BI Team -> Folder
    - prj-openai-weu-dev -> Project
      - vm-openai-linux-weu-dev -> Resource
      - aml-openai-weu-dev -> Resource
      - sqlsrv-openai-weu-dev -> Resource
      - bigquery-openai-weu-dve -> Resource
  - SRE Team -> Folder
    - prj-iac-weu-dev -> Project
      - jmpsrv-sre-weu-dev -> Resource
      - firewall-sre-weu-dev -> Resource
      - splunk-sre-weu-dev
- Graphical representation of above structure
![gcp-resource-hierarchy](./assets/img/gcp-resource-hierarchy.png "GCP Resource Hierarchy by GCP and Kodekloud")

<blockquote class="note">

**Note:** Resources can interact with each other even they are in different projects.

</blockquote>

<style>
.note {
  border: 1px solid #007bff;
  background-color: #f0faff;
  padding: 1em;
  margin: 1em 0;
}
</style>

#### Hierarchy based on application environment

- Following hierarchy is based on application environment.
![decide-resource-hierarchy-option1](./assets/img/decide-resource-hierarchy-option1.svg "Hierarchy Based on Application Environment")

Source: [GCP Documentation](https://cloud.google.com/architecture/landing-zones/decide-resource-hierarchy)

### Billing in GCP

- Personal account holder billed by the end of month.
- Organization billed by the end of year (or as per the agreement)

## Compute in GCP

- There are variety of options (services) available that provides compute services in GCP. Following are the some of them:
  - Compute Engine -> It comes under IaaS (Servers) - Virtual Machine in Azure
  - Cloud Functions -> It comes under IaaS (Servers) - Azure Function - App Service
  - Google Kubernetes Engine -> It comes under IaaS (Servers) - Azure Kubernetes Service
  - Cloud Run -> It comes under IaaS (Servers)
  - App Engine - App Service (App service plan)
  - Workflow - Logic App
  - Container Optimized OS

### Understanding Compute, Persistent Disk, and Firewall

- To understand the aforementioned, look at these things
  - Application (website)
    - Compute -> is a virtual machine that can run our application
  - Application (website) images/content
    - Persistent -> is a storage where we can store images and all content.
  - Client
    - Mobile and Browser etc -> that will use our application
  - Communication and Networking will be secure through HTTP/S
    - Cloud VPC Firewall -> to allow only HTTPS traffic
  - Location (region) where the clients will access the application from
    - Zone (in region) -> the data center which is nearest to client to minimize the latency.
- Graphical representation
![compute-persistent-networking](./assets/img/compute-persistent-networking.png)
Source: Kodekloud

### Compute demo points

<blockquote class="info">

**Info:** We need to enable the "Compute Engine API" for the first time we create compute instance. This is similar to *Resource Provider* in Microsoft Azure.

</blockquote>

<style>
.info {
  border: 1px solid #28a745;
  background-color: #f0fff0;
  padding: 1em;
  margin: 1em 0;
}
</style>

### Summary of compute, persistent disk, and firewall

#### Compute

- It is nothing but a Virtual machine
- It comes with different sizes
  - Scale-out workloads (T2D)
  - General purpose workloads (E2, N2, N2D, N1)
  - Ulta-high memory (M2, M1)
  - Compute-intensive workloads (C2, C2D)
  - ...etc
- Pricing
  - Pricing is based on per-second usage of the machine types.
- Operating systems
  - Public Images are available
    - Windows
    - Linux -> Ubuntu, Debian, CentOS, etc

#### Persistent Disk

- To host our source (or application) code to run our application, we need persistent disk.
- Persistent Disk have different types
  - HDD
  - SSD etc
- We can increase the size of disk according to our needs
- We can attach the disk to our instance and de-attach as well

#### VPC Firewall Rules

- To access the compute instance either via SSH or browsing the application, we need set of rules that we define in VPC firewall rules to allow or deny the traffic.
- This VPC is the guardian for that prevent unauthorized access to our systems behind these vpc.
- We can create following rules as an example
  - SSH to compute instance using port 22
  - HTTPS to web application using port 443
  - ...etc

There are three kinds of roles in Cloud IAM:

- Primitive roles
  - Roles existing in GCP before Cloud IAM was introduced and that continue to work: Owner, Editor, and Viewer
- Predefined roles - Roles managed by GCP granting granular access control to specific services.
- Custom roles - Roles created to tailor permissions to the needs of an organization when predefined roles cannot meet them.
