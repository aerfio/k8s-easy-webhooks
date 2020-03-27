# What this project is

<https://github.com/aerfio/k8s-easy-webhooks/issues/1>

# How to check how it works

Install:

```bash
kubectl apply -f https://raw.githubusercontent.com/aerfio/k8s-easy-webhooks/master/webhook.yaml
```

Wait a bit for it to start, look through logs till it stops reconciling etc. To observe defaulting webhook run:

```bash
kubectl apply -f https://raw.githubusercontent.com/aerfio/k8s-easy-webhooks/master/config/samples/test_v1alpha1_tester_defaulted.yaml
```

And to observe validating webhook, which will deny applying incorrect CR run this command:

```bash
kubectl apply -f https://raw.githubusercontent.com/aerfio/k8s-easy-webhooks/master/config/samples/test_v1alpha1_tester_incorrect.yaml
```
