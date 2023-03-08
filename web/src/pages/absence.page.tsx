import {FC, useEffect, useState} from "react";
import {
  Badge,
  Box,
  Breadcrumbs, Button,
  Card,
  Container,
  Link,
  Paper,
  Stack,
  Table, TableBody, TableCell,
  TableContainer,
  TableHead, TableRow,
  Typography
} from "@mui/material";
import * as api from '../api'
import {Absence} from "../model";

export const AbsencePage: FC = () => {
  const [absences, setAbsences] = useState<Absence[]>([]);
  async function load() {
    try {
      const req = await api.absences.getAbsences()
      console.log(req);
      // @ts-ignore
      setAbsences(req);
    } catch (e) {
      console.error("[APP] something went wrong");
    }
  }

  useEffect(() => {
    load().catch((e) => {
      console.error(e);
    })
  }, [])

  const getStatusColor = (status: string): any => {
    const colorMapping = {
      pending: "warning",
      declined: "error",
      approved: "success"
    }
    // @ts-ignore
    return colorMapping[status];
  }

  return(
      <Container>
        <Stack direction="row" justifyContent="space-between" mb={5}>
          <Stack>
            <Typography variant="h4" gutterBottom>Absences</Typography>
            <Breadcrumbs separator="›" aria-label="breadcrumb">
              <Link underline="hover" color="inherit" href="/">Dashboard</Link>
              <Typography color="text.primary">Absences</Typography>
            </Breadcrumbs>
          </Stack>

          <div>
            <Button color="info" size="medium" variant="contained">
              Register absence
            </Button>
          </div>
        </Stack>

        {absences && absences.length !== 0 && (
        <Card>
          <TableContainer component={Paper}>
            <Table sx={{ minWidth: 650 }} aria-label="simple table">
              <TableHead>
                <TableRow>
                  <TableCell>Name</TableCell>
                  <TableCell align="right">From</TableCell>
                  <TableCell align="right">To</TableCell>
                  <TableCell align="right">Status</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {absences.map((row: any) => (
                    <TableRow
                        key={row.name}
                        sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                    >
                      <TableCell component="th" scope="row">{row.name}</TableCell>
                      <TableCell align="right">{row.fromDate}</TableCell>
                      <TableCell align="right">{row.toDate}</TableCell>
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

      </Container>
  )
}