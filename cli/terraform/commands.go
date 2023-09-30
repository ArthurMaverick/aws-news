package terraform

import (
	"bytes"
	"fmt"
	"os/exec"
)

func NewTFCmd(bucket, key, region string) *Terraform {
	if bucket == "" || key == "" || region == "" {
		fmt.Println("Bucket, key or region is empty")
		return nil
	}

	return &Terraform{
		Backend: Backend{
			Bucket: bucket,
			Region: region,
			Key:    key,
		},
	}
}

func (t *Terraform) TFInit() {
	stateBucket := fmt.Sprintf("-backend-config=bucket=%s", t.Backend.Bucket)
	stateKey := fmt.Sprintf("-backend-config=key=%s", t.Backend.Key)
	stateRegion := fmt.Sprintf("-backend-config=region=%s", t.Backend.Region)

	cmd := exec.Command("terraform", "init", stateBucket, stateRegion, stateKey, "-reconfigure")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(out.String())
}

func (t *Terraform) TFPlan() {
	cmd := exec.Command("terraform", "plan", "-out=tfplan")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out.String())
}

func (t *Terraform) TFApply() {
	cmd := exec.Command("terraform", "apply", "tfplan", "-auto-approve")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out.String())
}

func (t *Terraform) TFDestroy() {
	cmd := exec.Command("terraform", "destroy", "-auto-approve")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out.String())
}

func (t *Terraform) TFFmt() {
	cmd := exec.Command("terraform", "fmt")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out.String())
}
