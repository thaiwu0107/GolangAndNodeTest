rankTable7CF = require("./RankTable7CF.json");
rankTable7CNF = require("./RankTable7CNF.json");

class PokerCalculator {
    constructor() {
    }

    getRankInfo(inputCards) {

        const cardMap = [
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
        ]

        let point;
        let suit;
        let suitMap = {};
        for (let i = 0; i < 7; i++) {
            suit = inputCards[i] % 10;
            point = (inputCards[i] - suit) / 10;

            cardMap[suit][point]++;
            cardMap[suit][0]++;
            cardMap[0][point]++;

            if (suitMap[point]) {
                suitMap[point].push(suit);
            }
            else {
                suitMap[point] = [suit];
            }

        }

        let isFlush = false;
        let selectSuit = 0;
        for (let i = 0; i < 5; i++) {
            if (cardMap[i][0] >= 5) {
                isFlush = true;
                selectSuit = i;
                break;
            }
        }

        let keyOfRank = this.getKey(cardMap[selectSuit]);
        let rankInfo;
        let selectCards = [];
        if (isFlush) {
            rankInfo = rankTable7CF[keyOfRank];
            for (let i = 0; i < 5; i++) {
                selectCards[i] = rankInfo.CardPoint[i] * 10 + selectSuit;
            }
        } else {
            rankInfo = rankTable7CNF[keyOfRank];
            for (let i = 0; i < 5; i++) {
                selectCards[i] = rankInfo.CardPoint[i] * 10 + suitMap[rankInfo.CardPoint[i]].pop();
            }
        }
        //rankInfo.SelectSuit = selectSuit;
        rankInfo.SelectCards = selectCards;
        return rankInfo;
    }

    getKey(cardArray) {
        return `${cardArray[14]}${cardArray[13]}${cardArray[12]}${cardArray[11]}${cardArray[10]}${cardArray[9]}${cardArray[8]}${cardArray[7]}${cardArray[6]}${cardArray[5]}${cardArray[4]}${cardArray[3]}${cardArray[2]}`;
    }

}

module.exports = PokerCalculator;
