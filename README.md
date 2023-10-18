# decentproof-backend
Backend for decentproof-app

For information, or raising issues, please checkout the decentproof-app [repository](https://github.com/Flajt/decentproof-app).

TODO: Implement docs, contribution tips

## Project Structure

- Functions
    - get_new_key => responsible for getting a new api key
    - has_new_key => responsible for checking if there is a new api key
    - sign => signs and submitts your data to originstamp
    - verify_hash => verifies your hash
    - webhook => used for originstamp webhooks
    - cronjob => refreshes api key

- Api Wrappers
    - originstamp => Originstamp.com api
    - scw_secret_wrapper => wraps scaleways secret manager

- Util
    - helper => has helper functions for authentication
    - util => currently only contains a script to generate public and private key pairs


## How to get started

### What do I need
- A PC, with ideally a Linux based OS or MacOS
- An .env file
- A scaleway account
- An originstamp account
- A firebase account

### .env file
Your .env file needs to contain the following data:
```Shell
SCW_ACCESS_KEY=<my-access-key-here>
SCW_SECRET_KEY=<my-secret-key-here>
SCW_DEFAULT_ORGANIZATION_ID=<my-org-id-here>
SCW_DEFAULT_PROJECT_ID=<my-project-id-here>
SCW_DEFAULT_REGION=<my-scaleway-defualt-region-here>
GOOGLE_ADMIN_SDK_CREDS=<my-firebase-admin-sdk-key-here>
ORIGINSTAMP_API_KEY=<my-api-key-here>
SECRET_KEY=<your-secret-key-for signatures> # use uti/generate_keys.go to generate it
SCW_EMAIL_SECRET=<your-secret-key-with-email-permissions>
WEBHOOK_URL=<the-url-for-the-webhook-callback> # if you don't set a domain you will need to deploy it first to get it.
PRIVATE_KEY=<the-private-key-for-signatures> # this one needs to be in the scaleway secret manager not .env file !
```
The issue is it's nearly needed everywhere, in every function, in every test folder, everywhere...
So please load it into your terminal enviroment. You can use my script in utils for that: `util/load_env.go`. This should load all env vars into your terminal (tested in VSCode), use the `--path` flag to pass the .env file path.

**NOTE:** Currently the Github Secret for Originstamp Api and E-Mail Secret are the same for DEV as for PROD. The latter should be changed at some point in time. 

## How to Test

### Test function implementation
    1. Go into every folder and download the deps with: `go get .` or use `go get -d ./...`
    2. Navigate to the `/cmd` folder of each function, there you can run: `go run main.go` this will deploy the function on a given port, be aware you need to load the env vars beforehand

### Automated Tests
Currently the following packages offer tests: <br>
- cronjob
- scw_secret_wrapper
- helper

More support for functions should be added in the future.
**Important:** Nearly all currently available tests are *E2E* and require access to scaleway 

<br>
**Note:** It's not planned to add teststing for originstamp related features due to the monthly limit of 5 free documents and to not clutter the blockchain.

## Why so many packages?
Yeah... about this: <br>
Scaleway requires every function to be it's own module, and since some scripts are used by multiple modules, I've moved these into their own modules, leading to even more modules. _moduleception_ .In the future, if this project userbase should increase from the .5 it currently has, I convert it from functions into a single api server. But that's in the far future.
