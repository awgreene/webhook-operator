# webhook-operator

This operator can be used to test validating, mutating, and conversion webhooks.

Built using [Kubebuilder](https://book.kubebuilder.io/)

### Install with Kubebuilder and Cert-Manager

### Install with Cert-Manager

The latest instructions can be found [here](https://cert-manager.io/docs/installation/kubernetes/).

```bash
# Kubernetes 1.16+
$ kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v1.0.1/cert-manager.yaml

# Kubernetes <1.16
$ kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v1.0.1/cert-manager-legacy.yaml
```

### Build and Deploy Webhook-Operator

```bash
# Build the operator Image
$ make docker-build IMG=quay.io/agreene/webhook-operator:latest

# Push the image to docker
$ docker push quay.io/agreene/webhook-operator:latest

# Install the operator
$ make deploy IMG=quay.io/agreene/webhook-operator:latest

# Check that the pods are up and running
$ watch kubectl get pods -n webhook-operator-system

# Try and create the resource that fails validation
$ kubectl apply -f config/samples/fails.validation.webhook_v1_webhooktest.yaml
Error from server (WebhookTest.test.operators.coreos.com "webhooktest-sample" is invalid: spec.schedule: Invalid value: false: Spec.Valid must be true): error when creating "config/samples/fails.validation.webhook_v1_webhooktest.yaml": admission webhook "vwebhooktest.kb.io" denied the request: WebhookTest.test.operators.coreos.com "webhooktest-sample" is invalid: spec.schedule: Invalid value: false: Spec.Valid must be true

# Check that mutate was set to true by the mutating webhook
kubectl apply -f config/samples/passes.validation.webhook_v1_webhooktest.yaml
webhooktest.webhook.operators.coreos.io/webhooktest-sample created

# Check that Spec.Mutate is set to true:
kubectl get webhooktest webhooktest-sample -n webhook-operator-system -o yaml | yq read - spec.mutate
true
```

## OLM Build and Installation

### Build a Bundle

```bash
$ make bundle
```

### Build a bundle Image

```bash
$ make bundle-build
$ docker push quay.io/agreene/webhook-operator-bundle:latest
```

### Build an Index

```bash
$ opm index add --bundles quay.io/agreene/webhook-operator-bundle:latest --tag quay.io/agreene/webhook-operator-index:latest -c docker

$ docker push quay.io/agreene/webhook-operator-index:latest
```

### Deploy on OLM

```bash
# Create the CatalogSource
$ kubectl apply -f olm/install/00_catsrc.yaml
catalogsource.operators.coreos.com/webhook-operator-catalog created

# Create a Subscription for the Operator
$ kubectl apply -f olm/install/01_sub.yaml
subscription.operators.coreos.com/webhook-operator-subscription created

# Check that the invalid webhookTest fails validation
$ kubectl apply -f olm/example-crs/fails.validation.webhook_v1_webhooktest.yaml
Error from server (WebhookTest.test.operators.coreos.com "webhooktest-sample" is invalid: spec.schedule: Invalid value: false: Spec.Valid must be true): error when creating "olm/example-crs/fails.validation.webhook_v1_webhooktest.yaml": admission webhook "vwebhooktest.kb.io" denied the request: WebhookTest.test.operators.coreos.com "webhooktest-sample" is invalid: spec.schedule: Invalid value: false: Spec.Valid must be true

$ kubectl apply -f olm/example-crs/passes.validation.webhook_v1_webhooktest.yaml
webhooktest.webhook.operators.coreos.io/webhooktest-sample unchanged

$ kubectl get -n olm webhooktest webhooktest-sample -o yaml | yq read - spec.mutate
true
```
