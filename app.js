"use strict";
const globalConfig = require('./env/config'),
    express = require('express'),
    route = require('./server/route'),
    bodyParser = require('body-parser'),
    mongo = require('./server/mongodb/config'),
    authMd = require('./server/middleware/auth');

let app = express();
app.set('x-powered-by', false);
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));


if(globalConfig.requireAuth){
    app.use(authMd.check);
}

route.bindAction(app);



mongo.initial();


app.listen(globalConfig.serverPort, globalConfig.serverIp, function () {
    console.log(`application is running at ${globalConfig.serverIp}:${globalConfig.serverPort}`);
    console.log(`Application worker ${process.pid} started...`);
});
