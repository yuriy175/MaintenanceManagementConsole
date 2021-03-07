const generateSessionUid = function () { // Public Domain/MIT
    var d = new Date().getTime();//Timestamp
    var r = Math.random() * 1000;//random number between 0 and 1000
    var d2 = (performance && performance.now && (performance.now()*1000)) || 0;//Time in microseconds since page-load or 0 if unsupported
    return `${d}_${r}`;
}

export const sessionUid = generateSessionUid();