//go:generate /bin/bash -c "profilegen -repo 'https://github.com/apple/device-management.git' -commit \"$(cat ../GENERATE_SHA)\" -repl ./repls.yaml"

package profiles
