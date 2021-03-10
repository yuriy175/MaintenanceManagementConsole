import React from 'react';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles((theme) => ({
  root: {
    display:"flex"
  },
}));

export default function MapComponent(props) {
  console.log("render MapComponent");

  const classes = useStyles();

  return (
    <div className={classes.root}>
    </div>
  );
}