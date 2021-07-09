import React from 'react';
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import Slide from '@material-ui/core/Slide';

export default function ConfirmDlg(props){
  const handleClose = () => {
    props?.onClose(false);
  };

  const handleCloseOK = () => {
    props?.onClose(true, props.context);
  };

  return (
    <Dialog
      open={props.open}
      onClose={handleClose}
      aria-labelledby="alert-dialog-title"
      aria-describedby="alert-dialog-description"
    >
      <DialogTitle id="alert-dialog-title">Подтверждение</DialogTitle>
      <DialogContent>
        <DialogContentText id="alert-dialog-description">
          {props.сonfirmMessage}
        </DialogContentText>
      </DialogContent>
      <DialogActions>
        <Button onClick={handleCloseOK} autoFocus>
          Да
        </Button>
        <Button onClick={handleClose} >
          Нет
        </Button>
      </DialogActions>
    </Dialog>
  );
}