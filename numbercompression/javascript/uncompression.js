// Ref. https://stackoverflow.com/questions/14347441/add-or-subtract-from-64bit-integer-in-javascript
// Indefinate length addition
function addAsString(x1, y1) { // x, y strings
    var x = String(x1)
    var y = String(y1)
    var s = '';
    if (y.length > x.length) { // always have x longer
        s = x;
        x = y;
        y = s;
    }
    s = (parseInt(x.slice(-9), 10) + parseInt(y.slice(-9), 10)).toString(); // add last 9 digits
    x = x.slice(0, -9); // cut off last 9 digits
    y = y.slice(0, -9);
    if (s.length > 9) { // if >= 10, add in the 1
        if (x === '') return s; // special case (e.g. 9+9=18)
        x = addAsString(x, '1');
        s = s.slice(1);
    } else if (x.length) { // if more recursions to go
        while (s.length < 9) { // make sure to pad with 0s
            s = '0' + s;
        }
    }
    if (y === '') return x + s; // if no more chars then done, return
    return addAsString(x, y) + s; // else recurse, next digit
}

// Indefinate length subtraction (x - y, |x| >= |y|)
function subtractAsString(x, y) {
    var s;
    s = (parseInt('1' + x.slice(-9), 10) - parseInt(y.slice(-9), 10)).toString(); // subtract last 9 digits
    x = x.slice(0, -9); // cut off last 9 digits
    y = y.slice(0, -9);
    if (s.length === 10 || x === '') { // didn't need to go mod 1000000000
        s = s.slice(1);
    } else { // went mod 1000000000, inc y
        if (y.length) { // only add if makes sense
            y = addAsString(y, '1');
        } else { // else set
            y = '1';
        }
        if (x.length) {
            while (s.length < 9) { // pad s
                s = '0' + s;
            }
        }
    }
    if (y === '') { // finished
        s = (x + s).replace(/^0+/, ''); // dont return all 0s
        return s;
    }
    return subtractAsString(x, y) + s;
}

// Indefinate length addition or subtraction (via above)
function addORsub(x, y) {
    var s = '';
    x = x.replace(/^(-)?0+/, '$1').replace(/^-?$/, '0'); // -000001 = -1
    y = y.replace(/^(-)?0+/, '$1').replace(/^-?$/, '0'); // -000000 =  0
    if (x[0] === '-') { // x negative
        if (y[0] === '-') { // if y negative too
            return '-' + addAsString(x.slice(1), y.slice(1)); // return -(|x|+|y|)
        }
        return addORsub(y, x); // else swap
    }
    if (y[0] === '-') { // x positive, y negative
        s = y.slice(1);
        if (s.length < x.length || (s.length === x.length && s < x)) return subtractAsString(x, s) || '0'; // if |x|>|y|, return x-y
        if (s === x) return '0'; // equal then 0
        s = subtractAsString(s, x); // else |x|<|y|
        s = (s && '-' + s) || '0';
        return s; // return -(|y|-x)
    }
    return addAsString(x, y); // x, y positive, return x+y
}

function multiplyAsString(w, q) {
    addedres = "0"
    for (var m = 0; m < q; m++) {
        var res1 = addAsString(addedres, w)
        addedres = res1
    }
    return addedres
}


var DecodeDic = {
    "A": 0,
    "B": 1,
    "C": 2,
    "D": 3,
    "E": 4,
    "F": 5,
    "G": 6,
    "H": 7,
    "I": 8,
    "J": 9,
    "K": 10,
    "L": 11,
    "M": 12,
    "N": 13,
    "O": 14,
    "P": 15,
    "Q": 16,
    "R": 17,
    "S": 18,
    "T": 19,
    "U": 20,
    "V": 21,
    "W": 22,
    "X": 23,
    "Y": 24,
    "Z": 25,
    "a": 26,
    "b": 27,
    "c": 28,
    "d": 29,
    "e": 30,
    "f": 31,
    "g": 32,
    "h": 33,
    "i": 34,
    "j": 35,
    "k": 36,
    "l": 37,
    "m": 38,
    "n": 39,
    "o": 40,
    "p": 41,
    "q": 42,
    "r": 43,
    "s": 44,
    "t": 45,
    "u": 46,
    "v": 47,
    "w": 48,
    "x": 49,
    "y": 50,
    "z": 51,
    "0": 52,
    "1": 53,
    "2": 54,
    "3": 55,
    "4": 56,
    "5": 57,
    "6": 58,
    "7": 59,
    "8": 60,
    "9": 61,
    "-": 62,
    "_": 63,
};

function UncompresNumber(encoded) {
    var b = Object.keys(DecodeDic).length
    var res = 0
    for (var i = encoded.length - 1; i >= 0; i--) {
        var ch = encoded.charAt(i)
        var v = DecodeDic[ch]
        res = addAsString(multiplyAsString(res, b), v) 
    }
    return res
}

