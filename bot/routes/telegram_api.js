const axios = require('axios');

const apiKey = require('./../package.json').telegramApiKey;
const telegramBaseUrl = "https://api.telegram.org/bot"+ apiKey;

function sendMessage(chatId, parseMode, content, reply, callback){    
    axios.post(telegramBaseUrl+'/sendMessage', 	{
        "chat_id" : chatId,
        "text" : content,
        "parse_mode" : parseMode,
        "reply_markup" : reply
    }).then(res => {
        callback(true, res);
    })
    .catch(error => {
        callback(false, error);
    });
}

module.exports = {
    sendMessage
}