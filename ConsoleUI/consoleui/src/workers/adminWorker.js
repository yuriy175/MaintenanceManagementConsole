import axios from 'axios';
import { EquipsServiceAddress } from '../model/constants'
import {HandlerWrapper, GetJsonHeader} from './commonWorker'

import {sessionUid} from '../utilities/utils'

const AdminController = '/equips';

export async function GetAllUsers() {
    return await HandlerWrapper('GetAllUsers', async () => {
        const path = EquipsServiceAddress + AdminController + '/GetAllUsers';
        console.log(path);
        const response = await axios.get(path);
        return response.data;
    });
};

export async function UpdateUser(user) {
    return await HandlerWrapper('UpdateUser', async () => {
        const response = await axios.post(EquipsServiceAddress + AdminController +
            '/UpdateUser',
            JSON.stringify(user),
            GetJsonHeader());
        return response.data;
    });
};
