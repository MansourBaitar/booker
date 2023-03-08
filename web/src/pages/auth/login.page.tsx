import {FC, useState} from "react";
import style from './login.module.scss';
import LoginIllustration from '../../assets/illustrations/illustration_login.png';
import {
  Button,
  Checkbox,
  TextField,
  Container,
  Divider,
  IconButton,
  InputAdornment,
  Link,
  Stack,
  Typography,
  Alert
} from "@mui/material";
import Iconify from "../../components/iconify/iconify";
import {useNavigate} from "react-router-dom";
import { LoadingButton } from '@mui/lab';
import useResponsive from "../../hooks/useResponsive";
import * as api from '../../api/index';

export const LoginPage: FC = (props) => {
  const mdUp = useResponsive('up', 'md');
  const navigate = useNavigate();
  const [showPassword, setShowPassword] = useState(false);
  const [email, setEmail] = useState('');
  const [pwd, setPwd] = useState('');
  const [error, setError] = useState('');


  const handleClick = async () => {
    const response = await api.users.login(email, pwd);
    if(response.status === 500) {
      setError('Email or password are wrong, please try again!');
      return;
    }
    if(response.status === 200) navigate('/', {} );
  };

  return (
      <div className={style.loginWrapper}>
        {mdUp && (
            <div className={style.welcomeMessage}>
              <Typography variant="h3" sx={{ px: 5, mt: 10, mb: 5 }}>
                Hi, Welcome Back
              </Typography>
              <img src={LoginIllustration} alt="Login illustration"/>
            </div>
        )}

        <Container maxWidth="sm">
          <div className={style.loginContent}>
            <Typography id={'cloudText'} variant="h4" gutterBottom>
              Sign in to Booker
            </Typography>
            { error !== '' &&
                <Stack spacing={3}>
                    <Alert severity="error">{error}</Alert>
                </Stack>
            }


            <Stack direction="row" spacing={2}>

              <Button fullWidth size="medium" color="inherit" variant="outlined">
                <Iconify icon="mdi:microsoft-office" color="#DF3E30" width={22} height={22} />
                Office 365
              </Button>

            </Stack>

            <Divider sx={{ my: 3 }}>
              <Typography variant="body2" sx={{ color: 'text.secondary' }}>
                OR
              </Typography>
            </Divider>


            <Stack spacing={3}>
              <TextField name="email" label="Email address" onChange={(e) => setEmail(e.target.value)}/>

              <TextField
                  name="password"
                  label="Password"
                  type={showPassword ? 'text' : 'password'}
                  onChange={(e) => setPwd(e.target.value)}
                  InputProps={{
                    endAdornment: (
                        <InputAdornment position="end">
                          <IconButton onClick={() => setShowPassword(!showPassword)} edge="end">
                            <Iconify icon={showPassword ? 'eva:eye-fill' : 'eva:eye-off-fill'} />
                          </IconButton>
                        </InputAdornment>
                    ),
                  }}
              />
            </Stack>

            <Stack direction="row" alignItems="center" justifyContent="space-between" sx={{ my: 2 }}>
              <Typography sx={{ color: 'text.secondary', display: 'flex', justifyContent: 'center', alignItems: 'center' }}>
                <Checkbox name="remember" />Remember me
              </Typography>

              <Link variant="subtitle2" underline="hover">
                Forgot password?
              </Link>
            </Stack>

            <LoadingButton id={'customButton'} sx={{ bgcolor: '#2e31b5', color: 'white'}} fullWidth size="large" type="submit" variant="contained" onClick={handleClick}>
              Login
            </LoadingButton>

          </div>
        </Container>
      </div>
  );
}