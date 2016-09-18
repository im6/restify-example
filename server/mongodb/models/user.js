'use strict';

const mongoose = require('mongoose'),
    Schema = mongoose.Schema,
    autoIncrementPlugin = require('../config').getAutoIncPlugin();

let UserSchema = new Schema({
    id: {
        type: Number,
        unique: true
    },

    name: {
        type: String,
        required: true
    },

    uidWeibo: {
        type: Number,
        unique: true
    },

    objWeibo: {
        type: Object
    },

    isAdmin: {
        type: Boolean,
        default: false
    }
});


UserSchema.pre('save', function(next){
    next();
});

UserSchema.plugin(autoIncrementPlugin, {
    model: 'User',
    field: 'id',
    startAt: 1,
    incrementBy: 1
});


mongoose.model('User', UserSchema);

