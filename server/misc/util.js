"use strict";
let privateFn = {

};

module.exports = {
    resSuccessObj: (data, msg = null)=>{
        return {
            error: false,
            data: null,
            msg: null
        };
    },
    resFailObj: (err, msg = null)=>{
        return {
            error: true,
            data: err,
            msg: msg
        };
    },

};