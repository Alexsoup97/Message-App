import { ChatNavigation } from "../components/chatNavigation";
import { Box } from "@mui/joy";

export function Dashboard() {
  return (
    <Box sx={{ height: "100vh", display: "flex", flexDirection: "row" }}>
      <ChatNavigation />
    </Box>
  );
}
