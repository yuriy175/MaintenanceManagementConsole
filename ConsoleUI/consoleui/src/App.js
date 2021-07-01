import React, { useContext, useState } from 'react';
import { NavLink, Route } from 'react-router-dom';

import { ThemeProvider, createMuiTheme } from '@material-ui/core/styles';

import { UseDarkTheme } from './model/constants'
import MainComponent from './components/mainComponents/MainComponent';
import AuthComponent from './components/mainComponents/AuthComponent';
import './App.css';


const theme = createMuiTheme({
  palette: {
    type: !UseDarkTheme ? "light" : "dark",
  }
});


function App() {
  return (  
    <>  
      <ThemeProvider theme={theme}>
        <div className="App">
          <Route exact path="/" component={AuthComponent} />
          <Route path="/auth" component={AuthComponent}/>
          <Route path="/info" component={MainComponent} />    
          {/* <MainComponent></MainComponent> */}
        </div>    
      </ThemeProvider>
    </>
  );
}

export default App;
