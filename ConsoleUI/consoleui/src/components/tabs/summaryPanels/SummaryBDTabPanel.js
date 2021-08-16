import React, {useState, useEffect, useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import InboxIcon from '@material-ui/icons/MoveToInbox';
import MailIcon from '@material-ui/icons/Mail';
import Drawer from '@material-ui/core/Drawer';
import Divider from '@material-ui/core/Divider';
import Box from '@material-ui/core/Box';
import Button from '@material-ui/core/Button';

import * as EquipWorker from '../../../workers/equipWorker'
import { CurrentEquipContext } from '../../../context/currentEquip-context';
import { UsersContext } from '../../../context/users-context';
import CommonTable from '../../tables/CommonTable'

const useStyles = makeStyles((theme) => ({
  root: {
    display:"table"
  },
  column:{
    width: "30%",
    marginRight: "12px",
  },
  listPanel:{
    width: "285px",
    // backgroundColor: 'white',
    backgroundColor: theme.palette.background.paper,
    // maxHeight: "900px",
    overflowY: 'auto',
    display: 'table-cell',
  },
  content: {
    // marginTop: 64, // `${drawerWidth}px`,
    // flexGrow: 1,
    width: '100%', 
    backgroundColor: theme.palette.background.default,
    paddingLeft: theme.spacing(1),
    display: 'table-cell',
  },
  button:{
    marginBottom: '1em',
    marginTop: '1em',
  }
}));

export default function SummaryBDTabPanel(props) {
  console.log("render SummaryBDTabPanel");

  const classes = useStyles();
  const [tableContent, setTableContent] = React.useState('');
  const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);
  const [usersState, usersDispatch] = useContext(UsersContext);

  const allDBs = currEquipState.allDBs;
  const allDBTables = currEquipState.allDBTables;

  useEffect(() => {
    if(Object.keys(currEquipState.allDBTables).length === 0){
      setTableContent('');
    }  
}, [currEquipState.allDBTables]);

  const getColumn = (key) => { 
    if(key.toLowerCase() === 'active')
    {
      return { 
        id: key, label: key, minWidth: 100, maxWidth: 100,
        format: (value) => value ? '+' : '-'
      }
    }

    return { id: key, label: key, minWidth: 100, maxWidth: 300 }
  }

  const token = usersState.token;
  const equipInfo = currEquipState.equipInfo;
  const handleListItemClick = async (event, index, type, text) => {
    const content = await EquipWorker.GetTableContent(token, equipInfo, type, text);
    let values = []
    if(Array.isArray(content)){
      values = content.map(c => JSON.parse(c)).flat(1);
    }
    else{
      values = content? JSON.parse(content) : null;
    }
    
    setTableContent(values);
  };

  const isRowBold = (row) =>
  {
    return row.Active
  }

  const onRecreate = async () => {
    const content = await EquipWorker.RecreateDBInfo(token, equipInfo);    
  }; 
    
  const columns = tableContent.length === 0 ? [] : Object.keys(tableContent[0]).map(k => getColumn(k));
  const rows = tableContent.length === 0 ? [] : tableContent;

  const hospTableMenu = allDBTables?.Hospital;
  const systemTableMenu = allDBTables?.System;
  const softwareTableMenu = allDBTables?.Software;
  const atlasTableMenu = allDBTables?.Atlas;
  return (
    <div className={classes.root}>
      {/* <div className={classes.listPanel}> */}
      {/* <Box className={classes.listPanel} height={'calc(100% - 500px)'}> */}
      <Box className={classes.listPanel} height={'100%'}>
        <Button className={classes.button} variant="contained" color="primary" onClick={onRecreate}>
              Пересоздать
        </Button>
        <Typography variant="h6" component="h2">ЛПУ</Typography>
        <List>
          {hospTableMenu?.map((text, index) => (
            <ListItem button key={text}
                // selected={selectedTab === index}
                onClick={(event) => handleListItemClick(event, index, "Hospital", text)}
            >
              <ListItemIcon>{index % 2 === 0 ? <InboxIcon /> : <MailIcon />}</ListItemIcon>
              <ListItemText primary={text} />
            </ListItem>
          ))}
        </List>
        <Typography variant="h6" component="h2">Система</Typography>
        <List>
          {systemTableMenu?.map((text, index) => (
            <ListItem button key={text}
                // selected={selectedTab === index}
                onClick={(event) => handleListItemClick(event, index, "System", text)}
            >
              <ListItemIcon>{index % 2 === 0 ? <InboxIcon /> : <MailIcon />}</ListItemIcon>
              <ListItemText primary={text} />
            </ListItem>
          ))}
        </List>
        <Typography variant="h6" component="h2">Общее ПО</Typography>
        <List>
          {softwareTableMenu?.map((text, index) => (
            <ListItem button key={text}
                // selected={selectedTab === index}
                onClick={(event) => handleListItemClick(event, index, "Software", text)}
            >
              <ListItemIcon>{index % 2 === 0 ? <InboxIcon /> : <MailIcon />}</ListItemIcon>
              <ListItemText primary={text} />
            </ListItem>
          ))}
        </List>
        <Typography variant="h6" component="h2">Атлас</Typography>
        <List>
          {atlasTableMenu?.map((text, index) => (
            <ListItem button key={text}
                // selected={selectedTab === index}
                onClick={(event) => handleListItemClick(event, index, "Atlas", text)}
            >
              <ListItemIcon>{index % 2 === 0 ? <InboxIcon /> : <MailIcon />}</ListItemIcon>
              <ListItemText primary={text} />
            </ListItem>
          ))}
        </List>
      {/* </div> */}
      </Box>
      <Box className={classes.content} height="100%">
        {/* <main className={classes.content} >
          <>   */}
                    <CommonTable 
                      columns={columns} 
                      rows={rows} 
                      isRowBold = {isRowBold}></CommonTable>
          {/* </>
        </main> */}
      </Box>
    </div>
  );
}