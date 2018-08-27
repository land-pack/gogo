# How to create a new package yourself


Pick up a dir and go into it
===

    cd /Users/frank/code/github/gogo/example/
    mkdir package_demo

Set go path
===
    
    cd package_demo
    export GOPATH=$(pwd)

Create a new model
===

    mkdir -p src/mymodel
    touch src/mymodel/xxx.go

Create a entrance of application
===
    
    vim main.go
    # ...


Test
===
    
    go run main.go
