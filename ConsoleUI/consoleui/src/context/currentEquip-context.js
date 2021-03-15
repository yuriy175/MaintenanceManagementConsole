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
  software:{},
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
    case 'SETCPU': {
      return {
        ...state,
        cpu: action.payload
      };
    }
    case 'SETORGANAUTO': {
      return {
        ...state,
        organAuto: action.payload
      };
    }
    case 'SETGENERATOR': {
      return {
        ...state,
        generator: {Id: action.payload.Id, State: {...state.generator.State, ...action.payload.State}}
      };
    }
    case 'SETDETECTOR': {
      return {
        ...state,
        detectors: [action.payload]
      };
    }
    case 'SETSTAND': {
      return {
        ...state,
        stand: {Id: action.payload.Id, State: {...state.stand.State, ...action.payload.State}}
      };
    }
    case 'SETDOSIMETER': {
      return {
        ...state,
        dosimeter: {Id: action.payload.Id, State: {...state.dosimeter.State, ...action.payload.State}}
      };
    }
    case 'SETCOLLIMATOR': {
      return {
        ...state,
        collimator: action.payload
      };
    }
    case 'SETSOFTWARE': {
      return {
        ...state,
        software: action.payload
      };
    }
    case 'SETDICOM': {
      return {
        ...state,
        dicom: action.payload
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