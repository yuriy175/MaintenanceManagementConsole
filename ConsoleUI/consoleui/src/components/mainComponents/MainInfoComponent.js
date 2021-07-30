import React, {useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import { AppContext } from '../../context/app-context';
import { SummaryTabIndex, EquipsTabIndex, MapTabIndex, EventsTabIndex,AdminTabIndex, MainTabPanelIndex,
  ControlTabIndex } from '../../model/constants';

import SummaryTab from '../tabs/SummaryTab';
import MapTab from '../tabs/MapTab';
import EventsTab from '../tabs/EventsTab';
import HistoryTab from '../tabs/HistoryTab';
import AdminTab from '../tabs/AdminTab';
import EquipsTab from '../tabs/EquipsTab';
import ControlTab from '../tabs/ControlTab';

export default function MainInfoComponent(props) {
  console.log("render MainInfoComponent");
  const [appState, appDispatch] = useContext(AppContext);

  const tabIndex = appState.currentTab?.tab ?? SummaryTabIndex;
  const tabPanelIndex = appState.currentTab?.panel ?? MainTabPanelIndex;
  return (
    <div>
      {tabIndex === SummaryTabIndex ? <SummaryTab panel={tabPanelIndex}></SummaryTab> : <></>}
      {tabIndex === EquipsTabIndex ? <EquipsTab></EquipsTab> : <></>}
      {tabIndex === MapTabIndex ? <MapTab></MapTab> : <></>}
      {tabIndex === EventsTabIndex ? <EventsTab></EventsTab> : <></>}
      {/* {tabIndex === HistoryTabIndex ? <HistoryTab></HistoryTab> : <></>} */}
      {tabIndex === ControlTabIndex ? <ControlTab panel={tabPanelIndex}></ControlTab> : <></>} 
      {tabIndex === AdminTabIndex ? <AdminTab panel={tabPanelIndex}></AdminTab> : <></>} 
    </div>
  );
}