# 🚀 Kubernetes Security Scan

## 📌 Overview
This folder contains the **Kubernetes Security Scan findings**, generated using **Kubescape** to assess security vulnerabilities in the cluster.

📄 **[View Security Scan Report](#)**

---

## 📌 Security Scan Findings - UI Preview

| **Cluster Overview** | **Vulnerabilities Found** |
|----------------------|-------------------------|
| ![Cluster Security Scan](https://github.com/user-attachments/assets/2ed1066a-c44a-44a0-a692-9f0315914e61) | ![NSA scan image](https://github.com/user-attachments/assets/bb54097b-7377-4102-a672-606638ca7e1d) |

| **Compliance Report** |
|----------------------|
| ![MITRE ATT&CK-based Scan](https://github.com/user-attachments/assets/5b2d4c2b-2870-4bf9-ac07-a2304546643e) |

---

## 📂 Files Included
- `results.json` - JSON file containing Kubernetes security scan findings.
- `Scan-Screenshots/` - Images displaying scan results.
- `Notion/` - Link to the **detailed security analysis report**.

---

## 📌 How to Perform the Security Scan

### **1️⃣ Install Kubescape**
```sh
iwr -useb https://raw.githubusercontent.com/armosec/kubescape/master/install.ps1 | iex
kubescape version
