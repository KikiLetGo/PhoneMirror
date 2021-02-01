GOOS="linux" GOARCH="arm" go build -o ./build/phonemirror src/service.go
adb -s 192.168.3.22 push ./build/phonemirror /data/local/tmp/
adb -s 192.168.3.22 shell "su -c '/data/local/tmp/phonemirror'"

# adb shell "cd /data/local/tmp/MyFileService/build/;./phonemirror"