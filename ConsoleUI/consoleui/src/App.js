import React, { useContext, useState } from 'react';
import { NavLink, Route } from 'react-router-dom';

import MainComponent from './components/mainComponents/MainComponent';
import AuthComponent from './components/mainComponents/AuthComponent';
import './App.css';

function App() {
  return (
    <div className="App">
      <Route exact path="/" component={AuthComponent} />
      <Route path="/auth" component={AuthComponent}/>
      <Route path="/info" component={MainComponent} />    
      {/* <MainComponent></MainComponent> */}
    </div>
  );
}

export default App;
