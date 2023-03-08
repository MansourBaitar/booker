import {FC, useEffect, useState} from "react";
import {
  Badge, Box,
  Breadcrumbs, Button,
  Card,
  Container, Divider, Drawer,
  Link,
  Paper,
  Stack,
  Table, TableBody, TableCell,
  TableContainer,
  TableHead, TableRow, TextField,
  Typography
} from "@mui/material";
import {Holiday} from "../model";
import * as api from "../api";
import {LocalizationProvider} from "@mui/x-date-pickers";
import dayjs, {Dayjs} from "dayjs";
import {AdapterDayjs} from "@mui/x-date-pickers/AdapterDayjs";
import {DateRange, DesktopDateRangePicker} from "@mui/x-date-pickers-pro";
// @ts-ignore
import dayjsBusinessDays from "dayjs-business-days";

export const HolidayPage: FC = () => {
  dayjs.extend(dayjsBusinessDays);

  const [holidays, setHolidays] = useState<Holiday[] | Record<any, any>>([]);
  const [drawerOpen, setDrawerOpen] = useState(false);
  const [dateRange, setDateRange] = useState<DateRange<Dayjs>>([null, null]);
  const [id, setId] = useState<number>(0);
  const [status, setStatus] = useState<string>('pending');

  useEffect(() => {
    load().catch((e) => {
      console.error(e);
    })
  }, [])
  const handleSubmit = async () => {
    try {
      if(dateRange[0] === null) {
        console.log('debug', 'daterange 0 is null', dateRange[0])
        return;
      }
      if(dateRange[1] === null) {
        console.log('debug', 'daterange 1 is null', dateRange[1])
        return;
      }
      const x = await api.holidays.createHolidays(id,'holiday', '34225', 'Mansour Baitar', status, dateRange[0], dateRange[1])
      console.log(x);
    } catch (e) {
      console.error(e);
    }
  }
  const calculateDays = () => {
    if(dateRange[0] !== null && dateRange[1] !== null ) {
      // @ts-ignore
      const x = dateRange[1].businessDiff(dateRange[0]);
      return x+1;
    }
  }
  async function load() {
    try {
      const req = await api.holidays.getHolidays()
      setHolidays(req);
    } catch (e) {
      console.error("[APP] something went wrong");
    }
  }
  const getStatusColor = (status: string): any => {
    const colorMapping = {
      pending: "warning",
      declined: "error",
      approved: "success"
    }
    // @ts-ignore
    return colorMapping[status];
  }
  const toggleDrawer = () => {
    setDrawerOpen(!drawerOpen);
  }

  return (
      <Container>
        <Stack direction="row" justifyContent="space-between" mb={5}>
          <Stack>
            <Typography variant="h4" gutterBottom>Holidays</Typography>
            <Breadcrumbs separator="›" aria-label="breadcrumb">
              <Link underline="hover" color="inherit" href="/">Dashboard</Link>
              <Typography color="text.primary">Holidays</Typography>
            </Breadcrumbs>
          </Stack>
          <div>
            <Button color="info" size="medium" variant="contained" onClick={() => toggleDrawer()}>
              Request holiday
            </Button>
          </div>
        </Stack>

        {holidays && holidays.length !== 0 && (
          <Card>
            <TableContainer component={Paper}>
              <Table sx={{minWidth: 650}} aria-label="simple table">
                <TableHead>
                  <TableRow>
                    <TableCell>Name</TableCell>
                    <TableCell align="center">From</TableCell>
                    <TableCell align="center">To</TableCell>
                    <TableCell align="center">Status</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {holidays.map((row: any) => (
                      <TableRow
                          key={row.id}
                          sx={{'&:last-child td, &:last-child th': {border: 0}}}
                      >
                        <TableCell component="th" scope="row">{row.name}</TableCell>
                        <TableCell align="center">{row.fromDate}</TableCell>
                        <TableCell align="center">{row.toDate}</TableCell>
                        <TableCell align="center">
                          <Badge badgeContent={row.status} color={getStatusColor(row.status)}></Badge>
                        </TableCell>
                      </TableRow>
                  ))}
                </TableBody>
              </Table>
            </TableContainer>
          </Card>
        )}

        <Drawer
            PaperProps={{
              sx: {
                width: 350,
                height: '100%',
                maxHeight: '100vh'
              }
            }}
            anchor={"right"}
            open={drawerOpen}
            onClose={() => toggleDrawer()}
        >
          <Stack sx={{
            margin: '2rem',
            height: '100%',
            overflow: 'hidden'
          }}>
            <Stack>
              <Typography variant="h6" gutterBottom>Request Holiday</Typography>
            </Stack>
            <Stack spacing={2}>
              <LocalizationProvider dateAdapter={AdapterDayjs} localeText={{start: 'Start date', end: 'End date'}}>
                <DesktopDateRangePicker
                    value={dateRange}
                    onChange={(newValue: any) => {
                      setDateRange(newValue);
                      console.log(newValue);
                    }}
                    renderInput={(startProps: any, endProps: any) => (
                        <>
                          <div className={"date-range-wrapper"}>
                            <TextField className={"date-picker"} {...startProps} />
                            <Divider sx={{ my: 3, mx: 2, color: 'pink' }}>
                              <Typography variant="body2" sx={{ color: '#000000' }}>
                                To
                              </Typography>
                            </Divider>
                            <TextField className={"date-picker"} {...endProps} />
                          </div>
                        </>
                    )}
                />
              </LocalizationProvider>
            </Stack>

            <Stack direction={'row'} alignItems="center" sx={{width: '100vw', maxWidth: '100vw'}}>
              <h1>
                {calculateDays()}
              </h1>
            </Stack>

            <Stack direction={'row'} alignItems="flex-end" sx={{height: '100vh', maxHeight: '100vh'}}>
              <Button variant="contained" onClick={() => handleSubmit()}>Submit</Button>
              <Button variant="text" onClick={() => toggleDrawer()}>Cancel</Button>
            </Stack>
          </Stack>
        </Drawer>
      </Container>
  );
}