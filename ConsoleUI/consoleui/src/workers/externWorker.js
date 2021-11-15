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

export async function GetRegionById(osm_id) {
    return await HandlerWrapper('GetRegionById', async () => {
        //  https://nominatim.openstreetmap.org/lookup?osm_ids=W88795833&format=json
        const path = `https://nominatim.openstreetmap.org/lookup?osm_ids=W${osm_id}&format=jsonv2`;
        const response = await axios.get(path);

        return response.data;
    });
};
