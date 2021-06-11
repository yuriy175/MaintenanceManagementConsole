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

export function getEquipFromTopic(topic ){
	const topicParts = topic.split("/");
    if(topicParts.length < 2){
        return null;
    }

	const equip = `${topicParts[0]}/${topicParts[1]}`;

	return equip;
}

export const parseLocalString = (value) => new Date(value).toLocaleString();

export const isToday = (value) =>
{
    const today = new Date();
    const date = new Date(value);
    return date.setHours(0,0,0,0) == today.setHours(0,0,0,0);
}

export const sessionUid = generateSessionUid();