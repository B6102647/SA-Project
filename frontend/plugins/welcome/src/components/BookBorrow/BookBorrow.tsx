import React,{ useState, useEffect }  from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,Header,Page,pageTheme,ContentHeader,
} from '@backstage/core';
import {
  Table,TableBody,TableCell,TableRow,Typography,
  TextField,Button,withStyles,makeStyles,
  Theme,FormControl,InputLabel,MenuItem,
  FormHelperText,Select,createStyles,
} from '@material-ui/core';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import { 
    EntBook,
    EntPurpose,
    EntUser,
} from '../../api/models/';
const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(3),
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
  }),
);
const username = { givenuser: 'Manuschanok Srikhrueadong' };
export default function Create() {
  const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบยืมหนังสือ' };
  const classes = useStyles();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [books, setBooks] = useState<EntBook[]>(Array);
  const [users, setUsers] = useState<EntUser[]>(Array);
  const [purposes, setPurposes] = useState<EntPurpose[]>(Array);
useEffect(() => {
    const getUser = async () => {
        const res = await api.listUser({ limit: 10, offset: 0 });
        setLoading(false);
        setUsers(res);
    };
    getUser();

    const getBooks = async () => {
        const res = await api.listBook({ limit: 10, offset: 0 });
        setLoading(false);
        setBooks(res);
    };
    getBooks();

    const getPurposes = async () => {
        const res = await api.listPurpose({ limit: 10, offset: 0 });
        setLoading(false);
        setPurposes(res);
    };
    getPurposes();
}, [loading]);

const AddedhandleChange = (event: any) => {
    setAdded(event.target.value as string);
  };
  const PurposeIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPurposeID(event.target.value as number);
  };
  const bookIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setBookID(event.target.value as number);
  };
  const UserIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUserID(event.target.value as number);
  };

  const [added, setAdded] = useState(String);
  const [bookID, setBookID] = useState(Number);
  const [purposeID, setPurposeID] = useState(Number);
  const [userID, setUserID] = useState(Number);
const createBookBorrow = async () => {
  const bookBorrow = {
    added : added + ":00+07:00",
    book : bookID,
    user : userID,
    purpose : purposeID,
  };
  console.log(bookBorrow);
  const res: any = await api.createBookborrow({ bookborrow : bookBorrow });
  setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }
};
return (
  <Page theme={pageTheme.home}>
    <Header
      title={`${profile.givenName}`}
    //subtitle="Some quick intro and links."
    ></Header>
    <Content>
      <ContentHeader title="เพิ่มข้อมูลการยืมหนังสือ">
        <Typography align="left" style={{ marginRight: 16, color: "#00eeff" }}>
          {username.givenuser}
        </Typography>
        <div>
          <Button variant="contained" color="primary">
            ออกจากระบบ
       </Button>
        </div>
        {status ? (
          <div>
            {alert ? (
              <Alert severity="success">
                This is a success alert — check it out!
              </Alert>
            ) : (
                <Alert severity="warning" style={{ marginTop: 20 }}>
                  This is a warning alert — check it out!
                </Alert>
              )}
          </div>
        ) : null}
      </ContentHeader>
      <div className={classes.root}>
        <form noValidate autoComplete="off">
        <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <InputLabel id="book-label">หนังสือ</InputLabel>
              <Select
                labelId="book-label"
                id="book"
                value={bookID}
                onChange={bookIDhandleChange}
                style={{ width: 400 }}
              >
              {books.map((item: EntBook) => (
                  <MenuItem value={item.id}>{item.bOOKNAME}</MenuItem>
                ))}
              </Select>
            </FormControl>
          <div>
            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <InputLabel id="user-label">สมาชิกห้องสมุด</InputLabel>
              <Select
                labelId="user-label"
                id="user"
                value={userID}
                onChange={UserIDhandleChange}
                style={{ width: 400 }}
              >
              {users.map((item: EntUser) => (
                <MenuItem value={item.id}>{item.uSERNAME}</MenuItem>
              ))}
              </Select>
            </FormControl>
          </div>
          <FormControl
            className={classes.margin}
            variant="outlined"
          >
            <InputLabel id="purpose">วัตถุประสงค์</InputLabel>
            <Select
              labelId="purpose"
              id="purpose"
              value={purposeID}
              onChange={PurposeIDhandleChange}
              style={{ width: 200 }}
            >
              {purposes.map((item: EntPurpose) => (
                <MenuItem value={item.id}>{item.pURPOSENAME}</MenuItem>
              ))}
            </Select>
          </FormControl>
          <div>
            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <TextField
                id="deathtime"
                label="ว/ด/ป เวลายืม"
                type="datetime-local"
                value={added}
                onChange={AddedhandleChange}
                className={classes.textField}
                InputLabelProps={{
                  shrink: true,
                }}
              />
            </FormControl>
          </div>
          <div className={classes.margin}>
            <Button
              onClick={() => {
                createBookBorrow();
              }}
              variant="contained"
              color="primary"
            >
              [บันทึกการยืม]
           </Button>
            <Button
              style={{ marginLeft: 20 }}
              component={RouterLink}
              to="/welcome"
              variant="contained"
            >
              กลับ
           </Button>
          </div>
        </form>
      </div>
    </Content>
  </Page>
);
}