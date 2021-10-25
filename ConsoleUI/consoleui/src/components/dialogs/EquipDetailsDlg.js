import React from 'react';
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import TextField from '@material-ui/core/TextField';

export default function EquipDetailsDlg(props){
  const equipName = props.equip?.EquipName;
  const equipAlias = props.equip?.EquipAlias;
  const hospitalLatitude = props.equip?.HospitalLatitude;
  const hospitalLongitude = props.equip?.HospitalLongitude;
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
      <DialogTitle id="alert-dialog-title">Карточка оборудования ({equipName})</DialogTitle>
      <DialogContent>
      <TextField
            autoFocus
            margin="dense"
            id="alias"
            label="Алиас"
            fullWidth
            variant="standard"
            value={equipAlias}
          />

        <TextField
            autoFocus
            margin="dense"
            id="latitude"
            label="Широта"
            fullWidth
            variant="standard"
            value={hospitalLatitude}
          />

          <TextField
            autoFocus
            margin="dense"
            id="longitude"
            label="Долгота"
            fullWidth
            variant="standard"
            value={hospitalLongitude}
          />
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