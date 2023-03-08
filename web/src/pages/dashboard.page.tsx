import {FC} from "react";
import {Container, Grid, Stack, Typography} from "@mui/material";
import {DashboardWidget} from "../components/widgets/dashboard-widget";

const INFO = {
  lighter: '#D0F2FF',
  light: '#74CAFF',
  main: '#1890FF',
  dark: '#0C53B7',
  darker: '#04297A',
  contrastText: '#fff',
};

const WARNING = {
  lighter: '#FFF7CD',
  light: '#FFE16A',
  main: '#FFC107',
  dark: '#B78103',
  darker: '#7A4F01',
  contrastText: '#212B36',
};

const ERROR = {
  lighter: '#FFE7D9',
  light: '#FFA48D',
  main: '#FF4842',
  dark: '#B72136',
  darker: '#7A0C2E',
  contrastText: '#fff',
};


export const DashboardPage: FC = () => {
  return(
      <Container maxWidth="xl">
        <Stack mb={5}>
          <Typography variant="h4" gutterBottom>Dashboard</Typography>
        </Stack>

        <Grid container spacing={3}>
          <Grid item xs={12} sm={6} md={3}>
            <DashboardWidget title="Total Agents" total={714000} icon={'ant-design:android-filled'} />
          </Grid>

          <Grid item xs={12} sm={6} md={3}>
            <DashboardWidget title="Total Applications" total={1352838} color={INFO} icon={'ant-design:apple-filled'} />
          </Grid>

          <Grid item xs={12} sm={6} md={3}>
            <DashboardWidget title="Enabled features" total={1723315} color={WARNING} icon={'ant-design:windows-filled'} />
          </Grid>

          <Grid item xs={12} sm={6} md={3}>
            <DashboardWidget title="Stopped applications" total={234} color={ERROR} icon={'ant-design:bug-filled'} />
          </Grid>
        </Grid>
      </Container>
  )
}