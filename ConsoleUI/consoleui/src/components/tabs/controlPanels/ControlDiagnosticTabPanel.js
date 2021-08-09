import React, {useState, useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import Typography from '@material-ui/core/Typography';
import TextareaAutosize from '@material-ui/core/TextareaAutosize';
import { ControlStateContext } from '../../context/controlState-context';

const useStyles = makeStyles((theme) => ({
  root: {
    width:'100%',
    borderColor: 'darkgray'
  },
}));

export default function ControlDiagnosticTabPanel(props) {
  console.log("render ControlDiagnosticTabPanel");

  const classes = useStyles();
  const [controlState, controlDispatch] = useContext(ControlStateContext); 

  return (
    <div>
      <TextareaAutosize className={classes.root}
        rowsMax={40}
        aria-label="maximum height"
        defaultValue={controlState.diagnostic}
      />
    </div>
  );
}
  