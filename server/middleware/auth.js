"use strict";
const mongoose = require('mongoose'),
    globalConfig = require('../../env/config');

const privateFn = {

};

module.exports = {
    check: function(req, res, next){
        if(globalConfig.isDev){
            next();
        }else{
            if(req.headers.sk === globalConfig.secretKey){
                next();
            }else{
                res.json({
                    error: true,
                    msg: 'unAuth'
                });
            }
        }

        return;
    }
};