import React, { useReducer, createContext } from 'react';

export const SystemVolatileContext = createContext();

const initialState = {
  currentVolatile: {},
};

function reducer(state, action) {
  switch (action.type) {    
    case 'RESET': {
      return initialState;
    }
    case 'SETVOLATILE': {
      return {
        ...state,
        currentVolatile: {...state.currentVolatile, ...action.payload}
      };
    }

    default:
      throw new Error();
  }
}

export const SystemVolatileContextProvider = props => {
  const [state, dispatch] = useReducer(reducer, initialState);
  const { children } = props;

  return (
    <SystemVolatileContext.Provider value={[state, dispatch]}>
      {children}
    </SystemVolatileContext.Provider>
  );
};