"use strict";

const testCtr = require('./middleware/test');

module.exports = {
    bindAction: function(server){
        server.get('/', testCtr.root);
        server.get('/echo/:name', testCtr.echo);
    }
};