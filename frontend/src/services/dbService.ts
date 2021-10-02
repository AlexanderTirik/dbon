import api from "../common/helpers/apiHelper"

export const fetchDbs = (): Promise<string[]> => api.get('/')

export const createDb = async (db: string) => api.post('/', { db })

export const deleteDb = async (db: string) => api.delete(`/${db}`)