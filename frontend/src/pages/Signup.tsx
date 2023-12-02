import {FormControl, Typography, FormLabel, Link, Button, Sheet, Input} from '@mui/joy'
import { APIConstants } from '../utils/constants'
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import {CheckCircle} from '@mui/icons-material';
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

export function Signup (){

    const [errorMessage, setErrorMessage] = useState("")
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const [confirmPassword, setConfirmPassword] = useState("")
    const [isSuccess, setSuccess] = useState(false)
    const navigate = useNavigate()

    async function handleFormSubmit(e:React.ChangeEvent<HTMLFormElement>){
        e.preventDefault() 
        if(password !== confirmPassword){
            setErrorMessage("The passwords you entered do not match") 
            return
        }
        let formData = new FormData(e.target);
        await fetch(APIConstants.BackendUrl +"/api/user/create", {
            method: 'POST',
            body: formData
        }).then( async (resp)=>{
            
            if(resp.ok){
               setSuccess(true) 
            }else{
                let temp = await resp.json()
                setErrorMessage(temp.message)

            }
        })
    }

    if(!isSuccess){
        return(
            <form onSubmit={handleFormSubmit}>
                <Sheet variant="soft"  sx={sheetStyle}>
                    <Typography sx={{my:2}} level="h2">Sign Up</Typography>
                    {
                        errorMessage && 
                        <Typography sx={{mx:2}} color="danger" level="title-sm">{errorMessage}</Typography> 
                    }
                    <FormControl>
                    <FormLabel>Username</FormLabel>
                    <Input name="username" sx={{width:250}} onChange={(e)=>{setUsername(e.target.value)}}> </Input>
                    </FormControl>
    
                    <FormControl>
                        <FormLabel>Password</FormLabel>
                        <Input name="password"  type="password" sx={{width:250}} onChange={(e)=>{setPassword(e.target.value)}}/>
                    </FormControl>
    
                    <FormControl>
                        <FormLabel>Confirm Password</FormLabel>
                        <Input type="password" sx={{width:250}} onChange={(e)=>{setConfirmPassword(e.target.value)}}/>
                    </FormControl>
                    <Typography level="body-sm">
                        Already have an account?{' '}
                        <Link href="/" level="title-sm">
                            Sign in!
                        </Link>
                    </Typography>
                    <Button  type="submit"variant="outlined" sx={{my:2, width:170}} disabled={confirmPassword.length ==0 || username.length ==0 || password.length ==0 }>Create Account</Button>
                </Sheet>
            </form>
        )
    }else{
        return(
            <Sheet variant="soft" sx={sheetStyle}>
             <CheckCircle style={{fontSize:75, color:'green', outline:'white'}}/>
            <Typography level="h3" textAlign="center" sx={{my:3 }} >Account created successfully</Typography>
            <Button variant="outlined" onClick={() =>navigate("/")}>Return to Sign In</Button>
            </Sheet>
        )
    }
}