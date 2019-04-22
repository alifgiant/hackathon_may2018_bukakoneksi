'use strict';
// Express.JS as framework
const express = require("express");
const app = express();

// Middleware
const cors = require("cors");
app.use(cors({ origin: true }));

// cookie parser
const cookieParser = require('cookie-parser');
app.use(cookieParser());

// body message parse
const bodyParser = require('body-parser');
// parse application/x-www-form-urlencoded
app.use(bodyParser.urlencoded({ extended: false }));
// parse application/json
app.use(bodyParser.json());

// routing webhook
const webhook = require('./routes/webhook');
app.use('/webhook', webhook);

// routing middleware
const index = require('./routes/api');
app.use('/api', index);

// catch 404 and forward to error handler
app.use((req, res, next) => {
    var err = new Error('Not Found');
    err.status = 404;
    next(err);
});

// error handler
app.use((err, req, res, next) => {
    // set locals, only providing error in development
    res.locals.message = err.message;
    res.locals.error = req.app.get('env') === 'development' ? err : {};

    // render the error page
    res.status(err.status || 500);
    res.send('error page');
});

module.exports = app