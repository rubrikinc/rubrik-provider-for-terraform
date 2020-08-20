#!/usr/bin/env bash

rm rubrikcdm/resource_rubrik_template.go

package_split=(${package//\// })
package_name="terraform-provider-rubrik"

platforms=("darwin/amd64" "linux/amd64" "windows/amd64")


for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$package_name'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi
    
    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name $package
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done
