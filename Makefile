# Makefile

# Determine the operating system for setting appropriate commands
ifeq ($(OS),Windows_NT)
	# Windows
	BINARY_EXT = .exe
	LD_FLAGS = -ldflags="-H=windowsgui"
	RM = del /Q
else
	# Unix-like
	BINARY_EXT =
	LD_FLAGS =
	RM = rm -f
endif

# Specify the name of your Go program's binary
BINARY_NAME = guipages2print$(BINARY_EXT)

# Default target to build the program
all: build

# Build the Go program using the go build command
build:
	go build $(LD_FLAGS) -o $(BINARY_NAME) .

# Clean up the generated binary
clean:
	$(RM) $(BINARY_NAME)

# Set the 'clean' target as a phony target (doesn't correspond to a file)
.PHONY: clean
