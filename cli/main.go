package main

import (
	clients "awsnews/aws"
	"awsnews/terraform"
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type AwsNewsConfig struct {
	Webhook   string `yaml:"webhook"`
	AvatarUrl string `yaml:"avatar_url"`
	Color     string `yaml:"color"`
	Bucket    string `yaml:"bucket"`
	Region    string `yaml:"region"`
	Key       string `yaml:"key"`
}

func main() {
	var command string

	flag.StringVar(&command, "command", "", "Command to execute")
	flag.Parse()

	switch command {
	case "deploy":
		Deploy(ReadYamlFile())
	case "update":
		Update(ReadYamlFile())
	case "destroy":
		Destroy(ReadYamlFile())
	default:
		fmt.Println("Command not found")
	}

	fmt.Println(ReadYamlFile()["discord"].Webhook)
}

// ReadYamlFile will read the yaml file
func ReadYamlFile() map[string]AwsNewsConfig {
	yfile, err := os.ReadFile(".aws-news-config.yaml")
	if err != nil {
		fmt.Println(err)
	}

	data := make(map[string]AwsNewsConfig)

	err2 := yaml.Unmarshal(yfile, &data)
	if err2 != nil {
		fmt.Println(err2)
	}

	return data
}

// Deploy will deploy the terraform code
func Deploy(yfile map[string]AwsNewsConfig) {
	webhookPayload := clients.NewServiceSSM()
	_, err := webhookPayload.PutParameter(yfile["discord"].Webhook, "/AWS_NEWS/DISCORD/WEBHOOK")
	if err != nil {
		fmt.Println(err.Error())
	}

	AvatarUrlPayload := clients.NewServiceSSM()
	_, err = AvatarUrlPayload.PutParameter(yfile["discord"].AvatarUrl, "/AWS_NEWS/DISCORD/AVATAR_URL")
	if err != nil {
		fmt.Println(err.Error())
	}

	ColorPayload := clients.NewServiceSSM()
	_, err = ColorPayload.PutParameter(yfile["discord"].Color, "/AWS_NEWS/DISCORD/COLOR")
	if err != nil {
		fmt.Println(err.Error())
	}

	s3 := clients.NewServiceS3()
	_, err = s3.GetBucket(yfile["terraform"].Bucket)
	if err != nil {
		fmt.Println(err.Error())
	}

	cmd := terraform.NewTFCmd(yfile["terraform"].Bucket, yfile["terraform"].Key, yfile["terraform"].Region)
	cmd.TFInit()
	cmd.TFFmt()
	cmd.TFPlan()
	cmd.TFApply()
}

// Update will update the terraform code or lambda function
func Update(yfile map[string]AwsNewsConfig) {
	cmd := terraform.NewTFCmd(yfile["terraform"].Bucket, yfile["terraform"].Key, yfile["terraform"].Region)
	cmd.TFInit()
	cmd.TFFmt()
	cmd.TFPlan()
	cmd.TFApply()
}

// Destroy will destroy the terraform code
func Destroy(yfile map[string]AwsNewsConfig) {
	cmd := terraform.NewTFCmd(yfile["terraform"].Bucket, yfile["terraform"].Key, yfile["terraform"].Region)
	cmd.TFInit()
	cmd.TFFmt()
	cmd.TFPlan()
	cmd.TFDestroy()
}
