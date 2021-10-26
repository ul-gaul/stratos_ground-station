# Stratos - Ground Station

_**TODO:** describe the project_

## Contributing

### Requirements

The following languages and dependencies need to be installed to be able to run and build the project.

- [Go / Golang](https://golang.org/doc/install) (min version: **v1.16**)<br>
  <sup>Run `go version` to get your current version</sup>
- [Docker](https://docs.docker.com/get-started/#download-and-install-docker) <sub><sup>**(optional)**</sup></sub><br>
  <sup>Only required to cross-compile (see [build](#for-the-raspberry-pi-cross-compile))</sup>
- [NodeJS](https://nodejs.org/en/download/current/)
- [Wails](https://wails.app/gettingstarted/)

### Run

To run the program in development mode:

1. Open a terminal in the project root directory.
2. Start the backend: `wails serve`
3. Open a second terminal in the project root directory.
4. Move to the frontend directory: `cd ./frontend`
5. Start the frontend: `npm run serve`
6. Wait for the application to be ready. <br>
   <img alt="frontend-serve-ready" width="200" src="https://user-images.githubusercontent.com/14944216/113919412-933a2900-97b1-11eb-8a6b-66f9bb0abd7f.png"/>
7. Open a browser and go to [localhost:8080](http://localhost:8080/).


### Build

#### For your OS

To build the binary (executable) for your operating system:

1. Open a terminal in the project root directory.
2. Build the project: `wails build`

#### For the Raspberry Pi (cross-compile)
TODO

## Structure

<dl>
<dt>/backend/acquisition</dt>
<dd>Contains the backend. (data acquisition, data transformation, etc.)</dd>

<dt>/backend/cmd <sup><sub>(TODO)</sub></sup></dt>
<dd>Contains the code reading arguements passed from the CLI.<br>
<sup><em>(CLI: Command Line Interface)</em></sup></dd>

<dt>/backend/config <sup><sub>(TODO)</sub></sup></dt>
<dd>Contains the code reading the configuration from a file.</dd>

<dt>/frontend</dt>
<dd>
Contains the frontend. (VueJS)
</dd>
</dl>

## TODOs

- [X] **[CLI][i2] <sup><sub>(/cmd)</sub></sup>:** read args from command line. (see [spf13/cobra](https://github.com/spf13/cobra))
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
  - [x] **[Dispatch:][i8]** create functions to send data to frontend.
  - [ ] **Tests:** create unit tests.
    - [ ] **Bindings:** test bindings functions.
    - [ ] **Dispatch:** test dispatch functions.
- [ ] **[Frontend][iFrontend] <sup><sub>(/dashboard)</sub></sup>**
  - [ ] **[Data reception:][i9]** create the functions receiving the data from the backend.
  - [ ] **Tests:** create unit tests.
    - [ ] **Data reception:** test the functions used to receive the data from the backend.
    - [ ] **Other:** test other functions of the frontend.


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