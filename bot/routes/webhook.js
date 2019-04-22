'use strict';

// expresss
const express = require('express');
const router = express.Router();

// my telegram api
const telegramApi = require('./telegram_api');
// backend api
const backendApi = require('./backend_api');
// string resource
const resString = require('./../resString.json');

function retrieveEventData(requestData) {
    const reqData = requestData.body;

    var event = "";
    var command = "";
    var data = reqData.message.chat;

    // console.log(reqData.message.text);

    if(reqData.message.new_chat_participant && reqData.message.new_chat_participant.id != "575704705"){
        event = "join group";
        data = reqData.message.new_chat_participant;
    } else if(reqData.message.left_chat_participant && reqData.message.left_chat_participant.id != "575704705"){
        event = "left group";
        data = reqData.message.left_chat_participant;
    }else if(reqData.message.text){
        const commandTextArr = reqData.message.text.split(" ");

        if(reqData.message.chat.type == "private" && commandTextArr[0].toLowerCase() == "/start"){
            event = "add bot";
        }else if(reqData.message.chat.type == "private" && commandTextArr[0].toLowerCase() == "/newfriend" && commandTextArr[1] && commandTextArr[2]){
            command = "new friend";
            data.name = commandTextArr;
        }else if(reqData.message.chat.type == "private" && commandTextArr[0].toLowerCase() == "/newfriend"){
            command = "new friend error";
        }else if(commandTextArr[0].toLowerCase() == "/where"){
            command = "where";
            data.name = commandTextArr;
        }else if(commandTextArr[0].toLowerCase() == "/maps"){
            command = "map";
            data.commandArr = commandTextArr;
        }else if(commandTextArr[0].toLowerCase() == "/help"){
            command = "help";
        }else if(commandTextArr[0].toLowerCase() == "/book" && commandTextArr[1] && commandTextArr[2] && commandTextArr[3] && commandTextArr[4]){
            command = "book room";
            data.city = commandTextArr[1].toLowerCase();
            data.capacity = commandTextArr[2].toLowerCase();
            data.start = commandTextArr[3].toLowerCase();
            data.stop = commandTextArr[4].toLowerCase();
        }else if(commandTextArr[0].toLowerCase() == "/book"){
            command = "book room error";            
        }
    }
    
    return {event, command, data};
}

function sendHelp(eventData, response){
    var helpMessage = resString.help + "\n";
    resString.commands.map(cmd => helpMessage += cmd+"\n");
    helpMessage += resString.help2 + "\n";
    resString.special.map(cmd => helpMessage += cmd+"\n");

    telegramApi.sendMessage(eventData.data.id, "Markdown", helpMessage, "", (isSuccess, result) => {
        if(isSuccess){
            response.send("success send help");
        }else{
            console.log(result);
            response.send("some error happened");
        }
    });
}

function createNewUser(eventData, response){
    backendApi.createNewUser(eventData, (isSuccess, result) => {
        if(isSuccess){
            response.send("success create user");
        }else{
            response.send("some error happened");
        }
    });
}

function deleteUser(eventData, response){
    backendApi.deleteUser(eventData, (isSuccess, result) => {
        if(isSuccess){
            response.send("success delete user");
        }else{
            response.send("some error happened");
        }
    });
}

function getOfficeAddress(eventData, response){
    function createKeyboard(data){
        return [{
            text : "/maps " + data.name + ", " + data.city
        }]
    }
    backendApi.getMaps((isSuccess, result) => {
        if(isSuccess){
            const mapsArr = result.data.data;
            const mapsKeyboards = mapsArr.map(map => createKeyboard(map));

            if (eventData.data.commandArr.length > 1){
                const addr = eventData.data.commandArr.slice(1).join(" ");
                const filtered = mapsArr.filter(map => addr == map.name + ", " + map.city);
    
                if (filtered.length == 1){
                    var address = "*" + filtered[0].name + ", " + filtered[0].city + "*\n";
                    address += filtered[0].address + "\n" + filtered[0].location_url;
                    telegramApi.sendMessage(eventData.data.id, "markdown", address, "", (isSuccess, result) => {
                        if(isSuccess){
                            response.send("send office address" + eventData);
                        }else{
                            console.log(result);
                            response.send("some error happaned");                
                        }
                    });
                    return true;
                }
            }
    
            const options = mapsArr.map(map => "- /maps " + map.name + ", " + map.city).join("\n");
            telegramApi.sendMessage(eventData.data.id, "markdown", "Which office do you want to know its location?\n" + options, {
                keyboard : mapsKeyboards
            }, (isSuccess, result) => {
                if (isSuccess){
                    response.send("get office success");
                }else{
                    console.log(result);
                    response.send("some error happened");
                }
            });
        }else{
            response.send("some error happened");
        }
    });
}

