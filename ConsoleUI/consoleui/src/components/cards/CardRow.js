import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import RemoveRedEyeIcon from '@material-ui/icons/RemoveRedEye';

const useStyles = makeStyles({
 root: {
   display: "flex",
  },
  descr: {
    width:'65%',
    margin: '0 0.5em',
    textAlign: 'left',
  },
  value: {
    marginRight: '0.5em',
    fontWeight: 'bold',
    width:'30%',
    textAlign: 'right',
  },
  errorDescr: {
    width:'10%',
  },
  errorValue: {
    width:'80%',
  },
});

export default function CardRow(props) {

  const classes = useStyles();
  const bull = <span className={classes.bullet}>•</span>;

  return (
    <div className={classes.root}>
        {props.icon !== 'none' ? <RemoveRedEyeIcon color="primary" size="small"></RemoveRedEyeIcon> : <></>}
        <Typography className={classes.descr} color="textSecondary">
          {props.descr}
        </Typography>
        <Typography className={classes.value} color="textSecondary" style = {{
            width: props.rightWidth ? props.rightWidth : classes.value.width,
            color: props.rightColor ? props.rightColor : "",
          }}>
          {props.value}
        </Typography>
    </div>
  );
}

export function CardErrorRow(props) {

  const classes = useStyles();

  return (
    <div className={classes.root}>
        <RemoveRedEyeIcon color="secondary" size="small"></RemoveRedEyeIcon>
        <Typography className={classes.descr, classes.errorDescr} color="secondary">
          {props.descr}
        </Typography>
        <Typography className={classes.value, classes.errorValue} color="secondary" style = {{
            width: props.rightWidth ? props.rightWidth : classes.errorValue.width,
            color: props.rightColor ? props.rightColor : "",
          }}>
          {props.value}
        </Typography>
    </div>
  );
}