function test() {
    testcases = [
        ["eVTbj9eg-CH", "8124073642988221790"],
        ["ZIAa1CC718B", "2248963214168031769"],
        ["biDGiFEtRkF", "6418109136891295899"],
        ["FT5A5Wd3biE", "5232019302122755269"],
        ["H_IAHrqhZcG", "7429117128365608903"],
        ["ywq0edPJKsC", "3101331938264394802"],
        ["MvUEXirt9mB", "1854839525471243212"],
        ["g5ZFCbwOt0C", "3255323005870775904"],
        ["_O4VcnbyMQ", "291829875975095231"],
        ["clRhcDA8EX", "415720952115304796"],
        ["B2QB5PketUD", "3831853290321415553"],
        ["an54-LdQYEE", "4690571386381638106"],
        ["MqAfU25yeCE", "4656382942409132684"],
        ["9FWmhgvBzTB", "1509557962884604285"],
        ["_z_tvstZXaH", "8545411909543263487"],
        ["Yne5kWtIVZC", "2762152247451576792"],
        ["t_5x4-3C3h", "609968917720834029"],
        ["VbRCAMpPZCD", "3501898986338981589"],
        ["3rKKrYY8_UE", "4989972470701796087"],
        ["cKG4dKqRdYH", "8511036535047479964"],
        ["pcP6oTtW1aD", "3942156919319885609"],
        ["D8z-bvxazLD", "3671395888636051203"],
        ["mxXUHpjeMsB", "1949067129467796582"],
        ["zAkjdb-qVBC", "2329957390282014771"],
        ["YriUGPaEm7G", "7991093983920138968"],
        ["OvECpAP3HeD", "4001409717980646350"],
        ["t2WOAS4IGhH", "8666653584915000749"],
        ["ofWsErhtegH", "8655555759938037736"],
        ["N9FCumb1WVE", "4996415830079594317"],
        ["pJaYLRMRZDH", "8131606212063109737"],
        ["dQQU2rnGGPH", "8342384475171259421"],
        ["9O9OKBUvDHF", "5891760821351732157"],
        ["9ISlCeGc4CD", "3510679499410055741"],
        ["tVImFlHrKs", "795637921042957677"],
        ["sL7IiIro91H", "9042562512790467308"],
        ["cbN2p9dDZED", "3537874235480921820"],
        ["9ooccrh-D4E", "5621611753163098685"],
        ["k-xPVvF69YD", "3908535532744810404"],
        ["8SFqsi_JwkE", "5273759112652936380"],
        ["kFhpO5IsxBB", "1184922302226501988"],
        ["Vq2CB8E6VQ", "294396776680221333"],
        ["5bItKSG4ys", "806954005681178361"],
        ["azDCZsKAA8G", "7998393673069378778"],
        ["KEpv_pmySIF", "5914011819449422090"],
        ["Dw7S6a8HrrC", "3092600507509488643"],
        ["fa3PfYI2_aB", "1639266859947619999"],
        ["YeBT-2WoNUE", "4975810656030431128"],
        ["XgWJuIuIVP", "276165306985310231"],
        ["9mfs85DzRnD", "4166335699046496701"],
        ["Gty0S4NyA6D", "4503820483494619974"],
        ["a7JeaH_-vvF", "6624790593056317146"],
        ["KKi62dq2jIG", "7071736252691260042"],
        ["LMCmfcm-qlH", "8749080546939249419"],
        ["YYE8DhvOsrD", "4245833386669655576"],
        ["2C9av7gRB5B", "2180100724584337590"],
        ["Q7CTb7Y6ZgE", "5195440444922408656"],
        ["raE4RXSWEGF", "5873917832967767723"],
        ["hmRy36nrOeC", "2850407473444624801"],
        ["7PFsnpkUiEC", "2387561231980450811"],
        ["q7HpBrSLJBD", "3479361848781078250"],
        ["vqF6qZXS20C", "3257872153397713583"],
        ["wFgSDbK6NfB", "1715282836036649328"],
        ["qTwUKd_Ei6B", "2207348720189768938"],
        ["p1kzm4UuXNB", "1393786355004099945"],
        ["f_574wiQioC", "3036061856502357983"],
        ["t7xxJpca-_H", "9222925404443451117"],
        ["DFfQM5ZMMoB", "1876929700660638019"],
        ["9GRPmWaHXXE", "5032523705905648061"],
        ["28gY5zgcRdH", "8597778563676114742"],
        ["K1J7XNdyw-E", "5742311434447789386"],
        ["PTieqjXf4XH", "8500682254945232079"],
        ["LE23N8fyPvB", "2004042456489746699"],
        ["pZ98kULwjxC", "3198612044604757609"],
        ["YFfVAhwYiUB", "1522888511094911320"],
        ["RmsubbISXoH", "8797580141140822417"],
        ["xzFtk4QsQWG", "7318544068780514545"],
        ["TZir3QlhCjC", "2937057603268519507"],
        ["hjnl2c4HFB", "19456438989912289"],
        ["vnVcj7HUslC", "2984849158951754223"],
        ["5_e5iDxJqNC", "2551895092338487289"],
        ["GO8cfFQC97B", "2232950891746673542"],
        ["4bEUSWFToTC", "2659459510372943608"],
        ["mLBfTbf0MoG", "7641713525781172966"],
        ["oRStGioeMmE", "5299745588302390376"],
        ["Uc1O9hATQPE", "4886489195041478420"],
        ["sO15Yvi2MEE", "4687361194043331500"],
        ["REFVHYLv2pH", "8824448009872101649"]
    ]

    for (let index = 0; index < testcases.length; index++) {
        var r = UncompresNumber(testcases[index][0])
        if (r !== testcases[index][1]) {
            console.log("error : ", testcases[index], r)
            return
        } else {
            console.log("correct : ", testcases[index], r)
        }
    }
    console.log("No Issues")
    console.log("No Issues")
    console.log("No Issues")

}