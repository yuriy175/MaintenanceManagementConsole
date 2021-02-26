import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter } from 'react-router-dom';
import './index.css';
import App from './App';

import { AllEquipsContextProvider } from './context/allEquips-context';
import { CurrentEquipContextProvider } from './context/currentEquip-context';

ReactDOM.render(
  <React.StrictMode>
    <AllEquipsContextProvider>
      <CurrentEquipContextProvider>
        <BrowserRouter>
          <App />
        </BrowserRouter>
      </CurrentEquipContextProvider>
    </AllEquipsContextProvider>
  </React.StrictMode>,
  document.getElementById('root')
);
