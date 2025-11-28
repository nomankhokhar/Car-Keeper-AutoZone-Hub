# How to Deploy Traefik with Kubernetes

To install Traefik (v3) directly using `kubectl` without Helm, you need to manually apply the Custom Resource Definitions (CRDs), Role-Based Access Control (RBAC), and the Deployment/Service manifests.

[Image of Kubernetes Ingress Controller architecture]

Here is the step-by-step procedure to get a standard instance running.

### 1\. Create the Namespace

First, create a dedicated namespace to keep things organized.

```bash
kubectl create namespace traefik
```

### 2\. Install Traefik CRDs

Traefik requires Custom Resource Definitions (like `IngressRoute`, `Middleware`) to function correctly. Apply these directly from the official Traefik repository.

```bash
kubectl apply -f https://raw.githubusercontent.com/traefik/traefik/v3.1/docs/content/reference/dynamic-configuration/kubernetes-crd-definition-v1.yml
```

### 3\. Create RBAC (Permissions)

Traefik needs permission to watch Services, Endpoints, and Ingresses in your cluster. Create a file named `traefik-rbac.yaml` (or copy the block below) and apply it.

**`traefik-rbac.yaml`**

```yaml
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: traefik-role
rules:
  - apiGroups:
      - ""
    resources:
      - services
      - endpoints
      - secret
      - nodes  # <--- Added to allow Traefik to access node information
    verbs:
      - get
      - list
      - watch
  - apiGroups:
      - extensions
      - networking.k8s.io
    resources:
      - ingresses
      - ingressclasses
    verbs:
      - get
      - list
      - watch
  - apiGroups:
      - extensions
      - networking.k8s.io
    resources:
      - ingresses/status
    verbs:
      - update
  - apiGroups:
      - traefik.io
    resources:
      - middlewares
      - middlewaretcps
      - ingressroutes
      - traefiktlsoptions
      - ingressroutetcps
      - traefikservices
      - ingressrouteudps
      - tlsoptions
      - tlsstores
      - serverstransports
    verbs:
      - get
      - list
      - watch

---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: traefik-role-binding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: traefik-role
subjects:
  - kind: ServiceAccount
    name: traefik-account
    namespace: traefik

---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: traefik-account
  namespace: default
```

**Apply it:**

```bash
kubectl apply -f traefik-rbac.yaml
```

### 4\. Deploy Traefik Service and Deployment

This manifest creates the LoadBalancer service (to expose it to the internet) and the Deployment (the actual Traefik pod).

Create a file named `traefik-deployment.yaml`:

**`traefik-deployment.yaml`**

```yaml
apiVersion: v1
kind: Service
metadata:
  name: traefik
  namespace: traefik
spec:
  ports:
    - protocol: TCP
      name: web
      port: 80
    - protocol: TCP
      name: websecure
      port: 443
    - protocol: TCP
      name: admin
      port: 8080
  selector:
    app: traefik
  type: LoadBalancer # Change to NodePort or ClusterIP if needed
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: traefik
  namespace: traefik
  labels:
    app: traefik
spec:
  replicas: 1
  selector:
    matchLabels:
      app: traefik
  template:
    metadata:
      labels:
        app: traefik
    spec:
      serviceAccountName: traefik-account
      containers:
        - name: traefik
          image: traefik:v3.1
          args:
            - "--api.insecure=true" # Enable Dashboard (Don't use insecure in Prod)
            - "--providers.kubernetescrd" # Enable CRD provider
            - "--entrypoints.web.address=:80"
            - "--entrypoints.websecure.address=:443"
          ports:
            - name: web
              containerPort: 80
            - name: websecure
              containerPort: 443
            - name: admin
              containerPort: 8080
```

**Apply it:**

```bash
kubectl apply -f traefik-deployment.yaml
```

-----

### Verification

1.  **Check Pod Status:**

    ```bash
    kubectl get pods -n traefik
    ```

2.  **Get Public IP:**

    ```bash
    kubectl get svc -n traefik
    ```

    *(Look for the EXTERNAL-IP under the `traefik` service)*.

3.  **Access Dashboard:**
    You can now access the Traefik dashboard by port-forwarding (since we didn't expose port 8080 via the LoadBalancer for security, usually):

    ```bash
    kubectl port-forward $(kubectl get pods --selector "app=traefik" --output=name -n traefik) -n traefik 9000:8080
    ```

    Open your browser to: `http://localhost:9000/dashboard/`
