import * as React from "react";
import AppBar from "@mui/material/AppBar";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import Container from "@mui/material/Container";
import Button from "@mui/material/Button";
import LocalHospitalIcon from "@mui/icons-material/LocalHospital";
import { Link as RouterLink } from "react-router-dom";

function ResponsiveAppBar_02() {
  const [anchorElNav, setAnchorElNav] = React.useState<null | HTMLElement>(
    null
  );
  const [anchorElUser, setAnchorElUser] = React.useState<null | HTMLElement>(
    null
  );

  return (
    <AppBar position="static" sx={{ bgcolor: "#024142" }}>
      <Container maxWidth="xl">
        <Toolbar>
          <LocalHospitalIcon
            fontSize="large"
            sx={{ display: { xs: "none", md: "flex" }, mr: 1 }}
          />
          <Typography
            variant="h6"
            noWrap
            component="a"
            href="/HomePage2"
            sx={{
              mr: 2,
              display: { xs: "none", md: "flex" },
              fontWeight: 700,
              color: "inherit",
              textDecoration: "none",
            }}
          >
            HOME
          </Typography>
          <Typography
            noWrap
            component="a"
            href="/SymptomCreate"
            sx={{
              mr: 2,
              display: { xs: "none", md: "flex" },
              fontWeight: 700,
              color: "inherit",
              textDecoration: "none",
            }}
          >
            ติดตามอาการคนไข้
          </Typography>
          <Typography
            noWrap
            component="a"
            href="/Manage"
            sx={{
              mr: 2,
              display: { xs: "none", md: "flex" },
              fontWeight: 700,
              color: "inherit",
              textDecoration: "none",
            }}
          >
            โภชนาการสำหรับคนไข้ใน
          </Typography>
          <Typography
            noWrap
            component="a"
            href="/BASKETCreate"
            sx={{
              mr: 2,
              display: { xs: "none", md: "flex" },
              fontWeight: 700,
              color: "inherit",
              textDecoration: "none",
            }}
          >
            การจ่ายยาสำหรับคนไข้ใน
          </Typography>
          <Typography
            sx={{
              mr: 2,
              display: { xs: "none", md: "flex" },
              fontWeight: 700,
              color: "inherit",
              textDecoration: "none",
              flexGrow: 1,
            }}
          />
          <Button
            sx={{ backgroundColor: "#003D2E" }}
            variant="contained"
            component={RouterLink}
            to="/"
          >
            LOGOUT
          </Button>
        </Toolbar>
      </Container>
    </AppBar>
  );
}
export default ResponsiveAppBar_02;
