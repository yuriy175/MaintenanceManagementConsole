import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';

import SummaryComponent from './SummaryComponent';
import MapComponent from './MapComponent';
import EventsComponent from './EventsComponent';
import HistoryComponent from './HistoryComponent';

const drawerWidth = 240;

const useStyles = makeStyles((theme) => ({
  appBar: {
    width: `calc(100% - ${drawerWidth}px)`,
    marginLeft: drawerWidth,
  },
}));

export default function MainInfoComponent(props) {
  console.log("render MainInfoComponent");

  const classes = useStyles();

  return (
    <div>
      <AppBar position="fixed" className={classes.appBar}>
        <Toolbar>
          <Typography variant="h6" noWrap>
            {props.Title}
          </Typography>
        </Toolbar>
      </AppBar>
      {props.Index === 0 ? <SummaryComponent></SummaryComponent> : <></>}
    </div>
  );
}