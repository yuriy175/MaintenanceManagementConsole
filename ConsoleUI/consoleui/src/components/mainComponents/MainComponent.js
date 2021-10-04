import React, {useContext, useEffect, useState} from 'react';
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
import { Redirect } from 'react-router-dom';

import { SummaryTabIndex, EquipsTabIndex, MapTabIndex, EventsTabIndex, MainTabPanelIndex,
  ControlTabIndex, CommonChat } from '../../model/constants';

import MainToolBar from './MainToolBar';
import MainInfoComponent from './MainInfoComponent';

import {AdminRole} from '../../model/constants'
import { UsersContext } from '../../context/users-context';
import { AppContext } from '../../context/app-context';
import { AllEquipsContext } from '../../context/allEquips-context';
import { EventsContext } from '../../context/events-context';
import { CommunicationContext } from '../../context/communication-context';
import { ControlStateContext} from '../../context/controlState-context';
import * as AdminWorker from '../../workers/adminWorker'
import * as EquipWorker from '../../workers/equipWorker'
import * as ControlWorker from '../../workers/controlWorker'
import {getUSFullDate} from '../../utilities/utils'

const drawerWidth = 240;
const menuItems = ['Обзор', 'Комплексы', 'Карта', 'Журнал событий', 'Панель управления']; // , 'Администрирование'];

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
    minWidth: '1300px',
    // width: `calc(100% - ${drawerWidth}px)`,
  },
}));

export default function MainComponent() {
  const classes = useStyles();  

  const [selectedIndex, setSelectedIndex] = React.useState(0);
  const [appState, appDispatch] = useContext(AppContext);
  const [usersState, usersDispatch] = useContext(UsersContext);
  const [allEquipsState, allEquipsDispatch] = useContext(AllEquipsContext);
  const [eventsState, eventsDispatch] = useContext(EventsContext);
  const [controlState, controlDispatch] = useContext(ControlStateContext);
  const [communicationState, communicationDispatch] = useContext(CommunicationContext);
  // const [redirect, setRedirect] = useState(false);

  useEffect(() => {
      (async () => {
        const token = usersState.token;
          if(!token)
          {
            return;
          }

          const users = await AdminWorker.GetAllUsers(token);
          usersDispatch({ type: 'SETUSERS', payload: users }); 
      })();
  }, [usersState.token]);

  const token = usersState.token;
  if (!token) {
    return <Redirect to="/" />;  
  }

  const handleListItemClick = async (event, index) => {
    if(index === EquipsTabIndex || index === MapTabIndex)
    {
      const allEquips = await EquipWorker.GetAllEquips(token);
      allEquipsDispatch({ type: 'SETALLEQUIPS', payload: allEquips });  
    }
    else if(index === EventsTabIndex)
    {
      const endDate = new Date();
      const allEvents = await EquipWorker.SearchEquip(token, 'Events', '', getUSFullDate(endDate), getUSFullDate(endDate));
      eventsDispatch({ type: 'SETEVENTS', payload: allEvents }); 
    }    
    else if(index === ControlTabIndex){  
      const state = await ControlWorker.GetServerState(token);
      controlDispatch({ type: 'SETSRVSTATE', payload: state });  

      const notes = await EquipWorker.GetCommunications(token, CommonChat);
      communicationDispatch({ type: 'SETCOMMONCHAT', payload: notes });     
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