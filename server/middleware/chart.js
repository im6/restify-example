
"use strict";
const mongoose = require('mongoose'),
    helper = require('../misc/helper');

const privateFn = {

};

module.exports = {
    createChart: function(req, res, next){
        var Chart = mongoose.model('Chart');

        var char1 = new Chart({});

        char1.save().then((data) =>{
            console.log(data);
            res.send({
                error: false
            });
        }, (data) =>{
            res.send({
                error: true,
                detail: err
            });
            console.error(err);
        });

        return next();
    },
    getChart: function(req, res, next){
        var Chart = mongoose.model('Chart');
        Chart.find({id: req.query.id}).then((data)=>{
            res.send(helper.resSuccessObj(data));
        }, (data)=>{
            res.send(helper.resFailObj(data));
        });

        return next();
    },
    updateChart: function(req, res, next){
        res.send(req.params);
        return next();
    },
    deleteChart: function(req, res, next){
        res.send(req.params);
        return next();
    }
};