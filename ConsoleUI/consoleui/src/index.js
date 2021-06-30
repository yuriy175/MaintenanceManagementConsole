import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter } from 'react-router-dom';
import './index.css';
import { UseDarkTheme } from './model/constants'
import App from './App';
import { ThemeProvider, createMuiTheme } from '@material-ui/core/styles';

import { AllEquipsContextProvider } from './context/allEquips-context';
import { CurrentEquipContextProvider } from './context/currentEquip-context';
import { UsersContextProvider} from './context/users-context';
import { AppContextProvider} from './context/app-context';
import { EventsContextProvider} from './context/events-context';
import { SystemVolatileContextProvider} from './context/systemVolatile-context';
import { CommunicationContextProvider} from './context/communication-context';

const theme = createMuiTheme({
  palette: {
    type: !UseDarkTheme ? "light" : "dark",
  }
});

ReactDOM.render(
  <React.StrictMode>
    <ThemeProvider theme={theme}>
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
    </ThemeProvider>
  </React.StrictMode>,
  document.getElementById('root')
);
