var fs = require("fs");
var readLine = require("readline");
const cidrTools = require('cidr-tools');

function readFileToArr(fReadName, cb) {

    var arr = [];
    var readObj = readLine.createInterface({
        input: fs.createReadStream(fReadName)
    });

    readObj.on('line', function (line) {
        arr.push(line);
    });
    readObj.on('close', function () {
        console.log('readLine close....');
        cb(arr);
    });
}

readFileToArr('whole.text', function (arr) {
require("fs").writeFile(
     './merge.text',
     cidrTools.merge(arr).map(function(v){ return v }).join('\n'),
     function (err) { console.log(err ? 'Error :'+err : 'ok') }
);

});
