"use strict";

const testCtr = require('./middleware/test');

module.exports = {
    bindAction: function(server){

        server.get('/echo/:name', testCtr.echo);
        server.get('test',testCtr.test);

    }
};