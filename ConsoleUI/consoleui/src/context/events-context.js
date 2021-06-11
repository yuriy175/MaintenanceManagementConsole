import React, { useReducer, createContext } from 'react';

export const EventsContext = createContext();

const initialState = {
  currentEvents: [],
};

function reducer(state, action) {
  switch (action.type) {
    case 'SETEVENTS': {
      return {
        ...state,
        currentEvents: action.payload
      };
    }   
    case 'ADDEVENT': {
      return {
        ...state,
        currentEvents: [...action.payload, ...state.currentEvents]
      };
    }   
    
    default:
      throw new Error();
  }
}

export const EventsContextProvider = props => {
  const [state, dispatch] = useReducer(reducer, initialState);
  const { children } = props;

  return (
    <EventsContext.Provider value={[state, dispatch]}>
      {children}
    </EventsContext.Provider>
  );
};