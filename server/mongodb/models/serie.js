'use strict';

const mongoose = require('mongoose'),
    Schema = mongoose.Schema,
    autoIncrementPlugin = require('../config').getAutoIncPlugin();

let SerieSchema = new Schema({
  id: {
      type: Number,
      unique: true
  },
  name: {
      type: String,
      required: true
  },

  category: {
      type: Array
  },

  data: {
    type: Array,
    required: true
  }
});


SerieSchema.pre('save', function(next){
  var me = this;
  if(me.category.length > 0 && me.category.length != me.data.length){
    var err = new Error('data and category mismatched.');
    console.error('data and category mismatched.');
    next(err);
  }else{
    next();
  }
});

SerieSchema.plugin(autoIncrementPlugin, {
    model: 'Serie',
    field: 'id',
    startAt: 1,
    incrementBy: 1
});


mongoose.model('Serie', SerieSchema);

