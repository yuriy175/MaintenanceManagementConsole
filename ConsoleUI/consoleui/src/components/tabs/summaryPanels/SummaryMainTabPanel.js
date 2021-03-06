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
        <SystemCard system={currEquipState.system} volatile={systemVolatileState.currentVolatile}></SystemCard>
        {/* <HddCard></HddCard> */}
      </div>
      <div className={classes.column}>
        <OrganAutoCard organAuto={currEquipState.organAuto}></OrganAutoCard>
        <ImagesCard images={currEquipState.images}></ImagesCard>
        <GeneratorCard generator={currEquipState.generator}></GeneratorCard>
        <DetectorCard detectors={currEquipState.detectors} aecs={currEquipState.aecs}></DetectorCard>
        <StandCard stand={currEquipState.stand}></StandCard>
        <DosimeterCard dosimeter={currEquipState.dosimeter}></DosimeterCard>
      </div>
      <div className={classes.column}>
        <RemoteAccessCard equipInfo={currEquipState.equipInfo} remoteaccess={currEquipState.remoteaccess} token={usersState.token}></RemoteAccessCard>
        <DicomCard dicom={currEquipState.dicom}></DicomCard>
        <SoftwareCard software={currEquipState.software} volatile={systemVolatileState.currentVolatile}></SoftwareCard>
      </div>
      {notifyText ? <NotifyDlg title='Данные FTP' text={'Данные посланы ' + (currEquipState.remoteaccess?.FtpSendResult ? 'успешно' : 'с ошибками') }></NotifyDlg> : <></>}
    </div>
  );
}