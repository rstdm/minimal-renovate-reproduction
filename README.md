# Current behavior

This project uses helm v3.8.0. Renovate logs

``
INFO: Unsupported go host - cannot look up versions (repository=ssob/install/k5-operator)
"packageName": "helm.sh/helm/v3"
``

and doesn't update the helm dependency.

# Expected behavior

Renovate updates the helm dependency to version v3.8.2.