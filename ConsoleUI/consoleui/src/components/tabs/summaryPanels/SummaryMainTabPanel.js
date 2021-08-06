import React, {useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import { CurrentEquipContext } from '../../../context/currentEquip-context';
import { AllEquipsContext } from '../../../context/allEquips-context';
import { SystemVolatileContext } from '../../../context/systemVolatile-context';
import { UsersContext } from '../../../context/users-context';

import DetectorCard from '../../cards/DetectorCard'
import GeneratorCard from '../../cards/GeneratorCard'
import SystemCard from '../../cards/SystemCard'
import OrganAutoCard from '../../cards/OrganAutoCard'
import EquipImageCard from '../../cards/EquipImageCard'
import DicomCard from '../../cards/DicomCard'
import RemoteAccessCard from '../../cards/RemoteAccessCard'
import StandCard from '../../cards/StandCard'
import DosimeterCard from '../../cards/DosimeterCard'
import SoftwareCard from '../../cards/SoftwareCard'
import ImagesCard from '../../cards/ImagesCard'
import NotifyDlg from '../../dialogs/NotifyDlg'

const useStyles = makeStyles((theme) => ({
  root: {
    display:"flex"
  },
  column:{
    width: "30%",
    marginRight: "12px",
  }
}));

export default function SummaryMainTabPanel(props) {
  console.log("render SummaryMainTabPanel");

  const classes = useStyles();
  const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);
  const [allEquipsState, allEquipsDispatch] = useContext(AllEquipsContext);
  const [systemVolatileState, systemVolatileDispatch] = useContext(SystemVolatileContext);
  const [usersState, usersDispatch] = useContext(UsersContext);

  const notifyText = currEquipState.remoteaccess?.FtpSendResult;
  const equipInfo = currEquipState.equipInfo;
  const isEquipConnected = allEquipsState.connectedEquips?.includes(equipInfo);
  const lastSeen = currEquipState.lastSeen;
  const hospital = currEquipState.locationInfo?.HospitalName;
  const address = currEquipState.locationInfo?.HospitalAddress;
  return (
    <div className={classes.root}>
      <div className={classes.column}>
        <EquipImageCard 
          equipInfo={equipInfo} 
          isConnected={isEquipConnected} 
          lastSeen={lastSeen}
          hospital={hospital}
          address={address}
        ></EquipImageCard>
        <SystemCard 
          system={currEquipState.system} 
          volatile={isEquipConnected ? systemVolatileState.currentVolatile : null}></SystemCard>
        
      </div>
      <div className={classes.column}>
        <OrganAutoCard organAuto={isEquipConnected ? currEquipState.organAuto : null}></OrganAutoCard>
        <ImagesCard images={isEquipConnected ? currEquipState.images : null}></ImagesCard>
        <GeneratorCard generator={isEquipConnected ? currEquipState.generator : null}></GeneratorCard>
        <DetectorCard detectors={isEquipConnected ? currEquipState.detectors : null} aecs={currEquipState.aecs}></DetectorCard>
        <StandCard stand={isEquipConnected ? currEquipState.stand : null}></StandCard>
        <DosimeterCard dosimeter={isEquipConnected ? currEquipState.dosimeter : null}></DosimeterCard>
      </div>
      <div className={classes.column}>
        <RemoteAccessCard equipInfo={isEquipConnected ? currEquipState.equipInfo : null} remoteaccess={currEquipState.remoteaccess} token={usersState.token}></RemoteAccessCard>
        <DicomCard dicom={isEquipConnected ? currEquipState.dicom : null}></DicomCard>
        <SoftwareCard 
          software={isEquipConnected ? currEquipState.software : null} 
          volatile={isEquipConnected ? systemVolatileState.currentVolatile : null}></SoftwareCard>
      </div>
      {notifyText ? <NotifyDlg title='Данные FTP' text={'Данные посланы ' + (currEquipState.remoteaccess?.FtpSendResult ? 'успешно' : 'с ошибками') }></NotifyDlg> : <></>}
    </div>
  );
}