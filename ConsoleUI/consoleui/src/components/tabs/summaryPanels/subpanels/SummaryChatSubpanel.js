import React, {useState, useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import Typography from '@material-ui/core/Typography';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';

import {parseLocalString} from '../../../../utilities/utils'

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
  },
  notesArea:{
    maxWidth: '100%',
    // overflowX: 'auto',
    overflowWrap: 'break-word',
  }
}));

export default function SummaryChatSubpanel(props) {
  console.log("render SummaryChatSubpanel");

  const classes = useStyles();
  const [newNote, setNewNote] = useState('');

  const onNoteChange = async (val)  => {
    setNewNote(val);
  }

  const type = props.type;
  const onSendNote = async () => {
    if(props.onSendNote){
      await props.onSendNote(type, newNote);
      setNewNote('');
    }
  };

  const title = props.title;
  const notes = props.notes;
  const currentUser = props.currentUser;
  const canSubmit = !!newNote;
  return (
      <div className={classes.column}>
        <Typography variant="h5">{title}</Typography>
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
        <Button variant="contained" color="primary" onClick={onSendNote} disabled={!canSubmit}>
              Послать
        </Button>
        <div className={classes.notesArea}>
          {notes?.length ? 
            notes.map((i, ind) => (
              <div key={ind.toString()} className={classes.fullRow}>
                <Typography variant="body1" align='left' color={currentUser === i.User? 'secondary' : 'primary'} className={classes.noteTitle}>
                    {i.User +" ("} {parseLocalString(i.DateTime) + ") - "}
                </Typography> 
                {
                  i.Message?.split("\n")?.map((s, ind) => 
                  {
                    return ind === 0? s : <Typography>{s}</Typography>
                  }
                  )
                }              
              </div>
              ))
              :
              <></>          
          }  
        </div>      
      </div>
      );
}
  