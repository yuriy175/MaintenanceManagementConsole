import React, {useEffect} from 'react';
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import TextField from '@material-ui/core/TextField';
import Divider from '@material-ui/core/Divider';


import * as ExternWorker from '../../workers/externWorker'

export default function EquipDetailsDlg(props){
  const equipName = props.equip?.EquipName;
  //const equipAlias = props.equip?.EquipAlias;
  //const hospitalLatitude = props.equip?.HospitalLatitude;
  //const hospitalLongitude = props.equip?.HospitalLongitude;
  //const hospName = props.equip?.HospitalName;
  //const hospAddress = props.equip?.HospitalAddress;
  //const hospitalZones = props.equip?.HospitalZones;  

  const [hospName, setHospName] = React.useState('');
  const [equipAlias, setEquipAlias] = React.useState('');
  const [hospitalLatitude, setHospitalLatitude] = React.useState('');
  const [hospitalLongitude, setHospitalLongitude] = React.useState('');
  const [hospAddress, setHospAddress] = React.useState('');
  const [hospitalZones, setHospitalZones] = React.useState('');

  const [city, setCity] = React.useState('Санкт-Петербург');
  const [street, setStreet] = React.useState('проспект Юрия Гагарина');
  const [house, setHouse] = React.useState('32');


  useEffect(() => {
    const equip = props.equip;
    setHospName(equip?.HospitalName);
    setEquipAlias(equip?.EquipAlias);
    setHospitalLatitude(equip?.HospitalLatitude);
    setHospitalLongitude(equip?.HospitalLongitude);
    setHospAddress(equip?.HospitalAddress);
    setHospitalZones(equip?.HospitalZones);
  }, [props.equip]);

  const handleClose = () => {
    props?.onClose(false);
  };

  const handleCloseOK = () => {
    props?.onClose(true, 
      {
        EquipName: equipName,
        EquipAlias: equipAlias,
        HospitalLatitude: hospitalLatitude,
        HospitalLongitude: hospitalLongitude,
        HospitalName: hospName,
        HospitalAddress: hospAddress,
        HospitalZones: hospitalZones,
      });
  };

  const handleCoordsQuery = async () => {
    //nominatim.openstreetmap.org/search.php?street=проспект+Юрия+Гагарина+32+к6&city=Санкт-Петербург&format=jsonv2

    const coords = await ExternWorker.GetEquipCoords(city, street, house);
    setHospitalLatitude(coords?.[0]?.lat);
    setHospitalLongitude(coords?.[0]?.lon);
  };  

  const onEquipAliasChange = async (val) =>{
    setEquipAlias(val.target?.value);
  }
  
  const onHospitalLongitudeChange = async (val) =>{
    setHospitalLongitude(val.target?.value);
  }

  const onHospitalLatitudeChange = async (val) =>{
    setHospitalLatitude(val.target?.value);
  }

  const onHospitalZonesChange = async (val) =>{
    setHospitalZones(val.target?.value);
  }

  const onHospNameChange = async (val) =>{
    setHospName(val.target?.value);
  }

  const onHospAddressChange = async (val) =>{
    setHospAddress(val.target?.value);
  }  

  const onCityChange = async (val) =>{
    setCity(val.target?.value);
  }  

  const onStreetChange = async (val) =>{
    setStreet(val.target?.value);
  }  

  const onHouseChange = async (val) =>{
    setHouse(val.target?.value);
  }  

  return (
    <Dialog
      open={props.open}
      onClose={handleClose}
      aria-labelledby="alert-dialog-title"
      aria-describedby="alert-dialog-description"
    >
      <DialogTitle id="alert-dialog-title">Карточка оборудования ({equipName})</DialogTitle>
      <DialogContent>
      <TextField
            autoFocus
            margin="dense"
            id="alias"
            label="Алиас"
            fullWidth
            variant="standard"
            value={equipAlias}
            onChange={onEquipAliasChange}
          />

        <TextField
            autoFocus
            margin="dense"
            id="alias"
            label="Название ЛПУ"
            fullWidth
            variant="standard"
            value={hospName}
            onChange={onHospNameChange}
          />

        <TextField
            autoFocus
            margin="dense"
            id="alias"
            label="Адрес ЛПУ"
            fullWidth
            variant="standard"
            value={hospAddress}
            onChange={onHospAddressChange}
          />
        <TextField
            autoFocus
            margin="dense"
            id="longitude"
            label="Зоны"
            fullWidth
            variant="standard"
            value={hospitalZones}
            onChange={onHospitalZonesChange}
          />

        <TextField
            autoFocus
            margin="dense"
            id="latitude"
            label="Широта"
            fullWidth
            variant="standard"
            value={hospitalLatitude}
            onChange={onHospitalLatitudeChange}
          />

          <TextField
            autoFocus
            margin="dense"
            id="longitude"
            label="Долгота"
            fullWidth
            variant="standard"
            value={hospitalLongitude}
            onChange={onHospitalLongitudeChange}
          />
          <TextField
            autoFocus
            margin="dense"
            id="city"
            label="Город"
            variant="standard"
            value={city}
            fullWidth onChange={onCityChange}
          />
          <TextField
            autoFocus
            margin="dense"
            id="street"
            label="Улица"
            variant="standard"
            value={street}
            fullWidth onChange={onStreetChange}
          />
          <TextField
            autoFocus
            margin="dense"
            id="house"
            label="Дом"
            variant="standard"
            value={house}
            fullWidth onChange={onHouseChange}
          />
          <Button onClick={handleCoordsQuery}>
            Запросить
          </Button>
        
      </DialogContent>
      <DialogActions>
        <Button onClick={handleCloseOK} autoFocus>
          Да
        </Button>
        <Button onClick={handleClose} >
          Нет
        </Button>
      </DialogActions>
    </Dialog>
  );
}