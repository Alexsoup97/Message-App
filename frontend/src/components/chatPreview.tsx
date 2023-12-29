import { Avatar, Box, Typography } from "@mui/joy";
import { useTheme } from "@mui/joy/styles";

type Props = {
  isSelected?: boolean;
  clickHandler: any;
  name: string;
  previewMessage: string;
};
export function ChatPreview({
  isSelected,
  clickHandler,
  name,
  previewMessage,
}: Props) {
  const theme = useTheme();
  const boxStyle = {
    display: "flex",
    mx: 1,
    my: 1,
    gap: 2,
    height: "3em",
    borderRadius: "16px",
    cursor: "pointer",
    overflow: "auto",
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
      <Box sx={{ marginTop: "3px" }}>
        <Typography level="title-md">{name}</Typography>
        <Typography level="body-sm" sx={{ marginTop: "-6px" }} noWrap>
          {previewMessage}
        </Typography>
      </Box>
    </Box>
  );
}
