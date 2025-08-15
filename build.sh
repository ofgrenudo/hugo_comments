#!/bin/bash

# Get Version from file, ensuring only the first line is used
version_string=$(grep -Eo '[0-9\.]+' version.rc | head -n 1)

# Check if a version string was found
if [ -z "$version_string" ]; then
    echo "Error: No version number found in version.rc"
    exit 1
fi

# Split the version string into major and minor parts
IFS='.' read -r major minor <<< "$version_string"

######################################################################
# Script Variables
######################################################################
distribution_dir="dist" # Fixed typo: distrubtion -> distribution
executable_name="hugo_comments"
target_entry_point="src/main.go"

echo "Building version $major.$minor"
echo ""

# Create the distribution directory if it doesn't exist
mkdir -p "${distribution_dir}"

######################################################################
# Build for Multiple Platforms
######################################################################

# Define our targets in "OS/ARCH" format
targets=("windows/amd64" "linux/amd64" "darwin/amd64")

for target in "${targets[@]}"; do
    # Split the target string into OS and Arch
    IFS='/' read -r os arch <<< "$target"

    echo "Building for ${os}..."
    export GOOS=${os}
    export GOARCH=${arch}

    # Set the output filename
    output_name="${distribution_dir}/${executable_name}_${os}-${major}.${minor}"
    
    # Add .exe for windows builds
    if [ "$GOOS" = "windows" ]; then
        output_name+=".exe"
    fi

    # Build the executable with corrected arguments
    go build -o "${output_name}" "${target_entry_point}"
    
    echo "Build complete: ${output_name}"
    echo ""
done