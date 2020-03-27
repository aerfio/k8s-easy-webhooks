# What this project is

<https://github.com/aerfio/k8s-easy-webhooks/issues/1>

# How to check how it works

Install:

```bash
kubectl apply -f
```

Wait a bit for it to start, look through logs till it stops reconciling etc. To observe defaulting webhook run:

```
kubectl apply -f
```

And to observe validating webhook, which will deny applying incorrect CR run this command:

```
kubectl apply -f
```
