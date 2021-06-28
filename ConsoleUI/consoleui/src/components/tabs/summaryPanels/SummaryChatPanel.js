import React, {useState, useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import Typography from '@material-ui/core/Typography';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';

import { CommunicationContext } from '../../../context/communication-context';
import * as EquipWorker from '../../../workers/equipWorker'

const useStyles = makeStyles((theme) => ({
  root: {
    display:"flex"
  },
  column:{
    width: "50%",
    marginRight: "12px",
  },
  fullRow:{
    width: '100%',
  }
}));

export default function SummaryChatPanel(props) {
  console.log("render SummaryChatPanel");

  const classes = useStyles();
  const [communicationState, communicationDispatch] = useContext(CommunicationContext);

  const equipName = props.equipName;

  const onSend = async () => {
    const note = 'blabla';
    const noteModel = await EquipWorker.SendNewNote(equipName, 'Note', note);
    /*const claims = parseJwt(data);
    if(data){
      usersDispatch({ type: 'SETUSER', payload: {Token: data, Claims: claims} }); 
      setRedirect(true);
    }
    else{
      setShowError(true);
    }*/
  };

  return (
    <div className={classes.root}>
      <div className={classes.column}>
        <Typography variant="h5">Заметки</Typography>
        <Typography variant="body2" gutterBottom>
            {communicationState.notes}
        </Typography>
        <TextField
          id="outlined-multiline-static"
          className={classes.fullRow}
          label="Новое сообщение"
          multiline
          rows={5}
          defaultValue=""
          variant="outlined"
        />
        <Button variant="contained" color="primary" onClick={onSend}>
              Послать
        </Button>
      </div>
      <div className={classes.column}>
        <Typography variant="h5">Чат</Typography>
      </div>    
    </div>
  );
}
  