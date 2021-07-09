import React, { useReducer, createContext } from 'react';

export const UsersContext = createContext();

const initialState = {
  currentUser: {},
  token: null,
  users: null,
};

function reducer(state, action) {
  switch (action.type) {    
    case 'RESET': {
      return initialState;
    }
    case 'SETUSER': {
      return {
        ...state,
        token: action.payload.Token,
        currentUser: {...action.payload.Claims, Surname: action.payload.Surname}
      };
    }
    case 'SETUSERS': {
      return {
        ...state,
        users: action.payload
      };
    }  

    default:
      throw new Error();
  }
}

export const UsersContextProvider = props => {
  const [state, dispatch] = useReducer(reducer, initialState);
  const { children } = props;

  return (
    <UsersContext.Provider value={[state, dispatch]}>
      {children}
    </UsersContext.Provider>
  );
};