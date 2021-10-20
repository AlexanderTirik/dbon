import React, { FC, useReducer, useState } from "react";
import { Button, Form, Modal } from "react-bootstrap";
import { IDictionary } from "../../common/models/IDictionary";
import { TableData } from "../../components/TableData";
import { join } from "../../services/dataService";
import { getTableCols } from "../../services/tableService";
import { joinReducer, initialJoinState, JoinAction } from "./utils";

interface IProps {
  onClose: () => void;
  db: string;
  tables: string[];
}

export const JoinTablesModal: FC<IProps> = ({ onClose, db, tables }) => {
  const [joinState, dispath] = useReducer(joinReducer, initialJoinState);
  const [data, setData] = useState<IDictionary[]>([]);

  const setFirstTable = async (v: string) => {
    if (v !== "Choose table") {
      const cols = await getTableCols(db, v);
      const colsArr = Object.keys(cols);
      dispath({ type: JoinAction.SetFirstCols, payload: colsArr });
      dispath({ type: JoinAction.SetFirstTable, payload: v });
    }
  };

  const setFirstCol = (v: string) =>
    dispath({ type: JoinAction.SetFirstCol, payload: v });

  const setSecondTable = async (v: string) => {
    if (v !== "Choose table") {
      const cols = await getTableCols(db, v);
      const colsArr = Object.keys(cols);
      dispath({ type: JoinAction.SetSecondCols, payload: colsArr });
      dispath({ type: JoinAction.SetSecondTable, payload: v });
    }
  };

  const setSecondCol = (v: string) =>
    dispath({ type: JoinAction.SetSecondCol, payload: v });

  const onJoin = async () => {
    const { firstTable, secondTable, firstCol, secondCol} = joinState;
    const joinedData = await join(db, firstTable, firstCol, secondTable, secondCol);
    setData(joinedData)
  };

  return (
    <Modal show={true} size="lg">
      <Modal.Header>
        <Modal.Title>Join Tables</Modal.Title>
      </Modal.Header>

      <Modal.Body>
        <Form>
          1
          <Form.Control
            as="select"
            data-test-id="first-join-table-select"
            onChange={(e) => setFirstTable(e.target.value)}
            value={joinState.firstTable}
          >
            <option>Choose table</option>
            {tables.map((table) => (
              <option value={table}>{table}</option>
            ))}
          </Form.Control>
          {joinState.firstTable ? (
            <Form.Control
              as="select"
              data-test-id="first-join-col-select"
              onChange={(e) => setFirstCol(e.target.value)}
              value={joinState.firstCol}
            >
              <option>Choose col</option>
              {joinState.firstCols.map((col: string) => (
                <option value={col}>{col}</option>
              ))}
            </Form.Control>
          ) : null}
          2
          <Form.Control
            as="select"
            data-test-id="second-join-table-select"
            onChange={(e) => setSecondTable(e.target.value)}
            value={joinState.secondTable}
          >
            <option>Choose table</option>
            {tables.map((table) => (
              <option value={table}>{table}</option>
            ))}
          </Form.Control>
          {joinState.secondTable ? (
            <Form.Control
              as="select"
              data-test-id="second-join-col-select"
              onChange={(e) => setSecondCol(e.target.value)}
              value={joinState.secondCol}
            >
              <option>Choose col</option>
              {joinState.secondCols.map((col: string) => (
                <option value={col}>{col}</option>
              ))}
            </Form.Control>
          ) : null}
        </Form>
        {data.length ? (
          <TableData testId="join-data" cols={Object.keys(data[0])} data={data} />
        ): null}
      </Modal.Body>

      <Modal.Footer>
        <Button onClick={onClose} variant="secondary">
          Close
        </Button>
        <Button onClick={onJoin} variant="primary">
          Join
        </Button>
      </Modal.Footer>
    </Modal>
  );
};
