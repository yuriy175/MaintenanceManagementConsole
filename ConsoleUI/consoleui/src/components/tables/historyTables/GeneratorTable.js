import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TablePagination from '@material-ui/core/TablePagination';
import TableRow from '@material-ui/core/TableRow';

const columns = [
  { id: 'EquipName', label: 'Комплекс', minWidth: 170 },
  { id: 'State', label: 'Состояние', minWidth: 100 },
  { id: 'DateTime', label: 'Время', minWidth: 100 },
  { id: 'Workstation', label: 'Раб. место', minWidth: 100 },
  { id: 'HeatStatus', label: 'Нагрев', minWidth: 100 },
  { id: 'ErrorDescriptions', label: 'Ошибки', minWidth: 100, 
    formatArray: (values) =>
    {
      if(!values || !values.length) {
        return values;
      }
      //const errors = values.reduce((accumulator, currentValue) => accumulator + `Code ${currentValue.Code}: ${currentValue. Description}`, '');
      return values.map((currentValue) => `Code ${currentValue.Code}: ${currentValue. Description}`);
    },
    hasErrors: (values) =>
    {
      if(!values || !values.length) {
        return false;
      }

      return true;
    }
  },
  { id: 'Mas', label: 'Ток', minWidth: 100 },
  { id: 'Kv', label: 'Напряжение', minWidth: 100 },  
];

const useStyles = makeStyles({
  root: {
    width: '100%',
  },
  container: {
    maxHeight: 440,
  },
  errorCell:{
    color: 'white',
    background: 'red',
  }
});

export default function GeneratorTable(props) {
  console.log("render GeneratorTable");

  const classes = useStyles();
  const [page, setPage] = React.useState(0);
  const [rowsPerPage, setRowsPerPage] = React.useState(10);

  const handleChangePage = (event, newPage) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (event) => {
    setRowsPerPage(+event.target.value);
    setPage(0);
  };

  const rows = props.data;
  return (
    <Paper className={classes.root}>
      <TableContainer className={classes.container}>
        <Table stickyHeader aria-label="sticky table">
          <TableHead>
            <TableRow>
              {columns.map((column) => (
                <TableCell
                  key={column.id}
                  align={column.align}
                  style={{ minWidth: column.minWidth }}
                >
                  {column.label}
                </TableCell>
              ))}
            </TableRow>
          </TableHead>
          <TableBody>
            {rows.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage).map((row) => {
              return (
                <TableRow hover role="checkbox" tabIndex={-1} key={row.code}>
                  {columns.map((column) => {
                    const value = row[column.id];
                    return (
                      <TableCell key={column.id} align={column.align} 
                        className={column.hasErrors && column.hasErrors(value) ? classes.errorCell : ''}>
                        {column.formatArray ? column.formatArray(value).map(v => <pre>{v}</pre>) : value}
                      </TableCell>
                    );
                  })}
                </TableRow>
              );
            })} 
          </TableBody>
        </Table>
      </TableContainer>
      <TablePagination
        rowsPerPageOptions={[10, 25, 100]}
        component="div"
        count={rows?.length}
        rowsPerPage={rowsPerPage}
        page={page}
        onChangePage={handleChangePage}
        onChangeRowsPerPage={handleChangeRowsPerPage}
      />
    </Paper>
  );
}