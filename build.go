//go:build mage
// +build mage

/*
 */
package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type (
	Proto     mg.Namespace
	Migration mg.Namespace
)

func (Proto) Generate() error {
	pwd, err := sh.Output("pwd")
	if err != nil {
		return err
	}
	user, err := userIDAndGroupID()
	if err != nil {
		return err
	}
	return sh.RunV("docker", "run", "--rm", "--user", user, "-v", fmt.Sprintf("%s:/defs", pwd), "namely/protoc-all:1.32_2", "-d", "/defs/proto", "-l", "go", "-o", "./proto/go")
}

func (Migration) Create(name string) error {
	pwd, err := sh.Output("pwd")
	if err != nil {
		return err
	}
	user, err := userIDAndGroupID()
	if err != nil {
		return err
	}
	migrationsDir := fmt.Sprintf("%s/database/migrations:/migrations", pwd)
	return sh.RunV("docker", "run", "--user", user, "-v", migrationsDir, "--network", "host", "migrate/migrate", "create", "-ext", "sql", "-dir", "/migrations", name)
}

func (Migration) Up() error {
	pwd, err := sh.Output("pwd")
	if err != nil {
		return err
	}
	databaseURI := devDatabaseURI()
	migrationsDir := fmt.Sprintf("%s/database/migrations:/migrations", pwd)
	return sh.RunV("docker", "run", "-v", migrationsDir, "--network", "host", "migrate/migrate", "-path", "/migrations", "-database", databaseURI, "up")
}

func devDatabaseURI() string {
	return "postgres://hypatia:hypatia@127.0.0.1:5432/library?sslmode=disable"
}

func userIDAndGroupID() (string, error) {
	uid, err := sh.Output("id", "-u")
	if err != nil {
		return "", err
	}
	gid, err := sh.Output("id", "-g")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s", uid, gid), nil
}
