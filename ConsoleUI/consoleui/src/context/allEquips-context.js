import React, { useReducer, createContext } from 'react';

export const AllEquipsContext = createContext();

const initialState = {
  equips: [],
};

function reducer(state, action) {
  switch (action.type) {
    case 'ADDEQUIP': {
      return {
        ...state,
        equipInfo: action.payload
      };
    }
    default:
      throw new Error();
  }
}

export const AllEquipsContextProvider = props => {
  const [state, dispatch] = useReducer(reducer, initialState);
  const { children } = props;

  return (
    <AllEquipsContext.Provider value={[state, dispatch]}>
      {children}
    </AllEquipsContext.Provider>
  );
};