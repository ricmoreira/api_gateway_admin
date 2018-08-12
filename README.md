# ecommerce-service
API Gateway for ecommerce UI

# Requirements
[Go v1.9 or upper](https://golang.org/doc/install)

# Dependencies
Packages of *ecommerce-service* are managed using dep

To install dep on a linux OS environment run
```
$ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```
[Check here for more information on how to install dep](https://golang.github.io/dep/docs/installation.html)

### Useful commands
```
$ dep ensure -add github.com/pkg/errors     # Get a package and add to dependencies
$ dep ensure                                # Install all packages and dependencies listed in Gopkg.toml
```

[Check here for an unofficial concise guide to dep](https://gist.github.com/subfuzion/12342599e26f5094e4e2d08e9d4ad50d)
[Check here for official documentation about dep](https://golang.github.io/dep/docs/introduction.html)


# Environment Variables
*ecommerce-service* uses Environment Variables(EV) to load some configurations
prior ro run the application, these EV should be provided to the environment, e.g.: 

```
export HOST=localhost:8089 ; go run main.go
```

# Task Runner
In order to perform some usefull tasks like running the server or running unit tests, a Makefile.dist is available.
Copy Makefile.dist file:

```
cp Makefile.dist Makefile
```
Now it's possible to run the server by typing command

```
make run
```

This file may evolve, for e.g. due to new environment variables needed to the app. Just edit this file according to your specific needs and use it 
to create new commands.

# Documentation
*products* Api Documentation is written with Api Blueprint

To generate the api documentation it's used aglio

Both these dependencies are installed through [npm](https://www.npmjs.com/package/npm). Install if it's not available in your OS. 
Once installed,

To install aglio

```
$ npm install
```

To generate api Documentation
```
$ npm run docs
```

or manualy
```
$ $(npm bin)/aglio -i blueprint/products-api.apib --theme-template triple -o public/docs.html
```

# Postman
Both a Postman collection and a Postman environment is provided under  

```
$ <project dir>/docs/postman
```

# Unit Tests
Unit Tests can be run using Go commands. [Check here for official documentation about Go testing](https://golang.org/pkg/testing/)

But the easiest way is using provided Makefile command.
To run all unit tests just type command

```
$ make test
```

To run all unit tests and get access to a coverage report

```
$ make cover
```
Coverage report will appear in a browser tab. Switch between files through the dropdown that will appear at the top of the page.

# Debug
A nice way to debug your application is using VSCode debug features.
A .vscode folder with launch configurations (launch.json) and IDE setting (settings.json) is provided. 

In order to runn your application in debug mode:

1. Hit Debug on VSCode left panel
2. Select "Debug Program" from dropdown on top bar of Debug panel
3. Hit "Play" Button

App will start in debug mode.

In order to debug unit tests, repeat the same steps but, on point 2., instead of choosing  "Debug Program", choose "Launch Tests".
An alternative to this is, on unit test file, click on actions suggested by VSCode. It will appear "run package tests | run file tests" on top of package name on function names. This will happen if Go extension is installed (ms-vscode.go).
