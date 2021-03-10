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

export async function SearchEquip(equipType, startDate, endDate) {
    return await HandlerWrapper('SearchEquip', async () => {
        const response = await axios.get(EquipsServiceAddress + EquipsController +
            '/SearchEquip?equipType=' + equipType+
            '&startDate=' + startDate+
            '&endDate=' + endDate);
        return response.data;
    });
};