# Stratos - Ground Station

_**TODO:** describe the project_

## Contributing

### Requirements

The following languages need to be installed to be able to run and build the project.

- [NodeJS](https://nodejs.org/en/download/current/)
- [Go / Golang](https://golang.org/doc/install) (min version: **v1.16**)<br>
  <sup>Run `go version` to get your current version</sup>

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
<sup>(run `mage -l` to list all available targets)</sup>

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

<dt>/cmd <sup><sub>(TODO)</sub></sup></dt>
<dd>Contains the code reading arguements passed from the CLI.<br>
<sup><em>(CLI: Command Line Interface)</em></sup></dd>

<dt>/config <sup><sub>(TODO)</sub></sup></dt>
<dd>Contains the code reading the configuration from a file.</dd>

<dt>/dashboard</dt>
<dd>
Contains the frontend. (VueJS)
</dd>
</dl>

## TODOs

- [ ] **[CLI][i2] <sup><sub>(/cmd)</sub></sup>:** read args from command line. (see [spf13/cobra](https://github.com/spf13/cobra))
- [ ] **[Config][i1] <sup><sub>(/config)</sub></sup>:** read configurations from files. (
  see [spf13/viper](https://github.com/spf13/viper))
- [ ] **[Backend][iBackend] <sup><sub>(/acquisition)</sub></sup>**
  - [ ] **[Data acquisition:][i3]** read data received by the RF module.
  - [ ] **[Transform data:][i4]** transform raw data into a human readable form.
  - [ ] **[Save data:][i5]** save transformed data into a file.
  - [ ] **[Read data:][i6]** add the possibility to read data from a file.
  - [ ] **Tests:** create unit tests.
    - [ ] **Data acquisition:** test functions of the data acquisition.
    - [ ] **Transfrom data:** test functions tranforming the data.
    - [ ] **Save data:** test functions saving data into files.
    - [ ] **Read data:** test functions reading data from files.
- [ ] **[Controller][iController] <sup><sub>(/controller)</sub></sup>**
  - [ ] **[Bindings:][i7]** create bindings allowing the frontend to call Go functions.
  - [ ] **[Dispatch:][i8]** create functions to send data to frontend.
  - [ ] **Tests:** create unit tests.
    - [ ] **Bindings:** test bindings functions.
    - [ ] **Dispatch:** test dispatch functions.
- [ ] **[Frontend][iFrontend] <sup><sub>(/dashboard)</sub></sup>**
  - [ ] **[Data reception:][i9]** create the functions receiving the data from the backend.
  - [ ] **Tests:** create unit tests.
    - [ ] **Data reception:** test the functions used to receive the data from the backend.
    - [ ] **Other:** test other functions of the frontend.
- [ ] **Other**
  - [ ] **Mage:** targets/commands to add to magefile
    - [ ] **Crossbuild:** add a target to build the project for Raspberry Pi (`GOOS=linux`, `GOARCH=arm`)


[i1]: https://github.com/ul-gaul/stratos_ground-station/issues/1
[i2]: https://github.com/ul-gaul/stratos_ground-station/issues/2
[i3]: https://github.com/ul-gaul/stratos_ground-station/issues/3
[i4]: https://github.com/ul-gaul/stratos_ground-station/issues/4
[i5]: https://github.com/ul-gaul/stratos_ground-station/issues/5
[i6]: https://github.com/ul-gaul/stratos_ground-station/issues/6
[i7]: https://github.com/ul-gaul/stratos_ground-station/issues/7
[i8]: https://github.com/ul-gaul/stratos_ground-station/issues/8
[i9]: https://github.com/ul-gaul/stratos_ground-station/issues/9
[iBackend]: https://github.com/ul-gaul/stratos_ground-station/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc+label%3Abackend
[iController]: https://github.com/ul-gaul/stratos_ground-station/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc+label%3Acontroller
[iFrontend]: https://github.com/ul-gaul/stratos_ground-station/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc+label%3Afrontend
