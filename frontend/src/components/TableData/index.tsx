import React, { FC } from "react";
import { Button, Table } from "react-bootstrap";
import { IDictionary } from "../../common/models/IDictionary";

interface IProps {
  data?: IDictionary[];
  cols: string[];
  onDelete?: (id: string) => void;
}

export const TableData: FC<IProps> = ({ data = [], cols, onDelete }) => (
  <Table striped bordered hover>
    <thead>
      <tr>
        <th>#</th>
        {cols.map((col) => (
          <th>{col}</th>
        ))}
        {onDelete ? <th>Delete</th> : null}
      </tr>
    </thead>
    <tbody>
      {data.map((el, i) => {
        return (
          <tr>
            <td>{i + 1}</td>
            {cols.map((col) => (
              <td>{el[col]}</td>
            ))}
            {onDelete ? (
              <td>
                <Button variant="danger" onClick={() => onDelete(el["id"])}>
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
