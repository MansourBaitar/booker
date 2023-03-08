import axios from 'axios';
import {config} from "../config";
import {User} from "../model";

const api = axios.create({
  baseURL: config.addr + '/api/users',
  validateStatus: () => true,
});

export async function login(email: string, pwd: string): Promise<any> {
  const data = new FormData();
  data.set('email', email);
  data.set('password', pwd);
  return await api.post<User>('/login', data, {
    headers: {
      'Accept': 'application/json'
    }
  });
}