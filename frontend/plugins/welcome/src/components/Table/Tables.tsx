import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntUser } from '../../api/models/EntUser';
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const http = new DefaultApi();
 const [users, setUsers] = useState<EntUser[]>([]);
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
   const createUsers = async () => {
     const res = await http.listUser({ limit: 10, offset: 0 });
     setLoading(false);
     setUsers(res);
   };
   createUsers();
 }, [loading]);
 
 const deleteUsers = async (id: number) => {
   const res = await http.deleteUser({ id: id });
   setLoading(true);
 };
 
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No.</TableCell>
           <TableCell align="center">Name</TableCell>
           <TableCell align="center">Email</TableCell>
           <TableCell align="center">Password</TableCell>
           {/* <TableCell align="center">Title</TableCell> */}
           <TableCell align="center">Manage</TableCell>

         </TableRow>
       </TableHead>
       <TableBody>
         {users.map(item => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             {/* <TableCell align="center">{item.title}</TableCell> */}
             <TableCell align="center">{item.name}</TableCell>
             <TableCell align="center">{item.email}</TableCell>
             <TableCell align="center">{item.password}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                   deleteUsers(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 Delete
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}