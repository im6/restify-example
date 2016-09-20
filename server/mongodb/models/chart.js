'use strict';

const mongoose = require('mongoose'),
    Schema = mongoose.Schema,
    autoIncrementPlugin = require('../config').getAutoIncPlugin();

let ChartSchema = new Schema({

    id: {
        type: Number,
        required: true
    },

    userId: {
        type: Number,
        required: true
    },

    obj: {
        type: Object,
        required: true
    }
});


ChartSchema.pre('save', function(next){
    next();
});

ChartSchema.plugin(autoIncrementPlugin, {
    model: 'Chart',
    field: 'id',
    startAt: 1,
    incrementBy: 1
});


mongoose.model('Chart', ChartSchema);