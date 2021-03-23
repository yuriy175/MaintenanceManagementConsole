import React, { useState, useEffect, useRef, useContext } from 'react';
import { WebSocketAddress } from '../model/constants'
import { CurrentEquipContext } from '../context/currentEquip-context';
import { AllEquipsContext } from '../context/allEquips-context';

import {sessionUid} from '../utilities/utils'

export function useWebSocket(props) {
    console.log(`useWebSocket `+sessionUid);

    const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);
    const [allEquipsState, allEquipsDispatch] = useContext(AllEquipsContext);
    const [connection, setConnection] = useState(null);
    
    useEffect(() => {
        try {            
            const socket = new WebSocket(WebSocketAddress + "/websocket?uid=" + sessionUid);
            setConnection(socket);
        } catch (e) {
            console.log(e);
        }
    }, []);

    useEffect(() => {
        if (connection) {
            connection.onopen = function () {
                console.log(`Status: Connected ${sessionUid}\n`);
                // connection.send("789 from ui");
            };
        
            connection.onclose = function(event) {
                console.log(`Status: Connected ${sessionUid}\n`);
              };

            connection.onmessage = function (e) {
                console.log("Server: " + e.data + "\n");
                const data = JSON.parse(e.data);
        
                // if(data?.Topic.includes('/ARM/Hardware/HDD'))
                // {
                //     const values = data? JSON.parse(data.Data) : null;
                //     currEquipDispatch({ type: 'SETHDDS', payload: values }); 
                // }
                // else if(data?.Topic.includes('/ARM/Hardware/Memory'))
                // {
                //     const values = data? JSON.parse(data.Data) : null;
                //     currEquipDispatch({ type: 'SETMEMORY', payload: values }); 
                // }
                // else if(data?.Topic.includes('/ARM/Hardware/Processor'))
                // {
                //     const values = data? JSON.parse(data.Data) : null;
                //     currEquipDispatch({ type: 'SETCPU', payload: values }); 
                // }   
                if(data?.Topic.includes('/ARM/Hardware'))
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
                else if(data?.Topic.includes('/organauto'))
                {
                    const values = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETORGANAUTO', payload: values }); 
                }                
                else if(data?.Topic.includes('/stand'))
                {
                    const values = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETSTAND', payload: values }); 
                }
                else if(data?.Topic.includes('/generator'))
                {
                    const values = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETGENERATOR', payload: values }); 
                }
                else if(data?.Topic.includes('/detector'))
                {
                    const values = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETDETECTOR', payload: values }); 
                }
                else if(data?.Topic.includes('/dosimeter'))
                {
                    const values = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETDOSIMETER', payload: values }); 
                }
                else if(data?.Topic.includes('/collimator'))
                {
                    const values = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETCOLLIMATOR', payload: values }); 
                }
                else if(data?.Topic.includes('/ARM/Software'))
                {
                    const values = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETSOFTWARE', payload: values }); 
                }
                else if(data?.Topic.includes('/dicom'))
                {
                    const values = data? JSON.parse(data.Data) : null;
                    currEquipDispatch({ type: 'SETDICOM', payload: values }); 
                }
                else if(data?.Topic.includes('Subscribe'))
                {
                    allEquipsDispatch({ type: 'CONNECTIONCHANGED', payload: data }); 
                }   
            };
        }
    }, [connection]);

    return connection;
}

