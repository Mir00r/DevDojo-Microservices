Here’s a simple and easy-to-understand `README.md` file for your `init_go_service_structure.sh` script.

---

# **Go Microservice Project Initializer**

This script initializes a **Golang microservice project** with a well-organized folder structure, sample files, and Go module setup. It also installs commonly used dependencies and dynamically creates files tailored to your service name.

---

## **Features**

- Automatically creates a **standard microservice folder structure**.
- Dynamically names files based on the service name.
- Installs common Go libraries such as `gin`, `gorm`, `viper`, etc.
- Sets up boilerplate code for configuration, routes, middleware, and more.
- Supports **cross-platform paths** (Linux, macOS, Windows).
- Provides a single entry point (`main.go`) with a working server setup.

---

## **Usage**

### **1. Prerequisites**

- Ensure you have the following installed:
    - **Go (Golang)** (version 1.18 or higher recommended)
    - **Git Bash** (on Windows, or a terminal on macOS/Linux)

- Make the script executable:
  ```bash
  chmod +x init_go_service_structure.sh
  ```

---

### **2. Running the Script**

Use the following command to run the script:

```bash
./init_go_service_structure.sh <service-name> <location> <github-username>
```

#### **Parameters**:
1. `<service-name>`: The name of your microservice (e.g., `auth`, `user`).
2. `<location>`: The absolute or relative path where the project should be created.
3. `<github-username>`: Your GitHub username to set up the Go module.

---

### **3. Example**

To create a project called `auth-service` in `/home/user/projects` with the GitHub username `Mir00r`:

```bash
./init_go_service_structure.sh auth /home/user/projects Mir00r
```

On **Windows**, you can use paths like:
```bash
./init_go_service_structure.sh user /c/Users/YourName/Desktop/projects Mir00r
```

---

### **4. After Running the Script**

The script will:
1. Create a project folder structure like this:
   ```
   <service-name>-service/
   ├── cmd/
   ├── config/
   ├── constants/
   ├── containers/
   ├── db/
   ├── errors/
   ├── internal/
   │   ├── api/
   │   ├── services/
   │   ├── models/
   │   ├── repositories/
   ├── middlewares/
   ├── routes/
   ├── scripts/
   ├── test/
   ├── utils/
   ├── build/
   ├── docs/
   └── logs/
   ```

2. Populate the project with:
    - A working `main.go` file in `cmd/`.
    - Sample configuration, middleware, and database files.
    - A ready-to-use Go module (`go.mod`).

3. Install common dependencies like:
    - `gin` (for HTTP routing)
    - `gorm` (for ORM)
    - `viper` (for configuration management)
    - `logrus` (for logging)
    - `jwt-go` (for JWT handling)

---

## **Troubleshooting**

### **1. Permission Denied**
If you see:
```
mkdir: cannot create directory ... Permission denied
```
It means the script does not have permission to create directories in the specified location. Try:
- Running with `sudo` (Linux/macOS):
  ```bash
  sudo ./init_go_service_structure.sh ...
  ```
- Choosing a directory where your user has write permissions (e.g., inside your home directory).

### **2. Incorrect Path on Windows**
If the project is created in the wrong location (e.g., `/c/Users/...` instead of `/Desktop/...`), provide the **full absolute path** for `<location>`:
```bash
/c/Users/YourName/Desktop/projects
```

### **3. Missing Go Installation**
Ensure Go is installed and available in your `PATH`:
```bash
go version
```
If not, download it from the [official Go website](https://go.dev/dl/).

---

## **Feedback**

If you encounter issues or have suggestions, feel free to create an issue or contribute to improving the script!

Happy coding! 🚀
