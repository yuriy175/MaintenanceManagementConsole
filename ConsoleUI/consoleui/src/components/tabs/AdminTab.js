import React, {useState, useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import { MainTabPanelIndex, AdminLogTabPanelIndex } from '../../model/constants';

import AdminMainTabPanel from './adminPanels/AdminMainTabPanel'
import AdminLogTabPanel from './adminPanels/AdminLogTabPanel'

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
      {tabPanelIndex === AdminLogTabPanelIndex ? <AdminLogTabPanel/> : <></>}     
    </div>
  );
}

