"use strict";

module.exports = {
    test: function(req, res, next){
        res.send({
            msg: "test success!"
        });
        return next();
    },
    echo: function(req, res, next){
        res.send(req.params);
        return next();
    }
};