import { ChatNavigation } from "../components/chatNavigation";
import { Box } from "@mui/joy";

export function Dashboard() {
  return (
    <Box sx={{ display: "flex", width: 1, flexDirection: "row" }}>
      <ChatNavigation />
      <Box sx={{ height: "10px" }}>
        <h1>hello</h1>
      </Box>
    </Box>
  );
}
