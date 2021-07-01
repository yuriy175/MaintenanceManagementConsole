import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter } from 'react-router-dom';
import './index.css';
import App from './App';

import { AllEquipsContextProvider } from './context/allEquips-context';
import { CurrentEquipContextProvider } from './context/currentEquip-context';
import { UsersContextProvider} from './context/users-context';
import { AppContextProvider} from './context/app-context';
import { EventsContextProvider} from './context/events-context';
import { SystemVolatileContextProvider} from './context/systemVolatile-context';
import { CommunicationContextProvider} from './context/communication-context';

ReactDOM.render(
  <React.StrictMode>
      <AppContextProvider>
        <UsersContextProvider>
          <AllEquipsContextProvider>
            <CurrentEquipContextProvider>
              <SystemVolatileContextProvider>
                <EventsContextProvider>
                  <CommunicationContextProvider>
                    <BrowserRouter>
                      <App />
                    </BrowserRouter>
                  </CommunicationContextProvider>
                </EventsContextProvider>
              </SystemVolatileContextProvider>
            </CurrentEquipContextProvider>
          </AllEquipsContextProvider>
        </UsersContextProvider>
      </AppContextProvider>
  </React.StrictMode>,
  document.getElementById('root')
);
