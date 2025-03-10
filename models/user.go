package models

import (
	"os/exec"
	"strings"

	"github.com/surbytes/gitusr/utils"
)

type User struct {
	Name  string
	Email string
}

// add new user into git config file
func (usr *User) AddUsr(name string, email string) (*User, error) {
	utils.PrintInfo("Adding user: %s <%s>", name, email)
	sectionName := "users." + usr.Name + ".name"
	sectionEmail := "users." + usr.Email + ".email"
	if err := exec.Command("git", "config", "--global", "--add", sectionEmail, usr.Email).Run(); err != nil {
		//utils.CheckErr(err)
		return &User{}, err
	}

	if err := exec.Command("git", "config", "--global", "--add", sectionName, usr.Name).Run(); err != nil {
		//utils.CheckErr(err)
		return &User{}, err
	}

	return &User{
		Name:  name,
		Email: email,
	}, nil
}

func SetUsr(name string, email string) {
	/*
		if err := exec.Command("git", "config", "--global", "", sectionName, usr.Name).Run(); err != nil {

		}*/
}

// get the current user on git config file
func GetCurrentUsr() User {

	//var email strings.Builder
	//var name strings.Builder
	name, err := exec.Command("git", "config", "--global", "user.name").Output()
	utils.CheckErr(err)

	email, err := exec.Command("git", "config", "--global", "user.email").Output()
	utils.CheckErr(err)

	nemail := strings.Trim(string(email), "\f\t\r\n ")
	nname := strings.Trim(string(name), "\f\t\r\n ")

	return User{
		Name:  nname,
		Email: nemail,
	}

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
