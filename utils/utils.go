package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/fatih/color"
	"github.com/surbytes/gitusr/models"
	"gopkg.in/ini.v1"
)

func PrintInfo(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// global .gitconfigfile path
func globalConfigFile() string {
	home, err := os.UserHomeDir()
	CheckErr(err)
	return filepath.Join(home, ".gitconfig")
}

// load config file
func loadGlobalConfigFile() *ini.File {
	cfg, err := ini.Load(globalConfigFile())
	CheckErr(err)
	return cfg
}

// get users section keys
func getGlobalUsersKeys() []string {
	file, err := os.Open(globalConfigFile())
	CheckErr(err)
	defer file.Close()

	re := regexp.MustCompile(`\[users\s+"(.*?)"\]`)
	scanner := bufio.NewScanner(file)
	var usersKeys []string

	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindStringSubmatch(line)
		if len(match) > 1 {
			usersKeys = append(usersKeys, match[1])
		}
	}

	err = scanner.Err()
	CheckErr(err)

	return usersKeys
}

// prepare users section
func prepareUsers(usersKeys []string) []models.User {
	var users []models.User
	for _, v := range usersKeys {
		u := "users \"" + v + "\""
		section := loadGlobalConfigFile().Section(u)
		name, err := section.GetKey("name")
		CheckErr(err)
		email, err := section.GetKey("email")
		CheckErr(err)
		user := models.User{
			Name:  name.String(),
			Email: email.String(),
		}
		users = append(users, user)
	}

	return append(users, models.GetCurrentUsr())
}

// render users
func RenderUsers() {
	if len(os.Args) > 1 {
		models.SetUsr(os.Args[1])
	}

	for _, usr := range prepareUsers(getGlobalUsersKeys()) {
		if usr.Name == models.GetCurrentUsr().Name && usr.Email == models.GetCurrentUsr().Email {
			color.Yellow("%s <%s> *", usr.Name, usr.Email)
		} else {
			fmt.Printf("%s <%s>\n", usr.Name, usr.Email)
		}
	}
}

// handle error
func CheckErr(err error) {
	if err == nil {
		return
	}
	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}
