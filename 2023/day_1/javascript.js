// Part 1
document.querySelector('pre').innerText
    .split('\n')
    .filter((row) => row.length > 0)
    .map((row) => row.match(/([0-9])/g))
    .map((row) => row.length === 1 ? parseInt(`${row[0]}${row[0]}`) :  parseInt(`${row[0]}${row[row.length-1]}`))
    .reduce((acc, value) => acc + value, 0)

// Part 2
let replaceObject = {1: 'one', 2: 'two', 3: 'three', 4: 'four', 5: 'five', 6: 'six', 7: 'seven', 8: 'eight', 9: 'nine'};
document.querySelector('pre').innerText.split('\n').filter((row) => row.length > 0).map((row) => {
    let y = [];
    for(let i = 1; i < 10; i++) {
        let x = row;
        let re = new RegExp(replaceObject[i], 'g');
        let matches = x.match(re)
        if (matches) {
            let p = 0;
            while (p < matches.length) {
                let idx = x.indexOf(matches[p]);
                y[x.indexOf(matches[p])] = `${i}`;
                let temp = Array.from(x);
                temp[idx] = 'x';
                x = temp.join('');
                p++;
            }
        } else {
            if (x.indexOf(replaceObject[i] > -1 )) {
                y[x.indexOf(replaceObject[i])] = `${i}`;
            }
        }
        re = new RegExp(i, 'g');
        matches = x.match(re)
        if (matches) {
            let p = 0;
            while (p < matches.length) {
                let idx = x.indexOf(matches[p]);
                y[x.indexOf(matches[p])] = `${i}`;
                let temp = Array.from(x);
                temp[idx] = 'x';
                x = temp.join('');
                p++;
            }
        } else {
            if (x.indexOf(i) > -1) {
                y[x.indexOf(i)] = `${i}`;
            }
        }
    }
    return y.flatMap((v) => v).join('')}
)
.map((row) => row.match(/([0-9])/g))
.map((row) => row.length === 1 ? parseInt(`${row[0]}${row[0]}`) :  parseInt(`${row[0]}${row[row.length-1]}`))
.reduce((acc, value) => acc + value, 0)
