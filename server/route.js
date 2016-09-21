"use strict";

const testCtr = require('./middleware/test');
const userCtr = require('./middleware/user');
const chartCtr = require('./middleware/chart');

module.exports = {
    bindAction: (server) => {
        server.get('/', testCtr.root);
        server.get('/echo/:name', testCtr.echo);

        server.get('/user', userCtr.getUser);
        server.delete('/user', userCtr.deleteUser);
        server.post('/user', userCtr.updateUser);
        server.post('/user/new', userCtr.createUser);


        server.get('/chart', chartCtr.getChart);
        server.delete('/chart', chartCtr.deleteChart);
        server.post('/chart', chartCtr.updateChart);
        server.post('/chart/new', chartCtr.createChart);

    }
};