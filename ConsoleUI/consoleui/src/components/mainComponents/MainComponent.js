import React, {useContext, useEffect} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Drawer from '@material-ui/core/Drawer';
import CssBaseline from '@material-ui/core/CssBaseline';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import List from '@material-ui/core/List';
import Typography from '@material-ui/core/Typography';
import Divider from '@material-ui/core/Divider';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import InboxIcon from '@material-ui/icons/MoveToInbox';
import MailIcon from '@material-ui/icons/Mail';

import MainToolBar from './MainToolBar';
import MainInfoComponent from './MainInfoComponent';

import { UsersContext } from '../../context/users-context';
import * as AdminWorker from '../../workers/adminWorker'

const drawerWidth = 240;
const mainMenu = ['Обзор', 'Карта', 'Журнал событий', 'История', 'Администрироание'];

const useStyles = makeStyles((theme) => ({
  root: {
    display: 'flex',
  },
  appBar: {
    width: `calc(100% - ${drawerWidth}px)`,
    marginLeft: drawerWidth,
  },
  drawer: {
    width: drawerWidth,
    flexShrink: 0,
  },
  drawerPaper: {
    width: drawerWidth,
  },
  // necessary for content to be below app bar
  toolbar: theme.mixins.toolbar,
  content: {
    marginTop: 64, // `${drawerWidth}px`,
    flexGrow: 1,
    backgroundColor: theme.palette.background.default,
    padding: theme.spacing(3),
  },
}));

export default function MainComponent() {
  const classes = useStyles();

  const [selectedIndex, setSelectedIndex] = React.useState(0);
  const [usersState, usersDispatch] = useContext(UsersContext);

  useEffect(() => {
      (async () => {
          if(usersState.users)
          {
            return;
          }

          const users = await AdminWorker.GetAllUsers();
          usersDispatch({ type: 'SETUSERS', payload: users }); 
      })();
  }, [usersState.users]);

  const handleListItemClick = (event, index) => {
    setSelectedIndex(index);
  };

  return (
    <div className={classes.root}>
      <CssBaseline />
      <MainToolBar></MainToolBar>
      <Drawer
        className={classes.drawer}
        variant="permanent"
        classes={{
          paper: classes.drawerPaper,
        }}
        anchor="left"
      >
        <div className={classes.toolbar} />
        <Divider />
        <List>
          {mainMenu.map((text, index) => (
            <ListItem button key={text}
                selected={selectedIndex === index}
                onClick={(event) => handleListItemClick(event, index)}
            >
              <ListItemIcon>{index % 2 === 0 ? <InboxIcon /> : <MailIcon />}</ListItemIcon>
              <ListItemText primary={text} />
            </ListItem>
          ))}
        </List>
      </Drawer>
      <main className={classes.content}>
          <MainInfoComponent Title={mainMenu[selectedIndex]} Index={selectedIndex}></MainInfoComponent>
      </main>
    </div>
  );
}