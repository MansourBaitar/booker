import {FC} from "react";
import {Breadcrumbs, Container, Link, Stack, Typography} from "@mui/material";

export const RequestPage: FC = () => {
  return(
      <Container>
        <Stack mb={5}>
          <Typography variant="h4" gutterBottom>Requests</Typography>
          <Breadcrumbs separator="›" aria-label="breadcrumb">
            <Link underline="hover" color="inherit" href="/">Dashboard</Link>
            <Typography color="text.primary">Requests</Typography>
          </Breadcrumbs>
        </Stack>
      </Container>
  );
}