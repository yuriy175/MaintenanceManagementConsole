import React from 'react';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles((theme) => ({
  root: {
    display:"flex"
  },
}));

export default function MapTab(props) {
  console.log("render MapTab");

  const classes = useStyles();

  return (
    <div className={classes.root}>
    </div>
  );
}