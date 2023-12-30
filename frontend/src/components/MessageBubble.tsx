import { Sheet, Typography } from "@mui/joy";

export function MessageBubble({ message }: any) {
  return (
    <Sheet variant="outlined" sx={{ px: 1.75, py: 1.25, borderRadius: "lg" }}>
      <Typography>{message}</Typography>
    </Sheet>
  );
}
