# Testing terraform-provider-rubrik

## Prerequisites

To run acceptance tests, Rubrik CDM must have the following:

* A VMware VM with no SLA assigned
* An SLA domain suitable to use for testing

## Running tests

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create/modify real resources as specified.

*Note:* Testing `resource_rubrik_configure_timezone` will set the CDM timezone setting to `UTC`.

```sh
$ make testacc
```

## Environment variables

There are several environment variables required for acceptance tests:

* `TF_ACC=1` enables the acceptance tests. It is also set when you run `make testacc`.
* `RUBRIK_CDM_NODE_IP` should be set to the IP or hostname of the CDM cluster/virtual appliance you are testing against.
* `RUBRIK_CDM_USERNAME` should be set to a username with access to the specified CDM instance.
* `RUBRIK_CDM_PASSWORD` should be set to the password of the user specified in `RUBRIK_CDM_USERNAME`.
* `RUBRIK_CDM_EXPECTED_VERSION` should be set to the version of CDM you are testing against. Ex: `5.0.2-1980`
* `RUBRIK_CDM_TEST_VM` should be set to the name of the VM with no SLA Domain assigned.
* `RUBRIK_CDM_TEST_SLA` should be set to the name of an SLA Domain suitable for testing.
