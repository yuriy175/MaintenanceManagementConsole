import React, {useEffect} from 'react';
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import TextField from '@material-ui/core/TextField';

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
    props?.onClose(true, props.context);
  };

  const handleCoordsQuery = async () => {
    //nominatim.openstreetmap.org/search.php?street=проспект+Юрия+Гагарина+32+к6&city=Санкт-Петербург&format=jsonv2

    const coords = await ExternWorker.GetEquipCoords('Санкт-Петербург', 'проспект Юрия Гагарина', '32');
    setHospitalLatitude(coords?.[0]?.lat);
    setHospitalLongitude(coords?.[0]?.lon);
  };  

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
          />

        <TextField
            autoFocus
            margin="dense"
            id="alias"
            label="Название ЛПУ"
            fullWidth
            variant="standard"
            value={hospName}
          />

        <TextField
            autoFocus
            margin="dense"
            id="alias"
            label="Адрес ЛПУ"
            fullWidth
            variant="standard"
            value={hospAddress}
          />

        <TextField
            autoFocus
            margin="dense"
            id="longitude"
            label="Зоны"
            fullWidth
            variant="standard"
            value={hospitalZones}
          />

        <TextField
            autoFocus
            margin="dense"
            id="latitude"
            label="Широта"
            fullWidth
            variant="standard"
            value={hospitalLatitude}
          />

          <TextField
            autoFocus
            margin="dense"
            id="longitude"
            label="Долгота"
            fullWidth
            variant="standard"
            value={hospitalLongitude}
          />
          
          <TextField
            autoFocus
            margin="dense"
            id="longitude"
            label="Запрос координат"
            fullWidth
            variant="standard"
            value={hospitalLongitude}
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