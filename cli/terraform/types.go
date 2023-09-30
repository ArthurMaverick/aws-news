package terraform

type Backend struct {
	Bucket string
	Region string
	Key    string
}

type Terraform struct {
	Backend Backend
}
