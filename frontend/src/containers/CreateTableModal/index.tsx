import React, { FC, useState } from "react";
import { Button, Form, FormControl, Modal } from "react-bootstrap";
import { Types } from "../../common/enums/Types";
import { IDictionary } from "../../common/models/IDictionary";
import { TableCols } from "../../components/TableCols";
import styles from "./styles.module.sass";

interface IProps {
  onCreate: (table: string, cols: IDictionary) => void;
  onClose: () => void;
}

export const CreateTableModal: FC<IProps> = ({ onCreate, onClose }) => {
  const [cols, setCols] = useState({});
  const [tableName, setTableName] = useState("");
  const [isAdding, setIsAdding] = useState(false);
  const [colName, setColName] = useState("");
  const [colType, setColType] = useState(Types.Integer);

  const onCancel = () => {
    setIsAdding(false);
    setColName("");
    setColType(Types.Integer);
  };

  const onAdd = () => {
    setCols({ ...cols, [colName]: colType });
    onCancel();
  };
  return (
    <Modal show={true}>
      <Modal.Header>
        <Modal.Title>Create table</Modal.Title>
      </Modal.Header>

      <Modal.Body>        
        <Form className={styles.formCreateTable}>
          <FormControl
            onChange={(e) => setTableName(e.target.value)}
            placeholder="Table name"
          />
          <TableCols cols={cols} />
          {isAdding ? (
            <div>
              <FormControl
                onChange={(e) => setColName(e.target.value)}
                placeholder="Column name"
                data-test-id="col-name"
              />
              <Form.Control
                as="select"
                onChange={(e) => setColType(e.target.value as Types)}
                value={colType}
              >
                <option value={Types.Integer}>{Types.Integer}</option>
                <option value={Types.Real}>{Types.Real}</option>
                <option value={Types.String}>{Types.String}</option>
                <option value={Types.Char}>{Types.Char}</option>
                <option value={Types.Color}>{Types.Color}</option>
                <option value={Types.ColorInvl}>{Types.ColorInvl}</option>
              </Form.Control>
              <Button onClick={onCancel}>cancel</Button>{" "}
              <Button onClick={onAdd}>add</Button>
            </div>
          ) : null}
        </Form>
        {!isAdding ? (
          <Button onClick={() => setIsAdding(true)}>+</Button>
        ) : null}
      </Modal.Body>

      <Modal.Footer>
        <Button onClick={onClose} variant="secondary">
          Close
        </Button>
        <Button onClick={() => onCreate(tableName, cols)} variant="primary">
          Save changes
        </Button>
      </Modal.Footer>
    </Modal>
  );
};
