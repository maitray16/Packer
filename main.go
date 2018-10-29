package main

import (
	"fmt"
	"os"
	"os/exec"
	"log"
	"strings"
	"github.com/spf13/viper"
)

func UI(project_name string){
	
	os.MkdirAll(project_name, os.ModePerm);
	os.Chdir(project_name)
	cmd:= exec.Command("create-react-app", strings.ToLower(project_name)+"_ui")
	out, err := cmd.Output()
		if err != nil {
    	log.Fatal(err)
	} else {
    	fmt.Printf("%s", out);
	}
    os.Chdir(strings.ToLower(project_name)+"_ui")



}

func main() {

	viper.SetConfigName("packer") 
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file not found...")
	} else {
		project_name := viper.GetString("meta.name")
		project_type := viper.GetString("meta.type")

		if strings.TrimRight(project_type, "\n") == "UI" { 
			ui_framework := viper.GetString("UI.framework")
			ui_library := viper.GetString("UI.library")
		}

	
		fmt.Printf("\nCreating project :\n Name = %s\n Type = %s\n",
			project_name,
			project_type)

		if strings.TrimRight(ui_framework, "\n") == "ReactJS" {
			UI(project_name, project_type)
		}
		
		
	}
}
