import {FC, useState} from "react";
import {Avatar, Box, Divider, IconButton, MenuItem, Popover, Stack, Typography} from "@mui/material";
import {useNavigate} from "react-router-dom";

const account = {
  displayName: 'Michael Jackson',
  email: 'michael@booker.io',
  photoURL: 'https://upload.wikimedia.org/wikipedia/commons/1/15/Michael_Jackson%2C_1988_%2846845017052%29.jpg',
};

const MENU_OPTIONS = [
  {
    label: 'Home',
    icon: 'eva:home-fill',
    location: '#'
  },
  {
    label: 'Profile',
    icon: 'eva:person-fill',
    location: '/account/profile'
  },
  {
    label: 'Settings',
    icon: 'eva:settings-2-fill',
    location: '#'
  },
];


export const AccountPopover: FC = (props: any) => {
  const navigate = useNavigate();
  const [open, setOpen] = useState(null);

  const handleOpen = (event: any) => {
    setOpen(event.currentTarget);
  };

  const handleClose = () => {
    setOpen(null);
  };

  const handleNavigate = (location: string) => {
    navigate(location, { replace: true });
  }

  const handleLogOut = () => {
    navigate('/auth/login', { replace: true });
  }
  return (
      <>
        <IconButton
            onClick={handleOpen}
            sx={{
              p: 0,
              zIndex: 99
            }}
        >
          <Avatar src={account.photoURL} alt="photoURL" />
        </IconButton>

        <Popover
            open={Boolean(open)}
            anchorEl={open}
            onClose={handleClose}
            anchorOrigin={{ vertical: 'bottom', horizontal: 'right' }}
            transformOrigin={{ vertical: 'top', horizontal: 'right' }}
            PaperProps={{
              sx: {
                p: 0,
                mt: 1.5,
                ml: 0.75,
                width: 180,
                '& .MuiMenuItem-root': {
                  typography: 'body2',
                  borderRadius: 0.75,
                },
              },
            }}
        >
          <Box sx={{ my: 1.5, px: 2.5 }}>
            <Typography variant="subtitle2" noWrap>
              {account.displayName}
            </Typography>
            <Typography variant="body2" sx={{ color: 'text.secondary' }} noWrap>
              {account.email}
            </Typography>
          </Box>

          <Divider sx={{ borderStyle: 'dashed' }} />

          <Stack sx={{ p: 1 }}>
            {MENU_OPTIONS.map((option) => (
                <MenuItem key={option.label} onClick={() => handleNavigate(option.location)}>
                  {option.label}
                </MenuItem>
            ))}
          </Stack>

          <Divider sx={{ borderStyle: 'dashed' }} />

          <MenuItem onClick={() => handleNavigate('/auth/login')} sx={{ m: 1 }}>
            Logout
          </MenuItem>
        </Popover>
      </>
  );
};