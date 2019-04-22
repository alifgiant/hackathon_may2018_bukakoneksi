// http client
const axios = require('axios');
// backend base url
const backendBaseUrl = "http://963221cf.ngrok.io";

function createNewUser(eventData, callback){
    axios.post(backendBaseUrl+'/members',{
        "name" : eventData.data.first_name + " " + eventData.data.last_name,
        "telegram_user_id" : eventData.data.id.toString(),
        "telegram_username" : eventData.data.username
    }).then(res => {
        callback(true, res);
    })
    .catch(error => {
        callback(false, error);
    });  
}

function deleteUser(eventData, callback){
    axios.delete(backendBaseUrl+'/members', {
        data : {
            "telegram_user_id" : eventData.data.id.toString()
        }
    }).then(res => {
        callback(true, res);
    })
    .catch(error => {
        callback(false, error);
    });  
}

function getMaps(callback){
    axios.get(backendBaseUrl+'/maps').then(res => {
        callback(true, res);
    })
    .catch(error => {
        callback(false, error);
    });  
}

function getTable(eventData, callback){
    const query = eventData.data.name.slice(1).join(" "); 
    axios.get(backendBaseUrl+"/tables", {
        params : {
            search : query,
            source : "bot"
        }
    }).then((res) => {
        // console.log(res.data);
        callback(true, res);
    })
    .catch((error) => {
        callback(false, error);
    });  
}

function newFriend(eventData, callback){
    axios.post(backendBaseUrl+'/friendship_events',{
        event : eventData.data.name[1],
        city : eventData.data.name[2],
        telegram_user_id : eventData.data.id.toString(),
        telegram_username : eventData.data.username
    }).then(res => {
        callback(true, res);
    })
    .catch(error => {
        callback(false, error);
    });
}

function getFriendPool(event, city, callback){
    axios.get(backendBaseUrl+'/friendship_events', {
        params: {
            city : city,
            event : event
        }
    }).then(res => {
        callback(true, res);
    })
    .catch(error => {
        callback(false, error);
    });
}

module.exports = {
    createNewUser, deleteUser, getMaps, getTable, newFriend, getFriendPool
}