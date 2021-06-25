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

import { SummaryTabIndex, EquipsTabIndex, EventsTabIndex, MainTabPanelIndex } from '../../model/constants';

import MainToolBar from './MainToolBar';
import MainInfoComponent from './MainInfoComponent';

import {AdminRole} from '../../model/constants'
import { UsersContext } from '../../context/users-context';
import { AppContext } from '../../context/app-context';
import { AllEquipsContext } from '../../context/allEquips-context';
import { EventsContext } from '../../context/events-context';
import * as AdminWorker from '../../workers/adminWorker'
import * as EquipWorker from '../../workers/equipWorker'
import {getUSFullDate} from '../../utilities/utils'

const drawerWidth = 240;
const menuItems = ['Обзор', 'Комплексы', 'Карта', 'Журнал событий']; // , 'Администрирование'];

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
  const [appState, appDispatch] = useContext(AppContext);
  const [usersState, usersDispatch] = useContext(UsersContext);
  const [allEquipsState, allEquipsDispatch] = useContext(AllEquipsContext);
  const [eventsState, eventsDispatch] = useContext(EventsContext);

  useEffect(() => {
      (async () => {
          if(!usersState.token)
          {
            return;
          }

          const users = await AdminWorker.GetAllUsers(usersState.token);
          usersDispatch({ type: 'SETUSERS', payload: users }); 
      })();
  }, [usersState.token]);

  const handleListItemClick = async (event, index) => {
    if(index === EquipsTabIndex)
    {
      const allEquips = await EquipWorker.GetAllEquips();
      allEquipsDispatch({ type: 'SETALLEQUIPS', payload: allEquips });  
    }
    else if(index === EventsTabIndex)
    {
      const endDate = new Date();
      const allEvents = await EquipWorker.SearchEquip('Events', '', getUSFullDate(endDate), getUSFullDate(endDate));
      eventsDispatch({ type: 'SETEVENTS', payload: allEvents }); 
    }

    appDispatch({ type: 'SETTAB', payload: {tab: index, panel: MainTabPanelIndex} }); 
    // setSelectedIndex(index);
  };

  const selectedTab = appState.currentTab?.tab ?? SummaryTabIndex;
  const isAdmin = usersState.currentUser?.Role === AdminRole;
  const mainMenu = isAdmin ? [...menuItems, 'Администрирование'] : menuItems; // , ];
  
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
                selected={selectedTab === index}
                onClick={(event) => handleListItemClick(event, index)}
            >
              <ListItemIcon>{index % 2 === 0 ? <InboxIcon /> : <MailIcon />}</ListItemIcon>
              <ListItemText primary={text} />
            </ListItem>
          ))}
        </List>
      </Drawer>
      <main className={classes.content}>
          <MainInfoComponent Title={mainMenu[selectedIndex]}></MainInfoComponent>
      </main>
    </div>
  );
}