import React, { FC, useState } from "react";
import { Button, Form, FormControl, Modal } from "react-bootstrap";
import { IDictionary } from "../../common/models/IDictionary";
import styles from "./styles.module.sass";

interface IProps {
  cols: IDictionary;
  onCreate: (data: IDictionary) => void;
  onClose: () => void;
}

export const AddDataModal: FC<IProps> = ({ cols, onCreate, onClose }) => {
  const [data, setData] = useState({});

  const onChangeData = (k: string, v: string) => setData({ ...data, [k]: v });

  return (
    <Modal show={true}>
      <Modal.Header>
        <Modal.Title>Add data</Modal.Title>
      </Modal.Header>

      <Modal.Body>
        <Form className={styles.formAddData}>
          <>
            {Object.entries(cols).map(([k, v]) => (
              <>
                <Form.Label htmlFor={`${k}-input`}>{k}</Form.Label>
                <FormControl
                  onChange={(e) => onChangeData(k, e.target.value)}
                  placeholder={v}
                  name={`${k}-input`}
                />
              </>
            ))}
          </>
        </Form>
      </Modal.Body>

      <Modal.Footer>
        <Button onClick={onClose} variant="secondary">
          Close
        </Button>
        <Button onClick={() => onCreate(data)} variant="primary">
          Save changes
        </Button>
      </Modal.Footer>
    </Modal>
  );
};
