ECHO "Building main.go"
sudo go build ./src/main.go
ECHO "Build complete"
mv ./main ./build/mac/main
ECHO "Moved executable to build/mac"
ECHO "Zipping the build"
zip -r -X build_mac.zip ./build/mac/
ECHO "Build for mac complete"
ECHO "Build for Windows now"
sudo GOOS=windows GOARCH=amd64 go build ./src/main.go
ECHO "Build complete"
mv ./main.exe ./build/win/main.exe
ECHO "Moved executable to build/win"
ECHO "Zipping the build"
zip -r -X build_windows.zip ./build/win/
ECHO "Build for Windows complete"
ECHO "Cleaning up"
rm -rf ./build/mac/main
rm -rf ./build/win/main.exe

mv ./build_windows.zip ./ready_builds/build_windows.zip
mv ./build_mac.zip ./ready_builds/build_mac.zip