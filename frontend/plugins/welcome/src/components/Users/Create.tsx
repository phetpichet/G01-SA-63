import React, { useEffect, FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import Grid from '@material-ui/core/Grid';
import { DefaultApi } from '../../api/apis';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Select from '@material-ui/core/Select';
import { IconButton, InputAdornment, InputLabel, MenuItem, OutlinedInput } from '@material-ui/core';
import { EntTitle } from '../../api/models/EntTitle';
import { EntPosition } from '../../api/models/EntPosition';
import { EntGender } from '../../api/models/EntGender';
import Swal from 'sweetalert2';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    formControl: {
      margin: theme.spacing(1),
      minWidth: 300,
      maxWidth: 300,
      marginTop: '1.5%',
    },
    withoutLabel: {
      marginTop: theme.spacing(4),
    },
    title: {
      flexGrow: 1,
    },

  }),
);
interface user {
  title: number;
  gender: number;
  position: number;
  name: string;
  email: string;
  password: string;
  showPassword: boolean;
}

const User: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();
  const [user, setUser] = React.useState<Partial<user>>({});
  const [titles, setTitles] = React.useState<EntTitle[]>([]);
  const [positions, setPositions] = React.useState<EntPosition[]>([]);
  const [genders, setGenders] = React.useState<EntGender[]>([]);
 

  const Toast = Swal.mixin({
    position: 'center',
    showConfirmButton: true,
    timer: 1500
  });

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,) => {
    const name = event.target.name as keyof typeof User;
    const { value } = event.target;
    setUser({ ...user, [name]: value });
    console.log(user);
  };

  const getTitles = async () => {
    const res = await http.listTitle({ limit: 10, offset: 0 });
    setTitles(res);
  };
  const getGenders = async () => {
    const res = await http.listGender({ limit: 10, offset: 0 });
    setGenders(res);
  };
  const getPositions = async () => {
    const res = await http.listPosition({ limit: 10, offset: 0 });
    setPositions(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getGenders();
    getPositions();
    getTitles();
  }, []);

  function clear() {
    setUser({});
  }

  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/users';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(user),
    };
    console.log(user); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          Toast.fire({
            title: 'บันทึกข้อมูลสำเร็จ',
            icon: 'success',
          });
        } else {
          Toast.fire({
            title: 'บันทึกข้อมูลไม่สำเร็จ',
            icon: 'error',
          });
        }
      });
  }

  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" className={classes.title}>
            ระบบบันทึกข้อมูลเภสัชกร
            </Typography>
          <Button color="inherit" component={RouterLink} to="/"> Login </Button>
        </Toolbar>
      </AppBar>
      <Content className={classes.withoutLabel}>
        <Typography variant="h5"
          className={classes.formControl} style={{
            marginLeft: 200
          }}>
          บันทึกข้อมูลเภสัชกร
        </Typography>
        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <Grid container spacing={3}>
              <Grid item xs={6}>
                <FormControl variant="outlined" className={classes.formControl}>
                  <InputLabel>คำนำหน้า</InputLabel>
                  <Select
                    name="title"
                    label="คำนำหน้า"
                    value={user.title || ''}
                    onChange={handleChange}
                  >
                    {titles.map(item => {
                      return (
                        <MenuItem key={item.id} value={item.id}>{item.title}</MenuItem>
                      );
                    })}
                  </Select>
                </FormControl>
              </Grid>
              <Grid item xs={6}>
                <FormControl
                  fullWidth
                  className={classes.formControl}
                  variant="outlined">
                  <TextField
                    name="name"
                    label="ชื่อ-นามสกุล"
                    variant="outlined"
                    size="medium"
                    value={user.name || ''}
                    onChange={handleChange}
                  />
                </FormControl>
              </Grid>
            </Grid>
            <Grid container spacing={3}>
              <Grid item xs={6}>
                <FormControl variant="outlined" className={classes.formControl}>
                  <InputLabel>ตำแหน่ง</InputLabel>
                  <Select
                    name="position"
                    type="string"
                    label="ตำแหน่ง"
                    value={user.position || ''}
                    onChange={handleChange}
                  >
                    {positions.map(item => {
                      return (
                        <MenuItem key={item.id} value={item.id}>{item.position}</MenuItem>
                      );
                    })}

                  </Select>
                </FormControl>
              </Grid>
              <Grid item xs={6}>
                <FormControl variant="outlined" className={classes.formControl}>
                  <InputLabel>เพศ</InputLabel>
                  <Select
                    name="gender"
                    label="เพศ"
                    type="string"
                    value={user.gender || ''}
                    onChange={handleChange}
                  >
                    {genders.map(item => {
                      return (
                        <MenuItem key={item.id} value={item.id}>{item.gender}</MenuItem>
                      );
                    })}
                  </Select>
                </FormControl>
              </Grid>
            </Grid>
            <Grid container spacing={3}>
              <Grid item xs={6}>
                <FormControl
                  fullWidth
                  className={classes.formControl}
                  variant="outlined">
                  <TextField
                    name="email"
                    label="Email"
                    variant="outlined"
                    type="string"
                    size="medium"
                    value={user.email || ''}
                    onChange={handleChange}
                  />
                </FormControl>
              </Grid>
              <Grid item xs={6}>
                <FormControl className={classes.formControl} variant="outlined">
                  <InputLabel>Password</InputLabel>
                  <OutlinedInput
                    labelWidth={80}
                    name="password"
                    type={user.showPassword ? 'text' : 'password'}
                    value={user.password}
                    onChange={handleChange}
                  />
                </FormControl>
              </Grid>
            </Grid>
            <div className={classes.formControl} style={{ marginLeft: 200 }}>
              <Button
                onClick={save}
                variant="contained"
                color="primary">
                SAVE
             </Button>
              <Button style={{ marginLeft: 10 }}
                onClick={clear}
                variant="contained"
                color="secondary">
                CLEAR
             </Button>
             <Button style={{ marginLeft: 10 }}
                component={RouterLink} to="/table"
                variant="contained"
                color="secondary">  
                BACK
             </Button>
            </div>
          </form>
        </div>
      </Content>
    </div>
  );
};
export default User;