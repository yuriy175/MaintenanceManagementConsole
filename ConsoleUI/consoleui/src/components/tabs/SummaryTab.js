import React, {useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import { SummaryDBTabPanelIndex, MainTabPanelIndex } from '../../model/constants';

import SummaryMainTabPanel from './summaryPanels/SummaryMainTabPanel'
import SummaryBDTabPanel from './summaryPanels/SummaryBDTabPanel'

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
    </div>
  );
}