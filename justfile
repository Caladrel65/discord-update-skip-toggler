default:
    @echo Hello! This is just. You got it working!

# Builds the artifact
build:
    mkdir -p bin
    @echo "Building project..."
    go build -ldflags="-w -s" -o bin/discord-update-toggler cmd/main.go
    @echo "Built!"

# Install the binary to /usr/local/bin/
install: build
    @echo "Installing to /usr/local/bin/..."
    @echo "This needs sudo access to modify that location."
    sudo cp bin/discord-update-toggler /usr/local/bin/
    @echo "Installed!"

# Remove the binary from this project
clean:
    rm -rf bin/

# Remove the binary from your system
uninstall:
    @echo "Removing the discord-update-toggler file from /usr/local/bin/..."
    sudo rm -f /usr/local/bin/discord-update-toggler
    @echo "Process finished."
