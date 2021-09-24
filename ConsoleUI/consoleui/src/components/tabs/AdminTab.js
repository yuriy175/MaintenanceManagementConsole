import React, {useState, useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import { MainTabPanelIndex } from '../../model/constants';

import AdminMainTabPanel from './adminPanels/AdminMainTabPanel'

const useStyles = makeStyles((theme) => ({
  root: {
    
  },
}));

export default function AdminTab(props) {
  console.log("render AdminTab");

  const classes = useStyles();

  const tabPanelIndex = props.panel ?? MainTabPanelIndex;
  return (
    <div>
      {tabPanelIndex === MainTabPanelIndex ? <AdminMainTabPanel/> : <></>} 
    </div>
  );
}

