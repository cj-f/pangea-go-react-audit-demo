import { createTheme } from '@mui/material/styles';
import { red } from '@mui/material/colors';

// A custom theme for this app
const theme = createTheme({
  palette: {
    primary: {
      main: '#556cd6',
    },
    secondary: {
      main: '#19857b',
    },
    error: {
      main: red.A400,
    },
  },
  components: {
    MuiGrid: {
      styleOverrides: {
        root: {
          ".LinedPangeaDataGrid-root.MuiDataGrid-root": {
            ".MuiDataGrid-columnHeaders": {
              backgroundColor: "#fff",
            },
            ".MuiDataGrid-row.Mui-selected": {
              backgroundColor: "#E9ECEB",
              ".PangeaDataGrid-Pinned-Right": {
                backgroundColor: "#E9ECEB",
              },
              ":hover": {
                backgroundColor: "#E9ECEB",
                ".PangeaDataGrid-Pinned-Right": {
                  backgroundColor: "#E9ECEB",
                },
              },
            },
            ".MuiDataGrid-row": {
              ":hover": {
                backgroundColor: "#E9ECEB",
                ".PangeaDataGrid-Pinned-Right": {
                  backgroundColor: "#E9ECEB",
                },
              },
            },
            ".PangeaDataGrid-ExpansionRow, .PangeaDataGrid-Chip": {
              backgroundColor: "#E9ECEB",
            },
          },
        },
      },
    },
  },
});

export default theme;