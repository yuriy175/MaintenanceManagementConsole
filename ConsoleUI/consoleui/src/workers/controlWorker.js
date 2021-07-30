import axios from 'axios';
import { EquipsServiceAddress } from '../model/constants'
import {HandlerWrapper, GetJsonHeader, GetTokenHeader} from './commonWorker'

import {sessionUid} from '../utilities/utils'

const ControlController = '/equips';

export async function GetServerState(token) {
    return await HandlerWrapper('GetServerState', async () => {
        const path = EquipsServiceAddress + ControlController + '/GetServerState';
        const header = GetTokenHeader(token);        
        const response = await axios.get(path, header);

        return response.data;
    });
};