import axios from 'axios';
import {Absence} from "../model/absence";
import {config} from "../config";

const api = axios.create({
    baseURL: config.addr + '/api/absences'
});


export async function getAbsences(): Promise<Absence> {
    const response = await api.get<Absence>('', {
        headers: {
            'Accept': 'application/json'
        }
    });

    return response.data
}
export async function getAbsence(id: number): Promise<Absence> {
    const response = await api.get<Absence>('/'+id, {
        headers: {
            'Accept': 'application/json'
        }
    });

    return response.data
}

export async function deleteAbsence(id: number): Promise<Absence> {
    const response = await api.delete<Absence>('/'+id, {
        headers: {
            'Accept': 'application/json'
        }
    });

    return response.data
}

export async function createAbsence(type: string, userId: string, name: string, status: string, fromDate: Date, toDate: Date) {
    const data = new FormData();
    data.set('type', type);
    data.set("userId", userId);
    data.set('name', name);
    data.set('status', status);
    data.set('fromDate', fromDate.toString());
    data.set('toDate', toDate.toString());
    data.set('creation', Date.now().toString());

    await api.post('', data, {
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    })
}

