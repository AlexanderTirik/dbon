import React, { useEffect, useReducer } from "react";
import { Button } from "react-bootstrap";
import { CreateDBModal } from "../../containers/CreateDBModal";
import { SchemaButtonGroup } from "../../components/SchemaButtonGroup";
import { createDb, deleteDb, fetchDbs } from "../../services/dbService";
import {
  createTable,
  deleteTable,
  fetchTables,
  getTableCols,
} from "../../services/tableService";
import {
  modalReducer,
  initialModalState,
  initialSchemaState,
  schemaReducer,
  SchemaAction,
  ModalAction,
} from "./utils";
import { CreateTableModal } from "../../containers/CreateTableModal";
import { IDictionary } from "../../common/models/IDictionary";
import styles from "./styles.module.sass";
import { TableCols } from "../../components/TableCols";
import { createData, deleteData, getData } from "../../services/dataService";
import { TableData } from "../../components/TableData";
import { AddDataModal } from "../../containers/AddDataModal";
import { JoinTablesModal } from "../../containers/JoinTablesModal";

const Dashboard = () => {
  const [schemaState, dispatchSchema] = useReducer(
    schemaReducer,
    initialSchemaState
  );
  const [modalState, dispatchModal] = useReducer(
    modalReducer,
    initialModalState
  );

  const setDbs = (dbs: string[]) =>
    dispatchSchema({ type: SchemaAction.SetDbs, payload: dbs });

  const getDbs = async () => {
    const fetchedDbs = await fetchDbs();
    setDbs(fetchedDbs);
  };

  useEffect(() => {
    getDbs();
  }, []);

  const setTables = (tables: string[]) =>
    dispatchSchema({ type: SchemaAction.SetTables, payload: tables });

  const setSelectedDb = (db?: string) =>
    dispatchSchema({ type: SchemaAction.SetSelectedDB, payload: db });

  const setSelectedTable = (table?: string) =>
    dispatchSchema({ type: SchemaAction.SetSelectedTable, payload: table });

  const showCreateDBModal = (show: boolean) =>
    dispatchModal({ type: ModalAction.CreateDB, payload: show });

  const showCreateTableModal = (show: boolean) =>
    dispatchModal({ type: ModalAction.CreateTable, payload: show });

  const showAddDataModal = (show: boolean) =>
    dispatchModal({ type: ModalAction.AddData, payload: show });

  const showJoinTablesModal = (show: boolean) =>
    dispatchModal({ type: ModalAction.JoinTables, payload: show });

  const setTableCols = (cols: IDictionary) =>
    dispatchSchema({ type: SchemaAction.SetTableCols, payload: cols });

  const setData = (data: IDictionary[]) =>
    dispatchSchema({ type: SchemaAction.SetData, payload: data });

  const getTables = async (db: string) => {
    const fetchedTables = await fetchTables(db);
    setTables(fetchedTables);
  };

  const onDbClick = (db: string) => {
    setSelectedDb(db);
    getTables(db);
  };

  const onCreateDb = async (db: string) => {
    await createDb(db);
    getDbs();
    showCreateDBModal(false);
  };

  const onCreateTable = async (table: string, cols: IDictionary) => {
    const db = schemaState.selectedDb;
    try {
      await createTable(db, table, cols);
    } finally {
      showCreateTableModal(false);
      getTables(db);
    }
  };

  const onDeleteDB = async () => {
    await deleteDb(schemaState.selectedDb);
    setSelectedDb();
    getDbs();
  };

  const onDeleteTable = async () => {
    const { selectedTable, selectedDb } = schemaState;
    await deleteTable(selectedDb, selectedTable);
    setSelectedTable();
    await getTables(selectedDb);
  };

  const onTableClick = async (table: string) => {
    const db = schemaState.selectedDb;
    setSelectedTable(table);
    const cols = await getTableCols(db, table);
    setTableCols(cols);
    const { data } = await getData(db, table);
    setData(data || []);
  };

  const onDeleteData = async (id: string) => {
    const { selectedDb, selectedTable } = schemaState;
    await deleteData(selectedDb, selectedTable, id);
    const { data } = await getData(selectedDb, selectedTable);
    setData(data || []);
  };

  const onAddData = async (createdData: IDictionary) => {
    const { selectedDb, selectedTable } = schemaState;
    await createData(selectedDb, selectedTable, createdData);
    const { data } = await getData(selectedDb, selectedTable);
    setData(data || []);
  };

  return (
    <div className={styles.dashboard}>
      <Button onClick={() => showCreateDBModal(true)}>Create database</Button>
      {modalState.showCreateDB ? (
        <CreateDBModal
          onClose={() => showCreateDBModal(false)}
          onCreate={onCreateDb}
        />
      ) : null}
      <SchemaButtonGroup
        selected={schemaState.selectedDb}
        data={schemaState.dbs}
        onClick={onDbClick}
      />
      {schemaState.selectedDb ? (
        <>
          <Button variant="danger" onClick={onDeleteDB}>
            Delete DB
          </Button>
          <Button onClick={() => showCreateTableModal(true)}>
            Create Table
          </Button>
        </>
      ) : null}

      {modalState.showCreateTable ? (
        <CreateTableModal
          onClose={() => showCreateTableModal(false)}
          onCreate={onCreateTable}
        />
      ) : null}
      {schemaState.selectedDb ? (
        <>
          <SchemaButtonGroup
            selected={schemaState.selectedTable}
            data={schemaState.tables}
            onClick={onTableClick}
          />
          {schemaState.selectedTable ? (
            <>
              <Button variant="danger" onClick={onDeleteTable}>
                Delete Table
              </Button>
              <TableCols cols={schemaState.tableCols} />
              <Button onClick={() => showAddDataModal(true)}>Add data</Button>
              {modalState.showAddData ? (
                <AddDataModal
                  onClose={() => showAddDataModal(false)}
                  onCreate={onAddData}
                  cols={schemaState.tableCols || {}}
                />
              ) : null}
              <TableData
                cols={Object.keys(schemaState.tableCols || {})}
                data={schemaState.data}
                onDelete={onDeleteData}
              />
              <Button onClick={() => showJoinTablesModal(true)}>
                Join tables
              </Button>
              {modalState.joinTables ? (
                <JoinTablesModal
                  db={schemaState.selectedDb}
                  tables={schemaState.tables}
                  onClose={() => showJoinTablesModal(false)}
                />
              ) : null}
            </>
          ) : null}
        </>
      ) : null}
    </div>
  );
};

export default Dashboard;
