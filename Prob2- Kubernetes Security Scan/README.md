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
---

## 📌 How to Perform the Security Scan

### **1️⃣ Install Kubescape**
```sh
iwr -useb https://raw.githubusercontent.com/armosec/kubescape/master/install.ps1 | iex
kubescape version

### **2️⃣ Run Kubernetes Security Scan**
sh
Copy
Edit
kubescape scan cluster --format json --output results.json

3️⃣ Review & Submit Findings
sh
Copy
Edit
notepad results.json  # Windows
cat results.json      # Linux/macOS
4️⃣ Move Findings to Submission Folder
sh
Copy
Edit
Move-Item -Path "results.json" -Destination "D:\projects\Accuknox-kubernetes-asset\Prob2-Kubernetes-Security-Scan"
📌 Submission Instructions
✅ Ensure results.json is included in this folder.
✅ Upload all required files to GitHub using the following commands:

sh
Copy
Edit
git add .
git commit -m "Added Kubernetes security scan findings"
git push origin main
✅ Provide the GitHub link where the JSON file is stored.

✅ Project Completion Checklist
✔ Kubescape installed successfully
✔ Security scan executed, JSON report generated
✔ Findings reviewed & stored in results.json
✔ Report submitted as per assessment requirements

🚀 Now, your Kubernetes security assessment is complete! 💯

yaml
Copy
Edit

---

This **README.md** file is now **formatted for GitHub** and **includes images, instructions, and submission steps**.  

📌 **Next Step:** Copy-paste this into `README.md` inside `Prob2-Kubernetes-Security-Scan/` and push it to GitHub! 🚀  

Let me know if you need any modifications! 🔥
