# gitusr
*gitusr* is a fast and simple CLI tool built in **Go** that allows you to effortlessly switch between multiple global Git users.  

This tool was inspired by [kubectx](https://github.com/ahmetb/kubectx) and [kubens](https://github.com/ahmetb/kubectx).


![terminal example](https://github.com/surbytes/gitusr/raw/refs/heads/main/gitusrgif.gif)

# Installation
You can install using the go command:  

```
go install github.com/surbytes/gitusr
```

or clone the project

```
git clone https://github.com/surbytes/gitusr.git
cd gitusr; go build -o gitusr main.go
./gitusr
```
# Configuration
To ensure *gitusr* works correctly, users should configure their .gitconfig file so that each `[users "<name>"]` section represents a global Git user, as follows:  

![image](https://github.com/user-attachments/assets/9f6f073d-acaa-4b96-bcad-5dc40423ab5b)
