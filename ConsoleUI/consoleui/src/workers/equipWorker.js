import axios from 'axios';
import { EquipsServiceAddress } from '../model/constants'
import {HandlerWrapper, GetJsonHeader, GetTokenHeader} from './commonWorker'

import {sessionUid} from '../utilities/utils'

const EquipsController = '/equips';

export async function GetConnectedEquips(token) {
    return await HandlerWrapper('GetConnectedEquips', async () => {
        const path = EquipsServiceAddress + EquipsController + '/GetConnectedEquips';
        console.log(path);
        const response = await axios.get(path, GetTokenHeader(token));
        return response.data;
    });
};

export async function Activate(token, activatedEquipInfo, deactivatedEquipInfo) {
    return await HandlerWrapper('Activate', async () => {
        const response = await axios.get(EquipsServiceAddress + EquipsController +
            '/Activate?sessionUid=' + sessionUid+
            '&activatedEquipInfo=' + activatedEquipInfo+
            '&deactivatedEquipInfo=' + deactivatedEquipInfo,
            GetTokenHeader(token));
        return response.data;
    });
};

export async function SearchEquip(token, currType, equipName, startDate, endDate) {
    return await HandlerWrapper('SearchEquip', async () => {
        const response = await axios.get(EquipsServiceAddress + EquipsController +
            '/SearchEquip?currType=' + currType+
            '&equipName=' + equipName+
            '&startDate=' + startDate+
            '&endDate=' + endDate,
            GetTokenHeader(token));
        return response.data;
    });
};

export async function GetPermanentData(token, currType, equipName) {
    return await HandlerWrapper('GetPermanentData', async () => {
        const response = await axios.get(EquipsServiceAddress + EquipsController +
            '/GetPermanentData?currType=' + currType+
            '&equipName=' + equipName,
            GetTokenHeader(token));
        return response.data;
    });
};

export async function RunTeamViewer(token, activatedEquipInfo) {
    return await HandlerWrapper('RunTeamViewer', async () => {
        const response = await axios.post(EquipsServiceAddress + EquipsController +
            '/RunTeamViewer?activatedEquipInfo=' + activatedEquipInfo,
            null, GetTokenHeader(token));
        return response.data;
    });
};

export async function RunTaskManager(token, activatedEquipInfo) {
    return await HandlerWrapper('RunTaskManager', async () => {
        const response = await axios.post(EquipsServiceAddress + EquipsController +
            '/RunTaskManager?activatedEquipInfo=' + activatedEquipInfo,
            null, GetTokenHeader(token));
        return response.data;
    });
};

export async function SendAtlasLogs(token, activatedEquipInfo) {
    return await HandlerWrapper('SendAtlasLogs', async () => {
        const response = await axios.post(EquipsServiceAddress + EquipsController +
            '/SendAtlasLogs?activatedEquipInfo=' + activatedEquipInfo,
            null, GetTokenHeader(token));
        return response.data;
    });
};

export async function XilibLogsOn(token, activatedEquipInfo, detailedXilib, verboseXilib) {
    return await HandlerWrapper('XilibLogsOn', async () => {
        const response = await axios.post(EquipsServiceAddress + EquipsController +
            '/XilibLogsOn?activatedEquipInfo=' + activatedEquipInfo+
            '&detailedXilib=' + detailedXilib +
            '&verboseXilib=' + verboseXilib,
            null, GetTokenHeader(token));
        return response.data;
    });
};

export async function GetAllEquips(token, withDisabled = false) {
    return await HandlerWrapper('GetAllEquips', async () => {
        const response = await axios.get(
            EquipsServiceAddress + EquipsController + '/GetAllEquips?withDisabled='+withDisabled,
            GetTokenHeader(token));
        return response.data;
    });
};

export async function GetAllTables(token, equipName) {
    return await HandlerWrapper('GetAllDBTableNames', async () => {
        const response = await axios.get(EquipsServiceAddress + EquipsController + 
            '/GetAllDBTableNames?equipName=' + equipName,
            GetTokenHeader(token));
        return response.data;
    });
};

export async function GetTableContent(token, equipName, tableType, tableName) {
    return await HandlerWrapper('GetTableContent', async () => {
        const response = await axios.get(EquipsServiceAddress + EquipsController +
            '/GetTableContent?equipName=' + equipName+
            '&tableType=' + tableType+
            '&tableName=' + tableName,
            GetTokenHeader(token));
        return response.data;
    });
};

export async function UpdateDBInfo(token, activatedEquipInfo) {
    return await HandlerWrapper('UpdateDBInfo', async () => {
        const response = await axios.post(EquipsServiceAddress + EquipsController +
            '/UpdateDBInfo?activatedEquipInfo=' + activatedEquipInfo,
            null, GetTokenHeader(token));
        return response.data;
    });
};

export async function DisableEquipInfo(token, equipName, disabled) {
    return await HandlerWrapper('DisableEquipInfo', async () => {
        const response = await axios.post(EquipsServiceAddress + EquipsController +
            '/DisableEquipInfo?equipName=' + equipName+
            '&disabled=' + disabled,
            null, GetTokenHeader(token));
        return response.data;
    });
};
