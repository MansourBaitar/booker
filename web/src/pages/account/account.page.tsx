import {Breadcrumbs, Container, Link, Stack, Typography} from "@mui/material";
import {FC} from "react";

export const AccountPage: FC = () => {
  return (
      <Container>
        <Stack mb={5} maxWidth="xl">
          <Typography variant="h4" sx={{ mb: 5 }} gutterBottom>Account</Typography>
          <Breadcrumbs separator="›" aria-label="breadcrumb">
            <Link underline="hover" color="inherit" href="/">Dashboard</Link>
            <Typography color="text.primary">Account</Typography>
          </Breadcrumbs>
        </Stack>
      </Container>
  );
}