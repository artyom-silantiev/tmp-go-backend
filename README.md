# preparation
## install golang env
https://golang.org/doc/install
https://golang.org/doc/tutorial/getting-started
## compile and install packages and dependencies
```sh
go install
```
## create build dir
```sh
mkdir build
```
## copy/paste app run script and view/edit config
```sh
cp dev.run-app.default.sh dev.run-app.sh
nano dev.run-app.sh # view, [edit,] save
```

# build
```sh
sh dev.build.sh
```

# run
```sh
sh dev.run-app.sh
```