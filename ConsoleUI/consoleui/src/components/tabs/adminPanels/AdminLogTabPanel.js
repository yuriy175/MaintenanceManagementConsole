import React, {useState, useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import Typography from '@material-ui/core/Typography';
import { CommunicationContext } from '../../../context/communication-context';

const useStyles = makeStyles((theme) => ({
  root: {
    width:'`100%'
  },
}));

export default function AdminLogTabPanel(props) {
  console.log("render AdminLogTabPanel");

  const classes = useStyles();
  const [communicationState, communicationDispatch] = useContext(CommunicationContext);

  return (
    <div>
        <Typography variant="body2" gutterBottom>
            {communicationState.logs}
        </Typography>
    </div>
  );
}
  