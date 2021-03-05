import React, { useState, useEffect, useRef, useContext } from 'react';
import { WebSocketAddress } from '../model/constants'
import { CurrentEquipContext } from '../context/currentEquip-context';

const generateSessionUid = function () { // Public Domain/MIT
    var d = new Date().getTime();//Timestamp
    var r = Math.random() * 1000;//random number between 0 and 1000
    var d2 = (performance && performance.now && (performance.now()*1000)) || 0;//Time in microseconds since page-load or 0 if unsupported
    return `${d}_${r}`;
}

export function useWebSocket(props) {
    console.log(`useWebSocket`);

    const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);
    const [connection, setConnection] = useState(null);
    
    useEffect(() => {
        const socket = new WebSocket(WebSocketAddress + "/websocket?uid=" + generateSessionUid());
        setConnection(socket);
    }, []);

    useEffect(() => {
        if (connection) {
            connection.onopen = function () {
                console.log("Status: Connected\n");
                // connection.send("789 from ui");
            };
        
            connection.onmessage = function (e) {
                console.log("Server: " + e.data + "\n");
                const data = JSON.parse(e.data);
        
                if(data?.Topic.includes('/ARM/Hardware/HDD'))
                {
                    const hdds = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETHDDS', payload: hdds }); 
                }
                else if(data?.Topic.includes('/ARM/Hardware/Memory'))
                {
                    const hdds = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETMEMORY', payload: hdds }); 
                }
                else if(data?.Topic.includes('/ARM/Hardware/Processor'))
                {
                    const hdds = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETCPU', payload: hdds }); 
                }                
                else if(data?.Topic.includes('/organauto'))
                {
                    const hdds = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETORGANAUTO', payload: hdds }); 
                }
                else if(data?.Topic.includes('/generator'))
                {
                    const hdds = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETGENERATOR', payload: hdds }); 
                }
                else if(data?.Topic.includes('/detector'))
                {
                    const hdds = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETDETECTOR', payload: hdds }); 
                }
            };
        }
    }, [connection]);

    return connection;
}

