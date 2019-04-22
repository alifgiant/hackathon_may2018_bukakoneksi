'use strict';

// expresss
const express = require('express');
const router = express.Router();

const axios = require('axios');

const apiKey = require('./../package.json').telegramApiKey;
const telegramBaseUrl = "https://api.telegram.org/bot"+ apiKey;

// backend api
const backendApi = require('./backend_api');
// my telegram api
const telegramApi = require('./telegram_api');

const flatMap = (f,xs) =>
  xs.reduce((acc,x) =>
    acc.concat(f(x)), []);

// /* GET ME. */
router.get("/", (request, response) => {
    axios.get(telegramBaseUrl+'/getMe').then(res => {
        response.send(res.data);
        return true;
    })
    .catch(error => {
        console.log(error);
    });
});

// /* GET PROFILE PIC - of specifi userId*/
router.get("/pic/:userId", (request, response) => {
    axios.post(telegramBaseUrl+'/getUserProfilePhotos', {
        user_id : request.params.userId
    }).then(res => {
        const photos = res.data.result.photos;
        const axiosReq = flatMap(photo => photo.map(version => 
            axios.post(telegramBaseUrl+"/getFile", {
                "file_id" : version.file_id
            })
        ), photos);

        if (axiosReq.length > 0){
            axios.all(axiosReq).then(axios.spread(function (acct, perms) {
                axios.post(telegramBaseUrl+'/getUserProfilePhotos', {
                    user_id : request.params.userId
                }).then(res2 => {
                    // res2.data.result.
                    response.send({
                        rootPath : "https://api.telegram.org/file/bot"+apiKey+"/",
                        expire : (new Date()).valueOf() + (1 /* hour */ *60*60*1000),
                        data : res2.data.result.photos
                    });
                }).catch(error => {
                    console.log(error);
                    response.send("error")
                });
            })).catch(error => {
                console.log(error);
                response.send("error")
            });
        }else {
            response.send({ msg: "no picture found" });
        }
        
    })
    .catch(error => {
        console.log(error);
        response.send("error")
    });
});

// /* SEND MESSAGE. */
router.post("/message/:type/:userId", (request, response) => {
    axios.post(telegramBaseUrl+'/sendMessage', 	{
        "chat_id" : request.params.userId,
        "parse_mode" : request.params.type,
        "text" : request.body.data
    }).then(res => {
        response.send("success send message");
        return true;
    })
    .catch(error => {
        console.log(error);
        response.send("some error happened");
    });      
});

function setupPools(users){
    const shuffledUser = users
        .map((a) => ({sort: Math.random(), value: a}))
        .sort((a, b) => a.sort - b.sort)
        .map((a) => a.value);
    
    var start = 0;
    var pools = []    
    while (start <= shuffledUser.length - 1){
        pools.push(shuffledUser.slice(start, start+4));
        start += 4;
        
    }
    if (pools.length > 2 && pools[pools.length-1] < 4){
        pools[pools.length-2].concat(pools[pools.length-1]);
        pools = pools.slice(0, pools.length - 1);
    }
    return pools;
}

function informPool(pools, event){
    pools.map(pool => {
        const users = pool.map(user => "@" + user.telegram_username).join(" ");            
        var message = "Here your " + event + " friends. " + users + " feel free to setup a meeting point.";        
        if (pool.length < 2){
            message = "We are sorry, currently no other user want to meet new friend for " + event;
        }
        pool.map(user => {
            telegramApi.sendMessage(user.telegram_user_id, "markdown", message, "", (isSuccess, result) => {
                if (isSuccess){
                    console.log("user " + user.telegram_user_id + " is informed");
                }else{
                    console.log("user " + user.telegram_user_id + " cant informed");
                }
            });
        });
    });
}

const cities = ["jakarta", "bandung"];
router.get("/lunchpool", (req, res) => {
    for (var j=0; j<cities.length; j++){
        backendApi.getFriendPool("lunch", cities[j], (isSuccess, result) => {
            if (isSuccess){
                const pools = setupPools(result.data.data);
                informPool(pools, "lunch");
            }else{
                console.log("cant create pool");
            }
        });
    }
    
    res.send("success");
});

router.get("/dinnerpool", (req, res) => {
    for (var j=0; j<cities.length; j++){
        backendApi.getFriendPool("dinner", cities[j], (isSuccess, result) => {
            if (isSuccess){
                const pools = setupPools(result.data.data);        
                informPool(pools, "dinner");
            }else{
                console.log("cant create pool");
            }
        });
    }
    
    res.send("success");
});

module.exports = router;