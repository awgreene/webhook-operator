# webhook-operator

This operator can be used to test validating, mutating, and conversion webhooks.

Built using [Kubebuilder](https://book.kubebuilder.io/)

## Install with Cert-Manager

The latest instructions can be found [here](https://cert-manager.io/docs/installation/kubernetes/).

```bash
# Kubernetes 1.16+
$ kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v1.0.1/cert-manager.yaml

# Kubernetes <1.16
$ kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v1.0.1/cert-manager-legacy.yaml
```

## Build and Deploy Webhook-Operator

```bash
# Build the operator Image
$ make docker-build IMG=quay.io/agreene/webhook-operator:latest

# Push the image to docker
$ docker push quay.io/agreene/webhook-operator:latest

# Install the operator
$ make deploy IMG=quay.io/agreene/webhook-operator:latest

# Check that the pods are up and running
$ watch oc get pods -n webhook-operator-system

# Try and create the resource that fails validation
$ oc apply -f config/samples/fails.validation.webhook_v1_webhooktest.yaml
Error from server (WebhookTest.test.operators.coreos.com "webhooktest-sample" is invalid: spec.schedule: Invalid value: false: Spec.Valid must be true): error when creating "config/samples/fails.validation.webhook_v1_webhooktest.yaml": admission webhook "vwebhooktest.kb.io" denied the request: WebhookTest.test.operators.coreos.com "webhooktest-sample" is invalid: spec.schedule: Invalid value: false: Spec.Valid must be true

# Check that mutate was set to true by the mutating webhook
oc apply -f config/samples/passes.validation.webhook_v1_webhooktest.yaml
webhooktest.webhook.operators.coreos.io/webhooktest-sample created

# Check that Spec.Mutate is set to true:
oc get webhooktest webhooktest-sample -n webhook-operator-system -o yaml | yq read - spec.mutate
true
```
