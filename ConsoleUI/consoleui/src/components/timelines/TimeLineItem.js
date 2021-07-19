import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import Divider from '@material-ui/core/Divider';
import ListItemText from '@material-ui/core/ListItemText';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import Avatar from '@material-ui/core/Avatar';
import Typography from '@material-ui/core/Typography';
import Collapse from '@material-ui/core/Collapse';
import ExpandLess from '@material-ui/icons/ExpandLess';
import ExpandMore from '@material-ui/icons/ExpandMore';

const useStyles = makeStyles((theme) => ({
  root: {
    width: '100%',
    // maxWidth: '36ch',
    backgroundColor: theme.palette.background.paper,
  },
  inline: {
    display: 'inline',
  },
  equip: {
    width: '10%',
  },
  isImportant:{
    backgroundColor: 'red',
  },
  collapse:{
    textAlign:'left',
  },
  outerLabel:{
    marginTop:'6px',
    marginRight:'12px',
  },
  ordinaryLabel:{
    borderRadius:'50%',
    backgroundColor: 'gray',
  },
  importantLabel:{    
    borderRadius:'50%',
    backgroundColor: 'red',
  },
  innerLabel:{
    width:'40px',
    height:'40px',
    display: 'flex',
    fontSize: '1.25rem',
    color: '#303030',
    alignItems: 'center',
    justifyContent: 'center',
  }
}));

export default function TimeLineItem(props) {
  const classes = useStyles();

  const [open, setOpen] = React.useState(false);

  const handleClick = () => {
    setOpen(!open);
  };

  const isImportant = props.isImportant;
  const equipName = props.equipName;
  const time = props.time;
  const title = props.title;
  const text = props.text;
  const details = props.details;
  return (
    <div>
        <ListItem alignItems="flex-start" button>
            <div className={classes.outerLabel}>
              <div className={isImportant ? classes.importantLabel : classes.ordinaryLabel}>
                <div className={classes.innerLabel}>
                  {isImportant ? '!' : 'O'}
                </div>
              </div>
            </div>
            {/* <ListItemAvatar >
                <Avatar className={isImportant ? classes.isImportant : null} alt={isImportant ? "!" : ""} src={'xz'}/>
            </ListItemAvatar>  */}
            <ListItemText
                primary={
                    <React.Fragment>
                        <Typography
                            component="span"
                            variant="body2"
                            className={classes.inline}
                            color="textPrimary"
                        >
                            {equipName}
                        </Typography>
                        {" " + time}
                    </React.Fragment>
                }
                secondary={
                    <React.Fragment>
                        <Typography
                            component="span"
                            variant="body2"
                            className={classes.inline}
                            color="textPrimary"
                        >
                            {title}
                        </Typography>
                        {" - " + text}
                    </React.Fragment>
                }                
            />
            {details && open ? <ExpandLess onClick={handleClick}/> : 
              details && !open ? <ExpandMore onClick={handleClick}/> : <></>}
        </ListItem>
        
            {details ? 
              <Collapse in={open} timeout="auto" unmountOnExit className={classes.collapse}>
                {details}
              </Collapse>: <></>}
        </div>
  );
}