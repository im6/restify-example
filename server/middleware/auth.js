"use strict";
const mongoose = require('mongoose'),
    globalConfig = require('../../env/config');

const privateFn = {

};

module.exports = {
    check: function(req, res, next){
      if(req.url=== '/'){
        next();
      }else if(req.authorization.basic &&
        req.authorization.basic.username + '|' + req.authorization.basic.password === globalConfig.credential
      ) {
          next();
      }else{
          res.json({
              error: true,
              msg: 'unAuth'
          });
      }
      return;
    }
};