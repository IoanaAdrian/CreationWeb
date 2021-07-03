let currentCode = [
    0,
    0,
    0,
    0
]
let k=0;
function handleCode(id) {
    currentCode[id[0]] = (k%2==0)?(parseInt(id[1]) + parseInt(currentCode[id[0]])).toString():(parseInt(id[1]) - parseInt(currentCode[id[0]])).toString()
    k++;
    document.getElementById('calcNumber').innerHTML = currentCode.join()
}

