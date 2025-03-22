# 🚀 **Kubernetes Security & Deployment Project - README**

Welcome to the **Kubernetes Security & Deployment Project!** This repository documents the step-by-step implementation of Kubernetes security scanning, GoLang application deployment, and product design documentation.

---

## 📌 **Project Overview**
This project consists of **three problem statements**, covering various aspects of **Kubernetes security, DevOps, and cloud-native application deployment**:

1️⃣ **Product Requirement Document (PRD) & Wireframes** - Defining and designing a **Container Image Vulnerability Scanner**.
2️⃣ **Kubernetes Security Scan** - Scanning Kubernetes clusters for security vulnerabilities using **Kubescape**.
3️⃣ **GoLang Kubernetes Deployment** - Deploying a **GoLang web application** on Kubernetes and exposing it publicly.

Each section provides a **detailed step-by-step implementation guide** with necessary commands and configurations.

---

# 1️⃣ **Product Requirements Document (PRD) & Wireframes**

### 📌 **Objective**
- Define the **product requirements** for a **Container Image Vulnerability Scanner**.
- Create **wireframes** to visualize the product’s user interface.

### 📌 **Key Features**
✅ **Container Image Scanning** (Supports tools like Trivy, Clair, etc.)  
✅ **Severity-Based Filtering** (Categorization of vulnerabilities)  
✅ **Detailed Reports & Fix Suggestions**  
✅ **CI/CD Pipeline Integration**  
✅ **Role-Based Access Control (RBAC)**  

### 📌 **Deliverables**
📂 `PRD.md` - A structured document defining features, workflows, and goals.  
📂 `Wireframes/` - Folder containing UI mockups and design elements.  

---

# 2️⃣ **Kubernetes Security Scan**

### 📌 **Objective**
- Install and configure a **local Kubernetes cluster**.
- Perform a **security scan** using **Kubescape**.
- Generate a **JSON file containing security findings**.

### 📌 **Step-by-Step Implementation**

#### **1️⃣ Set Up Kubernetes Cluster (Using Minikube)**
```sh
minikube start
kubectl get nodes
```

#### **2️⃣ Install Kubescape**
```sh
iwr -useb https://raw.githubusercontent.com/armosec/kubescape/master/install.ps1 | iex
kubescape version
```

#### **3️⃣ Run Kubernetes Security Scan**
```sh
kubescape scan cluster --format json --output results.json
```

#### **4️⃣ Review and Submit Findings**
```sh
notepad results.json  # View findings
Move-Item results.json D:\projects\Accuknox-kubernetes-asset\Prob2-Kubernetes-Security-Scan\
```

### 📌 **Deliverables**
📂 `results.json` - JSON file with security scan findings.  

---

# 3️⃣ **GoLang Kubernetes Deployment**

### 📌 **Objective**
- Develop, containerize, and deploy a **GoLang web application**.
- Deploy it on **Kubernetes with 2 replicas**.
- Expose the service **publicly using Ngrok**.

### 📌 **Step-by-Step Implementation**

#### **1️⃣ Create GoLang Web App (`main.go`)**
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

### **2️⃣ Final Submission**
Submit your GitHub repository link:  
📎 **`https://github.com/your-username/kubernetes-assessment`**

---

# ✅ **Project Completion Checklist**
✔ **Problem 1:** PRD Document & Wireframes Uploaded  
✔ **Problem 2:** Kubernetes Security Scan & `results.json` Generated  
✔ **Problem 3:** GoLang Kubernetes Deployment Completed & Public URL Available  
✔ **GitHub Repository Submitted**  

---

# 🚀 **Final Thoughts**
This README serves as a comprehensive guide to implementing Kubernetes security scans, GoLang deployments, and containerized workflows. By following this structured documentation, you ensure a **successful project execution and submission**! 🎯🔥

Let me know if you need any modifications! 🚀

