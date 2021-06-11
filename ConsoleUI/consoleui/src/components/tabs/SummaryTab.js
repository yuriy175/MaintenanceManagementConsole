import React, {useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import { SummaryDBTabPanelIndex, MainTabPanelIndex, SummaryHistoryTabPanelIndex } from '../../model/constants';

import SummaryMainTabPanel from './summaryPanels/SummaryMainTabPanel'
import SummaryBDTabPanel from './summaryPanels/SummaryBDTabPanel'
import SummaryHistoryTabPanel from './summaryPanels/SummaryHistoryTabPanel'

import { CurrentEquipContext } from '../../context/currentEquip-context';

const useStyles = makeStyles((theme) => ({
  root: {
    
  },
}));

export default function SummaryTab(props) {
  console.log("render SummaryTab");

  const classes = useStyles();
  const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);

  const tabPanelIndex = props.panel ?? MainTabPanelIndex;
  const equipName = currEquipState?.equipInfo;
  return (
    <div>
      {tabPanelIndex === MainTabPanelIndex ? <SummaryMainTabPanel/> : <></>}
      {tabPanelIndex === SummaryDBTabPanelIndex ? <SummaryBDTabPanel equipName={equipName}/> : <></>}
      {tabPanelIndex === SummaryHistoryTabPanelIndex ? <SummaryHistoryTabPanel equipName={equipName}/> : <></>}      
    </div>
  );
}