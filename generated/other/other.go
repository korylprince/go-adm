//go:generate /bin/bash -c "structgen -repo 'https://github.com/apple/device-management.git' -commit \"$(cat ../../GENERATE_SHA)\" -path 'other' -pkg other -tags json,plist -reqdef -out other.gen.go"

package other
