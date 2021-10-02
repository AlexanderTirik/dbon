import React, { FC } from "react";
import { Button, ButtonGroup } from "react-bootstrap";
import styles from "./styles.module.sass";

interface IProps {
  data: string[];
  selected?: string;
  onClick: (el: string) => void;
}

export const SchemaButtonGroup: FC<IProps> = ({ data, selected, onClick }) => (
  <ButtonGroup size="lg" className="mb-2">
    {data.map((el) => (
      <Button
        variant="info"
        className={selected === el ? styles.selected : undefined}
        onClick={() => onClick(el)}
      >
        {el}
      </Button>
    ))}
  </ButtonGroup>
);
