import React, { useReducer, createContext } from 'react';

export const CommunicationContext = createContext();

const initialState = {
  logs: [],
  notes: [],
};

function reducer(state, action) {
  switch (action.type) {
    case 'SETLOGS': {
      return {
        ...state,
        logs: action.payload
      };
    } 
    case 'SETCHATS': {
      return {
        ...state,
        notes: action.payload
      };
    } 
    case 'ADDNOTE': {
      return {
        ...state,
        notes: [action.payload, ...state.notes]
      };
    }
    
    default:
      throw new Error();
  }
}

export const CommunicationContextProvider = props => {
  const [state, dispatch] = useReducer(reducer, initialState);
  const { children } = props;

  return (
    <CommunicationContext.Provider value={[state, dispatch]}>
      {children}
    </CommunicationContext.Provider>
  );
};