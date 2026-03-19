//go:generate /bin/bash -c "structgen -repo 'https://github.com/apple/device-management.git' -commit \"$(cat ../../../GENERATE_SHA)\" -path 'mdm/errors' -pkg mdmerrors -tags json,plist -reqdef -out errors.gen.go"

package mdmerrors
