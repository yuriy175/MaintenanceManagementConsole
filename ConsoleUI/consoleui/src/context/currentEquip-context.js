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
  system:{},
  organAuto:{},
  software:{},
  remoteaccess:{},
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
    case 'SETSYSTEM': {
      return {
        ...state,
        system: action.payload
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
    case 'SETSOFTWAREMSG': {
      return {
        ...state,
        software: {...state.software, ...action.payload}
      };
    }
    case 'SETREMOTEACCESS': {
      return {
        ...state,
        remoteaccess: {...state.remoteaccess, ...action.payload}
      };
    }
    case 'SETDICOM': {
      let newWL = action.payload.WorkList;
      if(newWL){
        const oldWL = state.dicom.WorkList?.filter(e => !newWL.map(w => w.Name).includes(e.Name));
        if(oldWL){
          newWL = [...newWL, ...oldWL];
        }
      }
      else{
        newWL = state.dicom.WorkList;
      }

      let newPacs = action.payload.PACS;
      if(newPacs){
        const oldPacs = state.dicom.PACS?.filter(e => !newPacs.map(w => w.Name).includes(e.Name));
        if(oldPacs){
          newPacs = [...newPacs, ...oldPacs];
        }
      }
      else{
        newPacs = state.dicom.PACS;
      }

      const dicom = {PACS: newPacs, WorkList: newWL};
      return {
        ...state,
        dicom: dicom
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