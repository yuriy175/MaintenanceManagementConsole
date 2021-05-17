import React, { useReducer, createContext } from 'react';
import { SummaryTabIndex, MainTabPanelIndex } from '../model/constants';

export const AppContext = createContext();

const initialState = {
  currentTab: {tab: SummaryTabIndex, panel: MainTabPanelIndex},
};

function reducer(state, action) {
  switch (action.type) {    
    case 'RESET': {
      return initialState;
    }
    case 'SETTAB': {
      return {
        ...state,
        currentTab: action.payload
      };
    }

    default:
      throw new Error();
  }
}

export const AppContextProvider = props => {
  const [state, dispatch] = useReducer(reducer, initialState);
  const { children } = props;

  return (
    <AppContext.Provider value={[state, dispatch]}>
      {children}
    </AppContext.Provider>
  );
};