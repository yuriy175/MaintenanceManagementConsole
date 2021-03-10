const generateSessionUid = function () { // Public Domain/MIT
    var d = new Date().getTime();//Timestamp
    var r = Math.random() * 1000;//random number between 0 and 1000
    var d2 = (performance && performance.now && (performance.now()*1000)) || 0;//Time in microseconds since page-load or 0 if unsupported
    return `${d}_${r}`;
}

const leadZero = (val) => val < 10 ? '0' + val : val;
export function getUSFullDate(date)
{
    return date.getFullYear() + "-" + leadZero(date.getMonth() + 1) + "-"+ leadZero(date.getDate());
}

export const sessionUid = generateSessionUid();