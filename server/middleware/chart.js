
"use strict";
const mongoose = require('mongoose'),
    helper = require('../misc/helper');

const privateFn = {

};

module.exports = {
    createChart: function(req, res, next){
        var Chart = mongoose.model('Chart');
        let bodyObj = req.body;
        var char1 = new Chart({
            userId: bodyObj.userId,
            obj: bodyObj.obj
        });

        char1.save().then((data) =>{
            res.json(helper.resSuccessObj(data));
        }, (data) =>{
            res.json(helper.resFailObj(data));
        });
    },
    getChart: function(req, res, next){
        var Chart = mongoose.model('Chart');
        Chart.find({userId: req.query.userId}).then((data)=>{
            res.json(helper.resSuccessObj(data));
        }, (data)=>{
            res.json(helper.resFailObj(data));
        });

    },
    updateChart: function(req, res, next){
        res.send(req.params);
    },
    deleteChart: function(req, res, next){
        var Chart = mongoose.model('Chart');
        Chart.find({ userId: 0 }).remove().then((data) =>{
            res.json(helper.resSuccessObj(data));
        }, (data)=>{
            res.json(helper.resFailObj(data));
        });
    }
};