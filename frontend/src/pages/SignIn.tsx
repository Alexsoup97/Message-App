import { Sheet, FormControl, Input,
    FormLabel, Button, Typography, Link} from "@mui/joy"
import { APIConstants } from "../utils/constants"
import { useState } from "react"
import { useNavigate } from "react-router-dom"

const sheetStyle = {
            width: 300,
            mx: 'auto', // margin left & right
            my: 4, // margin top & bottom
            py: 3, // padding top & bottom
            px: 1  , // padding left & right
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
            justifyContent: 'center',
            gap: 2,
            borderRadius: 'sm',
            boxShadow: 'md',

}

export function SignIn(){

    const [errorMessage, setErrorMessage]  = useState("")
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")

    const nav = useNavigate()
    async function handleFormSubmit (e:React.ChangeEvent<HTMLFormElement>) {
        e.preventDefault()
        const formData = new FormData(e.target)

        await fetch(APIConstants.BackendUrl +"/api/user/", {
            method: 'POST',
            body: formData

        }).then((resp)=>{
            if(resp.ok){
                nav("/dashboard")
            }else{
                setErrorMessage("We didn't recognize the username or password you entered. Please try again.")
            }
        }).catch((error) =>{
            console.error(error)
        })
    }
    function updateUser(e: React.ChangeEvent<HTMLInputElement>){
        setUsername(e.target.value)
    }

    function updatePassword(e: React.ChangeEvent<HTMLInputElement>){
        setPassword(e.target.value)
    }

    return(
        <form onSubmit={handleFormSubmit}>
            <Sheet variant="soft"  sx={sheetStyle}>
                <Typography sx={{my:2}} level="h2">Sign In</Typography>
                {
                    errorMessage && 
                    <Typography sx={{mx:2}} color="danger" level="title-sm">{errorMessage}</Typography> 
                }
                <FormControl>
                <FormLabel>Username</FormLabel>
                <Input name="username" sx={{width:250}} onChange={updateUser}> </Input>
                </FormControl>
                <FormControl>
                    <FormLabel>Password</FormLabel>
                    <Input name="password" onChange={updatePassword} type="password" sx={{width:250}}></Input>
                    </FormControl>
                <Typography level="body-sm">
                    Don't have an account?{' '}
                    <Link href="signup" level="title-sm">
                        Sign up!
                    </Link>
                </Typography>
                <Button disabled={password === "" || username === ""} type="submit"variant="outlined" sx={{my:2, width:125}}>Sign In </Button>
            </Sheet>
        </form>

    )
}