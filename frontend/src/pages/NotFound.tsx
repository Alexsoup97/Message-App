import { Box, Typography } from "@mui/joy";

export function ErrorPage() {
  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        alignContent: "center",
        justifyContent: "center",
        height: "100vh",
        gap: 1,
      }}
    >
      <Typography textAlign="center" level="h1">
        Oops!
      </Typography>
      <Typography textAlign="center" level="h4">
        Sorry, an unexpected error has occurred.
      </Typography>
      <Typography textAlign="center">Page not found</Typography>
    </Box>
  );
}
