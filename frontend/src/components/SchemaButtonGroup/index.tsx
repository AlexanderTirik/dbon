import React, { FC } from "react";
import { Button, Form, FormControl, Modal } from "react-bootstrap";
import styles from "./styles.module.sass";

interface IProps {
  data: string[];
  selected?: string;
  onClick: (el: string) => void;
}

export const SchemaButtonGroup: FC<IProps> = ({ data, selected, onClick }) => (
  <Form>
    <FormControl
      as="select"
      onChange={(e) => onClick(e.target.value as string)}
      value={selected}
    >
      {data.map((el) => (
        <option value={el}>{el}</option>
      ))}
    </FormControl>
  </Form>
);
