import React, { useReducer, createContext } from 'react';
import { SummaryTabIndex, MainTabPanelIndex } from '../model/constants';

export const EquipLogContext = createContext();

const initialState = {
  currentLog: {},
};

function reducer(state, action) {
  switch (action.type) {    
    case 'RESET': {
      return initialState;
    }
    case 'ADDLOG': {
      return {
        ...state,
        currentLog: action.payload
      };
    }

    default:
      throw new Error();
  }
}

export const EquipLogContextProvider = props => {
  console.log("render EquipLogContextProvider");

  const [state, dispatch] = useReducer(reducer, initialState);
  const { children } = props;

  return (
    <EquipLogContext.Provider value={[state, dispatch]}>
      {children}
    </EquipLogContext.Provider>
  );
};