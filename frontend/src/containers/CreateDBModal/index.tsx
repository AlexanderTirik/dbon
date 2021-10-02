import React, { FC, useState } from "react";
import {
  Button,
  Form,
  FormControl,
  Modal,
} from "react-bootstrap";

interface IProps {
  onCreate: (db: string) => void;
  onClose: () => void;
}

export const CreateDBModal: FC<IProps> = ({ onCreate, onClose }) => {
  const [db, setDb] = useState("");
  return (
    <Modal show={true}>
      <Modal.Header>
        <Modal.Title>Create database</Modal.Title>
      </Modal.Header>

      <Modal.Body>
        <Form>
          <Form.Label>DB Name</Form.Label>
          <FormControl
            onChange={(e) => setDb(e.target.value)}
            placeholder="Database"
          />
        </Form>
      </Modal.Body>

      <Modal.Footer>
        <Button onClick={onClose} variant="secondary">
          Close
        </Button>
        <Button onClick={() => onCreate(db)} variant="primary">
          Save changes
        </Button>
      </Modal.Footer>
    </Modal>
  );
};
