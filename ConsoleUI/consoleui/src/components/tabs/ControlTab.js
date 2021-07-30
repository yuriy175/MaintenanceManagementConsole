import React, {useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import {CommonChat, ChatMessageType} from '../../model/constants'
import { CommunicationContext } from '../../context/communication-context';
import { UsersContext } from '../../context/users-context';
import { ControlStateContext } from '../../context/controlState-context';
import * as EquipWorker from '../../workers/equipWorker'

import ServerStateCard from '../cards/controlCards/ServerStateCard'
import UnasweredEquipsCard from '../cards/controlCards/UnasweredEquipsCard'

import SummaryChatSubpanel from './summaryPanels/subpanels/SummaryChatSubpanel'

const useStyles = makeStyles((theme) => ({
  root: {
    display:"flex"
  },
  column:{
    width: "30%",
    marginRight: "12px",
  }
}));

export default function ControlTab(props) {
  console.log("render ControlTab");

  const [communicationState, communicationDispatch] = useContext(CommunicationContext);
  const [usersState, usersDispatch] = useContext(UsersContext);
  const [controlState, controlDispatch] = useContext(ControlStateContext);

  const chatType = ChatMessageType;
  const classes = useStyles();

  const equipName = CommonChat;
  const token = usersState.token;

  const onSendNote = async (type, newnote) => {
    if(type === chatType){
      const note = await EquipWorker.SendNewNote(token, equipName, chatType, newnote.id, newnote.note);  
      if(newnote.id){
        communicationDispatch({ type: 'CHANGENOTE', payload: note}); 
      } 
    }
  };

  const onDeleteNote = async (type, noteId) => {
    await EquipWorker.DeleteNote(token, equipName, chatType, noteId); 
    communicationDispatch({ type: 'DELETENOTE', payload: noteId}); 
  };

  const chats = communicationState.commonNotes?.filter(n => n.Type === chatType);
  const serverState = controlState.serverState;

  return (
    <div className={classes.root}>
      <div className={classes.column}>
        <ServerStateCard serverState={serverState}></ServerStateCard>
        <UnasweredEquipsCard></UnasweredEquipsCard>
      </div>
      <SummaryChatSubpanel
        type={chatType}
        title='Чат'
        notes={chats}
        onSendNote={onSendNote}
        onDeleteNote = {onDeleteNote}
        currentUser=''
      >
      </SummaryChatSubpanel>
    </div>
  );
}