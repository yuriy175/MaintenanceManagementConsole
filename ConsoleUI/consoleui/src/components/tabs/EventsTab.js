import React from 'react';
import { makeStyles } from '@material-ui/core/styles';

import SummaryHistoryTabPanel from './summaryPanels/SummaryHistoryTabPanel'

const useStyles = makeStyles((theme) => ({
  root: {
    // display:"flex"
  },
}));

export default function EventsTab(props) {
  console.log("render EventsTab");

  const classes = useStyles();
  const equipName = '';

  return (
    <div className={classes.root}>
       <SummaryHistoryTabPanel equipName={equipName}/>
    </div>
  );
}