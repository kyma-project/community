# About
Here the results of the first threat modelling are documented.

The treat modelling was done in October 2018 in three sessions with Suleyman Akbas, Andreas Thaler and a representive of the kyma security team.

# What was done:
Bringing the different systems into one picture
Discussed every communication channel
Identified the assets being worth to protect
Identified the secrets in the game

# Assumptions
There is one prow cluster only. If there are more installations they need to be protected in a similar way.

# Working picture
![Working Picture](assets/landscape.JPG)

# Findings
- No blockers identified
- Dedicated cluster for job execution is recommended but not required for first setup
- Any external communication must be secured from the beginning, mainly the ingress for deck access and webhook call
- Secured internal communication is recommend but nor required for first setup
- Secret management is fine using KMS
- hmac token needs to be increased to 32 bytes
- tokens should be rotated at least at offboarding of people -> offboarding checklist
- Do not use a single technical service account for google cloud access, but introduce dedicated service accounts for the different scenarios
- There should be only few admin users mainly for operations, and no other roles or kind of access
-- Requires that all configuration is done in source code triggering a provisioning
-- No backup of configuration required by that and no access for developers needed
-- Anyone not being admin is treated as anonymous
- Use dedicated google project to avoid access to prow resources for project owners not being in relation to the prow topic
-- By default project owners have access to the prow cluster itself (can be avoided by proper RBAC setup)
-- Project owners will have access to the secrets in cloud storage
-- Project owners will have access to the build logs
-- Project owners will have access to the dynamicly created K8s clusters and VMs
-- VMs might have network access to other VMs not related to the prow setup
- Enable audit logs on prow cluster with storage on google cloud storage
- Assure that jobs have no access to API server
- Try to not inject any secrets to the jobs for example to avoid having them in logs. Use at least dedicated logic (plugin, script library of buildpack...) to process secrets so that risk of exposure gets lowered
- Dynamic cluster and VMs must be of temporary nature, assure cleanup of them even if related job gets killed unexpected for example by having periodic job running for cleanup
- As Docker-in-docker gets used for the jobs, a cleanup is required for images/containers, assure cleanup of them even if related job gets killed unexpected for example by having periodic job running for cleanup