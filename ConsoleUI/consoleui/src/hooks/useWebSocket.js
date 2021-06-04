import React, { useState, useEffect, useRef, useContext } from 'react';
import { WebSocketAddress } from '../model/constants'
import { CurrentEquipContext } from '../context/currentEquip-context';
import { SystemVolatileContext } from '../context/systemVolatile-context';
import { AllEquipsContext } from '../context/allEquips-context';
import * as EquipWorker from '../workers/equipWorker'

import {sessionUid, getEquipFromTopic} from '../utilities/utils'

export function useWebSocket(props) {
    console.log(`useWebSocket `+sessionUid);

    const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);
    const [allEquipsState, allEquipsDispatch] = useContext(AllEquipsContext);
    const [systemVolatileState, systemVolatileDispatch] = useContext(SystemVolatileContext);
    const [connection, setConnection] = useState(null);
    const equipInfo = useRef(currEquipState.equipInfo);
    
    function createSocket(){
        try {   
            console.log(`Status: Creating socket ${sessionUid}\n`);         
            const socket = new WebSocket(WebSocketAddress + "/websocket?uid=" + sessionUid);
            setConnection(socket);
        } catch (e) {
            console.log(e);
        }
    }

    useEffect(() => {
        createSocket();
    }, []);

    useEffect(() => {
        equipInfo.current = currEquipState.equipInfo;
    }, [currEquipState.equipInfo]);

    useEffect(() => {
        if (connection) {
            connection.onopen = async function () {
                console.log(`Status: Connected ${sessionUid}\n`);

                const equips = await EquipWorker.GetConnectedEquips();
                allEquipsDispatch({ type: 'SETCONNECTEDEQUIPS', payload: equips ? equips : [] });     

                // Send a ping every 10s
                // to keep the connection live
                setInterval(function(){
                    console.log(`pinging...\n`);
                    connection.send('ping');
                }, 10000);
            };
        
            connection.onclose = function(event) {
                console.log(`Status: Closed ${sessionUid}\n`);
                currEquipDispatch({ type: 'RESET', payload: true });    
                setTimeout(function() {
                    console.log(`Status: Reconnecting ${sessionUid}\n`);
                    createSocket();
                  }, 1000);
              };
              
            connection.onerror = function(err) {
                console.error('Socket encountered error: ', err.message, 'Closing socket');
                connection.close();
              };

            connection.onmessage = function (e) {
                console.log("Server: " + e.data + "\n");
                const data = JSON.parse(e.data);
        
                const topic = data?.Topic;
                if(!topic){
                    return;
                }

                if(topic.startsWith('Subscribe'))
                {
                    allEquipsDispatch({ type: 'CONNECTIONCHANGED', payload: data }); 
                    return;
                }   

                const equip = getEquipFromTopic(data?.Topic);
                if(!equip || equip !== equipInfo.current){
                    return;
                }

                const path = data.Topic.replace(equip, '');
                if(path.startsWith('/ARM/Hardware/Complex'))
                {
                    try
                    {
                        const values = data? JSON.parse(data.Data) : null;
                        currEquipDispatch({ type: 'SETSYSTEM', payload: values }); 
                    }
                    catch(e)
                    {
                        console.log(e);
                    }                    
                }
                else if(path.startsWith('/images'))
                {
                    const values = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETIMAGES', payload: values }); 
                }            
                else if(path.startsWith('/organauto'))
                {
                    const values = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETORGANAUTO', payload: values }); 
                }                
                else if(path.startsWith('/stand'))
                {
                    const values = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETSTAND', payload: values }); 
                }
                else if(path.startsWith('/generator'))
                {
                    const values = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETGENERATOR', payload: values }); 
                }
                else if(path.startsWith('/detector'))
                {
                    const values = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETDETECTOR', payload: values }); 
                }
                else if(path.startsWith('/dosimeter'))
                {
                    const values = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETDOSIMETER', payload: values }); 
                }
                else if(path.startsWith('/collimator'))
                {
                    const values = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETCOLLIMATOR', payload: values }); 
                }
                else if(path.startsWith('/aec/'))
                {
                    const values = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETAEC', payload: values }); 
                }
                else if(path.startsWith('/ARM/Software/Complex'))
                {
                    const values = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETSOFTWARE', payload: values }); 
                }
                else if(path.startsWith('/ARM/Software/msg'))
                {
                    const values = data? JSON.parse(data.Data) : null;
                    // currEquipDispatch({ type: 'SETSOFTWAREMSG', payload: values }); 
                    systemVolatileDispatch({ type: 'SETVOLATILE', payload: values }); 
                }
                else if(path.startsWith('/dicom'))
                {
                    const values = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETDICOM', payload: values }); 
                }
                else if(path.startsWith('/remoteaccess'))
                {
                    const values = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETREMOTEACCESS', payload: values }); 
                }
                else if(path.startsWith('/ARM/Hardware/Processor') ||
                        path.startsWith('/ARM/Hardware/HDD') || 
                        path.startsWith('/ARM/Hardware/Memory'))
                {
                    const values = data? JSON.parse(data.Data) : null;
                    systemVolatileDispatch({ type: 'SETVOLATILE', payload: values }); 
                }
            };
        }
    }, [connection]);

    return connection;
}

