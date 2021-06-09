str="
name: Dummy-golangci-lint
"
echo "$str"

encoded=`echo $str | base64`
echo $encoded
