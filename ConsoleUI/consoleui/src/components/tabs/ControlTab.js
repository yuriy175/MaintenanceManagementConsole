import React, {useState, useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import { MainTabPanelIndex, ControlDiagnosticTabPanelIndex } from '../../model/constants';

import ControlMainTabPanel from './controlPanels/ControlMainTabPanel'
import ControlDiagnosticTabPanel from './controlPanels/ControlDiagnosticTabPanel'

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
    </div>
  );
}