# Runtime Go
The Runtime Script programming language written in Golang

Learn more about Runtime Script:  
https://github.com/yjlo123/runtime-script

## Usage  
```shell
$ runtime <file_path>
```


## Build for CLI
```
cd cmd/run
go build .

#Build for other OS & arch, e.g.
env GOOS=linux GOARCH=amd64 go build .
```

## Packaging GUI (Deprecated)
https://developer.fyne.io/started/packaging  
```
cd ui
go build .
fyne package -os windows -icon ../icon.png -executable ui.exe
fyne package -os darwin -icon ../icon.png -executable ui
```

<img src="https://github.com/yjlo123/runtime-go/blob/main/screenshot_vm.png">
