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
            <ListItemAvatar >
                <Avatar className={isImportant ? classes.isImportant : null} alt="Remy Sharp" src="/static/images/avatar/1.jpg" />
            </ListItemAvatar>
            {/* <ListItemText className={classes.equip} primary={equipName} secondary={time} /> */}
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