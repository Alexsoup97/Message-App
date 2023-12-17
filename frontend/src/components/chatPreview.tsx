import { Avatar, Box, Typography } from "@mui/joy";
import { useTheme } from "@mui/joy/styles";

type Props = {
  isSelected?: boolean;
  clickHandler: any;
};
export function ChatPreview({ isSelected, clickHandler }: Props) {
  const theme = useTheme();
  const boxStyle = {
    display: "flex",
    mx: 1,
    my: 1,
    gap: 2,
    borderRadius: "16px",
    cursor: "pointer",
    bgcolor: isSelected
      ? theme.vars.palette.primary.solidActiveBg
      : theme.vars.palette.neutral.softBg,
    "&:hover": {
      background: isSelected
        ? theme.vars.palette.primary.solidActiveBg
        : theme.vars.palette.neutral.solidHoverBg,
    },
  };

  return (
    <Box sx={boxStyle} onClick={clickHandler}>
      <Avatar sx={{ postion: "relative", top: 3, left: 2 }} size="md" />
      <Box>
        <Typography level="title-md">Group Chat Name</Typography>
        <Typography level="body-md">
          This is a preview message for the chat
        </Typography>
      </Box>
    </Box>
  );
}
