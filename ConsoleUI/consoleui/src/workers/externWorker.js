import axios from 'axios';
import {HandlerWrapper} from './commonWorker'

export async function GetEquipCoords(city, street, house) {
    return await HandlerWrapper('GetEquipCoords', async () => {
        const addr = street.split(' ').join('+')+'+'+house;
        const path = `https://nominatim.openstreetmap.org/search.php?street=${addr}&city=${city}&format=jsonv2`;
        const response = await axios.get(path);

        return response.data;
    });
};
