
"use strict";
const mongoose = require('mongoose'),
    helper = require('../misc/helper');

const privateFn = {

};


module.exports = {
    createUser: function(req, res, next){

        var User = mongoose.model('User');
        let bodyObj = req.body;

        var userOne = new User({
            name: bodyObj.objWeibo.name,
            uidWeibo: bodyObj.objWeibo.id,
            objWeibo: bodyObj.objWeibo
        });

        userOne.save().then((data) =>{
            res.send(helper.resSuccessObj(data));
        }, (err) => {
            res.send(helper.resFailObj(err));
            console.error(err);
        });

        return next();

    },
    getUser: function(req, res, next){
        var User = mongoose.model('User');
        User.find({id: req.query.id}).then((data)=>{
            res.send(helper.resSuccessObj(data));
        }, (data)=>{
            res.send(helper.resFailObj(data));
        });
        return next();
    },
    updateUser: function(req, res, next){
        res.send(req.params);
        return next();
    },
    deleteUser: function(req, res, next){
        res.send(req.params);
        return next();
    }
};