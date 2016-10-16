
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
        return next();
    },
    getChart: function(req, res, next){
      var Chart = mongoose.model('Chart');
      var Serie = mongoose.model('Serie');

      var filter = {};
      if(req.query.userId) filter['userId'] = req.query.userId;
      if(req.query.id) filter['id'] = req.query.id;
      Chart.find(filter).then((data)=>{

        var sid = [];
        data.forEach((v,k)=>{
          sid = sid.concat(v.series);
        });

        Serie.find({id: {
          $in: sid
        }}).then((data1) => {
          var result = {
            charts: data,
            series: data1
          };
          res.json(helper.resSuccessObj(result));

        });
      }, (data)=>{
          res.json(helper.resFailObj(data));
      });

      return next();

    },
    updateChart: function(req, res, next){
        res.send(req.params);
        return next();
    },
    deleteChart: function(req, res, next){
        var Chart = mongoose.model('Chart');
        Chart.find({ userId: 0 }).remove().then((data) =>{
            res.json(helper.resSuccessObj(data));
        }, (data)=>{
            res.json(helper.resFailObj(data));
        });
        return next();
    }
};