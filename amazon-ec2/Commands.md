## Amazon EC2 (Elastic Compute Cloud) instances can be managed and interacted with using the AWS Command Line Interface (CLI) or various SDKs (Software Development Kits). Here are some common AWS CLI commands for managing EC2 instances

### Listing Instances => aws ec2 describe-instances

### Creating an EC2 Instance => aws ec2 run-instances --image-id <my-ami-id> --instance-type <instance-type> --key-name <my-key-pair-name>

### Terminating an EC2 Instance => aws ec2 terminate-instances --instance-ids <instance-id>

### Start and Stop an EC2 Instance => aws ec2 start-instances --instance-ids <instance-id>
###                                   aws ec2 stop-instances --instance-ids <instance-id>

### Rebooting an EC2 Instance => aws ec2 reboot-instances --instance-ids <instance-id>

### Creating a Security Group => aws ec2 create-security-group --group-name <security-group-name> --description "My security group"
