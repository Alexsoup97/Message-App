import { ConversationWindow } from "../components/ConversationWindow";
import { ChatNavigation } from "../components/chatNavigation";
import { Box } from "@mui/joy";

export function Dashboard() {
  return (
    <Box sx={{ display: "flex", width: 1, flexDirection: "row" }}>
      <ChatNavigation />
      <Box sx={{ height: "10px" }}>
        <ConversationWindow></ConversationWindow>
      </Box>
    </Box>
  );
}
