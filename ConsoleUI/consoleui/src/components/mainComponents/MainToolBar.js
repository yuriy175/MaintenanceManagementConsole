import React, { useState, useEffect, useRef, useContext } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import FormHelperText from '@material-ui/core/FormHelperText';
import FormControl from '@material-ui/core/FormControl';
import NativeSelect from '@material-ui/core/NativeSelect';
import Select from '@material-ui/core/Select';
import MenuItem from '@material-ui/core/MenuItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import LocationOnOutlinedIcon from '@material-ui/icons/LocationOnOutlined';
import ListSubheader from '@material-ui/core/ListSubheader';
import LocationOffOutlinedIcon from '@material-ui/icons/LocationOffOutlined';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import Button from '@material-ui/core/Button';

import "../../styles/styles.css";
import { SummaryTabIndex, SummaryDBTabPanelIndex, MainTabPanelIndex, SummaryHistoryTabPanelIndex, SummaryChatTabPanelIndex,
  AdminTabIndex, AdminLogTabPanelIndex } from '../../model/constants';

import { AppContext } from '../../context/app-context';
import { AllEquipsContext } from '../../context/allEquips-context';
import { EventsContext } from '../../context/events-context';
import { CurrentEquipContext } from '../../context/currentEquip-context';
import { UsersContext } from '../../context/users-context';
import { CommunicationContext } from '../../context/communication-context';
import {useSetCurrEquip} from '../../hooks/useSetCurrEquip'

import * as EquipWorker from '../../workers/equipWorker'
import * as AdminWorker from '../../workers/adminWorker'
// import * as WebSocket from '../../workers/webSocket'
import {sessionUid} from '../../utilities/utils'
import { useWebSocket } from '../../hooks/useWebSocket'
import { SettingsBackupRestore } from '@material-ui/icons';
import {getUSFullDate} from '../../utilities/utils'
import AdminLogTabPanel from '../tabs/adminPanels/AdminLogTabPanel';
  
const drawerWidth = 240;

const useStyles = makeStyles((theme) => ({
  appBar: {
    width: `calc(100% - ${drawerWidth}px)`,
    marginLeft: drawerWidth,
  },
  formControl: {
    margin: theme.spacing(1),
    minWidth: 270,    
  },
  tabControl: {
    margin: theme.spacing(1),
    minWidth: 640,    
  },
  selectEmpty: {
    // marginTop: theme.spacing(2),
    color: "white",
    display: 'flex',
  },
  optionStyle:{
    backgroundColor: "#3f51b5",
    color:"white",
  },
  userName:{
    // textAlign: "end",
    width: `calc(100% - 180px)`,
  },
  button: {
    marginRight: '0.5em',
    width:'30%',
  },
}));

