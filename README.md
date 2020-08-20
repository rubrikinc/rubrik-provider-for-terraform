Terraform `Rubrik` Provider
=========================

- Website: https://www.rubrik.com
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Maintainers
-----------

This provider plugin is maintained by the Terraform team at [Rubrik](https://www.rubrik.com/).

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x
-	[Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/rubrikinc/terraform-provider-rubrik/rubrik`

```sh
$ git clone git@github.com:rubrikinc/terraform-provider-rubrik $GOPATH/src/github.com/rubrikinc/terraform-provider-rubrik
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/rubrikinc/terraform-provider-rubrik
$ make build
```

Using the provider
----------------------

Here are some resources to get you started! If you find any challenges from this project are not properly documented or are unclear, please [raise an issue](https://github.com/rubrikinc/rubrik-provider-for-terraform/issues/new/choose) and let us know! This is a fun, safe environment - don't worry if you're a GitHub newbie! :heart:

* [Quick Start Guide](https://github.com/rubrikinc/rubrik-provider-for-terraform/blob/master/docs/quick-start.md)
* [Rubrik Provider for Terraform Documentation](https://rubrik.gitbook.io/terraform-provider-for-rubrik/)
* [Rubrik API Documentation](https://github.com/rubrikinc/api-documentation)
* [VIDEO: Getting Started with the Rubrik Provider for Terraform](https://www.youtube.com/watch?v=kV1xiP1tHY0)
* [BLOG: Using Terraform with Rubrik Just Got Easier!](https://www.rubrik.com/blog/rubrik-provider-terraform/)

## :mag: Example 

```hcl
provider "rubrik" {}

resource "rubrik_configure_timezone" "LA-Timezone" {
  timezone = "America/Los_Angeles"
}
```

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make bin
...
$ $GOPATH/bin/terraform-provider-rubrik
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```