package main

import (
	"fmt"
	appAdmin "geektime-gocamp/week4/homework/internal/app/admin"
	"geektime-gocamp/week4/homework/internal/pkg/app"
	"os"

	"github.com/BurntSushi/toml"
)

func main() {
	configPath := "./week4/homework/cmd/admin/config.ini"
	var config config
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		fmt.Printf("Invalid config, path=[%+v], error=[%+v]\n", configPath, err)
		os.Exit(-1)
	}
	adminServer := appAdmin.NewServer(
		appAdmin.WebAddress(config.Web.Address),
		appAdmin.DataSourceName(config.MySQL.DataSourceName),
	)
	a := app.New(
		app.Name("student.admin"),
		app.Version("1.0.0"),
		adminServer,
	)
	if err := a.Run(); err != nil {
		fmt.Printf("Exit with error=[%+v]\n", err)
		os.Exit(-1)
	}
	fmt.Printf("Finished")
}
