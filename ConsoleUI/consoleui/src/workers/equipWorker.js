import axios from 'axios';
import { EquipsServiceAddress } from '../model/constants'
import {HandlerWrapper, GetJsonHeader} from './commonWorker'

import {sessionUid} from '../utilities/utils'

const EquipsController = '/equips';

/*export async function GetStudyList(filter) {
    return await HandlerWrapper('StudyList', async () => {
        const response = await axios.post(JournalServiceAddress + JournalController +
            '/StudyList',
            JSON.stringify(filter),
            GetJsonHeader());
        return response.data;
    });
};*/

export async function GetConnectedEquips() {
    return await HandlerWrapper('GetConnectedEquips', async () => {
        const path = EquipsServiceAddress + EquipsController + '/GetConnectedEquips';
        console.log(path);
        const response = await axios.get(path);
        return response.data;
    });
};

export async function Activate(activatedEquipInfo, deactivatedEquipInfo) {
    return await HandlerWrapper('Activate', async () => {
        const response = await axios.get(EquipsServiceAddress + EquipsController +
            '/Activate?sessionUid=' + sessionUid+
            '&activatedEquipInfo=' + activatedEquipInfo+
            '&deactivatedEquipInfo=' + deactivatedEquipInfo);
        return response.data;
    });
};

export async function SearchEquip(currType, equipName, startDate, endDate) {
    return await HandlerWrapper('SearchEquip', async () => {
        const response = await axios.get(EquipsServiceAddress + EquipsController +
            '/SearchEquip?currType=' + currType+
            '&equipName=' + equipName+
            '&startDate=' + startDate+
            '&endDate=' + endDate);
        return response.data;
    });
};

export async function GetPermanentData(currType, equipName) {
    return await HandlerWrapper('GetPermanentData', async () => {
        const response = await axios.get(EquipsServiceAddress + EquipsController +
            '/GetPermanentData?currType=' + currType+
            '&equipName=' + equipName);
        return response.data;
    });
};

export async function RunTeamViewer(activatedEquipInfo) {
    return await HandlerWrapper('RunTeamViewer', async () => {
        const response = await axios.post(EquipsServiceAddress + EquipsController +
            '/RunTeamViewer?activatedEquipInfo=' + activatedEquipInfo);
        return response.data;
    });
};

export async function RunTaskManager(activatedEquipInfo) {
    return await HandlerWrapper('RunTaskManager', async () => {
        const response = await axios.post(EquipsServiceAddress + EquipsController +
            '/RunTaskManager?activatedEquipInfo=' + activatedEquipInfo);
        return response.data;
    });
};

export async function SendAtlasLogs(activatedEquipInfo) {
    return await HandlerWrapper('SendAtlasLogs', async () => {
        const response = await axios.post(EquipsServiceAddress + EquipsController +
            '/SendAtlasLogs?activatedEquipInfo=' + activatedEquipInfo);
        return response.data;
    });
};

export async function XilibLogsOn(activatedEquipInfo, detailedXilib, verboseXilib) {
    return await HandlerWrapper('XilibLogsOn', async () => {
        const response = await axios.post(EquipsServiceAddress + EquipsController +
            '/XilibLogsOn?activatedEquipInfo=' + activatedEquipInfo+
            '&detailedXilib=' + detailedXilib +
            '&verboseXilib=' + verboseXilib);
        return response.data;
    });
};

export async function GetAllEquips() {
    return await HandlerWrapper('GetAllEquips', async () => {
        const response = await axios.get(EquipsServiceAddress + EquipsController + '/GetAllEquips');
        return response.data;
    });
};

export async function GetAllTables(equipName) {
    return await HandlerWrapper('GetAllDBTableNames', async () => {
        const response = await axios.get(EquipsServiceAddress + EquipsController + 
            '/GetAllDBTableNames?equipName=' + equipName);
        return response.data;
    });
};

export async function GetTableContent(equipName, tableType, tableName) {
    return await HandlerWrapper('GetTableContent', async () => {
        const response = await axios.get(EquipsServiceAddress + EquipsController +
            '/GetTableContent?equipName=' + equipName+
            '&tableType=' + tableType+
            '&tableName=' + tableName);
        return response.data;
    });
};

export async function UpdateDBInfo(activatedEquipInfo) {
    return await HandlerWrapper('UpdateDBInfo', async () => {
        const response = await axios.post(EquipsServiceAddress + EquipsController +
            '/UpdateDBInfo?activatedEquipInfo=' + activatedEquipInfo);
        return response.data;
    });
};