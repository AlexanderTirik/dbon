import React, { FC } from "react";
import { Button, Table } from "react-bootstrap";
import { IDictionary } from "../../common/models/IDictionary";
import dsStyles from "../../scenes/Dashboard/styles.module.sass";
import styles from "./styles.module.sass";

interface IProps {
  data?: IDictionary[];
  cols: string[];
  testId?: string;
  onDelete?: (id: string) => void;
}

export const TableData: FC<IProps> = ({ data = [], cols, testId, onDelete }) => (
  <Table striped bordered hover data-test-id={testId}>
    <thead>
      <tr>
        <th className={styles.tableCell}>#</th>
        {cols.map((col) => (
          <th className={styles.tableCell}>{col}</th>
        ))}
        {onDelete ? <th>Delete</th> : null}
      </tr>
    </thead>
    <tbody>
      {data.map((el, i) => {
        return (
          <tr>
            <td className={styles.tableCell}>{i + 1}</td>
            {cols.map((col) => (
              <td className={styles.tableCell}>{el[col]}</td>
            ))}
            {onDelete ? (
              <td>
                <Button variant="danger" onClick={() => onDelete(el["id"])} className={dsStyles.deleteButton}>
                  Delete
                </Button>
              </td>
            ) : null}
          </tr>
        );
      })}
    </tbody>
  </Table>
);
