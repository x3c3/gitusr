package models

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/surbytes/gitusr/utils"
)

type User struct {
	Name  string
	Email string
}

// add new user into git config file
func (usr *User) AddUsr(name string, email string) *User {
	utils.PrintInfo("Adding user: %s <%s>", name, email)
	if err := exec.Command("git", "config", "--global", "--add", "users.email", usr.Email).Run(); err != nil {
		utils.CheckErr(err)
	}

	if err := exec.Command("git", "config", "--global", "--add", "users.name", usr.Name).Run(); err != nil {
		utils.CheckErr(err)
	}

	return &User{
		Name:  name,
		Email: email,
	}
}

// get the current user on git config file
func GetCurrentUsr() {

	//var email strings.Builder
	//var name strings.Builder
	name, err := exec.Command("git", "config", "--global", "user.name").Output()
	utils.CheckErr(err)

	email, err := exec.Command("git", "config", "--global", "user.email").Output()
	utils.CheckErr(err)

	nemail := strings.Trim(string(email), "\f\t\r\n ")
	nname := strings.Trim(string(name), "\f\t\r\n ")

	utils.PrintInfo("Current user: %s <%s>", nname, nemail)

}

// delete user from git config file
func (usr *User) DelUsr(name string, email string) {

	utils.PrintInfo("Deleting user: %s <%s>", name, email)
	if err := exec.Command("git", "config", "--global", "--unset-all", "users.email", usr.Email).Run(); err != nil {
		utils.CheckErr(err)
	}

	if err := exec.Command("git", "config", "--global", "--unset-all", "users.name", usr.Name).Run(); err != nil {
		utils.CheckErr(err)
	}
}

// list users
func listusrs() {
	utils.PrintInfo("Listing users: ")
	usrsemail, err := exec.Command("git", "config", "--global", "--get-all", "users.email").Output()
	utils.CheckErr(err)
	usrsname, err := exec.Command("git", "config", "--global", "--get-all", "users.name").Output()
	utils.CheckErr(err)

	utils.PrintInfo("%s", usrsemail)
	utils.PrintInfo("%s", usrsname)
	fmt.Println(usrsemail)
}
