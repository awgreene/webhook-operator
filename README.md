# webhook-operator

This operator can be used to test validating, mutating, and conversion webhooks.

Built using [Kubebuilder](https://book.kubebuilder.io/)

## Deploying the Webhook Operator with Kubebuilder and Cert-Manager

> Note that the pod will crash until OLM matches the certs in the location expected by Kubebuilder. See [this issue](https://github.com/operator-framework/operator-lifecycle-manager/issues/1315)

## Building the Webhook Operator with Kubebuilder

```bash
# Build the operator Image
$ make docker-build IMG=quay.io/agreene/webhook-operator:latest

# Push the image to docker
$ docker push quay.io/agreene/webhook-operator:latest
```

## Deploy the Webhook Operator with Cert-Manager and Kubebuilder

### Deploying Cert-Manager

The latest instructions can be found [here](https://cert-manager.io/docs/installation/kubernetes/).

```bash
# Kubernetes 1.16+
$ kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v1.0.1/cert-manager.yaml

# Kubernetes <1.16
$ kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v1.0.1/cert-manager-legacy.yaml
```

### Deploying the Webhook Operator with Kubebuilder

```bash
# Deploy the Webhook Operator with Kubebuilder
$ make deploy IMG=quay.io/agreene/webhook-operator:latest

# Check that the pods are up and running
$ watch kubectl get pods -n webhook-operator-system

# Try and create the resource that fails validation
$ kubectl apply -f config/samples/invalid.cr.yaml
Error from server (WebhookTest.test.operators.coreos.com "webhooktest-sample" is invalid: spec.schedule: Invalid value: false: Spec.Valid must be true): error when creating "config/samples/fails.validation.webhook_v1_webhooktest.yaml": admission webhook "vwebhooktest.kb.io" denied the request: WebhookTest.test.operators.coreos.com "webhooktest-sample" is invalid: spec.schedule: Invalid value: false: Spec.Valid must be true

# Check that mutate was set to true by the mutating webhook
kubectl apply -f config/samples/passes.validation.webhook_v1_webhooktest.yaml
webhooktest.webhook.operators.coreos.io/webhooktest-sample created

# Check that Spec.Mutate is set to true:
kubectl get webhooktest webhooktest-sample -n webhook-operator-system -o yaml | yq read - spec.mutate
true
```

## Deploying the Webhook Operator with OLM

### Build a Bundle Image

```bash
$ make bundle-build BUNDLE_IMG=quay.io/agreene/webhook-operator-bundle:latest
$ docker push quay.io/agreene/webhook-operator-bundle:latest
```

### Build an Index

```bash
$ opm index add --bundles quay.io/agreene/webhook-operator-bundle:latest --tag quay.io/agreene/webhook-operator-index:latest -c docker
$ docker push quay.io/agreene/webhook-operator-index:latest
```

### Deploy with OLM on Vanilla Kubernetes

```bash
# Create the CatalogSource
$ kubectl apply -f olm/upstream/install/00_catsrc.yaml
catalogsource.operators.coreos.com/webhook-operator-catalog created

# Create a Subscription for the Operator
$ kubectl apply -f olm/upstream/install/01_sub.yaml
subscription.operators.coreos.com/webhook-operator-subscription created

# Check that the invalid webhookTest is rejected by the Validating webhook.
$ kubectl apply -f olm/upstream/example-crs/invalid.cr.yaml
Error from server (WebhookTest.test.operators.coreos.com "webhooktest-sample" is invalid: spec.schedule: Invalid value: false: Spec.Valid must be true): error when creating "olm/upstream/example-crs/invalid.cr.yaml": admission webhook "vwebhooktest.kb.io" denied the request: WebhookTest.test.operators.coreos.com "webhooktest-sample" is invalid: spec.schedule: Invalid value: false: Spec.Valid must be true

# Check that the valid webhookTest is approved by the Validating webhook
$ kubectl apply -f olm/upstream/example-crs/valid.cr.yaml
webhooktest.webhook.operators.coreos.io/webhooktest-sample created

# Check that the Conversion Webhook can serve v1 of the webhookTest CR and that the spec.mutate field is true
$ kubectl get webhooktests.v1.webhook.operators.coreos.io webhooktest-sample -n operators -o yaml | yq read - spec
mutate: true
valid: true

# Check that the Conversion Webhook can serve v2 of the webhookTest CR and that the spec.conversion.mutate field is true
$ kubectl get webhooktests.v2.webhook.operators.coreos.io webhooktest-sample -n operators -o yaml | yq read - spec
conversion:
  mutate: true
  valid: true
```

### Deploy with OLM on OpenShift

```bash
# Create the CatalogSource
$ kubectl apply -f olm/ocp/install/00_catsrc.yaml
catalogsource.operators.coreos.com/webhook-operator-catalog created

# Create a Subscription for the Operator
$ kubectl apply -f olm/ocp/install/01_sub.yaml
subscription.operators.coreos.com/webhook-operator-subscription created

# Check that the invalid webhookTest is rejected by the Validating webhook.
$ kubectl apply -f olm/ocp/example-crs/invalid.cr.yaml
Error from server (WebhookTest.test.operators.coreos.com "webhooktest-sample" is invalid: spec.schedule: Invalid value: false: Spec.Valid must be true): error when creating "olm/ocp/example-crs/invalid.cr.yaml": admission webhook "vwebhooktest.kb.io" denied the request: WebhookTest.test.operators.coreos.com "webhooktest-sample" is invalid: spec.schedule: Invalid value: false: Spec.Valid must be true

# Check that the valid webhookTest is approved by the Validating webhook
$ kubectl apply -f olm/ocp/example-crs/valid.cr.yaml
webhooktest.webhook.operators.coreos.io/webhooktest-sample created

# Check that the Conversion Webhook can serve v1 of the webhookTest CR and that the spec.mutate field is true
$ kubectl get webhooktests.v1.webhook.operators.coreos.io webhooktest-sample -n openshift-operators -o yaml | yq read - spec
mutate: true
valid: true

# Check that the Conversion Webhook can serve v2 of the webhookTest CR and that the spec.conversion.mutate field is true
$ kubectl get webhooktests.v2.webhook.operators.coreos.io webhooktest-sample -n openshift-operators -o yaml | yq read - spec
conversion:
  mutate: true
  valid: true
```
