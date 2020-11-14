GOOS="linux" GOARCH="arm" go build -o ./build/service service.go
adb push ../MyFileService /data/local/tmp
adb shell "cd /data/local/tmp/MyFileService; ./build/service"