import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import DialogTitle from '@material-ui/core/DialogTitle';
import Dialog from '@material-ui/core/Dialog';
import Typography from '@material-ui/core/Typography';
import { blue } from '@material-ui/core/colors';

const useStyles = makeStyles({
  text: {
    margin: '1em',
  },
});

export default function NotifyDlg(props) {
  const classes = useStyles();  
  const [open, setOpen] = React.useState(!!props.text);

//   if(!!props.text && !open){
//     setOpen(true);
//   }

  const handleClose = (value) => {
    setOpen(false);
  };

  return (
    <Dialog onClose={handleClose} aria-labelledby="simple-dialog-title" open={open}>
      <DialogTitle id="simple-dialog-title">{props.title}</DialogTitle>
      <Typography className={classes.text}>{props.text}</Typography>
    </Dialog>
  );
}
