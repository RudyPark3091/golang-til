const fs = require("fs")
const assert = require("assert")

var jsonData;

fs.readFile("./example2.json", (err, data) => {
	jsonData = JSON.parse(data)
	console.log(jsonData)

	assert.equal(jsonData.id, "0001")
	assert.equal(jsonData.type, "donut")
	assert.equal(jsonData.ppu, 0.55)
	assert.equal(typeof jsonData.batters, "object")
	assert.equal(typeof jsonData.batters.batter[0], "object")
	assert.equal(jsonData.batters.batter[0].id, "1001")
	assert.equal(jsonData.batters.batter[1].type, "Chocolate")
	assert.equal(typeof jsonData.topping[0], "object")
	assert.equal(jsonData.topping[0].id, "5001")
	assert.equal(jsonData.topping[2].type, "Sugar")
});

