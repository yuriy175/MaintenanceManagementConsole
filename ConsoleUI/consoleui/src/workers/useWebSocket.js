import React, { useState, useEffect, useRef, useContext } from 'react';
import { WebSocketAddress } from '../model/constants'
import { CurrentEquipContext } from '../context/currentEquip-context';
import { AllEquipsContext } from '../context/allEquips-context';

import {sessionUid, getEquipFromTopic} from '../utilities/utils'

export function useWebSocket(props) {
    console.log(`useWebSocket `+sessionUid);

    const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);
    const [allEquipsState, allEquipsDispatch] = useContext(AllEquipsContext);
    const [connection, setConnection] = useState(null);
    const equipInfo = useRef(currEquipState.equipInfo);
    
    function createSocket(){
        try {            
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
            connection.onopen = function () {
                console.log(`Status: Connected ${sessionUid}\n`);
                // connection.send("789 from ui");
            };
        
            connection.onclose = function(event) {
                console.log(`Status: Closed ${sessionUid}\n`);
                setTimeout(function() {
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
        
                const topic = getEquipFromTopic(data?.Topic);
                if(!topic || topic !== equipInfo.current){
                    return;
                }

                const path = data.Topic.replace(topic, '');
                if(path.startsWith('/ARM/Hardware'))
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
                else if(path.startsWith('/ARM/Software'))
                {
                    const values = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETSOFTWARE', payload: values }); 
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
                else if(path.startsWith('Subscribe'))
                {
                    allEquipsDispatch({ type: 'CONNECTIONCHANGED', payload: data }); 
                }   
            };
        }
    }, [connection]);

    return connection;
}

