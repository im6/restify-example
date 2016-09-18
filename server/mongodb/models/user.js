'use strict';

const mongoose = require('mongoose'),
    Schema = mongoose.Schema;

let UserSchema = new Schema({

    id: {
        type: String,
        required: true,
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


mongoose.model('User', UserSchema);

