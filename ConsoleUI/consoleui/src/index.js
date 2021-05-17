import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter } from 'react-router-dom';
import './index.css';
import App from './App';

import { AllEquipsContextProvider } from './context/allEquips-context';
import { CurrentEquipContextProvider } from './context/currentEquip-context';
import { UsersContextProvider} from './context/users-context';
import { AppContextProvider} from './context/app-context';

ReactDOM.render(
  <React.StrictMode>
    <AppContextProvider>
      <UsersContextProvider>
        <AllEquipsContextProvider>
          <CurrentEquipContextProvider>
            <BrowserRouter>
              <App />
            </BrowserRouter>
          </CurrentEquipContextProvider>
        </AllEquipsContextProvider>
      </UsersContextProvider>
    </AppContextProvider>
  </React.StrictMode>,
  document.getElementById('root')
);
