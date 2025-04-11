#!/bin/bash

# Configuration variables
APP_NAME="SecureFiles"           # Name of your application
BINARY_NAME="out/bin/secure-files-go-gui-darwin_arm64"        # Name of your compiled Go binary
ICNS_FILE="internal/gui/assets/app-icon.icns"
VERSION="1.0"             # Version number
BUNDLE_ID="com.darshana.securefilesgogui"  # Unique bundle identifier

# Check if binary exists
if [ ! -f "$BINARY_NAME" ]; then
    echo "Error: Binary '$BINARY_NAME' not found. Please compile your Go program first."
    exit 1
fi

# Create the .app directory structure
APP_DIR="out/app/${APP_NAME}.app"
CONTENTS_DIR="${APP_DIR}/Contents"
MACOS_DIR="${CONTENTS_DIR}/MacOS"
RESOURCES_DIR="${CONTENTS_DIR}/Resources"

mkdir -p "$MACOS_DIR"
mkdir -p "$RESOURCES_DIR"

# Copy the binary
cp "$BINARY_NAME" "$MACOS_DIR/$APP_NAME"

# Copy the icns file
cp "$ICNS_FILE" "$RESOURCES_DIR/$APP_NAME.icns"

# Create the Info.plist file
cat > "${CONTENTS_DIR}/Info.plist" << EOL
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>CFBundleExecutable</key>
    <string>${APP_NAME}</string>
    <key>CFBundleIconFile</key>
    <string>${APP_NAME}.icns</string>
    <key>CFBundleIdentifier</key>
    <string>${BUNDLE_ID}</string>
    <key>CFBundleName</key>
    <string>${APP_NAME}</string>
    <key>CFBundleVersion</key>
    <string>${VERSION}</string>
    <key>CFBundlePackageType</key>
    <string>APPL</string>
    <key>CFBundleShortVersionString</key>
    <string>${VERSION}</string>
    <key>NSPrincipalClass</key>
	<string>NSApplication</string>
    <key>NSHighResolutionCapable</key>
    <true/>
</dict>
</plist>
EOL

# cat > "$MACOS_DIR/launcher.sh" << EOL
# #!/bin/bash

# SCRIPT_DIR="\$(cd "\$(dirname "\$0")" && pwd)"
# BINARY="\$SCRIPT_DIR/$APP_NAME"

# # Log message to Console with a custom tag
# logger -t "$APP_NAME" "Launching binary: $BINARY"

# # Run the binary in a pseudo-terminal (no visible window)
# script -q /dev/null "\$BINARY"
# EOL

codesign --force --deep --sign - ${APP_DIR}

# Make the binary executable
chmod +x "$MACOS_DIR/$APP_NAME"
#chmod +x "$MACOS_DIR/launcher.sh"

echo "Created ${APP_NAME}.app successfully!"
echo "You can move it to /Applications or run it from here."