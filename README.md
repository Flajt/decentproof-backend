# decentproof-backend
Backend for decentproof-app

For information, or raising issues, please checkout the decentproof-app [repository](https://github.com/Flajt/decentproof-app).


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
    - helper => has helper functions for authentication & key generation
    - util => currently only contains a script to generate public and private key pairs

- Other
    - encryption_service => For en- & decryption of data with AES encryption


## How to get started

### What do I need
- A PC, with ideally a Linux based OS or MacOS
- An .env file
- A scaleway account
- An originstamp account
- A firebase account
- A running instance of mCaptcha if you want to verify w. the website

### .env file
Your .env file needs to contain the following data:
```Shell
SCW_ACCESS_KEY=<my-access-key-here> # can be empty in local mode
SCW_SECRET_KEY=<my-secret-key-here> # can be empty in local mode
SCW_DEFAULT_ORGANIZATION_ID=<my-org-id-here> # can be empty in local mode
SCW_DEFAULT_PROJECT_ID=<my-project-id-here> # can be empty in local mode
SCW_DEFAULT_REGION=<my-scaleway-defualt-region-here> # can be empty in local mode
GOOGLE_ADMIN_SDK_CREDS=<my-firebase-admin-sdk-key-here> # can be empty in local mode
ORIGINSTAMP_API_KEY=<my-api-key-here> # can be empty in local mode
SECRET_KEY=<your-secret-key-for signatures> # use uti/generate_keys.go to generate it
EMAIL_SECRET=<your-secret-key-with-email-permissions> # can be empty in local mode
WEBHOOK_URL=<the-url-for-the-webhook-callback> # if you don't set a domain you will need to deploy it first to get it.
PRIVATE_KEY=<the-private-key-for-signatures> # this one needs to be in the scaleway secret manager, or local mode!
ENCRYPTION_KEY=<encryption-key-32bytes-for-mail> # only in secret manager, or local mode!
DEBUG=<TRUE-or-anything-else> # if set to TRUE you can run the functions in local mode, which should improve testing capabilities and local development
API_KEY=<base64-encoded-32byte-long-key> 
MCAPTCHA_SECRET=<your-mcaptcha-secret> # needs to be set if you want to verify with the website, in PROD it's an encrypted secret
MCAPTCHA_SITEKEY=<your-mcaptcha-site-key> # needs to be set if you want to verify with the website, in PROD it's an encrypted secret
MCAPTCHA_INSTANCE_URL=<your-mcpatcha-instance-url> # needs to be set if you want to verify w. website
```
#### Env vars per function
- sign: `WEBHOOK_URL` & `ORIGINSTAMP_API_KEY`
- get-new-key: `GOOGLE_ADMIN_SDK_CREDS`
- cron-job: ``
- webhook: `ORIGINSTAMP_API_KEY` & `EMAIL_SECRET`
- verify-hash: `ORIGINSTAMP_API_KEY` & `MCAPTCHA_SECRET` & `MCAPTCHA_SITEKEY` & `MCAPTCHA_INSTANCE_URL`

Note to self:
Currently you need to update the functions secrets and env vars via `scw function function update`

The issue is it's nearly needed everywhere, in every function, in every test folder, everywhere...
So please load it into your terminal enviroment. You can use my script in utils for that: `util/load_env.go`. This should load all env vars into your terminal (tested in VSCode), use the `--path` flag to pass the .env file path.

**NOTE**: Currently the Github Secret for Originstamp Api and E-Mail Secret are the same for DEV as for PROD. The latter should be changed at some point in time. 

## "Local Mode"
What I've mentioned as local mode is just running the functitions locally from the `cmd` folder as explained below in testing (this can also be used for local development). If you set `DEBUG=TRUE` a local implementation of the scw_secret_wrapper is used which loads the three most important api keys into memory. 

It will also disable app check for the time beeing, so you don't need to create a firebase account either.

If you think it makes sense to sperate firebase app check "dis- & enablement" into a sperate var please open a discussion.

> This doesn't work for cronjobs and webhooks will only work partially as it depends on both an originstamp account as well as a working SMTP service to send E-Mails which in turn uses the `SCW_DEFAULT_PROJECT_ID` & `EMAIL_SECRET` as credentials

### SMTP 
If you know an easy way to mock / setup a local SMTP service, feel free to open an issue / discussion.
In the meantime if you want to deploy a SMTP service to test things locally, you need to replace the `smtpServer` variable in `webhook.go` 

## How to Test

### Test function implementation
    1. Go into every folder and download the deps with: `go get .` or use `go get -d ./...`
    2. Navigate to the `/cmd` folder of each function, there you can run: `go run main.go` this will deploy the function on a given port, be aware you need to load the env vars beforehand

### Automated Tests
Currently the following packages offer tests: <br>
- cronjob (currently not applicable)
- scw_secret_wrapper
- helper
- webhook
- encryption_service
- sign (currently not applicable)

More support for functions should be added in the future.
**Important:** Nearly all currently available tests are *E2E* and require access to scaleway 

<br>

**Note:** It's not planned to add teststing for originstamp related features due to the monthly limit of 5 free documents and to not clutter the blockchain.

**Note2:** Currently tests which use the secret manager wipe all secrets after they are done, which needs to be resolved!

## Why so many packages?
Yeah... about this: <br>
Scaleway requires every function to be it's own module, and since some scripts are used by multiple modules, I've moved these into their own modules, leading to even more modules. _moduleception_ .In the future, if this project userbase should increase from the .5 it currently has, I convert it from functions into a single api server. But that's in the far future.