export default function MainToolBar() {
  console.log(`! render MainToolBar ` + sessionUid);

  const classes = useStyles();
  const [appState, appDispatch] = useContext(AppContext);
  const [allEquipsState, allEquipsDispatch] = useContext(AllEquipsContext);
  const [eventsState, eventsDispatch] = useContext(EventsContext);
  const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);
  const [usersState, usersDispatch] = useContext(UsersContext);
  const [communicationState, communicationDispatch] = useContext(CommunicationContext);
  
  // const [currEquip, setCurrEquip] = useState('none');
  const [userName, setUserName] = useState('');
  //const [tabIndex, setTabIndex] = useState(0);
  const setCurrEquip = useSetCurrEquip();

  const token = usersState.token;
  const handleEquipsChange = async (event) => {
    const select = event.target;
    const val = select.value;// select.options[select.selectedIndex].value;

    await onEquipChanged(val);
  };

  const onEquipChanged = async equipInfo =>
  {
    setCurrEquip(equipInfo, 'SETEQUIPINFO');
    getEvents(equipInfo);
  }

  const getEvents = async (equipInfo) =>
  {
    const endDate = new Date();
    const allEvents = await EquipWorker.SearchEquip(token, 'Events', equipInfo, getUSFullDate(endDate), getUSFullDate(endDate));
    eventsDispatch({ type: 'SETEVENTS', payload: allEvents }); 
  }

  const getLogs = async () =>
  {
    const logs = await AdminWorker.GetServerLogs(token);
    communicationDispatch({ type: 'SETLOGS', payload: logs }); 
  }

  const getChats = async () =>
  {
    const notes = await EquipWorker.GetCommunications(token, equipInfo);
    communicationDispatch({ type: 'SETCHATS', payload: notes }); 
  }

  useEffect(() => {
    (async () => {
      setUserName(usersState?.currentUser?.Surname);
    })();
  }, [usersState.currentUser]);


  const webSocket = useWebSocket(
    {
    }
  );

  const equipInfo = currEquipState.equipInfo;
  const selectedTab = appState.currentTab?.tab ?? SummaryTabIndex;
  const selectedTabPanel = appState.currentTab?.panel ?? MainTabPanelIndex;

  const onTabIndexChange = async (event, newValue) => {
    if(SummaryTabIndex === selectedTab && SummaryDBTabPanelIndex === newValue && equipInfo){
      const allTables = await EquipWorker.GetAllTables(token, equipInfo);
      currEquipDispatch({ type: 'SETALLDBTABLES', payload: allTables }); 
    }
    else if(SummaryTabIndex === selectedTab && SummaryHistoryTabPanelIndex === newValue){        
      getEvents(equipInfo);
    }
    else if(SummaryTabIndex === selectedTab && SummaryChatTabPanelIndex === newValue){        
      getChats(equipInfo);
    }
    else if(AdminTabIndex === selectedTab && AdminLogTabPanelIndex === newValue){        
      getLogs();
    }

    appDispatch({ type: 'SETTAB', payload: {tab: selectedTab, panel: newValue} }); 
  };

  const onUpdateDBInfo = async () =>{
    if(equipInfo){
      const res = await EquipWorker.UpdateDBInfo(token, equipInfo);
    }
  }

  const isDBInfoStateUpdating = currEquipState.remoteaccess?.DBInfoStateUpdating;
  return (    
    <AppBar position="fixed" className={classes.appBar}>
        <Toolbar>
            <Typography variant="h6" noWrap>
            </Typography>
            <FormControl className={classes.formControl}>
              <Select
                labelId="demo-simple-select-label"
                id="mainToolbarCombobox"
                value={currEquipState.equipInfo}
                onChange={handleEquipsChange}
                className={classes.selectEmpty}
                variant="outlined"
              >
                <ListSubheader className={classes.optionStyle}>Выбрано</ListSubheader>
                {allEquipsState.selectedEquips?.map((i, ind) => (
                    <MenuItem key={ind.toString()} value={i} className={classes.optionStyle}>
                      <ListItemIcon>
                        {/* <LocationOffOutlinedIcon fontSize="small" /> */}
                        <LocationOnOutlinedIcon fontSize="large" style={{ color: 'white' }}/>
                      </ListItemIcon>
                      <Typography variant="inherit">{i}</Typography>                      
                    </MenuItem>
                    ))
                }
                <ListSubheader className={classes.optionStyle}>Активно</ListSubheader>
                {allEquipsState.connectedEquips?.map((i, ind) => (
                    <MenuItem key={ind.toString()} value={i} className={classes.optionStyle}>
                      <ListItemIcon>
                        {/* <LocationOffOutlinedIcon fontSize="small" /> */}
                        <LocationOnOutlinedIcon fontSize="large" style={{ color: 'white' }}/>
                      </ListItemIcon>
                      <Typography variant="inherit">{i}</Typography>                      
                    </MenuItem>
                    ))
                }
              </Select>
            </FormControl>
            <Tabs value={selectedTabPanel} onChange={onTabIndexChange} aria-label="simple tabs example" className={classes.tabControl}>
              <Tab label="Главная" id= "mainTabPanel" />
              {selectedTab === SummaryTabIndex?
                  <Tab label="БД" id= "dbTabPanel" /> : <></>
              }
              {selectedTab === SummaryTabIndex?
                  <Tab label="История" id= "histTabPanel" /> : <></>
              }
              {selectedTab === SummaryTabIndex?
                  <Tab label="Коммуникации" id= "chatTabPanel" /> : <></>
              }
              {selectedTab === AdminTabIndex?
                  <Tab label="Логи" id= "logsTabPanel" /> : <></>              
              }              
            </Tabs>
            {selectedTab === SummaryTabIndex?
                <Button variant="contained" 
                    color={isDBInfoStateUpdating ? "secondary" : "primary"}
                    className={classes.button} 
                    onClick={onUpdateDBInfo} 
                >
                  {isDBInfoStateUpdating ? 'Обновляется' : 'Обновить'}
                </Button> : <></>
            }
            <Typography variant="h6" noWrap align="right"  className={classes.userName}> 
              {userName}
            </Typography>
        </Toolbar>
    </AppBar>
  );
}