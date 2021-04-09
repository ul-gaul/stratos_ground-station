# Stratos - Ground Station

_**TODO:** describe the project_

## Contributing

### Requirements

The following languages need to be installed to be able to run and build the project.

- [Go / Golang](https://golang.org/doc/install) (min version: **v1.16**)<br>
  <sup>Run `go version` to get your current version</sup>
- [Docker](https://docs.docker.com/get-started/#download-and-install-docker) <sub><sup>**(optional)**</sup></sub><br>
  <sup>Only required to build the binaries (executable files).</sup>
- [NodeJS](https://nodejs.org/en/download/current/)

### Run
There are 3 different ways to run the program. You can either run only the frontend, run only the backend with pre-built frontend files or run the backend with hot reload of the frontend.

<dl>
<dt>
  1. With hot reload of frontend (recommended)<br>
</dt>
<dd>
<sup>The application will run in a dedicated window.</sup>
<details>
  <summary>Steps</summary>

1. Follow steps 1 to 4 descibed in [_Frontend only_](#run-frontend-steps).
2. Open a second terminal in the project root directory.
3. Start the application: `go run -tags dev .`

<sub><em>
<strong>Note 1:</strong> changes made to the frontend (in /dashboard) will automatically be reflected in the application window. <br>
<strong>Note 2:</strong> only repeat steps 2 and 3 to see changes made to the backend.
</em></sub>
</details>
</dd>

<dt>
  2. Without hot reload
</dt>
<dd>
<sup>The application will run in a dedicated window.</sup>
<details>
  <summary>Steps</summary>

1. Open a terminal in the project's root directory.
2. Move to the frontend directory: `cd ./dashboard`
3. Build the frontend: `npm run build`
4. Move back to the project's root directory: `cd ..`
6. Start the application: `go run .`

<sub><em>
<strong>Note 1:</strong> only repeat step 6 to see changes made to the backend. <br>
<strong>Note 2:</strong> repeat all steps to see changes made to the frontend (in /dashboard).
</em></sub>
</details>
</dd>

<dt>
  3. Frontend only
</dt>
<dd>
<sup>Only the frontend development server will be running so no data will be provided from the backend.</sup>
<details>
  <summary id="run-frontend-steps">Steps</summary>

1. Open a terminal in the project root directory.
2. Move to the frontend directory: `cd ./dashboard`
3. Start the development server: `npm run serve`
4. Wait for the application to be ready. <br>
   <img alt="frontend-serve-ready" width="200" src="https://user-images.githubusercontent.com/14944216/113919412-933a2900-97b1-11eb-8a6b-66f9bb0abd7f.png"/>
5. Open a browser and go to [localhost:8080](http://localhost:8080/).
</details>
</dd>

</dl>


### Mage

Once this is done, you need to install [Mage](https://magefile.org/). To do so open a terminal and execute the following
commands or follow the instructions [here](https://magefile.org/#installation):

On Windows:
```sh
git clone https://github.com/magefile/mage
cd mage
go run bootstrap.go
```
On MacOS:
```sh
brew install mage
```


### Build

The project uses [Mage](https://magefile.org) to make it easier to run and build the project. <br>
Here are the principal mage targets/commands:<br>
<sup>(run `mage -l` to list all available targets)</sup>

<dl>
<dt>
  <code>mage</code> or <code>mage build</code>
</dt>
<dd>
  Builds and compiles the project into a binary.
</dd>

<dt>
  <code>mage clean</code>
</dt>
<dd>
  Clean the project by removing built binaries, <em>node_modules</em> directory and other generated files.
</dd>
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