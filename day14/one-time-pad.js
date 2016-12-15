var salt = "ahsbgdzn";
// var salt = "abc";
var md5 = require('md5');
var i = 0;
var threeRe = /(.)\1\1/;
var hashes;
var keys = [];
var found = 0;

var findFiveRepeats = function(character) {
    var fiveRe = new RegExp(character.repeat(5));
    var hasFive = false;
    hashes.forEach(function(e) {
        if (fiveRe.test(e)) {
            hasFive = true;
        }
    });

    return hasFive;
}

var getHash = function(index) {
    return md5(salt + i);
}

var getHash2 = function(index) {
    var hash = getHash(index);
    for (var x = 0; x < 2016; x++) {
        hash = md5(hash);
    }

    return hash;
}

var findPad = function(hashMethod) {
    hashes = [];

    for (i = 0; i < 1001; i++) {
        hashes.push(hashMethod(i));
    }

    while (true) {
        hash = hashes.shift();
        md = hash.match(threeRe);

        if (md && findFiveRepeats(md[1])) {
            if (++found >= 64) {
                console.log(i - 1000);
                break;
            }
        }

        hashes.push(hashMethod(++i));
    }
}


findPad(getHash);
