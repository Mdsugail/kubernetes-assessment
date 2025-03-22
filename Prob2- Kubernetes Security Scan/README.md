# 🚀 Kubernetes Security Scan

## 📌 Overview
This folder contains the **Kubernetes Security Scan findings**, generated using **Kubescape** to assess security vulnerabilities in the cluster.

📄 **[View Security Scan Report](#)**

---

## 📌 Security Scan Findings - UI Preview

| **Cluster Overview** | **NSA Overview** |
|----------------------|-------------------------|
| ![Cluster Security Scan](https://github.com/user-attachments/assets/2ed1066a-c44a-44a0-a692-9f0315914e61) | ![NSA scan image](https://github.com/user-attachments/assets/bb54097b-7377-4102-a672-606638ca7e1d) |

| **MITRE ATT&CK Report** |
|----------------------|
| ![MITRE ATT&CK-based Scan](https://github.com/user-attachments/assets/5b2d4c2b-2870-4bf9-ac07-a2304546643e) |

---

## 📂 Files Included
- `results.json` - JSON file containing Kubernetes security scan findings.
- `Scan-Screenshots/` - Images displaying scan results.
---

## 📌 How to Perform the Security Scan

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
```sh
Copy
Edit
git add .
git commit -m "Added Kubernetes security scan findings"
git push origin main#### **1️⃣ Set Up Kubernetes Cluster (Using Minikube)**
```sh
minikube start
kubectl get nodes
```

✅ Provide the GitHub link where the JSON file is stored.

✅ Project Completion Checklist
✔ Kubescape installed successfully
✔ Security scan executed, JSON report generated
✔ Findings reviewed & stored in results.json
✔ Report submitted as per assessment requirements

🚀 Now, your Kubernetes security is complete! 💯

