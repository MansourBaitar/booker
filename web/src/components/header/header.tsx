import {FC} from "react";
import {Box, IconButton, Stack} from "@mui/material";
import Iconify from "../iconify/iconify";
import {AccountPopover} from "../popover/account.popover";
import {LanguagePopover} from "../popover/language.popover";

interface NavHeaderProps {
  onOpen: () => void;
}



export const NavHeader: FC<NavHeaderProps> = (props) => {
  return (
      <>
        <IconButton
            onClick={props.onOpen}
            sx={{
              mr: 1,
              color: 'white',
              display: { lg: 'none' },
              zIndex: 99
            }}
        >
          <Iconify icon="eva:menu-2-fill" />
        </IconButton>

        <Box sx={{ flexGrow: 1 }} />
        <Stack direction="row" alignItems="center" spacing={{xs: 0.5, sm: 1 }}>
          <LanguagePopover/>
          <AccountPopover />
        </Stack>
      </>
  );
}