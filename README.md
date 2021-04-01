# Stratos - Ground Station

_**TODO:** describe the project_

## Contributing

### Requirements

The following languages need to be installed to be able to run and build the project.

- [NodeJS](https://nodejs.org/en/download/current/)
- [Go / Golang](https://golang.org/doc/install) (Min version: **v1.16**)<br>
  <small>Run `go version` to get your current version</small>

Once this is done, you need to install [Mage](https://magefile.org/). To do so open a terminal and execute the following
commands or follow the instructions [here](https://magefile.org/#installation):

```sh
go get -u -d github.com/magefile/mage
cd $GOPATH/src/github.com/magefile/mage
go run bootstrap.go
```

### Run & Build

The project uses [Mage](https://magefile.org) to make it easier to run and build the project. <br>
Here are the principal mage targets/commands:<br>
<small>(run `mage -l` to list all available targets)</small>

<dl>
<dt><code>mage</code> or <code>mage run</code></dt>
<dd>Starts the project in development mode.</dd>

<dt><code>mage build</code></dt>
<dd>Builds and compiles the project into a binary.</dd>

<dt><code>mage clean</code></dt>
<dd>Clean the project by removing built binaries, <em>node_modules</em> directory and other compiled code.</dd>
</dl>

## Structure

<dl>
<dt>/acquisition</dt>
<dd>Contains the backend. (data acquisition, data transformation, etc.)</dd>

<dt>/cmd <small>(TODO)</small></dt>
<dd>Contains the code reading arguements passed from the CLI.<br>
<small><em>(CLI: Command Line Interface)</em></small></dd>

<dt>/config <small>(TODO)</small></dt>
<dd>Contains the code reading the configuration from a file.</dd>

<dt>/dashboard</dt>
<dd>
Contains the frontend. (VueJS)
</dd>
</dl>

## TODOs

- [ ] **CLI <small>(/cli)</small>:** read args from command line. (see [spf13/cobra](https://github.com/spf13/cobra))
- [ ] **Config <small>(/config)</small>:** read configurations from files. (
  see [spf13/viper](https://github.com/spf13/viper))
- [ ] **Backend <small>(/acquisition)</small>**
  - [ ] **Data acquisition:** read data received by the RF module.
  - [ ] **Transform data:** transform raw data into a human readable form.
  - [ ] **Save data:** save transformed data into a file.
  - [ ] **Read data:** add the possibility to read data from a file.
  - [ ] **Tests:** create unit tests.
    - [ ] **Data acquisition:** test functions of the data acquisition.
    - [ ] **Transfrom data:** test functions tranforming the data.
    - [ ] **Save data:** test functions saving data into files.
    - [ ] **Read data:** test functions reading data from files.
- [ ] **Controller <small>(/controller)</small>**
  - [ ] **Bindings:** create bindings allowing the frontend to call Go functions.
  - [ ] **Dispatch:** create functions to send data to frontend.
  - [ ] **Tests:** create unit tests.
    - [ ] **Bindings:** test bindings functions.
    - [ ] **Dispatch:** test dispatch functions.
- [ ] **Frontend <small>(/dashboard)</small>**
  - [ ] **Data reception:** create the functions receiving the data from the backend.
  - [ ] **Tests:** create unit tests.
    - [ ] **Data reception:** test the functions used to receive the data from the backend.
    - [ ] **Other:** test other functions of the frontend.