function getSomeone(eventData, response){
    function mapTables (table){
        return "- "+ table.office_floor.name + ", Kantor " + table.office.name + ", " + table.office.address  + ", " + table.office.city;        
    }
    function createQueryKeyboard(data){
        return [{
            text : "/where " + data.name
        }]
    }

    backendApi.getTable(eventData, (isSuccess, result) => {
        if(isSuccess){
            const query = eventData.data.name.slice(1).join(" "); 
            const datas = result.data.data;
            const exportData = datas.map(data => mapTables(data));
            console.log(datas);
            if (exportData.length > 1 && query[0] != "@"){
                const message = "There is mutiple table named *" + query + "* please select which one you want";
                const keyboards = datas.map(data => createQueryKeyboard(data));
                telegramApi.sendMessage(eventData.data.id, "markdown", message, {                
                    keyboard : keyboards                    
                }, (isSuccess, result) => {
                    if(isSuccess){
                        response.send("success user found");
                    }else{
                        console.log(result);
                        response.send("some error happened");
                    }
                });
            }else if (exportData.length > 0){
                const message = "*You can find *" + query + "* at:*\n";
                const tables = exportData.join("\n");
                telegramApi.sendMessage(eventData.data.id, "markdown", message + tables, "", (isSuccess, result) => {
                    if(isSuccess){
                        response.send("success user found");
                    }else{
                        console.log(result);
                        response.send("some error happened");
                    }
                });
            }else{
                response.send("success no user found");
            }
        }else{
            console.log(result);
            response.send("some error happened");
        }
    });
}

function getFriend(eventData, response){    
    backendApi.newFriend(eventData, (isSuccess, result) => {
        if(isSuccess){
            var message = "Hi there, we will find you friend(s) for *"+ eventData.data.name[1] + " in " + eventData.data.name[2] + "*.\n";
            if(eventData.data.name[1].toLowerCase() == "lunch"){
                message += "You will be inform again before lunch at 10:00 WIB";
            }else if(eventData.data.name[1].toLowerCase() == "dinner"){
                message += "You will be inform again before dinner at 16:00 WIB";
            }
            telegramApi.sendMessage(eventData.data.id, "markdown", message, "", (isSuccess, result) => {
                if(isSuccess) {
                    response.send("registered for the event");
                }else{
                    response.send("cant plot the event");
                }
            });
        }else{
            console.log("error");
            response.send("some error happened");
        }
    });
}

// /* EVENT HANDLER ROUTE */
router.post("/", (request, response) => {
    var eventData = retrieveEventData(request);

    if(eventData.command == "help"){
        sendHelp(eventData, response);
    }else if(eventData.event == "left group"){
        deleteUser(eventData, response);
    }else if(eventData.event){
        createNewUser(eventData, response);           
    }else if(eventData.command == "map"){
        getOfficeAddress(eventData, response);           
    }else if(eventData.command == "where"){
        getSomeone(eventData, response);         
    }else if(eventData.command == "new friend"){        
        getFriend(eventData, response);
    }else if(eventData.command == "new friend error"){
        // Send Other Request To Alif
        telegramApi.sendMessage(eventData.data.id, "markdown", "please specify event and city (ex: /newfriend lunch bandung)", "", (isSuccess, result) => {
            if(isSuccess){
                response.send("send un recognize help");
            }else{
                console.log(result);
                response.send("not recognize command");                
            }
        });    
    }else{
        // Send Other Request To Alif
        telegramApi.sendMessage(486532003, "markdown", request.body, "", (isSuccess, result) => {
            if(isSuccess){
                response.send("send unknown event to alif");
            }else{
                console.log(result);
                response.send("not recognize command");                
            }
        });
    }
});

module.exports = router;