echo
# /home/cirobispo/go/
# _gopath="$(go env GOPATH):$(pwd)"
_gopath="$(pwd)"
echo "SET/EXPORT GOPATH"
export GOPATH="$_gopath"
go env -w GOPATH="$_gopath"
echo "GOPATH=$GOPATH"
go env GOPATH
echo

_go111="auto"
echo "SET/EXPORT GO111MODULE"
export GO111MODULE="$_go111"
echo "GO111MODULE=$GO111MODULE"
go env -w GO111MODULE="$_go111"
go env GO111MODULE
echo
# code --folder-uri ./src

# export PATH=$PATH:$GOPATH/bin
