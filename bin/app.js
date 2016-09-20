"use strict";
const env = process.env,
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
server.use(authMd.check);

route.bindAction(server);

mongo.initial();


server.listen(env.NODE_PORT || 3000, env.NODE_IP || 'localhost', function () {
    console.log(`Application worker ${process.pid} started...`);
});
