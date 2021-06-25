import axios from 'axios';
import { EquipsServiceAddress } from '../model/constants'
import {HandlerWrapper, GetJsonHeader, GetTokenHeader} from './commonWorker'

import {sessionUid} from '../utilities/utils'

const AdminController = '/equips';

export async function GetAllUsers(token) {
    return await HandlerWrapper('GetAllUsers', async () => {
        const path = EquipsServiceAddress + AdminController + '/GetAllUsers';
        const header = GetTokenHeader(token);        
        const response = await axios.get(path, header);

        return response.data;
    });
};

export async function UpdateUser(user, token) {
    return await HandlerWrapper('UpdateUser', async () => {
        const response = await axios.post(EquipsServiceAddress + AdminController + '/UpdateUser',
            user, //json,
            {
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded",
                    "Accept": "application/json",
                    "Authorization": "Bearer " + token
                }
            });

        return response.data;
    });
};

export async function Login(login) {
    return await HandlerWrapper('Login', async () => {
        const response = await axios.post(EquipsServiceAddress + AdminController + '/Login',
            login, 
            {
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded",
                    "Accept": "application/json",
                }
            });

        return response.data;
    });
};
