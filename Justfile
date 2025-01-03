set windows-shell := ["pwsh.exe", "-NoLogo", "-Command"]

alias r:=run
alias b:=build
alias s:=spellcheck
alias t:=test
alias ta:=test-all

# run the project
run DAY PART:
    go run . -d {{DAY}} -p {{PART}}

# build with tectonic
build:
    go build .

# run test on specified day
test DAY:
    go test ./solutions/day{{DAY}}

# run all tests
test-all:
    go test ./...

# run aspell on the specified target
spellcheck FILE:
    aspell -c {{FILE}}
