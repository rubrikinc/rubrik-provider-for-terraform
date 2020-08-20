
# Quick Start Guide: Terraform Provider for Rubrik

## Introduction to the Terraform Provider for Rubrik

Rubrik's API-first architecture enables organizations to embrace and integrate Rubrik functionality into their existing automation processes. While Rubrik APIs can be consumed natively, companies are at various stages in their automation journey with different levels of automation knowledge on staff. The Rubrik Terraform Provider transforms the Rubrik RESTful API functionality into easy-to-consume Terraform configuration, which eliminates the need to understand how to consume raw Rubrik APIs and extends upon one of Rubrik's main design centers - simplicity.

## Installation

Requirements: Terraform has been successfully [installed](https://learn.hashicorp.com/terraform/getting-started/install.html).

1. Download the latest compiled binary from [GitHub releases](https://github.com/rubrikinc/rubrik-provider-for-terraform/releases).

   ```
   macOS: terraform-provider-rubrik-darwin-amd64
   Linux: terraform-provider-rubrik-linux-amd64
   Windows: terraform-provider-rubrik-windows-amd64.exe
   ```

2. Move the Rubrik provider into the correct Terraform plugin directory
   
   ```
   macOS: ~/.terraform.d/plugins/darwin_amd64
   Linux: ~/.terraform.d/plugins/linux_amd64
   Windows: %APPDATA%\terraform.d\plugins\windows_amd64
   ```
   
   _You may need to manually create the `plugin` directory._

3. Rename the the Rubrik provder to `terraform-provider-rubrik`

4. On Linux and MacOS ensure that the binary has the appropriate permissions by running `chmod 744 terraform-provider-rubrik`

5. Run `terraform init` in the directory that contains your Terraform configuration fiile (`main.tf`)

## Authentication

The Rubrik provider offers a flexible means of providing credentials for
authentication. The following methods are supported, in this order, and
explained below:

- Environment variables
- Static credentials

### Environment variables

Storing credentials in environment variables is a more secure process than storing them in your source code, and it ensures that your credentials are not accidentally shared if your code is uploaded to an internal or public version control system such as GitHub. 

* **RUBRIK_CDM_NODE_IP** (Contains the IP/FQDN of a Rubrik node)
* **RUBRIK_CDM_USERNAME** (Contains a username with configured access to the Rubrik cluster)
* **RUBRIK_CDM_PASSWORD** (Contains the password for the above user).



```hcl
provider "rubrik" {}
```

#### Setting Environment Variables in Microsoft Windows

For Microsoft Windows-based operating systems, the environment variables can be set utilizing the setx command as follows:

```
setx RUBRIK_CDM_NODE_IP "192.168.0.100"
setx RUBRIK_CDM_USERNAME "user@domain.com"
setx RUBRIK_CDM_PASSWORD "SecretPassword"
```

Run set without any other parameters to view current environment variables. Using setx saves the environment variables permanently, and the variables defined in the current shell will not be available until a new shell is opened. Using set instead of setx will define variables in the current shell session, but they will not be saved between sessions.

#### Setting Environment Variables in macOS and \*nix

For macOS and \*nix based operating systems the environment variables can be set utilizing the export command as follows:

```
export RUBRIK_CDM_NODE_IP=192.168.0.100
export RUBRIK_CDM_USERNAME=user@domain.com
export RUBRIK_CDM_PASSWORD=SecretPassword
```

Run export without any other parameters to view current environment variables. In order for the environment variables to persist across terminal sessions, add the above three export commands to the `~\.bash_profile` or `~\.profile` file.

### Static credentials 

Static credentials can be provided by adding a `node_ip`, `username` and `password` in-line in the
Rubrik provider block:

Usage:

```hcl
provider "rubrik" {
  node_ip     = "10.255.41.201"
  username    = "admin"
  password    = "RubrikTFDemo2019"
}
```

## Sample Syntax

This section provides sample syntax to help you get started. For additional information and examples, see the [Rubrik Provider for Terraform Documentation](https://rubrik.gitbook.io/terraform-provider-for-rubrik/). 


### Cluster Timezone Configuration

The following demonstrates configuring the time zone on a Rubrik cluster: 


```
provider "rubrik" {}

resource "rubrik_configure_timezone" "LA-Timezone" {
  timezone = "America/Los_Angeles"
}
```



### Cluster Bootstrap

The following demonstrates an example of bootstrapping a new Rubrik cluster:


```
resource "rubrik_bootstrap" "example" {
cluster_name = "tf-demo"
admin_email = "tf@demo.com"
admin_password = "RubrikTFDemo2019"
management_gateway = "10.167.8.1"
management_subnet_mask = "255.255.252.0"
dns_name_servers = ["10.167.8.2"]
ntp_servers = ["8.8.8.8"]
node_config = {
tf-node01 = "10.167.8.180"
}
}
```

## Rubrik Provider for Terraform Documentation

This guide acts only as a quick start to get up and running with the Terraform Provider for Rubrik. For detailed information on all of the functions and features included, see the complete [Terraform Provider for Rubrik documentation](https://rubrik.gitbook.io/terraform-provider-for-rubrik/).


## API Documentation

The Rubrik Provider for Terraform supports much of the configuration for deploying the Rubrik CDM software as well as preparing AWS and Azure for CloudOut and CloudOn. However, keep in mind that the release cycles between the Provider and Rubrik CDM are not simultaneous. This means there may be times when new features or enhancements are added to the product but resources and functions to utilize them may be missing from the SDK. In these situations Terraform may be used to make native calls to Rubrik's RESTful API.

Rubrik prides itself upon its API-first architecture, ensuring everything available within the HTML5 interface, and more, is consumable via a RESTful API. For more information on Rubrik's API architecture and its documentation, please see the [Rubrik API Documentation](https://github.com/rubrikinc/api-documentation).


## Contributing to the Rubrik Provider for Terraform

The Rubrik Provider for Terraform is hosted on a public repository on GitHub. If you would like to get involved and contribute to the Rubrik Provider please follow the below guidelines.


### Common Environment Setup



1.  Clone the Rubrik Provider for Terraform repository

    `git clone https://github.com/rubrikinc/rubrik-provider-for-terraform.git`


1.  Change to the repository root directory

    `cd rubrik-provider-for-terraform`


1.  Switch to the devel branch

    `git checkout devel`


### New Module Development

The` /rubrik-provider-for-terraform/rubrikcdm` directory contains all of the Rubrik Terraform resources. You can also utilize the following file as a template for all new resource functions:


```
rubrik-modules-for-ansible/blob/master/docs/rubrik_module_template.py
```


To add parameters specific to the new resource you can update the following section which starts on `line 67`:


```
func resourceRubrikAWSNativeAccountCreate(d *schema.ResourceData, meta interface{}) error {
```


After the new variables have been defined you can start adding any new required logic after the code block section.


```
##################################
######### Code Block #############
##################################
##################################
```


Once the resource and functions have been fully coded, please update or add to Rubrik Provider for Terraform documentation. The directory is located [here](https://github.com/rubrikinc/rubrik-provider-for-terraform/tree/master/docs). 


## Further Reading

*   [Rubrik Provider for Terraform GitHub Repository](https://github.com/rubrikinc/rubrik-provider-for-terraform)
*   [Rubrik Provider for Terraform Official Documentation](https://rubrik.gitbook.io/terraform-provider-for-rubrik/)
*   [Rubrik CDM API Documentation](https://github.com/rubrikinc/api-documentation)
*   [BLOG: Using Terraform with Rubrik Just Got Easier!](https://www.rubrik.com/blog/rubrik-provider-terraform/)
*   [VIDEO: Getting Started with the Rubrik Provider for Terraform](https://www.youtube.com/watch?v=kV1xiP1tHY0)
