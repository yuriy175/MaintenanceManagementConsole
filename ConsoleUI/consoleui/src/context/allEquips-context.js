import React, { useReducer, createContext } from 'react';

export const AllEquipsContext = createContext();

const initialState = {
  allEquips: null,
  connectedEquips: null,
  selectedEquips: [],
};

function reducer(state, action) {
  switch (action.type) {
    case 'SETALLEQUIPS': {
      return {
        ...state,
        allEquips: action.payload?.filter(p => p)
      };
    }    
    case 'SETCONNECTEDEQUIPS': {
      return {
        ...state,
        connectedEquips: action.payload?.filter(p => p)
      };
    }    
    case 'ADDSELECTEDEQUIPS': {
      return {
        ...state,
        selectedEquips: [...state.selectedEquips, action.payload]
      };
    }    
    case 'ADDEQUIP': {
      return {
        ...state,
        equipInfo: action.payload
      };
    }
    case 'CONNECTIONCHANGED': {
      let connectedEquips = state.connectedEquips ?? [];
      const equipName = action.payload.State.Name;
      if(action.payload.State.Connected && !connectedEquips.includes(equipName)){
        connectedEquips = [...connectedEquips, equipName]
      }
      else if(!action.payload.State.Connected && connectedEquips.includes(equipName)){
        connectedEquips = connectedEquips.filter(p => p != equipName)
      }
      
      return {
        ...state,
        connectedEquips: connectedEquips
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