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
  images:{},
  aecs:{},
  allDBs:{},
  allDBTables:{},
  lastSeen:undefined,
  locationInfo: '',
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
    case 'SETALLDB': {
      return {
        ...state,
        allDBs: action.payload
      };
    }    
    case 'SETALLDBTABLES': {
      return {
        ...state,
        allDBTables: action.payload
      };
    }    
        
    case 'SETORGANAUTO': {
      return {
        ...state,
        organAuto: action.payload
      };
    }

    case 'SETLASTSEEN': {
      return {
        ...state,
        lastSeen: action.payload
      };
    }
    case 'SETGENERATOR': {
      const newState = {
        ...state,
        generator: {Id: action.payload.Id, State: {...state.generator.State, ...action.payload.State}}
      };
      //
      if(newState.generator?.State?.ErrorDescriptions && 
        (!newState.generator?.State?.Error || newState.generator?.State?.Error[0] === 0)){
        newState.generator.State.ErrorDescriptions = [];
      }

      return newState;
    }
    case 'SETDETECTOR': {
      const newDetectorId = action.payload.DetectorId;
      let oldDetector = state.detectors?.filter(d => d.DetectorId === newDetectorId)[0];
      let newDetectors = state.detectors;
      if(oldDetector){
        oldDetector = {...oldDetector, ...action.payload}
        newDetectors = [...state.detectors?.filter(d => d.DetectorId !== newDetectorId), oldDetector]
      }
      else{
        newDetectors = [...state.detectors, action.payload]
      }

      return {
        ...state,
        detectors: newDetectors //[...state.detectors, ...action.payload]
      };
    }
    case 'SETAEC': {
      return {
        ...state,
        aecs: {Id: action.payload.Id, Type: action.payload.Type, State: {...state.aecs.State, ...action.payload.State}}
      };
    }    
    case 'SETSTAND': {
      // return {
      //   ...state,
      //   stand: {Id: action.payload.Id, State: {...state.stand.State, ...action.payload.State}}
      // };
      const newState = {
        ...state,
        stand: {Id: action.payload.Id, State: {...state.stand.State, ...action.payload.State}}
      };
      //
      if(newState.stand?.State?.ErrorDescriptions && 
        (!newState.stand?.State?.Error || newState.stand?.State?.Error[0] === 0)){
        newState.stand.State.ErrorDescriptions = [];
      }

      return newState;
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
    case 'SETIMAGES': {
      return {
        ...state,
        images: action.payload
      };
    }
    // case 'ADDIMAGE': {
    //   return {
    //     ...state,
    //     images: {...state.images, ...action.payload}
    //   };
    // }
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

    case 'SETFULLINFO': {
      return {
        ...state,
        system: action.payload?.SystemInfo[0],
        software: action.payload?.SoftwareInfo[0],
        lastSeen: action.payload?.LastSeen,
        locationInfo: action.payload?.LocationInfo
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