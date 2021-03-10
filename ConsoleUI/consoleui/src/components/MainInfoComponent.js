import React from 'react';
import { makeStyles } from '@material-ui/core/styles';

import SummaryComponent from './SummaryComponent';
import MapComponent from './MapComponent';
import EventsComponent from './EventsComponent';
import HistoryComponent from './HistoryComponent';

export default function MainInfoComponent(props) {
  console.log("render MainInfoComponent");

  return (
    <div>
      {props.Index === 0 ? <SummaryComponent></SummaryComponent> : <></>}
      {props.Index === 1 ? <MapComponent></MapComponent> : <></>}
      {props.Index === 2 ? <EventsComponent></EventsComponent> : <></>}
      {props.Index === 3 ? <HistoryComponent></HistoryComponent> : <></>}
    </div>
  );
}