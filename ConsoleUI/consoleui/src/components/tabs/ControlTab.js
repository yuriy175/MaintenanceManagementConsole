import React, {useState, useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import { MainTabPanelIndex, ControlDiagnosticTabPanelIndex, ControlLogTabPanelIndex } from '../../model/constants';

import ControlMainTabPanel from './controlPanels/ControlMainTabPanel'
import ControlDiagnosticTabPanel from './controlPanels/ControlDiagnosticTabPanel'
import ControlLogTabPanel from './controlPanels/ControlLogTabPanel'

const useStyles = makeStyles((theme) => ({
  root: {
    
  },
}));


export default function ControlTab(props) {
  console.log("render ControlTab");

  const classes = useStyles();

  const tabPanelIndex = props.panel ?? MainTabPanelIndex;
  return (
    <div>
      {tabPanelIndex === MainTabPanelIndex ? <ControlMainTabPanel/> : <></>}
      {tabPanelIndex === ControlDiagnosticTabPanelIndex ? <ControlDiagnosticTabPanel/> : <></>}         
      {tabPanelIndex === ControlLogTabPanelIndex ? <ControlLogTabPanel/> : <></>}      
    </div>
  );
}