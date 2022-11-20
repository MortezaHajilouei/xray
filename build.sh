rm ./xray.tar.xz
go build -o xray -trimpath -ldflags "-s -w -buildid=" ./main
tar -I 'xz -9 -T0' -cf ./x-ui.tar.xz xray
rm xray
