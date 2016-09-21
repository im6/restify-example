"use strict";
const globalConfig = require('../env/config'),
    restify = require('restify'),
    route = require('../server/route'),
    mongo = require('../server/mongodb/config'),
    authMd = require('../server/middleware/auth');

var server = restify.createServer({
    name: 'myapp',
    version: '1.0.0'
});

server.use(restify.acceptParser(server.acceptable));
server.use(restify.queryParser());
server.use(restify.bodyParser());
if(globalConfig.requireAuth){
    server.use(authMd.check);
}

route.bindAction(server);

mongo.initial();


server.listen(globalConfig.serverPort, globalConfig.serverIp, function () {
    console.log(`application is running at ${globalConfig.serverIp}:${globalConfig.serverPort}`);
    console.log(`Application worker ${process.pid} started...`);
});