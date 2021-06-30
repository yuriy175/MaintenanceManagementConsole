import React, {useState, useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import Typography from '@material-ui/core/Typography';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';

import { CommunicationContext } from '../../../context/communication-context';
import { UsersContext } from '../../../context/users-context';
import * as EquipWorker from '../../../workers/equipWorker'
import {parseLocalString} from '../../../utilities/utils'

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
    textAlign: 'left',
  },
  textField:{
    marginTop: theme.spacing(1),
    marginBottom: theme.spacing(1),
    width: '100%',

  } ,
  noteTitle:{
    width: '100%',
    fontWeight: 'bolder',
    textAlign: 'left',
    display: 'inline',
  }
}));

export default function SummaryChatPanel(props) {
  console.log("render SummaryChatPanel");

  const classes = useStyles();
  const [communicationState, communicationDispatch] = useContext(CommunicationContext);
  const [usersState, usersDispatch] = useContext(UsersContext);
  const [newNote, setNewNote] = useState('');

  const equipName = props.equipName;
  const token = usersState.token;

  const onNoteChange = async (val)  => {
    setNewNote(val);
  }

  const onSend = async () => {
    const note = await EquipWorker.SendNewNote(token, equipName, 'Note', newNote);    
    communicationDispatch({ type: 'ADDNOTE', payload: note}); 
    setNewNote('');
  };

  const notes = communicationState.notes;
  return (
    <div className={classes.root}>
      <div className={classes.column}>
        <Typography variant="h5">Заметки</Typography>
        {notes?.length ? 
          notes.map((i, ind) => (
            <div key={ind.toString()} className={classes.fullRow}>
              <Typography variant="body1" align='left' color='primary' className={classes.noteTitle}>
                  {i.User +" ("} {parseLocalString(i.DateTime) + ") - "}
              </Typography> 
              {i.Message}
            </div>
            ))
            :
            <></>          
        }
        <TextField
          id="outlined-multiline-static"
          className={classes.textField}
          label="Новое сообщение"
          multiline
          rows={5}
          defaultValue=""
          variant="outlined"
          value={newNote}
          onChange={e => onNoteChange(e.target.value)}
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
  