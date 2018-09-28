const CalClass = require("./PokerCalculator");
const config = require("./config.json");

let time;
let diff;


//計時
const NS_PER_SEC = 1e9;
const MS_PER_NS = 1e-6
const mysim = new CalClass();
//單場
time = process.hrtime();
let inputCards = [144, 131, 121, 122, 91, 24, 132];
for (let index = 0; index < config.testloop; index++) {
    mysim.getRankInfo(inputCards);
}
diff = process.hrtime(time);
console.log(`Nodejs - Benchmark: ${config.testloop} times -> took: ${diff[0] * NS_PER_SEC + diff[1]} nanoseconds`);
console.log(`Nodejs - Benchmark: ${config.testloop} times -> took: ${(diff[0] * NS_PER_SEC + diff[1]) * MS_PER_NS} milliseconds`);
console.log(`Nodejs - Benchmark each times -> took: ${~~((diff[0] * NS_PER_SEC + diff[1]) / config.testloop)} ns/op`);
console.log(`Nodejs - Benchmark each times -> took: ${~~(((diff[0] * NS_PER_SEC + diff[1]) / config.testloop) / 1000)} µs/op`);
