const fs = require("fs")

var jsonData;

fs.readFile("./example2.json", (err, data) => {
	jsonData = JSON.parse(data)
	console.log(jsonData)
});

