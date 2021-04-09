import React from 'react';
import { makeStyles } from '@material-ui/core/styles';

import SummaryTab from '../tabs/SummaryTab';
import MapTab from '../tabs/MapTab';
import EventsTab from '../tabs/EventsTab';
import HistoryTab from '../tabs/HistoryTab';
import AdminTab from '../tabs/AdminTab';

export default function MainInfoComponent(props) {
  console.log("render MainInfoComponent");

  return (
    <div>
      {props.Index === 0 ? <SummaryTab></SummaryTab> : <></>}
      {props.Index === 1 ? <MapTab></MapTab> : <></>}
      {props.Index === 2 ? <EventsTab></EventsTab> : <></>}
      {props.Index === 3 ? <HistoryTab></HistoryTab> : <></>}
      {props.Index === 4 ? <AdminTab></AdminTab> : <></>}
    </div>
  );
}