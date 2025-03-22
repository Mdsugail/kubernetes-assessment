# 🚀 GoLang Kubernetes Deployment

## 📌 Overview
This folder contains the **GoLang Web Application**, containerized using **Docker**, deployed on **Kubernetes**, and exposed to the internet.

📄 **[View Live Deployment](#)** (Replace with actual public URL)

---

## 📌 Deployment Architecture

| **Kubernetes Deployment** | **Service Exposure (Ngrok)** |
|---------------------------|-----------------------------|
| ![Kubernetes Deployment](https://github.com/user-attachments/assets/c774638c-ede2-417f-9821-82fcfd97f464) | ![Ngrok Deployment](https://github.com/user-attachments/assets/a02d4486-4eae-4a7d-8719-7136dbbe747b) |

| **App Running Locally** |
|-------------------------|
| ![Local Output](https://github.com/user-attachments/assets/8af055c7-39a1-44c1-b131-6a1606ce5227) |

---

## 📂 Files Included
- `main.go` - GoLang Web Server Code.
- `Dockerfile` - Containerization Configuration.
- `deployment.yaml` - Kubernetes Deployment Configuration.
- `service.yaml` - Kubernetes Service Configuration.
- `Images` - Screenshots.

---

## 📌 How to Build & Deploy the GoLang App

### **1️⃣ Create the GoLang Web App**
Create `main.go` and add the following code:

```go
package main

import (
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Current Time: %s", time.Now().Format(time.RFC1123))
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

```

#### **2️⃣ Create Dockerfile**
```dockerfile
FROM golang:1.19
WORKDIR /app
COPY main.go .
RUN go build -o server main.go
CMD ["/app/server"]
```
```sh
docker build -t timeapp .
docker run -p 8080:8080 timeapp
```

#### **3️⃣ Deploy to Kubernetes (`deployment.yaml`)**
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: golang-app
spec:
  replicas: 2
  selector:
    matchLabels:
      app: golang-app
  template:
    metadata:
      labels:
        app: golang-app
    spec:
      containers:
      - name: golang-app
        image: your-dockerhub-username/timeapp:latest
        ports:
        - containerPort: 8080
```
```sh
kubectl apply -f deployment.yaml
kubectl get pods
```

#### **4️⃣ Expose App Publicly (`service.yaml`)**
```yaml
apiVersion: v1
kind: Service
metadata:
  name: golang-app-service
spec:
  selector:
    app: golang-app
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080
  type: LoadBalancer
```
```sh
kubectl apply -f service.yaml
minikube service golang-app-service --url
```

#### **5️⃣ Make It Public Using Ngrok**
```sh
ingrok http 32382  # Replace with actual port from Minikube
```

### 📌 **Deliverables**
📂 `main.go` - GoLang Web App Source Code.  
📂 `Dockerfile` - Container Configuration.  
📂 `deployment.yaml` - Kubernetes Deployment File.  
📂 `service.yaml` - Kubernetes Service File.  
📂 `Ngrok URL` - Publicly accessible link.  

---

# 📌 **GitHub Submission Guide**

### **1️⃣ Upload Everything to GitHub**
```sh
git add .
git commit -m "Added all Kubernetes assessment files"
git push origin main
```

🚀 Now, your Golang Kubernetes Deployment is complete! 💯
---
