echo "开始打包---"
rm -rf pkg
mkdir pkg
cp -r static pkg/
cp -r configs pkg/
mkdir pkg/log
go build
cp frame pkg/