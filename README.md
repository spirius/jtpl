# jtpl
Simple JSON based templating engine. You can provide JSON content as standard input to script and render golang templates using as a variables data from input.

# Build

    GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o jtpl_linux_x86_64
    GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -o jtpl_darwin_x86_64
