rm ./xray.tar.xz
go build -o xray -trimpath -ldflags "-s -w -buildid=" ./main
tar -I 'xz -9 -T0' -cf ./xray.tar.xz xray geoip.dat geosite.dat LICENSE README.md
rm xray
