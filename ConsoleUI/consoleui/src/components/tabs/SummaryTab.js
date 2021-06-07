import React, {useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import { SummaryDBTabPanelIndex, MainTabPanelIndex, SummaryHistoryTabPanelIndex } from '../../model/constants';

import SummaryMainTabPanel from './summaryPanels/SummaryMainTabPanel'
import SummaryBDTabPanel from './summaryPanels/SummaryBDTabPanel'
import SummaryHistoryTabPanel from './summaryPanels/SummaryHistoryTabPanel'

const useStyles = makeStyles((theme) => ({
  root: {
    
  },
}));

export default function SummaryTab(props) {
  console.log("render SummaryTab");

  const classes = useStyles();
  const tabPanelIndex = props.panel ?? MainTabPanelIndex;
  return (
    <div>
      {tabPanelIndex === MainTabPanelIndex ? <SummaryMainTabPanel/> : <></>}
      {tabPanelIndex === SummaryDBTabPanelIndex ? <SummaryBDTabPanel/> : <></>}
      {tabPanelIndex === SummaryHistoryTabPanelIndex ? <SummaryHistoryTabPanel/> : <></>}      
    </div>
  );
}