interface IJoinState {
  firstTable?: string;
  firstCol?: string;
  secondTable?: string;
  secondCol?: string;
  firstCols: string[];
  secondCols: string[];
}

export const initialJoinState = {
  firstCols: [],
  secondCols: [],
};

export enum JoinAction {
  SetFirstTable,
  SetSecondTable,
  SetFirstCols,
  SetFirstCol,
  SetSecondCols,
  SetSecondCol,
}

export const joinReducer = (
  state: IJoinState,
  { type, payload }: { type: JoinAction; payload: any }
) => {
  switch (type) {
    case JoinAction.SetFirstTable:
      return { ...state, firstTable: payload };
    case JoinAction.SetSecondTable:
      return { ...state, secondTable: payload };
    case JoinAction.SetFirstCols:
      return { ...state, firstCols: payload };
    case JoinAction.SetSecondCols:
      return { ...state, secondCols: payload };
    case JoinAction.SetFirstCol:
      return { ...state, firstCol: payload };
    case JoinAction.SetSecondCol:
      return { ...state, secondCol: payload };
    default:
      throw new Error();
  }
};
