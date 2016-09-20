"use strict";

const testCtr = require('./middleware/test');

module.exports = {
    bindAction: (server) => {
        server.get('/', testCtr.root);
        server.get('/echo/:name', testCtr.echo);
    }
};