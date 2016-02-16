# Emerald
Find the best Gems in DCC

## Installation

 > **Prerequisites**: Emerald needs a Mongo-DB for storing its data. Please install it before running Emerald.

1. Download and extract the release according to your environment
2. Edit the configuration file `emerald.conf`, which should look like this:
  ```
  [server]
  port = 8080
  logFile = emerald.log
  mongodb = localhost
  ```


3. run the executable with the flag "conf" indication the location of the config-file:

  **For Linux:**
  ```sh
  emerald -conf emerald.conf
  ```
  **For Windows:**
  ```sh
  emerald.exe -conf emerald.conf
  ```


#### Configuration-File
| Value         |Description                                     | default-value |
| ------------- |------------------------------------------------|---------------|
| port          | The Port on which Emerald should run.          | `8080`        |
| logFile       | The location of the logfile                    | `emerald.log` |
| mongodb       | the host and port where the MongoDB is running | `localhost`   |
