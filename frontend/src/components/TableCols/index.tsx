import React, { FC } from "react";
import { Table } from "react-bootstrap";
import { IDictionary } from "../../common/models/IDictionary";

interface IProps {
  cols?: IDictionary;
}

export const TableCols: FC<IProps> = ({ cols = {} }) => (
  <Table striped bordered hover data-test-id="table-cols">
    <thead>
      <tr>
        <th>#</th>
        <th>Column Name</th>
        <th>Column Type</th>
      </tr>
    </thead>
    <tbody>
      {Object.entries(cols).map(([v, k], i) => (
        <tr>
          <td>{i + 1}</td>
          <td>{v}</td>
          <td>{k}</td>
        </tr>
      ))}
    </tbody>
  </Table>
);
