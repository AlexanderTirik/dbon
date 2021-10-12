import React, { FC, useEffect } from "react";
import { Button, Form, FormControl } from "react-bootstrap";

interface IProps {
  data: string[];
  selected?: string;
  onSelect: (el: string) => void;
}

export const SchemaSelector: FC<IProps> = ({ data, selected, onSelect }) => (
  <div>
    <Form>
      <FormControl
        as="select"
        onChange={(e) => onSelect(e.target.value as string)}
        value={selected}
      >
        {data.map((el) => (
          <option value={el}>{el}</option>
        ))}
      </FormControl>
    </Form>
  </div>
);
