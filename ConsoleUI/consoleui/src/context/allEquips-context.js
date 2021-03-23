import React, { useReducer, createContext } from 'react';

export const AllEquipsContext = createContext();

const initialState = {
  equips: null,
};

function reducer(state, action) {
  switch (action.type) {
    case 'SETEQUIPS': {
      return {
        ...state,
        equips: action.payload?.filter(p => p)
      };
    }    
    case 'ADDEQUIP': {
      return {
        ...state,
        equipInfo: action.payload
      };
    }
    case 'CONNECTIONCHANGED': {
      let equips = state.equips ?? [];
      const equipName = action.payload.State.Name;
      if(action.payload.State.Connected && !equips.includes(equipName)){
        equips = [...equips, equipName]
      }
      else if(!action.payload.State.Connected && equips.includes(equipName)){
        equips = equips.filter(p => p != equipName)
      }
      
      return {
        ...state,
        equips: equips
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