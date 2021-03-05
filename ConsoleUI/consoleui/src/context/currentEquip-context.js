import React, { useReducer, createContext } from 'react';

export const CurrentEquipContext = createContext();

const initialState = {
  equipInfo: '',
  detectors: [],
  generator: {},
  collimator:{},
  dosimeter:{},
  stand:{},
  dicom:{},
  hdd:[],
  memory:{},
  cpu:{},
  organAuto:{},
};

function reducer(state, action) {
  switch (action.type) {    
    case 'RESET': {
      return initialState;
    }
    case 'SETEQUIPINFO': {
      return {
        ...state,
        equipInfo: action.payload
      };
    }
    case 'SETHDDS': {
      return {
        ...state,
        hdd: action.payload
      };
    }
    case 'SETMEMORY': {
      return {
        ...state,
        memory: action.payload
      };
    }
    case 'SETORGANAUTO': {
      return {
        ...state,
        organAuto: action.payload
      };
    }
    default:
      throw new Error();
  }
}

export const CurrentEquipContextProvider = props => {
  const [state, dispatch] = useReducer(reducer, initialState);
  const { children } = props;

  return (
    <CurrentEquipContext.Provider value={[state, dispatch]}>
      {children}
    </CurrentEquipContext.Provider>
  );
};