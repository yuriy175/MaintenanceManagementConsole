import React, {useState, useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import Typography from '@material-ui/core/Typography';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import Grid from '@material-ui/core/Grid';

import IconButton from '@material-ui/core/IconButton';
import Menu from '@material-ui/core/Menu';
import MenuItem from '@material-ui/core/MenuItem';
import MoreVertIcon from '@material-ui/icons/MoreVert';

import {parseLocalString} from '../../../../utilities/utils'

const useStyles = makeStyles((theme) => ({
  root: {
    display:"flex"
  },
  column:{
    width: "calc(50% - 6px)",
    marginRight: "12px",
  },
  fullRow:{
    display: 'grid',
    width: '100%',
    textAlign: 'left',
    gridTemplateColumns: '40px 1fr',
    // maxWidth: 'calc(50% - 46px)',
    // overflowWrap: 'break-word',
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
    // maxWidth: 'calc(50% - 0px)',
    overflowWrap: 'break-word',
    alignSelf: 'center',
    width: '100%',
    margin: '0',
    display: 'block',
  }
}));

export default function SummaryChatSubpanel(props) {
  console.log("render SummaryChatSubpanel");

  const classes = useStyles();
  const [newNote, setNewNote] = useState({id: '', note:''});

  const [anchorEl, setAnchorEl] = React.useState(null);
  const openEditMenu = Boolean(anchorEl);

  const handleMenuClick = (event, row) => {
    setAnchorEl(event.currentTarget);
  };

  const handleMenuClose = () => {
    setAnchorEl(null);
  };

  const onNoteChange = async (val)  => {
    setNewNote({id: newNote.id, note: val});
  }

  const type = props.type;
  const onSendNote = async () => {
    if(props.onSendNote){
      await props.onSendNote(type, newNote);
      setNewNote({id: '', note:''});
    }
  };

  const handleDeleteNote = async (ev) => {
    const noteId = anchorEl.id;
    if(props.onDeleteNote){
      await props.onDeleteNote(type, noteId);
    }
    handleMenuClose();
  };

  const title = props.title;
  const notes = props.notes;
  const currentUser = props.currentUser;
  const canSubmit = !!newNote?.note;
  const handleEditNote = async (ev) => {
    const noteId = anchorEl.id; // ev.target.id;
    const note = notes.find(n => n.ID == noteId);
    if(note){
      setNewNote({id: noteId, note:note.Message});
    }
    handleMenuClose();
  };

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
          value={newNote?.note}
          onChange={e => onNoteChange(e.target.value)}
        />
        <Button variant="contained" color="primary" onClick={onSendNote} disabled={!canSubmit}>
              Послать
        </Button>
        <div className={classes.fullRow}>
          {notes?.length ?
            notes.map((i, ind) => 
            {
              const row = i;
              return (
              <>
                <div key={ind.toString()}>
                    {
                      currentUser === i.User?
                            <>
                              <IconButton
                                id={i.ID}
                                aria-label="more"
                                aria-controls="long-menu"
                                aria-haspopup="true"
                                onClick={handleMenuClick}
                              >
                                <MoreVertIcon />
                              </IconButton>
                              <Menu
                                id="long-menu"
                                anchorEl={anchorEl}
                                keepMounted
                                open={openEditMenu}
                                onClose={handleMenuClose}
                              >
                                  <MenuItem onClick={handleEditNote}>Редактировать</MenuItem>
                                  <MenuItem onClick={handleDeleteNote}>Удалить</MenuItem>
                              </Menu>
                            </>
                            : <></>
                    }
                </div>
                 <div className={classes.notesArea}>
                    <Typography variant="body1" align='left' color={currentUser === i.User? 'secondary' : 'primary'} className={classes.noteTitle}>
                        {i.User +" ("} {parseLocalString(i.DateTime) + ") - "}
                     </Typography>
                     {
                       i.Message?.split("\n")?.map((s, ind) =>
                       {
                         return ind === 0? s : <Typography>{s}</Typography>
                       })
                     }
                 </div> 
               </>
              )})
              :
              <></>
          }
        </div>
        {/* <div className={classes.notesArea}>
          {notes?.length ?
            notes.map((i, ind) => (
              <div key={ind.toString()} className={classes.fullRow}>
                {currentUser === i.User?
                    <IconButton
                      aria-label="more"
                      aria-controls="long-menu"
                      aria-haspopup="true"
                      onClick={handleMenuClick}
                    >
                      <MoreVertIcon />
                    </IconButton>
                  : <></>
                }
                <Typography variant="body1" align='left' color={currentUser === i.User? 'secondary' : 'primary'} className={classes.noteTitle}>
                    {i.User +" ("} {parseLocalString(i.DateTime) + ") - "}
                </Typography>
                {
                  i.Message?.split("\n")?.map((s, ind) =>
                  {
                    return ind === 0? s : <Typography>{s}</Typography>
                  })
                }
              </div>
              ))
              :
              <></>
          }
        </div>  */}
      </div>
      );
}
