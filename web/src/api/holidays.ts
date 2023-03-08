import axios from 'axios';
import {Holiday} from "../model";
import {config} from "../config";
import {Dayjs} from "dayjs";

const api = axios.create({
    baseURL: config.addr + '/api/holidays'
});


export async function getHolidays(): Promise<Holiday> {
    const response = await api.get<Holiday>('', {
        headers: {
            'Accept': 'application/json'
        }
    });

    return response.data
}
export async function getAHoliday(id: number): Promise<Holiday> {
    const response = await api.get<Holiday>('/'+id, {
        headers: {
            'Accept': 'application/json'
        }
    });

    return response.data
}

export async function deleteHolidays(id: number): Promise<Holiday> {
    const response = await api.delete<Holiday>('/'+id, {
        headers: {
            'Accept': 'application/json'
        }
    });

    return response.data
}

export async function createHolidays(id: number, type: string, userId: string, name: string, status: string, fromDate: Dayjs, toDate: Dayjs) {
    const data = new FormData();
    data.set('id', id.toString());
    data.set('type', type);
    data.set("userId", userId);
    data.set('name', name);
    data.set('status', status);
    data.set('fromDate', fromDate.toString());
    data.set('toDate', toDate.toString());
    data.set('creation', Date.now().toString());

    const response = await api.post('', data, {
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    });
    console.log(response.status);

    return response;
}

