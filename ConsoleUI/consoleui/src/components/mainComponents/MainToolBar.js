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
  CommonChat,
  AdminTabIndex, AdminLogTabPanelIndex,
  ControlTabIndex, ControlDiagnosticTabPanelIndex } from '../../model/constants';

import { AppContext } from '../../context/app-context';
import { AllEquipsContext } from '../../context/allEquips-context';
import { EventsContext } from '../../context/events-context';
import { CurrentEquipContext } from '../../context/currentEquip-context';
import { UsersContext } from '../../context/users-context';
import { CommunicationContext } from '../../context/communication-context';
import { ControlStateContext } from '../../context/controlState-context';
import {useSetCurrEquip} from '../../hooks/useSetCurrEquip'

import * as EquipWorker from '../../workers/equipWorker'
import * as AdminWorker from '../../workers/adminWorker'
import * as ControlWorker from '../../workers/controlWorker'
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
  const [controlState, controlDispatch] = useContext(ControlStateContext);  
  
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

  
  const selectedTab = appState.currentTab?.tab ?? SummaryTabIndex;
  const selectedTabPanel = appState.currentTab?.panel ?? MainTabPanelIndex;

  const onEquipChanged = async equipInfo =>
  {
    setCurrEquip(equipInfo, 'SETEQUIPINFO');
    allEquipsDispatch({ type: 'ADDSELECTEDEQUIPS', payload: equipInfo }); 
    //getEvents(equipInfo);
    
    if(SummaryTabIndex === selectedTab && SummaryHistoryTabPanelIndex === selectedTabPanel){        
      getEvents(equipInfo);
    }
    else if(SummaryTabIndex === selectedTab && SummaryChatTabPanelIndex === selectedTabPanel){        
      getChats(equipInfo);
    }
    else if(SummaryTabIndex === selectedTab && SummaryDBTabPanelIndex === selectedTabPanel && equipInfo){
      getAllDBTables(equipInfo);
    }
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

  const getDiagnostics = async () =>
  {
    const metrics = await ControlWorker.GetServerMetrics(token);
    controlDispatch({ type: 'SETDIAGNOSTIC', payload: metrics }); 
  }

  const getAllDBTables = async (equipInfo) =>
  {
    const allTables = await EquipWorker.GetAllTables(token, equipInfo);
    currEquipDispatch({ type: 'SETALLDBTABLES', payload: allTables }); 
  }

  const getChats = async (equipInfo) =>
  {
    const notes = await EquipWorker.GetCommunications(token, equipInfo);
    communicationDispatch({ 
      type: 'SETCHATS', 
      payload: notes 
    }); 
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

  const onTabIndexChange = async (event, newValue) => {
    if(SummaryTabIndex === selectedTab && SummaryDBTabPanelIndex === newValue && equipInfo){
      getAllDBTables(equipInfo);
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
    else if(ControlTabIndex === selectedTab && ControlDiagnosticTabPanelIndex === newValue){        
      getDiagnostics();
    }

    appDispatch({ type: 'SETTAB', payload: {tab: selectedTab, panel: newValue} }); 
  };

  const onUpdateDBInfo = async () =>{
    if(equipInfo){
      const res = await EquipWorker.UpdateDBInfo(token, equipInfo);
    }
  }

  const isDBInfoStateUpdating = currEquipState.remoteaccess?.DBInfoStateUpdating;
  const isValidSummaryTab = selectedTab === SummaryTabIndex && equipInfo;
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
              <Tab label="Главная" id= "mainTabPanel" key="mainTabPanel"/>
              {isValidSummaryTab?
                  <Tab label="БД" id= "dbTabPanel" key="dbTabPanel"/> : <div key="dbTabPanel"></div>
              }
              {isValidSummaryTab?
                  <Tab label="История" id= "histTabPanel" key="histTabPanel"/> : <div></div>
              }
              {isValidSummaryTab?
                  <Tab label="Коммуникации" id= "chatTabPanel" key="chatTabPanel"/> : <div></div>
              }
              {selectedTab === AdminTabIndex?
                  <Tab label="Логи" id= "logsTabPanel" key="logsTabPanel"/> : <div></div>              
              }            
              {selectedTab === ControlTabIndex?
                  <Tab label="Диагностика" id= "diagnosticTabPanel" key="diagnosticTabPanel"/> : <div></div>              
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