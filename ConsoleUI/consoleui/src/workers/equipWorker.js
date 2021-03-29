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

export async function GetAllEquips() {
    return await HandlerWrapper('GetAllEquips', async () => {
        const path = EquipsServiceAddress + EquipsController + '/GetAllEquips';
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

export async function XilibLogsOn(activatedEquipInfo) {
    return await HandlerWrapper('XilibLogsOn', async () => {
        const response = await axios.post(EquipsServiceAddress + EquipsController +
            '/XilibLogsOn?activatedEquipInfo=' + activatedEquipInfo);
        return response.data;
    });
};