import React, {useState, useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import Typography from '@material-ui/core/Typography';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import Divider from '@material-ui/core/Divider';

import {ChatMessageType} from '../../../model/constants'

import { CommunicationContext } from '../../../context/communication-context';
import { UsersContext } from '../../../context/users-context';
import * as EquipWorker from '../../../workers/equipWorker'
import {parseLocalString} from '../../../utilities/utils'
import SummaryChatSubpanel from './subpanels/SummaryChatSubpanel'


const useStyles = makeStyles((theme) => ({
  root: {
    display:"flex"
  },
  divider: {
    marginRight: "12px"
  },
}));

export default function SummaryChatTabPanel(props) {
  console.log("render SummaryChatTabPanel");

  const classes = useStyles();
  const [communicationState, communicationDispatch] = useContext(CommunicationContext);
  const [usersState, usersDispatch] = useContext(UsersContext);

  const equipName = props.equipName;
  const token = usersState.token;

  const noteType = 'Note';
  const chatType = ChatMessageType;

  const onSendNote = async (type, note) => {
    if(type === noteType){
      await onSendNoteNote(note);
    }
    else if(type === chatType){
      await onSendChatNote(note);
    }
  };

  const onDeleteNote = async (type, noteId) => {
    await EquipWorker.DeleteNote(token, equipName, chatType, noteId); 
    communicationDispatch({ type: 'DELETENOTE', payload: noteId}); 
  };

  const onSendNoteNote = async (newnote) => {
    const note = await EquipWorker.SendNewNote(token, equipName, noteType, newnote.id, newnote.note);    
    communicationDispatch({ type: newnote.id? 'CHANGENOTE' : 'ADDNOTE', payload: note}); 
  };

  const onSendChatNote = async (newnote) => {
    const note = await EquipWorker.SendNewNote(token, equipName, chatType, newnote.id, newnote.note);  
    if(newnote.id){
      communicationDispatch({ type: 'CHANGENOTE', payload: note}); 
    } 
  };

  const notes = communicationState.notes?.filter(n => n.Type === noteType);
  const chats = communicationState.notes?.filter(n => n.Type === chatType);
  return (
    <div className={classes.root}>     
      <SummaryChatSubpanel
        type={noteType}
        title='Заметки'
        notes={notes}
        onSendNote={onSendNote}
        onDeleteNote = {onDeleteNote}
        currentUser={usersState.currentUser?.Login}
      >
      </SummaryChatSubpanel>
      <Divider orientation="vertical" flexItem className={classes.divider}/>
      <SummaryChatSubpanel
        type={chatType}
        title='Чат'
        notes={chats}
        onSendNote={onSendNote}
        onDeleteNote = {onDeleteNote}
        currentUser={usersState.currentUser?.Login}
      >
      </SummaryChatSubpanel>
    </div>
  );
}
  