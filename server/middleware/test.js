"use strict";
const mongoose = require('mongoose');

const privateFn = {
    getMongoStatus: (stateId) => {
        var stateStr = null;
        switch (stateId) {
            case 0:
                stateStr = "disconnected";
                break;
            case 1:
                stateStr = "connected";
                break;
            case 2:
                stateStr = "connecting";
                break;
            case 3:
                stateStr = "disconnecting";
                break;
            default:
                stateStr = "unknown";
        }
        return stateStr;
    }
};

module.exports = {
    root: function(req, res, next){
        var stateId = mongoose.connection.readyState;
        res.send({
            msg: "test success!",
            db: privateFn.getMongoStatus(stateId),

        });
        return next();
    },
    echo: function(req, res, next){
        console.log('come to echo');
        res.send(req.params);
        return next();
    }
};