#!/bin/sh

VERSION="1.0.0"
APP="acli"

echo "Cleaning dist..."
rm -rf dist
mkdir -p dist

echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o dist/${APP}-linux .

echo "Building for macOS..."
GOOS=darwin GOARCH=amd64 go build -o dist/${APP}-mac .

echo "Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o dist/${APP}.exe .

echo "Preparing .deb package structure..."

PKGDIR="dist/${APP}_deb"
mkdir -p ${PKGDIR}/DEBIAN
mkdir -p ${PKGDIR}/usr/local/bin

echo "Copying Linux binary..."
cp dist/${APP}-linux ${PKGDIR}/usr/local/bin/${APP}
chmod 755 ${PKGDIR}/usr/local/bin/${APP}

echo "Creating control file..."

cat > ${PKGDIR}/DEBIAN/control <<EOF
Package: ${APP}
Version: ${VERSION}
Section: utils
Priority: optional
Architecture: amd64
Maintainer: Vicky Chhetri <vickychhetri4@gmail.com>
Description: acli - Advanced daily log CLI tool
 A CLI tool to manage daily logs, priorities, categories, weekly summary and more.
EOF

echo "Building .deb package..."
dpkg-deb --build ${PKGDIR}

echo "Renaming final file..."
mv dist/${APP}_deb.deb dist/${APP}_${VERSION}_amd64.deb

echo "----------------------------------------"
echo "Build Completed!"
echo "Linux binary: dist/${APP}-linux"
echo "macOS binary: dist/${APP}-mac"
echo "Windows binary: dist/${APP}.exe"
echo "Debian package: dist/${APP}_${VERSION}_amd64.deb"
echo "----------------------------------------"
