import React, {useState, useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import Typography from '@material-ui/core/Typography';
import TextareaAutosize from '@material-ui/core/TextareaAutosize';
import { CommunicationContext } from '../../../context/communication-context';

const useStyles = makeStyles((theme) => ({
  root: {
    width:'100%',
    borderColor: 'darkgray'
  },
}));

export default function AdminLogTabPanel(props) {
  console.log("render AdminLogTabPanel");

  const classes = useStyles();
  const [communicationState, communicationDispatch] = useContext(CommunicationContext);

  return (
    <div>
      <TextareaAutosize className={classes.root}
        rowsMax={40}
        aria-label="maximum height"
        defaultValue={communicationState.logs}
      />
        {/* <Typography variant="body2" gutterBottom>
            {communicationState.logs}
        </Typography> */}
    </div>
  );
}
  