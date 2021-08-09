import React, { useReducer, createContext } from 'react';
import { SummaryTabIndex, MainTabPanelIndex } from '../model/constants';

export const ControlStateContext = createContext();

const initialState = {
  serverState: {},
  diagnostic: {}
};

function reducer(state, action) {
  switch (action.type) {    
    case 'RESET': {
      return initialState;
    }
    case 'SETSRVSTATE': {
      return {
        ...state,
        serverState: action.payload
      };
    }
    case 'SETDIAGNOSTIC': {
      return {
        ...state,
        diagnostic: action.payload
      };
    }

    default:
      throw new Error();
  }
}

export const ControlStateContextProvider = props => {
  const [state, dispatch] = useReducer(reducer, initialState);
  const { children } = props;

  return (
    <ControlStateContext.Provider value={[state, dispatch]}>
      {children}
    </ControlStateContext.Provider>
  );
};