import api from "../common/helpers/apiHelper"
import { IDictionary } from "../common/models/IDictionary"

export const getData = async (db: string, table: string): Promise<{ data: IDictionary[], colNames: string[] }> => api.get(`/${db}/table/${table}/data`)

export const createData = async (db: string, table: string, data: IDictionary) => api.post(`/${db}/table/${table}/data`, data)

export const deleteData = async (db: string, table: string, id: string) => api.delete(`/${db}/table/${table}/data/${id}`)

export const join = async (db: string, table1: string, on1: string, table2: string, on2: string): Promise<IDictionary[]> => api.get(`/${db}/join/${table1}/on/${on1}/with/${table2}/on/${on2}`);