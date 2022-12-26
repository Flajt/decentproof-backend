const {app,Deta} = require("deta");
const updateApiKeys = require("./updateApiKeys");

const deta = Deta(process.env.PROJECT_KEY)

app.lib.cron((event) =>updateApiKeys(deta))

module.exports = app;