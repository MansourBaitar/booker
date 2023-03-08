import {Container, Typography, Box, Button} from "@mui/material";
import NotFoundIllustration from '../assets/illustrations/illustration_404.svg';
import { Link as RouterLink } from 'react-router-dom';

export const NotFoundPage = () => {
  return (
    <Container>
      <div className={'NotFoundWrapper'}>
        <Typography variant="h3" paragraph>
          Sorry, page not found!
        </Typography>

        <Typography sx={{ color: 'text.secondary' }}>
          Sorry, we couldn’t find the page you’re looking for. Perhaps you’ve mistyped the URL? Be sure to check your
          spelling.
        </Typography>

        <Box component="img" src={NotFoundIllustration} sx={{ height: 260, mx: 'auto', my: { xs: 5, sm: 10 } }}/>

        <Button to="/" size="large" variant="contained" component={RouterLink}>
          Go to Home
        </Button>
      </div>
    </Container>
  );
}