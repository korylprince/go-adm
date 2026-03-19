//go:generate /bin/bash -c "structgen -repo 'https://github.com/apple/device-management.git' -commit \"$(cat ../../../GENERATE_SHA)\" -path 'mdm/checkin' -pkg checkin -tags json,plist -reqdef -out checkin.gen.go"

package checkin
