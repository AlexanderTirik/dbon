import { IDictionary } from "../../common/models/IDictionary";

interface IModalState {
  showCreateDB: boolean;
  showCreateTable: boolean;
  showAddData: boolean;
  joinTables: boolean;
}

export const initialModalState = {
  showCreateDB: false,
  showCreateTable: false,
  showAddData: false,
  joinTables: false,
};

export enum ModalAction {
  CreateDB,
  CreateTable,
  AddData,
  JoinTables,
}

export const modalReducer = (
  state: IModalState,
  { type, payload }: { type: ModalAction; payload: any }
) => {
  switch (type) {
    case ModalAction.CreateDB:
      return { ...state, showCreateDB: payload };
    case ModalAction.CreateTable:
      return { ...state, showCreateTable: payload };
    case ModalAction.AddData:
      return { ...state, showAddData: payload };
    case ModalAction.JoinTables:
      return { ...state, joinTables: payload };
    default:
      throw new Error();
  }
};

interface ISchemaState {
  dbs: string[];
  tables: string[];
  selectedDb?: string;
  selectedTable?: string;
  tableCols?: IDictionary;
  data: IDictionary[];
}

export const initialSchemaState = {
  dbs: [],
  tables: [],
  data: [],
};

export enum SchemaAction {
  SetDbs,
  SetTables,
  SetSelectedTable,
  SetSelectedDB,
  SetTableCols,
  SetData,
}

export const schemaReducer = (
  state: ISchemaState,
  { type, payload }: { type: SchemaAction; payload: any }
) => {
  switch (type) {
    case SchemaAction.SetDbs:
      return { ...state, dbs: payload };
    case SchemaAction.SetTables:
      return { ...state, tables: payload };
    case SchemaAction.SetSelectedDB:
      return { ...state, selectedDb: payload };
    case SchemaAction.SetSelectedTable:
      return { ...state, selectedTable: payload };
    case SchemaAction.SetTableCols:
      return { ...state, tableCols: payload };
    case SchemaAction.SetData:
      return { ...state, data: payload };
    default:
      throw new Error();
  }
};
