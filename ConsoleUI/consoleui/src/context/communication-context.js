import React, { useReducer, createContext } from 'react';

export const CommunicationContext = createContext();

const initialState = {
  logs: [],
  notes: [],
  commonNotes: [],
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
    case 'SETCOMMONCHAT': {
      return {
        ...state,
        commonNotes: action.payload
      };
    } 
    case 'ADDNOTE': {
      return {
        ...state,
        notes: [action.payload, ...state.notes]
      };
    }
    case 'ADDCOMMONNOTE': {
      return {
        ...state,
        commonNotes: [action.payload, ...state.commonNotes]
      };
    }
    case 'CHANGENOTE': {
      const oldNote = state.notes.filter(n => n.ID === action.payload.ID);
      if(oldNote?.length > 0){
        oldNote[0].Message = action.payload.Message;
      }

      return {
        ...state,
        notes: [...state.notes]
      };
    }
    case 'DELETENOTE': {
      return {
        ...state,
        notes: [...state.notes.filter(n => n.ID !== action.payload)]
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