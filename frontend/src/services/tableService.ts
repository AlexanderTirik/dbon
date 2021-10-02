import api from "../common/helpers/apiHelper"
import { IDictionary } from "../common/models/IDictionary"

export const fetchTables = async (db: string): Promise<string[]> => api.get(`/${db}`)

export const createTable = async (db: string, table:string, cols: IDictionary) => 
    api.post(`/${db}/table`, { table, colNames: Object.keys(cols), colTypes: cols })

export const deleteTable = async (db: string, table: string) => api.delete(`/${db}/table/${table}`);

export const getTableCols = async (db: string, table: string): Promise<IDictionary> => api.get(`/${db}/table/${table}`)